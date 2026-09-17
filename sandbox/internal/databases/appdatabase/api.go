package appdatabase

import (
	database "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/database"
)

type AppDatabase struct {
	InnerDatabase database.DatabaseHandle
}
