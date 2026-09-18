package shortner

import (

	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/globals"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	body, status := ReadBody(sandbox, route, response)
	if status != api.StatusOk {
		return status
	}

	if globals.DB == nil {
		response.SetStatus(api.StatusFailure)
		response.Write([]byte("Database not initialized\n"))
		return api.StatusFailure
	}

	alias := sandbox.Deps.Stringsdeps.FormatInt(sandbox.Deps.Std.Now(), 16)

	err := globals.DB.AddUrlLink(alias, body.Url)
	if err != nil {
		response.SetStatus(api.StatusFailure)
		response.Write([]byte("Failed to create short link\n"))
		return api.StatusFailure
	}

	response.SetHeader("Content-Type", "text/plain")
	response.SetStatus(api.StatusOk)
	response.Write([]byte(alias))
	return api.StatusOk
}
