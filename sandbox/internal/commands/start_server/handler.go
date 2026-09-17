package start_server

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/config"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/globals"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/appdatabase"
	server "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/server"
)

func CommandHandler(sandbox *api.Sandbox, command *api.Command) int {
	config.RootPassword = command.GetString("root_password")
	globals.DB = appdatabase.NewDataBase("app_db_data", sandbox.Deps.Database.Databases.New)
	err := server.ServerMain(sandbox, api.ServeProps{
		Addr:           command.GetString("addr"),
		ReadTimeoutMs:  command.GetInt("read_timeout_ms"),
		WriteTimeoutMs: command.GetInt("write_timeout_ms"),
	})
	if err != nil {
		sandbox.Deps.Std.Error("server stopped: %s \n", err.Error())
		return api.ExitFailure
	}
	return api.ExitOk
}
