package route

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func NewRoute(sandbox *api.Sandbox) *api.Route {
	self := api.Route{}

	self.IsActionable = func(request *serverdeps.Request) bool {
		return IsActionable(sandbox, &self, request)
	}
	self.RequestHandler = func(request *serverdeps.Request, response *serverdeps.Response) error {
		return Requesthandler(sandbox, &self, request, response)
	}

	return &self
}
