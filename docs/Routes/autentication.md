# `ANY /admin/* !(/admin/login)`

Requires a valid backoffice session on /admin, except /admin/login

## Try it

Only what is required:

```bash
curl localhost:3000/admin
```

With every value it reads:

```bash
curl localhost:3000/admin \
  -H 'x-client-ip: my-x-client-ip' \
  -b 'admin_token=my-admin-token'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| the whole address | text that must not be exactly `/admin/login` | — | the login route is reachable without a session |

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `admin_token` | cookie | text | no | `my-admin-token` | the session JWT set by POST /admin/login |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip the request came from, set by the server and never by the client |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/autentication/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
