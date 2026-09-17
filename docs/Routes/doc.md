# Routes

Every route this server answers, generated from
`sandbox/internal/routes/<name>/route.yaml` ([RouteYaml](../RouteYaml/doc.md)) on each build.
Hidden routes are not listed.

A path is matched segment by segment, most specific route first. A path nothing matches is
`404`; one matched under another method is `405`. Every field below is bound, converted and
range-checked before the handler runs — a failure there is `400`, never the handler's call.

## Server

### `GET /health`

Reports that the server is up

```bash
curl localhost:8080/health
```

### `GET /redirect`

makes the redirection

### `POST /shortner`

shorts the url

## Assets

### `GET /static/{item...}`

Serves one file from the embedded static assets

| Field | In | Type | Default | Description |
| --- | --- | --- | --- | --- |
| `item` | path | string, the rest of the path, required |  | the asset path under assets/frontend/static, one or more segments |
| `sha` | query | string |  | the digest staticref stamped on the url; when it matches the asset the answer is cacheable forever |

Statuses and who answers each one are in [RouteYaml](../RouteYaml/doc.md#dispatch).
