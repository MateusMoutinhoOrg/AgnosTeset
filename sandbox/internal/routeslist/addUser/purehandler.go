package adduser

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

func PureHandler(sandbox *api.Sandbox, route *api.Route, entries *Entries, response *serverdeps.Response) error {
	response.SetStatus(200)
	response.Write([]byte("Hello World"))
	return nil
}
