package routes

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func NewRoute() *api.Route {
	self := api.Route{}
	self.Handler = func(entries *serverdeps.Request, response *serverdeps.Response) error {
		//based on PureHandler and paths and paramthers call the Pure Handler
	}
	self.IsActionable = func(entries *serverdeps.Request) bool {
		//based on paths and paramthers check if the route is actionable
	}

	return &self
}
