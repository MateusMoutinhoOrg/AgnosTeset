package render

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
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
func roleName(sandbox *api.Sandbox, role backofficeauth.Role) string {
	switch role {
	case backofficeauth.RoleRoot:
		return "root"
	case backofficeauth.RoleViewer:
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
func Home(sandbox *api.Sandbox, response *serverdeps.Response, user *maindatabase.BackofficeuserItem, sessionMinutes int) error {
	return Html(sandbox, response, api.StatusOk, "templates/home.html", HomePage{
		Id:             sandbox.Deps.Stringsdeps.FormatInt(user.Id, 10),
		Username:       user.Username,
		Initial:        initialOf(sandbox, user.Username),
		Email:          user.Email,
		Role:           roleName(sandbox, backofficeauth.Role(user.Role)),
		SessionMinutes: sessionMinutes,
	})
}
