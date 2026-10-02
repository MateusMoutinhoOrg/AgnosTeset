# `ANY /api/admin/*`

Requires a valid API token on every /api/admin route

Every /api/admin route needs an API token, sent as Authorization: Bearer <token>. Tokens are created and revoked on the backoffice page /admin/list-backoffice-api-tokens; the api never issues one. A token may expire on a date or never, and may be limited to a list of client ips. A missing, unknown, revoked or expired token, or one used from an ip it does not allow, is answered 401 in JSON.

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

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `authorization` | header | text | no | `my-authorization` | an API token created on /admin/list-backoffice-api-tokens, as Bearer <token> |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip the request came from, set by the server and never by the client |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

---

For developers: `sandbox/internal/routeslist/api_autentication/` · Middleware · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
