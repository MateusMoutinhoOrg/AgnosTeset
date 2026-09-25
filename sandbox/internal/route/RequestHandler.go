package route

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func Requesthandler(sandbox *api.Sandbox, self *api.Route, request *serverdeps.Request, response *serverdeps.Response) error {

	// these function must use introspection over self.InternalPureHandler to retrive args, and covert the entries .
	return nil
}
