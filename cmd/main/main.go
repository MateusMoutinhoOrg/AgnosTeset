package main

import (
	"os"

	agnosadapter "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/availables/standard"

	agnoslib "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox"
)

func main() {

	deps := agnosadapter.New()

	lib := agnoslib.New(&deps)
	argslist := os.Args[1:]
	result := lib.Cli.CliMain(argslist)
	os.Exit(result)
}
