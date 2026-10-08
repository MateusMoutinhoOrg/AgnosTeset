package main

import (
	"os"

	agnosadapter "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/bindings/standard"

	agnoslib "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox"
)

func main() {

	deps := agnosadapter.New()

	lib := agnoslib.New(&deps)
	argslist := os.Args[1:]
	result := lib.Cli.Main(argslist)
	os.Exit(result)
}
