package same_origin

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeguard"
)

// InternalPureHandler runs in front of every ANY /admin/{*Rest}, login
// included, on a lower rung of the chain than the authentication middleware:
// it is a middleware.
//
// A request whose Origin header names another host than the one it was sent
// to was sent by another site's page, and is refused with a 403 before
// anything reads its session cookie: that is what a cross-site form post is.
// Anything else declines, so the route after it runs. It backs up the
// SameSite=Strict of the session cookie, which keeps another site's requests
// from carrying a session to begin with.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if backofficeguard.SameOrigin(sandbox, entries.Origin, entries.Host) {
		return nil
	}
	return sandbox.Deps.OpinatedAgnosServer.Fail(api.StatusForbidden, "origin", "a request sent by another site's page is refused")
}
