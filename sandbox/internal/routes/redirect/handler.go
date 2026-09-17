package redirect

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	response.SetHeader("Content-Type", "text/plain")
	response.SetStatus(api.StatusOk)
	response.Write([]byte("redirect called\n"))
	return api.StatusOk
}
