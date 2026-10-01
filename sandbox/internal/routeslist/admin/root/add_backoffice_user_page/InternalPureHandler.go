package add_backoffice_user_page

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeusers"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/render"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers GET /admin/root/add-backoffice-user with the
// empty form that POST /admin/root/add-backoffice-user reads, the viewer role
// selected.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	fields := backofficeusers.Fields{Role: int64(backofficeauth.RoleViewer)}
	return render.AddBackofficeUserForm(sandbox, response, api.StatusOk, props.User, fields, "")
}
