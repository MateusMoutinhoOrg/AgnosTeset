# `GET /admin/create-backoffice-api-token`

Shows the form that creates an API token

## Try it

Only what is required:

```bash
curl localhost:3000/admin/create-backoffice-api-token
```

With every value it reads:

```bash
curl localhost:3000/admin/create-backoffice-api-token \
  -H 'x-client-ip: my-x-client-ip'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip the request came from, set by the server and never by the client |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/create_backoffice_api_token_page/` · Backoffice API Tokens · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
