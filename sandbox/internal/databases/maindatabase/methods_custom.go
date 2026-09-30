package maindatabase

import (
	api "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/databaseio"
)

// UpdateBackofficeuserHostsMincreation writes a new mincreation on one hosts
// record of one backofficeuser record. The declaration generates Add and List
// for a nested collection but no Update, so it is written here.
func UpdateBackofficeuserHostsMincreation(sandbox *api.Sandbox, self *Maindatabase, parent_id int64, id int64, value int64) error {
	schema, err := databaseio.Schema(sandbox, self.handle, "backofficeuser")
	if err != nil {
		return err
	}
	parent, ok := schema.FindById(parent_id)
	if !ok {
		return sandbox.Deps.Std.Errorf("backofficeuser %d not found", parent_id)
	}
	for _, item := range parent.ListAll("hosts") {
		if item.Id == id {
			return databaseio.Fail(sandbox, item.Update("mincreation", value))
		}
	}
	return sandbox.Deps.Std.Errorf("hosts %d of backofficeuser %d not found", id, parent_id)
}
