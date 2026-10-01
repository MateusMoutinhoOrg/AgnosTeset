package api_logout

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeapi"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /api/admin/logout for the session the
// api-autentication middleware put on props.Session: the session is closed,
// so its token is refused from here on. Every other session of the user, on
// this client or another, stays open.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil || props.Session == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	err := backofficeauth.Logout(sandbox, *props.User, *props.Session)
	if err != nil {
		return err
	}
	return routeio.WriteJSON(sandbox, *response, api.StatusOk, backofficeapi.Ok(sandbox))
}
