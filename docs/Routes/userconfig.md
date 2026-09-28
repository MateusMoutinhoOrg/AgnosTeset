# `GET /admin/user/{User:integer}`

## Try it

```bash
curl localhost:3000/admin/user/1
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{User:integer}` | whole number | `1` |  |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |
| `404` | `{User:integer}` is not a whole number, so this route does not answer the address. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`adminmiddlware`](adminmiddlware.md) | always |

---

For developers: `sandbox/internal/routeslist/userconfig/` · frontend · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
