package api_list_backup_snapshots

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeapi"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/snapshots"
)

// Handle answers GET /api/admin/root/list-backup-snapshots: every
// snapshot, newest first, as {"snapshots": [...], "busy": ...}, busy telling
// whether a snapshot or a restore is running.
func Handle(sandbox *api.Sandbox, props *routeprops.RouteProps, input *Input, response *serverdeps.Response) error {
	if props.User == nil {
		return sandbox.Deps.OpinionatedAgnosServer.Fail(api.StatusUnauthorized, "", "no authenticated user")
	}

	listed, err := snapshots.List(sandbox)
	if err != nil {
		return err
	}
	return sandbox.Deps.OpinionatedAgnosServer.WriteJSON(sandbox.Deps.SerializableDeps, *response, api.StatusOK, backofficeapi.SnapshotListJSON(sandbox, listed, snapshots.Busy(sandbox)))
}
