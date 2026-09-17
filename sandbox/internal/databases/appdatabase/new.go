


func NewDataBase(path string) *AppDatabase {
	props = api.Props{
		Path: path,
		Schemas: []api.Schema{
			{
				Name: "url",
				Itens: []api.Item{
					{Name: "alias", Type: api.Key, Required: true},
					{Name: "link", Type: api.String, Required: true},
					{Name: "creation", Type: api.Int, Required: true},
					{Name: "redirects", Type: api.Int, Required: true},
				},
			},
		},
	}
	self := AppDatabase{}
	self.InnerDatabase = api.NewDatabase(props)
	self.FindUrlLinkByAlias = func(alias string) *UrlItem {
		return FindUrlLinkByAlias(&self, alias)
	}

	return &self
}