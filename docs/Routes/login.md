# `ANY /admin/login`

## Try it

```bash
curl localhost:3000/admin/login
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/plain`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/login/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
