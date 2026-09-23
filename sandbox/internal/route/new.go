package route

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func NewRoute(sandbox *api.Sandbox, props *api.RouteProps) *api.Route {
	self := api.Route{}
	self.Name = props.Name
	self.AcceptMethods = props.AcceptMethods
	self.Priority = props.Priority
	self.Category = props.Category
	self.Help = props.Help
	self.LongDescription = props.LongDescription
	self.Parameters = props.Parameters
	self.Paths = props.Paths
	self.Examples = props.Examples

	self.Handler = func(request *serverdeps.Request, response *serverdeps.Response) error {
		return Handler(sandbox, &self, request, response)
	}
	self.IsActionable = func(request *serverdeps.Request) bool {
		return IsActionable(sandbox, &self, request)
	}

	return &self
}
