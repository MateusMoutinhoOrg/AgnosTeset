package snapshots

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/backup"
)

// A snapshot is every file of every database under DataDir — the backup
// database itself, under BackupDir, aside — as it was at one instant, kept in
// the backup database: the content of each file once in the blob table,
// indexed by its SHA-256, and the snapshot a list of (path, sha) pairs under
// its content. A file that did not change between two snapshots is stored
// once, whatever the number of snapshots holding it.
//
// Creating one, restoring one and removing the blobs no snapshot holds are
// slow, so a route only starts them: the work runs on a goroutine of its own
// once the route has answered. A snapshot is recorded the moment it is
// started, under StatusCreating, and turns StatusReady or StatusFailed when
// its job ends. Every job — those three, an upload and a removal — runs one
// at a time.

// DataDir is the directory every database of the project lives under, one
// directory per database.
const DataDir = "data"

// BackupDir is the directory of DataDir the backup database lives in. It is
// never put in a snapshot, and a restore never touches it.
const BackupDir = "backup"

// The status of a snapshot.
const (
	// StatusCreating is a snapshot whose files are still being stored.
	StatusCreating = "creating"
	// StatusReady is a snapshot holding every file it was taken with.
	StatusReady = "ready"
	// StatusFailed is a snapshot whose job failed or was interrupted: it
	// holds part of its files at most, so it is never restored nor
	// downloaded.
	StatusFailed = "failed"
)

// The outcomes StartRestore, Export, Import and Remove answer with.
const (
	// OutcomeOk is a restore started, an archive built or stored, or a
	// snapshot removed.
	OutcomeOk = "ok"
	// OutcomeBusy is a job refused because another one runs.
	OutcomeBusy = "busy"
	// OutcomeRefused is an archive that is not one a snapshot can be read
	// from.
	OutcomeRefused = "refused"
	// OutcomeNotFound is an id no snapshot holds.
	OutcomeNotFound = "not-found"
	// OutcomeNotReady is a snapshot that is not StatusReady.
	OutcomeNotReady = "not-ready"
)

// MaxArchiveBytes is the largest archive an upload may send.
const MaxArchiveBytes = 256 << 20

// MaxUnpackedBytes is the most an uploaded archive may hold once unpacked.
const MaxUnpackedBytes = 1 << 30

// NameLayout is how the instant a snapshot was taken at is spelled in its
// name, after its prefix.
const NameLayout = "20060102-150405"

// keepLockSuffix ends the lock files the store leaves beside a key while it
// writes it.
const keepLockSuffix = ".keeplock"

// jobs holds a token while a job runs, so a second one is refused rather than
// run over the first: a create and a restore read and write every database
// of DataDir, and an optimize removes every blob no (path, sha) pair names —
// one a create or an upload is about to name included.
var jobs = make(chan struct{}, 1)

// acquire takes the job token and reports whether it was free.
func acquire(sandbox *api.Sandbox) bool {
	select {
	case jobs <- struct{}{}:
		return true
	default:
		return false
	}
}

// release gives the job token back.
func release(sandbox *api.Sandbox) {
	<-jobs
}

// Busy tells whether a job is running.
func Busy(sandbox *api.Sandbox) bool {
	return len(jobs) > 0
}

// nowSeconds is the current instant in Unix seconds.
func nowSeconds(sandbox *api.Sandbox) int64 {
	return sandbox.Deps.StdDeps.Now() / 1_000_000_000
}

// StartCreate records a new snapshot under StatusCreating and answers it at
// once, leaving its files to a goroutine that runs after the caller returns.
// started is false, and nothing is recorded, when another job runs.
func StartCreate(sandbox *api.Sandbox) (snapshot backup.SnapshotRecord, started bool, err error) {
	if !acquire(sandbox) {
		return backup.SnapshotRecord{}, false, nil
	}
	db := backup.New(sandbox)
	snapshot, err = begin(sandbox, db, "snapshot")
	if err != nil {
		release(sandbox)
		return backup.SnapshotRecord{}, true, err
	}
	go func() {
		defer release(sandbox)
		finish(sandbox, db, snapshot)
	}()
	return snapshot, true, nil
}

// create takes one snapshot named after prefix and waits for it. The caller
// holds the job token.
func create(sandbox *api.Sandbox, db *backup.Backup, prefix string) (backup.SnapshotRecord, error) {
	snapshot, err := begin(sandbox, db, prefix)
	if err != nil {
		return snapshot, err
	}
	if err := finish(sandbox, db, snapshot); err != nil {
		return snapshot, err
	}
	snapshot.Status = StatusReady
	return snapshot, nil
}

// begin records a new snapshot under StatusCreating, named prefix followed
// by the current instant.
func begin(sandbox *api.Sandbox, db *backup.Backup, prefix string) (backup.SnapshotRecord, error) {
	now := nowSeconds(sandbox)
	name := uniqueName(sandbox, db, prefix+"-"+sandbox.Deps.TimeDeps.FormatUnix(now, NameLayout))
	return db.AddSnapshot(backup.SnapshotInput{Name: name, Data: now, Status: StatusCreating})
}

// finish stores every file of DataDir under snapshot and records how that
// ended: StatusReady, or StatusFailed with the cause on stderr. A panic is
// caught and fails the snapshot, so a job never takes the server down.
func finish(sandbox *api.Sandbox, db *backup.Backup, snapshot backup.SnapshotRecord) (err error) {
	files := 0
	defer func() {
		if recovered := recover(); recovered != nil {
			err = sandbox.Deps.StdDeps.Errorf("%v", recovered)
		}
		status := StatusReady
		if err != nil {
			status = StatusFailed
			sandbox.Deps.StdDeps.Eprintf("snapshot %s failed: %s\n", snapshot.Name, err.Error())
		} else {
			sandbox.Deps.StdDeps.Logf("snapshot %s ready: %d files\n", snapshot.Name, files)
		}
		if setErr := db.SetSnapshotStatus(snapshot.Id, status); setErr != nil {
			sandbox.Deps.StdDeps.Eprintf("snapshot %s: recording it %s failed: %s\n", snapshot.Name, status, setErr.Error())
			if err == nil {
				err = setErr
			}
		}
	}()
	files, err = fill(sandbox, db, snapshot.Id)
	return err
}

// fill stores every file collect finds under the snapshot of id, one file at
// a time, and answers how many it stored. A file removed between the listing
// and its reading is left out: the store removes one whenever a record goes.
func fill(sandbox *api.Sandbox, db *backup.Backup, id int64) (int, error) {
	io := sandbox.Deps.IoDeps
	files := 0
	for _, path := range collect(sandbox) {
		content, err := io.ReadFile(path)
		if err != nil {
			if !io.Exists(path) {
				continue
			}
			return files, err
		}
		sha, err := storeBlob(sandbox, db, content)
		if err != nil {
			return files, err
		}
		if _, err := db.AddSnapshotContent(id, backup.SnapshotContentInput{Path: relative(sandbox, path), Sha: sha}); err != nil {
			return files, err
		}
		files++
	}
	return files, nil
}

// collect is the path of every file a snapshot holds: every file under every
// database directory of DataDir but BackupDir, in lexical order, leaving out
// the temporary and lock files the store keeps while it writes.
func collect(sandbox *api.Sandbox) []string {
	io := sandbox.Deps.IoDeps
	backupDir := io.Join(DataDir, BackupDir)
	paths := []string{}
	for _, dir := range io.ListDirs(DataDir) {
		if dir == backupDir {
			continue
		}
		for _, path := range io.ListFilesRecursively(dir) {
			if kept(sandbox, path) {
				paths = append(paths, path)
			}
		}
	}
	return paths
}

// kept tells whether the file at path belongs in a snapshot: the store names
// its temporary files with a leading "." and its locks with keepLockSuffix,
// which no file of a record ever carries.
func kept(sandbox *api.Sandbox, path string) bool {
	strings := sandbox.Deps.StringsDeps
	slashed := strings.ReplaceAll(path, "\\", "/")
	name := slashed[strings.LastIndex(slashed, "/")+1:]
	return !strings.HasPrefix(name, ".") && !strings.HasSuffix(name, keepLockSuffix)
}

// relative is path, a host path under DataDir, as a snapshot stores it:
// slash-separated and relative to DataDir.
func relative(sandbox *api.Sandbox, path string) string {
	strings := sandbox.Deps.StringsDeps
	return strings.TrimPrefix(strings.ReplaceAll(path, "\\", "/"), DataDir+"/")
}

// storeBlob stores content once, under its SHA-256, and answers that sha.
// Content already stored is not stored again; an insert that lost a race
// against another one storing the same content reads as stored.
func storeBlob(sandbox *api.Sandbox, db *backup.Backup, content []byte) (string, error) {
	sha := sandbox.Deps.HashDeps.Sha256Hex(content)
	if _, ok := db.FindBlobBySha(sha); ok {
		return sha, nil
	}
	if _, err := db.AddBlob(backup.BlobInput{Sha: sha, Value: content}); err != nil {
		if _, ok := db.FindBlobBySha(sha); ok {
			return sha, nil
		}
		return "", err
	}
	return sha, nil
}

// uniqueName is base when no snapshot holds that name, base-2, base-3 and on
// otherwise: the first one free.
func uniqueName(sandbox *api.Sandbox, db *backup.Backup, base string) string {
	name := base
	for count := 2; ; count++ {
		if _, ok := db.FindSnapshotByName(name); !ok {
			return name
		}
		name = base + "-" + sandbox.Deps.StringsDeps.FormatInt(int64(count), 10)
	}
}

// List is every snapshot, newest first.
func List(sandbox *api.Sandbox) ([]backup.SnapshotRecord, error) {
	listed, err := backup.New(sandbox).ListSnapshots(backup.SnapshotFilter{})
	if err != nil {
		return nil, err
	}
	sandbox.Deps.SortDeps.SliceStable(listed, func(i int, j int) bool {
		if listed[i].Data != listed[j].Data {
			return listed[i].Data > listed[j].Data
		}
		return listed[i].Id > listed[j].Id
	})
	return listed, nil
}

// StartRecover runs RecoverInterrupted on a goroutine holding the job token,
// so no snapshot or restore starts before the ones a stopped process left
// are marked, and writes on stderr what it marked. It does not wait: a
// process killed while writing leaves its write lease behind, and the first
// write after it waits until that lease expires — up to a minute the server
// would otherwise spend not listening.
func StartRecover(sandbox *api.Sandbox) {
	if !acquire(sandbox) {
		return
	}
	go func() {
		defer release(sandbox)
		interrupted, err := RecoverInterrupted(sandbox)
		if err != nil {
			sandbox.Deps.StdDeps.Eprintf("warning: the snapshots left creating by the last run could not be marked failed: %s\n", err.Error())
		} else if interrupted > 0 {
			sandbox.Deps.StdDeps.Eprintf("warning: %d snapshot(s) were still being created when the last run stopped, and are marked failed\n", interrupted)
		}
	}()
}

// RecoverInterrupted marks StatusFailed every snapshot left StatusCreating by
// a process that stopped before its job ended, and answers how many it
// marked. The caller holds the job token, so none of them is being created.
func RecoverInterrupted(sandbox *api.Sandbox) (int, error) {
	db := backup.New(sandbox)
	stuck, err := db.ListSnapshots(backup.SnapshotFilter{StatusEquals: StatusCreating})
	if err != nil {
		return 0, err
	}
	for _, snapshot := range stuck {
		if err := db.SetSnapshotStatus(snapshot.Id, StatusFailed); err != nil {
			return 0, err
		}
	}
	return len(stuck), nil
}
