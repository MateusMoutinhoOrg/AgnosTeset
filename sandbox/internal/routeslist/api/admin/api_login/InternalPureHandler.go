package api_login

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeapi"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /api/admin/login. The username field takes
// a username or an email; a match opens a session bound to the client ip the
// request came from and answers its token, to send as
// `Authorization: Bearer <token>`, anything else is refused with a 401.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	user, ok, err := backofficeauth.Authenticate(sandbox, entries.Body.Username, entries.Body.Password)
	if err != nil {
		return err
	}
	if !ok {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "invalid username or password")
	}

	token, err := backofficeauth.IssueToken(sandbox, user, entries.XClientIp)
	if err != nil {
		return err
	}
	return routeio.WriteJSON(sandbox, *response, api.StatusOk, backofficeapi.Session(sandbox, token, user))
}
