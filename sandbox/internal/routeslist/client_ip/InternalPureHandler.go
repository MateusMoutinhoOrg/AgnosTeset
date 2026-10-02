package client_ip

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server/httpguard"
)

// InternalPureHandler runs in front of every ANY /*, on the lowest rung of the
// chain: it is a middleware. It puts the ip the request came from on
// props.ClientIp — the connection's own, or, when start-server trusts
// X-Forwarded-For, the one the reverse proxy in front appended — and declines,
// so every route after it reads one ip, worked out once.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	props.ClientIp = httpguard.ClientIp(sandbox, entries.XClientIp, entries.XForwardedFor)
	return nil
}
