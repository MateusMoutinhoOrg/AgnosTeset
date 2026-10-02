package root_guard

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/render"
)

// InternalPureHandler runs in front of every ANY /admin/root/{*Rest}, after the
// autentication middleware put the signed-in user on props.User: it is a
// middleware. A root declines, so the route after it runs; anyone else is
// answered the forbidden page under a 403.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	if backofficeauth.Role(props.User.Role) != backofficeauth.RoleRoot {
		return render.Forbidden(sandbox, response, props.User)
	}
	return nil
}
