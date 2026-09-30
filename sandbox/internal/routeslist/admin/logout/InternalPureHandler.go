package logout

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /admin/logout for the session the
// autentication middleware put on props.Session: the session is closed, so its
// token is refused from here on, the session cookie is cleared, and the
// browser is sent back to /admin/home, which answers the login page. Every
// other session of the user, on this device or another, stays open.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil || props.Session == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	err := backofficeauth.Logout(sandbox, *props.User, *props.Session)
	if err != nil {
		return err
	}

	response.AddHeader("Set-Cookie", backofficeauth.ClearedCookie(sandbox))
	response.SetHeader("Location", "/admin/home")
	response.SetStatus(api.StatusSeeOther)
	return nil
}
