# `GET /health`

Reports that the server is up

## Try it

```bash
curl localhost:3000/health
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/health/` · Server · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
