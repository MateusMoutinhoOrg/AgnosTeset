package list_links

import (

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

	filtrage := appdatabase.FiltrageProps{
		LinkStartsWith: route.GetString("starts_with"),
		Creation:       int64(route.GetInt("min_creation")),
		Redirects:      int64(route.GetInt("min_redirects")),
	}
	links := globals.DB.ListUrls(filtrage)
	if links == nil {
		links = []appdatabase.UrlItem{} // ensure we return an array, not null
	}
	
	array := sandbox.Deps.Serializables.CreateArray()
	for _, link := range links {
		obj := sandbox.Deps.Serializables.CreateObject()
		obj.AddItemToObject("Alias", sandbox.Deps.Serializables.CreateString(link.Alias))
		obj.AddItemToObject("Link", sandbox.Deps.Serializables.CreateString(link.Link))
		obj.AddItemToObject("Creation", sandbox.Deps.Serializables.CreateInt(link.Creation))
		obj.AddItemToObject("Redirects", sandbox.Deps.Serializables.CreateInt(link.Redirects))
		array.AddItemToArray(obj)
	}
	dataStr := sandbox.Deps.Serializables.SerializeToJson(array)
	data := []byte(dataStr)

	response.SetHeader("Content-Type", "application/json")
	response.SetStatus(api.StatusOk)
	response.Write(data)
	return api.StatusOk
}
