package adminmiddleware

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// InternalPureHandler runs in front of every ANY /adimin/{*Rest} on a
// lower rung of the chain: it is a middleware. Returning nil without answering
// hands the request to the next route; answering — a status, or a byte —
// ends the chain here.
//
// Refuse a request with routeio.Fail, which answers it through the project's
// own handler for that status:
//
//	return routeio.Fail(sandbox, route, api.StatusUnauthorized, "authorization", "invalid token")
//
// Hand what you learned to the routes after it through route.Locals:
//
//	routeio.SetLocal(route, "user", user)
func InternalPureHandler(sandbox *api.Sandbox, route *api.Route, entries *Entries, response *serverdeps.Response) error {
	return nil
}
