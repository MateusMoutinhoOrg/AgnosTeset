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

	res := "<html><head><title>Login Success</title></head><body><h1>Login form data received.</h1><p>Authentication is not implemented.</p><p>Username: " + username + "</p><p>Password: " + password + "</p></body></html>"

	response.SetHeader("Content-Type", "text/html")
	response.Write([]byte(res))

	return nil
}
