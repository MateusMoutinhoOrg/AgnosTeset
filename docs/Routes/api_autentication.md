# `ANY /api/admin/* !(/api/admin/login)`

Requires a valid Bearer token on /api/admin, except /api/admin/login

Every /api/admin route but POST /api/admin/login needs the token that login answers, sent as Authorization: Bearer <token>. The token is bound to the client ip it was issued to and lasts 30 minutes; a missing, invalid or expired one is answered 401 in JSON.

## Try it

Only what is required:

```bash
curl localhost:3000/api/admin
```

With every value it reads:

```bash
curl localhost:3000/api/admin \
  -H 'authorization: my-authorization' \
  -H 'x-client-ip: my-x-client-ip'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| the whole address | text that must not be exactly `/api/admin/login` | — | the login route is reachable without a token |

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `authorization` | header | text | no | `my-authorization` | the session JWT answered by POST /api/admin/login, as Bearer <token> |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip the request came from, set by the server and never by the client |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/api_autentication/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
