package add_backoffice_user

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficerender"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeusers"
)

// InternalPureHandler answers POST /admin/root/add-backoffice-user. A user the
// form describes well is added and the browser is sent to the list; anything
// else answers the form again, filled with what was sent but the password,
// under a 400 with the reason above it.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	fields := backofficeusers.Fields{
		Username: entries.Body.Username,
		Email:    entries.Body.Email,
		Password: entries.Body.Password,
		Role:     int64(entries.Body.Role),
	}
	_, message, err := backofficeusers.Add(sandbox, fields)
	if err != nil {
		return err
	}
	if message != "" {
		return backofficerender.AddBackofficeUserForm(sandbox, response, api.StatusBadRequest, props.User, fields, message)
	}
	return routeio.Redirect(*response, api.StatusSeeOther, backofficeusers.ListLocation(sandbox, backofficeusers.NoticeAdded))
}
