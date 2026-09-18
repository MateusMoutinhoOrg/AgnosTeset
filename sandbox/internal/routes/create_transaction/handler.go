package create_transaction

import (
	"encoding/json"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/finance"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/routeio"
)

func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	db := finance.New(sandbox)
	request := routeio.RequestOf(route)
	body, err := request.ReadBody(-1)
	if err != nil {
		response.SetStatus(api.StatusBadRequest)
		return api.StatusBadRequest
	}

	var props finance.TransactionsNew
	if err := json.Unmarshal(body, &props); err != nil {
		response.SetStatus(api.StatusBadRequest)
		return api.StatusBadRequest
	}

	item, err := db.AddTransactions(props)
	if err != nil {
		response.SetStatus(api.StatusFailure)
		return api.StatusFailure
	}
	bytes, _ := json.Marshal(item)
	response.SetHeader("Content-Type", "application/json")
	response.SetStatus(api.StatusOk)
	response.Write(bytes)
	return api.StatusOk
}
