# `GET /admin/root/add-user`

Shows the form that adds a backoffice user

## Try it

```bash
curl localhost:3000/admin/root/add-user
```

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
| [`root-guard`](root_guard.md) | always |

---

For developers: `sandbox/internal/routeslist/add_user_page/` · Admin · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
