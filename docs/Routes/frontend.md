# `GET /{*Rest}`

Serves any file of the embedded assets/frontend tree

| Entries | Read from | In | Type | Default | Description | From |
| --- | --- | --- | --- | --- | --- | --- |
| `Rest` | segments 0..-1 | path | string |  | the file under assets/frontend; none is its index.html | — |

| Runs in front of it | When |
| --- | --- |
| [`adminmiddlware`](adminmiddlware.md) | may run — `explain-command` gives the exact answer |
| [`health`](health.md) | may run — `explain-command` gives the exact answer |
| [`page1`](page1.md) | may run — `explain-command` gives the exact answer |

`sandbox/internal/routeslist/frontend/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
