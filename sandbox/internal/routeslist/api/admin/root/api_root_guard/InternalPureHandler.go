package api_root_guard

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeauth"
)

// InternalPureHandler runs in front of every ANY /api/admin/root/{*Rest},
// after the api-autentication middleware put the token's user on props.User:
// it is a middleware. A root declines, so the route after it runs; anyone
// else is refused with a 403 in JSON.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	if backofficeauth.Role(props.User.Role) != backofficeauth.RoleRoot {
		return routeio.Fail(sandbox, api.StatusForbidden, "", "only root users may do this")
	}
	return nil
}
