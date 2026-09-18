# `GET /static/{item...}`

Serves one file from the embedded static assets

| Field | In | Type | Default | Description |
| --- | --- | --- | --- | --- |
| `item` | path | string, the rest of the path, required |  | the asset path under assets/frontend/static, one or more segments |
| `sha` | query | string |  | the digest staticref stamped on the url; when it matches the asset the answer is cacheable forever |

`sandbox/internal/routes/static/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
