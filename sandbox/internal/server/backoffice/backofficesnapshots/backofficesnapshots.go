package backofficesnapshots

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/backup"
)

// The backups of the backoffice: the pages and the api through which a root
// takes, downloads, uploads and restores the snapshots of sandbox/internal/snapshots.
// Every one of them is under /admin/root or /api/admin/root, so only a root
// reaches it: a restore replaces every database, and an archive carries
// every password hash and session.

// ListPath is the page every action on a snapshot sends the browser back to.
const ListPath = "/admin/root/list-backup-snapshots"

// The notices an action hands the list page through ListPath?notice=<code>.
// Each one is a fixed word the page words itself, so nothing the client sent
// is ever shown back. NoticeBusy, NoticeNotFound and NoticeNotReady spell
// the snapshots outcome of the same meaning.
const (
	// NoticeCreating follows a snapshot started.
	NoticeCreating = "creating"
	// NoticeRestoring follows a restore started.
	NoticeRestoring = "restoring"
	// NoticeUploaded follows an archive uploaded.
	NoticeUploaded = "uploaded"
	// NoticeBusy follows a create or a restore refused while another runs.
	NoticeBusy = "busy"
	// NoticeNotFound follows an action on a snapshot that does not exist.
	NoticeNotFound = "not-found"
	// NoticeNotReady follows a download or a restore of a snapshot that is
	// not ready.
	NoticeNotReady = "not-ready"
)

// StatusAccepted is the status of a job the api started and does not wait
// for; api has no constant of its own for it.
const StatusAccepted = 202

// ListLocation is the list page showing the notice code notice.
func ListLocation(sandbox *api.Sandbox, notice string) string {
	return ListPath + "?notice=" + notice
}

// WriteArchive answers archive, the zip of snapshot, as a file to save under
// the snapshot's name. A snapshot name holds no character a header or a file
// name has to escape: snapshots names every one of them.
func WriteArchive(sandbox *api.Sandbox, response *serverdeps.Response, snapshot backup.SnapshotRecord, archive []byte) error {
	response.SetHeader("Content-Type", "application/zip")
	response.SetHeader("Content-Disposition", "attachment; filename=\""+snapshot.Name+".zip\"")
	response.SetStatus(api.StatusOK)
	return response.Write(archive)
}
