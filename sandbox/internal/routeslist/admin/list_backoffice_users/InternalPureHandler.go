package list_backoffice_users

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficeusers"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/render"
)

// InternalPureHandler answers GET /admin/list-backoffice-users, open to every
// backoffice user, with templates/backoffice_users.html: one page of the users
// the search and role filters keep, and, for a root, the controls that add,
// edit and remove them.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	listing, err := backofficeusers.List(sandbox, backofficeusers.Query{
		Search: entries.Search,
		Role:   entries.Role,
		Page:   entries.Page,
		Limit:  entries.Limit,
	})
	if err != nil {
		return err
	}
	return render.BackofficeUsers(sandbox, response, props.User, listing, entries.Notice)
}
