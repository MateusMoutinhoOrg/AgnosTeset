package create_backup_snapshot_form

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficesnapshots"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/snapshots"
)

// Handle answers POST /admin/root/create-backup-snapshot: a snapshot
// of every database is recorded as creating and the browser is sent to the
// list at once, while its files are stored in the background. One already
// running — a snapshot or a restore — refuses it with a notice.
func Handle(sandbox *api.Sandbox, props *routeprops.RouteProps, input *Input, response *serverdeps.Response) error {
	if props.User == nil {
		return sandbox.Deps.OpinionatedAgnosServer.Fail(api.StatusUnauthorized, "", "no authenticated user")
	}

	_, started, err := snapshots.StartCreate(sandbox)
	if err != nil {
		return err
	}
	notice := backofficesnapshots.NoticeCreating
	if !started {
		notice = backofficesnapshots.NoticeBusy
	}
	return sandbox.Deps.OpinionatedAgnosServer.Redirect(*response, api.StatusSeeOther, backofficesnapshots.ListLocation(sandbox, notice))
}
