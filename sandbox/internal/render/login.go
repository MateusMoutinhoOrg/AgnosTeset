package render

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// LoginPage is what templates/login.html is rendered with.
type LoginPage struct {
	// Error is shown above the form, "" for none.
	Error string
	// Username refills the login field after a failed attempt.
	Username string
}

// Login answers the login page under status, with message above the form.
func Login(sandbox *api.Sandbox, response *serverdeps.Response, status int, message string, username string) error {
	return Html(sandbox, response, status, "templates/login.html", LoginPage{Error: message, Username: username})
}
