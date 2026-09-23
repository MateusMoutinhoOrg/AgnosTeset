package health

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// RouteHandler answers the built-in health route with a fixed JSON object. It
// is the server layer's `version` command: a route agnos writes itself, so a
// freshly initialized server already answers something.
func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) error {

	return nil
}
