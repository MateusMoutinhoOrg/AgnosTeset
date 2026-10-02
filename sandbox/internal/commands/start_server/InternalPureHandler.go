package start_server

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/commandprops"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/cliio"
	server "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/server/server"
)

// InternalPureHandler backs `start-server`: it serves until the process is
// asked to stop, and answers the command line once it has.
func InternalPureHandler(sandbox *api.Sandbox, props *commandprops.CommandProps, entries *Entries, response *api.CommandResponse) error {
	err := server.ServerMain(sandbox, api.ServeProps{
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
