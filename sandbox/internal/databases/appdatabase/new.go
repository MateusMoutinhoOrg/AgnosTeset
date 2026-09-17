


func NewDataBase(path string) *AppDatabase {
	props = api.Props{
		Path: path,
		Schemas: []api.Schema{
			{
				Name: "urls",
				Itens: []api.Item{
					{Name: "alias", Type: api.Key, Required: true},
					{Name: "link", Type: api.String, Required: true},
					{Name: "creation", Type: api.Int, Required: true},
					{Name: "redirects", Type: api.Int, Required: true},
				},
			},
		},
	}

}