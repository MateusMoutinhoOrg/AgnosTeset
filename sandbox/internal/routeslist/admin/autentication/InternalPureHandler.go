package autentication

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/render"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler runs in front of every ANY /admin/{*Rest} except
// /admin/login, on a lower rung of the chain: it is a middleware.
//
// A valid session cookie, issued on the host this request was sent to and
// after that host's last logout, puts its user on props.User and declines, so
// the route after it runs. Anything else answers the login page under a 401, and
// clears a cookie that no longer holds a valid session.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	user, ok := backofficeauth.UserOfToken(sandbox, entries.AdminToken, entries.Host)
	if ok {
		props.User = &user
		return nil
	}

	message := ""
	if entries.AdminToken != "" {
		response.AddHeader("Set-Cookie", backofficeauth.ClearedCookie(sandbox))
		message = "Your session has expired. Please sign in again."
	}
	return render.Login(sandbox, response, api.StatusUnauthorized, message, "")
}
