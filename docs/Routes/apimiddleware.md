# `ANY /api/admin/{*Rest}`

## Try it

```bash
curl localhost:3000/api/admin/my-rest
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{*Rest}` | the rest of the address — one part or more, like `a/b.png` | `my-rest` |  |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/plain`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/apimiddleware/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
