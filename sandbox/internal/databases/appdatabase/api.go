package appdatabase

import (
	database "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/database"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

type UrlItem struct {
	Alias     string
	Link      string
	Creation  int64
	Redirects int64
}

type FiltrageProps struct {
	LinkStartsWith string
	Creation       int64
	Redirects      int64
}

type AppDatabase struct {
	deps               *deps.Deps
	InnerDatabase      database.DatabaseHandle
	AddUrlLink         func(alias string, link string) error
	FindUrlLinkByAlias func(alias string) *UrlItem
	FindUrlLinkByLink  func(link string) *UrlItem
	ListUrls           func(filtrage FiltrageProps) []UrlItem
}
