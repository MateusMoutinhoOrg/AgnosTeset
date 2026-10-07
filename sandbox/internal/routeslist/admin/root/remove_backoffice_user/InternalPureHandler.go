package remove_backoffice_user

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeusers"
)

// InternalPureHandler answers POST /admin/root/remove-backoffice-user/{id}: the
// user is removed, with every session of it, so its token is refused from here
// on, and the browser is sent to the list with the outcome. A root may not
// remove its own account.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return sandbox.Deps.OpinatedAgnosServer.Fail(api.StatusUnauthorized, "", "no authenticated user")
	}

	notice, err := backofficeusers.Remove(sandbox, *props.User, int64(entries.Id))
	if err != nil {
		return err
	}
	return sandbox.Deps.OpinatedAgnosServer.Redirect(*response, api.StatusSeeOther, backofficeusers.ListLocation(sandbox, notice))
}
