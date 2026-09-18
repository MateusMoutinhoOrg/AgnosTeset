package get_transactions

import (
	"encoding/json"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/finance"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	db := finance.New(sandbox)
	items, err := db.ListTransactions(finance.TransactionsFiltrage{})
	if err != nil {
		response.SetStatus(api.StatusFailure)
		return api.StatusFailure
	}
	bytes, err := json.Marshal(items)
	if err != nil {
		response.SetStatus(api.StatusFailure)
		return api.StatusFailure
	}
	response.SetHeader("Content-Type", "application/json")
	response.SetStatus(api.StatusOk)
	response.Write(bytes)
	return api.StatusOk
}
