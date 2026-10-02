# `POST /api/admin/get-backoffice-user`

Answers one backoffice user by id

## Try it

Only what is required:

```bash
curl -X POST localhost:3000/api/admin/get-backoffice-user \
  -H 'Content-Type: application/json' \
  -d '{"id":1}'
```

With every value it reads:

```bash
curl -X POST localhost:3000/api/admin/get-backoffice-user \
  -H 'authorization: my-authorization' \
  -H 'x-client-ip: my-x-client-ip' \
  -H 'x-forwarded-for: my-x-forwarded-for' \
  -H 'Content-Type: application/json' \
  -d '{"id":1}'
```

More examples:

```bash
curl -X POST localhost:3000/api/admin/get-backoffice-user -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"id":2}'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `authorization` | header | text | no | `my-authorization` | an API token created on /admin/list-backoffice-api-tokens, as Bearer <token> — read by [`api-authentication`](api_authentication.md), which runs first |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client — read by [`client-ip`](client_ip.md), which runs first |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for — read by [`client-ip`](client_ip.md), which runs first |

## Body

Send JSON with the header `Content-Type: application/json`, up to 1 MB. The body is required.

| Field | What goes there | Required | Rules |
| --- | --- | --- | --- |
| `id` | whole number | yes |  |

Example:

```json
{
  "id": 1
}
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |
| `400` | Something you sent is missing or has the wrong type or format. The answer's `field` names it. |
| `413` | The body is larger than 1 MB. |
| `415` | The body was not sent with `Content-Type: application/json`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`api-authentication`](api_authentication.md) | always |
| [`client-ip`](client_ip.md) | always |
| [`security-headers`](security_headers.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/api_get_backoffice_user/` · Backoffice Users API · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
