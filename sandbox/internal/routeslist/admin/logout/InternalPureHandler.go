package logout

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /admin/logout for the user the
// autentication middleware put on props.User: the mincreation of the
// request's host becomes now, invalidating every token issued on it so far,
// the session cookie is cleared, and the browser is sent back to /admin/home,
// which answers the login page.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	err := backofficeauth.Logout(sandbox, *props.User, entries.Host)
	if err != nil {
		return err
	}

	response.AddHeader("Set-Cookie", backofficeauth.ClearedCookie(sandbox))
	response.SetHeader("Location", "/admin/home")
	response.SetStatus(api.StatusSeeOther)
	return nil
}
