# `POST /admin/revoke-backoffice-api-token/{Id:integer}`

Revokes an API token: your own, or anyone's for a root

## Try it

```bash
curl -X POST localhost:3000/admin/revoke-backoffice-api-token/1
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{Id:integer}` | whole number | `1` |  |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |
| `404` | `{Id:integer}` is not a whole number, so this route does not answer the address. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/revoke_backoffice_api_token/` · Backoffice API Tokens · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
