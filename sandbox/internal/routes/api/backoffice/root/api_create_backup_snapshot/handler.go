package api_create_backup_snapshot

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeapi"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficesnapshots"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/snapshots"
)

// Handle answers POST /api/admin/root/create-backup-snapshot: a
// snapshot of every database is recorded as creating and answered 202 at
// once, while its files are stored in the background; list-backup-snapshots
// tells when it is ready. Another backup job running is answered 409.
func Handle(sandbox *api.Sandbox, props *routeprops.RouteProps, input *Input, response *serverdeps.Response) error {
	if props.User == nil {
		return sandbox.Deps.OpinionatedAgnosServer.Fail(api.StatusUnauthorized, "", "no authenticated user")
	}

	snapshot, started, err := snapshots.StartCreate(sandbox)
	if err != nil {
		return err
	}
	if !started {
		return sandbox.Deps.OpinionatedAgnosServer.WriteError(sandbox.Deps.SerializableDeps, *response, api.StatusConflict, "", backofficesnapshots.BusyMessage)
	}
	return sandbox.Deps.OpinionatedAgnosServer.WriteJSON(sandbox.Deps.SerializableDeps, *response, backofficesnapshots.StatusAccepted, backofficeapi.SnapshotCreatingJSON(sandbox, snapshot))
}
