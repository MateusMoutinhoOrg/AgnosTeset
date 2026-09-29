package login

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// InternalPureHandler answers POST /admin/login.
func InternalPureHandler(sandbox *api.Sandbox, props *api.RouteProps, entries *Entries, response *serverdeps.Response) error {
	response.SetStatus(api.StatusOk)
	
	username := entries.Body.Username
	password := entries.Body.Password

	res := "Login form data received. Authentication is not implemented.\nUsername: " + username + "\nPassword: " + password

	response.SetHeader("Content-Type", "text/plain")
	response.Write([]byte(res))

	return nil
}
