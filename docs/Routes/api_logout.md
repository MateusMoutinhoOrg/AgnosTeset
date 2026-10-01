# `POST /api/admin/logout`

Ends the session of the Bearer token

## Try it

```bash
curl -X POST localhost:3000/api/admin/logout
```

More examples:

```bash
curl -X POST localhost:3000/api/admin/logout -H "Authorization: Bearer $TOKEN"
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
| [`api-autentication`](api_autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/api_logout/` · Backoffice API · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
