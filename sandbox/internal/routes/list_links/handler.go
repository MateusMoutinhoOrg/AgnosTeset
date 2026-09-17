package list_links

import (
	"encoding/json"

	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/config"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/appdatabase"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/globals"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	password := route.GetString("password")
	if password != config.RootPassword {
		response.SetStatus(api.StatusNotFound)
		response.Write([]byte("Unauthorized\n"))
		return api.StatusNotFound
	}

	if globals.DB == nil {
		response.SetStatus(api.StatusFailure)
		response.Write([]byte("DB not initialized\n"))
		return api.StatusFailure
	}

	links := globals.DB.ListUrls(appdatabase.FiltrageProps{})
	if links == nil {
		links = []appdatabase.UrlItem{} // ensure we return an array, not null
	}
	
	data, err := json.Marshal(links)
	if err != nil {
		response.SetStatus(api.StatusFailure)
		response.Write([]byte("Failed to encode response\n"))
		return api.StatusFailure
	}

	response.SetHeader("Content-Type", "application/json")
	response.SetStatus(api.StatusOk)
	response.Write(data)
	return api.StatusOk
}
