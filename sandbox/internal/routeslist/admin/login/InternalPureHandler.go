package login

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/render"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /admin/login. The username field takes a
// username or an email; a match sets a session cookie bound to the request's
// host and redirects to /admin/home, anything else answers the login page again under a 401.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	username := entries.Body.Username

	user, ok, err := backofficeauth.Authenticate(sandbox, username, entries.Body.Password)
	if err != nil {
		return err
	}
	if !ok {
		return render.Login(sandbox, response, api.StatusUnauthorized, "Invalid username or password.", username)
	}

	token, err := backofficeauth.IssueToken(sandbox, user, entries.Host)
	if err != nil {
		return err
	}

	response.AddHeader("Set-Cookie", backofficeauth.SessionCookie(sandbox, token))
	response.SetHeader("Location", "/admin/home")
	response.SetStatus(api.StatusSeeOther)
	return nil
}
