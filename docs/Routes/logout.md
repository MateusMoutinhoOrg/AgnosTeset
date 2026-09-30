# `POST /admin/logout`

Ends the session on this host

## Try it

Only what is required:

```bash
curl -X POST localhost:3000/admin/logout
```

With every value it reads:

```bash
curl -X POST localhost:3000/admin/logout \
  -H 'host: my-host'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `host` | header | text | no | `my-host` | the host the request was sent to |

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

For developers: `sandbox/internal/routeslist/logout/` · Routes · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
