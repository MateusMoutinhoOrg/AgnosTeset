package security_headers

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/httpguard"
)

// InternalPureHandler runs in front of every ANY /admin and /api/admin path,
// on a lower rung of the chain than the autentication middlewares: it is a
// middleware. It sets httpguard.SecurityHeaders and declines; setting a header
// answers nothing, so whichever route or Handle* file answers the request
// next — a page, a JSON document, a 401 — carries them.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	httpguard.SecurityHeaders(sandbox, response)
	return nil
}
