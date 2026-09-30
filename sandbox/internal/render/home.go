package render

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// HomePage is what templates/home.html is rendered with.
type HomePage struct {
	Id       string
	Username string
	// Initial is the first letter of Username, shown in the avatar.
	Initial        string
	Email          string
	Role           string
	SessionMinutes int
}

// roleName is the display name of a role.
func roleName(sandbox *api.Sandbox, role api.UserRole) string {
	switch role {
	case api.UserRoleRoot:
		return "root"
	case api.UserRoleViewer:
		return "viewer"
	}
	return "unknown"
}

// initialOf is the first letter of name, "?" for an empty one.
func initialOf(sandbox *api.Sandbox, name string) string {
	for _, letter := range name {
		return string(letter)
	}
	return "?"
}

// Home answers the home page for user, whose session lasts sessionMinutes.
func Home(sandbox *api.Sandbox, response *serverdeps.Response, user api.User, sessionMinutes int) error {
	return Html(sandbox, response, api.StatusOk, "templates/home.html", HomePage{
		Id:             user.Id,
		Username:       user.Username,
		Initial:        initialOf(sandbox, user.Username),
		Email:          user.Email,
		Role:           roleName(sandbox, user.Role),
		SessionMinutes: sessionMinutes,
	})
}
