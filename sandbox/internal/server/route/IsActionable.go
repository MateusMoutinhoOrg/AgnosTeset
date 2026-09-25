package route

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func IsActionable(sandbox *api.Sandbox, route *api.Route, request *serverdeps.Request) bool {
	return true
}
