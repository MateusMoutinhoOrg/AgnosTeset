# `ANY /~(^/(api/)?admin(/|$))`

Sends the security headers on every /admin and /api/admin response

## Try it

Only what is required:

```bash
curl localhost:3000/
```

With every value it reads:

```bash
curl localhost:3000/ \
  -H 'x-client-ip: my-x-client-ip' \
  -H 'x-forwarded-for: my-x-forwarded-for'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| the whole address | text that must match the pattern `^/(api/)?admin(/\|$)` | — |  |

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client — read by [`client-ip`](client_ip.md), which runs first |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for — read by [`client-ip`](client_ip.md), which runs first |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/plain`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`client-ip`](client_ip.md) | always |

---

For developers: `sandbox/internal/routeslist/security_headers/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
