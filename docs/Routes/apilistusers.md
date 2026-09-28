# `GET /api/admin/list-users`

## Try it

```bash
curl localhost:3000/api/admin/list-users
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`apimiddleware`](apimiddleware.md) | always |

---

For developers: `sandbox/internal/routeslist/apilistusers/` · api · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
