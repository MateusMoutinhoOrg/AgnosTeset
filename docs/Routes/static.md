# `GET /static/{*Item}`

Serves one file from the embedded static assets

| Entries | Read from | In | Type | Default | Description |
| --- | --- | --- | --- | --- | --- |
| `Mount` | segments 0..0 | path | string, equal `/static` |  | the directory every static asset is served under |
| `Item` | segments 1..-1 | path | string |  | the asset path under assets/frontend/static, one or more segments |
| `Sha` | `sha` | query | string |  | the digest staticref stamped on the url; when it matches the asset the answer is cacheable forever |

`sandbox/internal/routeslist/static/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
