package adduser

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
)

func InterfaceHandler(sandbox *api.Sandbox, entries *Entries, response *api.Response) error {
	response.SetStatus(200)
	response.WriteBody([]byte("Hello World"))
	return nil
}
