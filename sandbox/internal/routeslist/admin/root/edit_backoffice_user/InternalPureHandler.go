package edit_backoffice_user

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficerender"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeusers"
)

// InternalPureHandler answers POST /admin/root/edit-backoffice-user/{id}. The
// user's username, email and role are written, and their password when one was
// given, then the browser is sent to the list. A new password ends every
// session of the user and revokes their API tokens — the session of this
// request survives when a root edits their own account — and the list says
// so. A refused edit answers the form
// again, filled with what was sent but the password, under a 400 with the
// reason above it. A user that does not exist sends the browser back to the
// list.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	id := int64(entries.Id)
	_, ok := backofficeusers.Find(sandbox, id)
	if !ok {
		return routeio.Redirect(*response, api.StatusSeeOther, backofficeusers.ListLocation(sandbox, backofficeusers.NoticeNotFound))
	}

	fields := backofficeusers.Fields{
		Username: entries.Body.Username,
		Email:    entries.Body.Email,
		Password: entries.Body.Password,
		Role:     int64(entries.Body.Role),
	}
	message, notice, err := backofficeusers.Update(sandbox, *props.User, props.Session, id, fields)
	if err != nil {
		return err
	}
	if message != "" {
		return backofficerender.EditBackofficeUserForm(sandbox, response, api.StatusBadRequest, props.User, id, fields, message)
	}
	return routeio.Redirect(*response, api.StatusSeeOther, backofficeusers.ListLocation(sandbox, notice))
}
