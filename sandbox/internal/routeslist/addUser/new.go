package adduser

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/route"
)

func NewRoute(sandbox *api.Sandbox) *api.Route {
	self := route.NewRoute(sandbox)

	self.Name = "addUser"
	self.AcceptMethods = []string{"POST"}
	self.Priority = 1
	self.Category = "Server"
	self.Help = "Add a user to the system"
	self.LongDescription = "Add a user to the system"
	self.Examples = []string{"curl -X POST -H \"Content-Type: application/json\" -d '{\"name\":\"John Doe\",\"email\":\"[EMAIL_ADDRESS]\",\"password\":\"password\"}' http://localhost:8080/addUser"}
	self.InternalPurehandler = func(entries *Entries, response *serverdeps.Response) error {
		return InternalPureHandler(sandbox, self, entries, response)
	}

	return self
}
