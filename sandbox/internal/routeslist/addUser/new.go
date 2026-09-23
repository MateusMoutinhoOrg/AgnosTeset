package adduser

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/route"
)

func NewRoute(sandbox *api.Sandbox) *api.Route {
	props := api.RouteProps{}

	props.Name = "addUser"
	props.AcceptMethods = []string{"POST"}
	props.Priority = 1
	props.Category = "Server"
	props.Help = "Add a user to the system"
	props.LongDescription = "Add a user to the system"
	props.Examples = []string{"curl -X POST -H \"Content-Type: application/json\" -d '{\"name\":\"John Doe\",\"email\":\"[EMAIL_ADDRESS]\",\"password\":\"password\"}' http://localhost:8080/addUser"}
	props.PureHandler = func(entries *Entries, response *serverdeps.Response) error {
		return PureHandler(sandbox, routes, entries, response)
	}

	return route.NewRoute(sandbox, &props)
}
