# `GET /{*Rest}`

Serves any file of the embedded assets/frontend tree

## Try it

```bash
curl localhost:3000/my-rest
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{*Rest}` | the rest of the address — one part or more, like `a/b.png` | `my-rest` | the file under assets/frontend; none is its index.html |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html; charset=utf-8`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`home`](home.md) | depends on the address — `explain-route` gives the exact answer |
| [`health`](health.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/frontend/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
