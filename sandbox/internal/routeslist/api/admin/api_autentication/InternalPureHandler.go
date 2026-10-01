package api_autentication

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler runs in front of every ANY /api/admin/{*Rest} except
// /api/admin/login, on a lower rung of the chain: it is a middleware.
//
// An `Authorization: Bearer <token>` header holding a valid session token,
// issued to the client ip this request came from for a session no logout has
// closed, puts its user on props.User and its session on props.Session and
// declines, so the route after it runs. Anything else is refused with a 401
// in JSON. The session cookie is never read here, so a browser carrying one
// cannot be made to call these routes by another site.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	token := backofficeauth.BearerToken(sandbox, entries.Authorization)
	user, session, ok := backofficeauth.SessionOfToken(sandbox, token, entries.XClientIp)
	if ok {
		props.User = &user
		props.Session = &session
		return nil
	}

	response.SetHeader("WWW-Authenticate", "Bearer")
	if token == "" {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "authorization", "send the token POST /api/admin/login answers in the Authorization header, after Bearer")
	}
	return routeio.Fail(sandbox, api.StatusUnauthorized, "authorization", "the token is invalid or expired, sign in again")
}
