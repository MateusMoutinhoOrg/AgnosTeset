package adduser

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func PureHandler(sandbox *api.Sandbox, entries any, response *serverdeps.Response) error {
	response.SetStatus(200)
	response.WriteBody([]byte("Hello World"))
	return nil
}
