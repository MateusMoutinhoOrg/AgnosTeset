# Routes

Every route this server answers, one page each, generated from
`sandbox/internal/routeslist/<name>/route.yaml` ([RouteYaml](../RouteYaml/doc.md)) on each build —
open the one you need rather than this whole page. Hidden routes are not listed.

Routes run lowest `priority` first, the `after` phase once the request is answered; the segment
count, every path (its type and its trigger) and every parameter trigger of a route have to match
for it to run. A path nothing answers is `404`; one matched under another method is `405`.
`agnos list-routes` is the chain in run order, and
`agnos explain-route <METHOD> <path>` says which routes one request reaches.
Every parameter of a route is bound and converted before the handler runs — a failure there is
`400`, never the handler's call.

## Server

| Route | Answers | Package |
| --- | --- | --- |
| [`GET /health`](health.md) | Reports that the server is up | `health` |

## Routes

| Route | Answers | Package |
| --- | --- | --- |
| [`GET /{*Rest}`](home.md) |  | `home` |

Statuses and who answers each one are in [RouteYaml](../RouteYaml/doc.md#failures).
