package api_upload_backup_snapshot

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeapi"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/snapshots"
)

// Handle answers POST /api/admin/root/upload-backup-snapshot, the
// zip archive a download built as the whole body: it is stored as a new
// ready snapshot, answered 201 as {"snapshot": ...}. An archive that is not
// one a download built is answered 400 with why.
func Handle(sandbox *api.Sandbox, props *routeprops.RouteProps, input *Input, response *serverdeps.Response) error {
	if props.User == nil {
		return sandbox.Deps.OpinionatedAgnosServer.Fail(api.StatusUnauthorized, "", "no authenticated user")
	}

	snapshot, refused, err := snapshots.Import(sandbox, input.Body)
	if err != nil {
		return err
	}
	if refused != "" {
		return sandbox.Deps.OpinionatedAgnosServer.Fail(api.StatusBadRequest, "body", refused)
	}
	return sandbox.Deps.OpinionatedAgnosServer.WriteJSON(sandbox.Deps.SerializableDeps, *response, api.StatusCreated, backofficeapi.SnapshotResponseJSON(sandbox, snapshot))
}
