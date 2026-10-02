package list_backoffice_api_tokens

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/backoffice/backofficetokens"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/render"
)

// InternalPureHandler answers GET /admin/list-backoffice-api-tokens, open to
// every backoffice user, with templates/backoffice_api_tokens.html: the API
// tokens of the signed-in user — of every user, for a root — newest first,
// each with the control that revokes it.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	listed, err := backofficetokens.List(sandbox, *props.User)
	if err != nil {
		return err
	}
	return render.BackofficeApiTokens(sandbox, response, api.StatusOk, props.User, listed, entries.Notice, render.CreatedToken{})
}
