package edit_backoffice_user_page

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeusers"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/render"
)

// InternalPureHandler answers GET /admin/root/edit-backoffice-user/{id} with
// the form that POST /admin/root/edit-backoffice-user/{id} reads, filled with
// the user's current username, email and role. A user that does not exist sends
// the browser back to the list.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	user, ok := backofficeusers.Find(sandbox, int64(entries.Id))
	if !ok {
		return routeio.Redirect(*response, api.StatusSeeOther, backofficeusers.ListLocation(sandbox, backofficeusers.NoticeNotFound))
	}
	fields := backofficeusers.Fields{Username: user.Username, Email: user.Email, Role: user.Role}
	return render.EditBackofficeUserForm(sandbox, response, api.StatusOk, props.User, user.Id, fields, "")
}
