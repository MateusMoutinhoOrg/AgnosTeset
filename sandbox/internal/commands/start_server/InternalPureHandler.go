package start_server

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/commandprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/cliio"
	server "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/server/server"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/httpguard"
)

// InternalPureHandler backs `start-server`: it serves until the process is
// asked to stop, and answers the command line once it has. The secret that
// signs the backoffice sessions comes from the environment, never from the
// command line; without one the server does not start.
func InternalPureHandler(sandbox *api.Sandbox, props *commandprops.CommandProps, entries *Entries, response *api.CommandResponse) error {
	secret, err := backofficeauth.ReadSecret(sandbox)
	if err != nil {
		return cliio.Fail(sandbox, api.ExitFailure, "", err.Error())
	}

	// Make the settings available to every route handler through sandbox.Config.
	sandbox.Config.Secret = secret
	sandbox.Config.AllowXForwardedFor = entries.AllowXForwardedFor
	sandbox.Config.InsecureHttp = entries.InsecureHttp

	if entries.AllowXForwardedFor && httpguard.ListensEverywhere(sandbox, entries.Addr) {
		sandbox.Deps.Std.Error("warning: X-Forwarded-For is trusted but the server listens on every interface: bind it to the address only the proxy reaches (--addr 127.0.0.1:3000) or firewall the port, or anyone reaching it can forge their ip\n")
	}

	err = server.ServerMain(sandbox, api.ServeProps{
		Addr:              entries.Addr,
		ReadTimeoutMs:     entries.ReadTimeoutMs,
		WriteTimeoutMs:    entries.WriteTimeoutMs,
		ShutdownTimeoutMs: entries.ShutdownTimeoutMs,
	})
	if err != nil {
		return cliio.Fail(sandbox, api.ExitFailure, "", "server stopped: "+err.Error())
	}
	response.SetStatus(api.ExitOk)
	return nil
}
