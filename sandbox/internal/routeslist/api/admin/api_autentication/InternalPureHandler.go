package api_autentication

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficetokens"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler runs in front of every ANY /api/admin/{*Rest}, on a
// lower rung of the chain: it is a middleware.
//
// An `Authorization: Bearer <token>` header holding an API token — created on
// the backoffice page, not revoked, not expired, and allowed from the client
// ip this request came from — puts the user it belongs to on props.User and
// the token on props.ApiToken and declines, so the route after it runs.
// Anything else is refused with a 401 in JSON. The session cookie is never
// read here, so a browser carrying one cannot be made to call these routes by
// another site, and the api issues no token of its own.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	token := backofficeauth.BearerToken(sandbox, entries.Authorization)
	user, apiToken, ok, err := backofficetokens.Resolve(sandbox, token, entries.XClientIp)
	if err != nil {
		return err
	}
	if ok {
		props.User = &user
		props.ApiToken = &apiToken
		return nil
	}

	response.SetHeader("WWW-Authenticate", "Bearer")
	if token == "" {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "authorization", "send an API token, created on "+backofficetokens.ListPath+", in the Authorization header, after Bearer")
	}
	return routeio.Fail(sandbox, api.StatusUnauthorized, "authorization", "the token is invalid, expired, revoked or not allowed from this ip")
}
