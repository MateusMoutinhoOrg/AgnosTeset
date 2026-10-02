# `GET /admin/list-backoffice-api-tokens`

Lists your API tokens, or every user's for a root

## Try it

Only what is required:

```bash
curl localhost:3000/admin/list-backoffice-api-tokens
```

With every value it reads:

```bash
curl 'localhost:3000/admin/list-backoffice-api-tokens?notice=my-notice'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `notice` | query string | text | no | `my-notice` | the outcome code of the last revoke |

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

For developers: `sandbox/internal/routeslist/list_backoffice_api_tokens/` · Backoffice API Tokens · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
