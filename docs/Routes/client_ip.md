# `ANY /*`

Works out the client ip every route after it reads, from the connection or the reverse proxy in front

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

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/plain`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/client_ip/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
