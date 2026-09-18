
package appdatabase

import (
	database "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/database"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

func NewDataBase(deps *deps.Deps, path string, buildDatabase func(props database.Props) database.DatabaseHandle) *AppDatabase {
	props := database.Props{
		Path: path,
		Schemas: []database.Schema{
			{
				Name: "url",
				Itens: []database.Item{
					{Name: "alias", Type: database.Key, Required: true},
					{Name: "link", Type: database.String, Required: true},
					{Name: "creation", Type: database.Int, Required: true},
					{Name: "redirects", Type: database.Int, Required: true},
				},
			},
		},
	}
	self := AppDatabase{deps: deps}
	self.InnerDatabase = buildDatabase(props)
	self.AddUrlLink = func(alias string, link string) error {
		return AddUrlLink(&self, alias, link)
	}
	self.FindUrlLinkByAlias = func(alias string) *UrlItem {
		return FindUrlLinkByAlias(&self, alias)
	}
	self.FindUrlLinkByLink = func(link string) *UrlItem {
		return FindUrlLinkByLink(&self, link)
	}
	self.ListUrls = func(filtrage FiltrageProps) []UrlItem {
		return ListUrls(&self, filtrage)
	}

	return &self
}