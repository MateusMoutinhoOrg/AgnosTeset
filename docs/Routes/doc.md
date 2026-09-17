# Routes

Every route this server answers, one page each, generated from
`sandbox/internal/routes/<name>/route.yaml` ([RouteYaml](../RouteYaml/doc.md)) on each build —
open the one you need rather than this whole page. Hidden routes are not listed.

A path is matched segment by segment, most specific route first. A path nothing matches is
`404`; one matched under another method is `405`. Every field of a route is bound, converted and
range-checked before the handler runs — a failure there is `400`, never the handler's call.

## Server

| Route | Answers | Package |
| --- | --- | --- |
| [`GET /health`](health.md) | Reports that the server is up | `health` |
| [`GET /redirect`](redirect.md) | redirect the url | `redirect` |
| [`POST /shortner`](shortner.md) | shortner a url | `shortner` |

## Pages

| Route | Answers | Package |
| --- | --- | --- |
| [`GET /`](homepage.md) | THe homepage of the url shortner | `homepage` |

## Assets

| Route | Answers | Package |
| --- | --- | --- |
| [`GET /static/{item...}`](static.md) | Serves one file from the embedded static assets | `static` |

Statuses and who answers each one are in [RouteYaml](../RouteYaml/doc.md#dispatch).
