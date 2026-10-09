package backofficerender

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/backoffice_db"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/backup"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficesnapshots"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/snapshots"
)

// BackupSnapshotsPage is what backoffice/backup_snapshots.html is rendered
// with.
type BackupSnapshotsPage struct {
	Viewer Viewer
	// Notice is shown above the list, its Text "" for none.
	Notice    Notice
	Snapshots []BackupSnapshotRow
	// Busy disables what would start a second job: a snapshot or a restore
	// runs.
	Busy bool
	// Refresh reloads the page every few seconds, while a job runs or a
	// snapshot is still being created.
	Refresh bool
	// MaxUploadMegabytes is the largest archive the upload form takes.
	MaxUploadMegabytes int
}

// BackupSnapshotRow is one snapshot of the list.
type BackupSnapshotRow struct {
	Id   string
	Name string
	// Date is the instant the snapshot was taken at.
	Date       string
	Status     string
	IsReady    bool
	IsCreating bool
	IsFailed   bool
}

// backupSnapshotNoticeOf is how the list page words a backofficesnapshots
// notice code; an unknown code shows nothing.
func backupSnapshotNoticeOf(sandbox *api.Sandbox, code string) Notice {
	switch code {
	case backofficesnapshots.NoticeCreating:
		return Notice{Text: "Snapshot started. It is listed as creating until every database is copied, and this page refreshes until then.", Kind: "ok"}
	case backofficesnapshots.NoticeRestoring:
		return Notice{Text: "Restore started. A pre-restore snapshot of the current data is taken first, then every database is replaced: you may be signed out when it ends.", Kind: "ok"}
	case backofficesnapshots.NoticeUploaded:
		return Notice{Text: "Snapshot uploaded. It can be downloaded or restored now.", Kind: "ok"}
	case backofficesnapshots.NoticeBusy:
		return Notice{Text: "Another snapshot or restore is running. Try again once it ends.", Kind: "error"}
	case backofficesnapshots.NoticeNotFound:
		return Notice{Text: "That snapshot no longer exists.", Kind: "error"}
	case backofficesnapshots.NoticeNotReady:
		return Notice{Text: "That snapshot is not ready: only a finished snapshot can be downloaded or restored.", Kind: "error"}
	}
	return Notice{}
}

// backupSnapshotRow is the row of snapshot on the list page.
func backupSnapshotRow(sandbox *api.Sandbox, snapshot backup.SnapshotRecord) BackupSnapshotRow {
	return BackupSnapshotRow{
		Id:         sandbox.Deps.StringsDeps.FormatInt(snapshot.Id, 10),
		Name:       snapshot.Name,
		Date:       instantOf(sandbox, snapshot.Data),
		Status:     snapshot.Status,
		IsReady:    snapshot.Status == snapshots.StatusReady,
		IsCreating: snapshot.Status == snapshots.StatusCreating,
		IsFailed:   snapshot.Status == snapshots.StatusFailed,
	}
}

// RenderBackupSnapshotsPage answers, under status, the snapshot list page
// for user: listed, newest first, with the notice code notice above it, and
// what starts a job disabled while busy.
func RenderBackupSnapshotsPage(sandbox *api.Sandbox, response *serverdeps.Response, status int, user *backoffice_db.BackofficeUserRecord, listed []backup.SnapshotRecord, busy bool, notice string) error {
	rows := []BackupSnapshotRow{}
	refresh := busy
	for _, snapshot := range listed {
		row := backupSnapshotRow(sandbox, snapshot)
		refresh = refresh || row.IsCreating
		rows = append(rows, row)
	}

	return RenderHTML(sandbox, response, status, "backoffice/backup_snapshots.html", BackupSnapshotsPage{
		Viewer:             viewerOf(sandbox, user),
		Notice:             backupSnapshotNoticeOf(sandbox, notice),
		Snapshots:          rows,
		Busy:               busy,
		Refresh:            refresh,
		MaxUploadMegabytes: snapshots.MaxArchiveBytes >> 20,
	})
}
