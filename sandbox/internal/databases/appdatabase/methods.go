package appdatabase

import (


	database "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/database"
)

func AddUrlLink(self *AppDatabase, alias string, link string) error {
	schema, ok := self.InnerDatabase.GetSchema("url")
	if !ok {
		return self.deps.Std.Errorf("schema not found")
	}
	_, err := schema.NewItem(map[string]any{
		"alias":     alias,
		"link":      link,
		"creation":  self.deps.Std.Now(),
		"redirects": 0,
	})
	if err != nil {
		return self.deps.Std.Errorf(err.Message)
	}
	return nil
}

func FindUrlLinkByAlias(self *AppDatabase, alias string) *UrlItem {
	schema, ok := self.InnerDatabase.GetSchema("url")
	if !ok {
		return nil
	}
	item, ok := schema.FindByKey("alias", alias)
	if !ok {
		return nil
	}
	return buildUrlItem(item)
}

func FindUrlLinkByLink(self *AppDatabase, link string) *UrlItem {
	schema, ok := self.InnerDatabase.GetSchema("url")
	if !ok {
		return nil
	}
	items, err := schema.ListAll()
	if err != nil {
		return nil
	}
	for _, item := range items {
		linkVal, _ := item.Get("link")
		if linkVal.(string) == link {
			return buildUrlItem(item)
		}
	}
	return nil
}

func ListUrls(self *AppDatabase, filtrage FiltrageProps) []UrlItem {
	schema, ok := self.InnerDatabase.GetSchema("url")
	if !ok {
		return nil
	}
	items, err := schema.ListAll()
	if err != nil {
		return nil
	}
	var res []UrlItem
	for _, item := range items {
		ui := buildUrlItem(item)
		if ui == nil {
			continue
		}
		if filtrage.LinkStartsWith != "" && !self.deps.Stringsdeps.HasPrefix(ui.Link, filtrage.LinkStartsWith) {
			continue
		}
		if filtrage.Creation > 0 && ui.Creation < filtrage.Creation {
			continue
		}
		if filtrage.Redirects > 0 && ui.Redirects < filtrage.Redirects {
			continue
		}
		res = append(res, *ui)
	}
	return res
}

func buildUrlItem(item database.SchemaItem) *UrlItem {
	aliasVal, _ := item.Get("alias")
	linkVal, _ := item.Get("link")
	creationVal, _ := item.Get("creation")
	redirectsVal, _ := item.Get("redirects")
	
	var alias, link string
	if aliasVal != nil {
		alias = aliasVal.(string)
	}
	if linkVal != nil {
		link = linkVal.(string)
	}
	var creation, redirects int64
	if creationVal != nil {
		creation = creationVal.(int64)
	}
	if redirectsVal != nil {
		redirects = redirectsVal.(int64)
	}

	return &UrlItem{
		Alias:     alias,
		Link:      link,
		Creation:  creation,
		Redirects: redirects,
	}
}
