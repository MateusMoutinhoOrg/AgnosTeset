package redirect

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/globals"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	alias := route.GetString("alias")

	if globals.DB == nil {
		response.SetStatus(api.StatusFailure)
		response.Write([]byte("Database not initialized\n"))
		return api.StatusFailure
	}

	urlItem := globals.DB.FindUrlLinkByAlias(alias)
	if urlItem == nil {
		response.SetStatus(api.StatusNotFound)
		response.Write([]byte("Alias not found\n"))
		return api.StatusNotFound
	}

	response.SetHeader("Location", urlItem.Link)
	response.SetStatus(302)
	return 302
}
