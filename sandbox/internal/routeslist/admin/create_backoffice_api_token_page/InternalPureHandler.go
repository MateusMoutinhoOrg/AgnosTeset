package create_backoffice_api_token_page

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficetokens"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/routeio"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/render"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
)

// InternalPureHandler answers GET /admin/create-backoffice-api-token with the
// empty form that POST /admin/create-backoffice-api-token reads, the default
// expiration selected and the client ip of the browser offered for the ips
// field.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	fields := backofficetokens.Fields{Expiration: backofficetokens.DefaultExpiration}
	return render.CreateBackofficeApiTokenForm(sandbox, response, api.StatusOk, props.User, fields, props.ClientIp, "")
}
