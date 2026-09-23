package server

import (
	api "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
)

func NewServer(sandbox *api.Sandbox) api.Server {
	server := api.Server{}

	server.Routes = []api.Route{}

	server.Serve = func(props api.ServeProps) error {
		return ServerMain(sandbox, props)
	}

	return server
}
