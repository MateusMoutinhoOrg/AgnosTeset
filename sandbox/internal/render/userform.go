package render

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeusers"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
)

// UserFormPage is what templates/user_form.html is rendered with: the form
// that adds a user, or the one that edits one.
type UserFormPage struct {
	Viewer Viewer
	// Title heads the page.
	Title string
	// Action is where the form posts.
	Action string
	// Submit labels the submit button.
	Submit string
	// Error is shown above the form, "" for none.
	Error    string
	Username string
	Email    string
	Roles    []RoleOption
	// PasswordRequired is true when adding: an edit keeps the current
	// password when the field is left blank.
	PasswordRequired bool
}

// RoleOption is one role the form offers.
type RoleOption struct {
	Value    int64
	Name     string
	Selected bool
}

// roleOptions are every role, the one of fields selected.
func roleOptions(sandbox *api.Sandbox, fields backofficeusers.Fields) []RoleOption {
	options := []RoleOption{}
	for _, role := range backofficeauth.Roles(sandbox) {
		options = append(options, RoleOption{
			Value:    int64(role),
			Name:     backofficeauth.RoleName(sandbox, role),
			Selected: int64(role) == fields.Role,
		})
	}
	return options
}

// AddUserForm answers, under status, the form that adds a user, filled with
// fields — never their password — and with message above it.
func AddUserForm(sandbox *api.Sandbox, response *serverdeps.Response, status int, user *maindatabase.BackofficeuserItem, fields backofficeusers.Fields, message string) error {
	return Html(sandbox, response, status, "templates/user_form.html", UserFormPage{
		Viewer:           viewerOf(sandbox, user),
		Title:            "Add user",
		Action:           "/admin/root/add-user",
		Submit:           "Add user",
		Error:            message,
		Username:         fields.Username,
		Email:            fields.Email,
		Roles:            roleOptions(sandbox, fields),
		PasswordRequired: true,
	})
}

// EditUserForm answers, under status, the form that edits the user with id
// id, filled with fields — never their password — and with message above it.
func EditUserForm(sandbox *api.Sandbox, response *serverdeps.Response, status int, user *maindatabase.BackofficeuserItem, id int64, fields backofficeusers.Fields, message string) error {
	return Html(sandbox, response, status, "templates/user_form.html", UserFormPage{
		Viewer:   viewerOf(sandbox, user),
		Title:    "Edit user #" + sandbox.Deps.Stringsdeps.FormatInt(id, 10),
		Action:   "/admin/root/edit-user/" + sandbox.Deps.Stringsdeps.FormatInt(id, 10),
		Submit:   "Save changes",
		Error:    message,
		Username: fields.Username,
		Email:    fields.Email,
		Roles:    roleOptions(sandbox, fields),
	})
}
