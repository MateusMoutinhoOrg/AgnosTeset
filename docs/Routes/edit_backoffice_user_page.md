# `GET /admin/root/edit-backoffice-user/{Id:integer}`

Shows the form that edits a backoffice user

## Try it

Only what is required:

```bash
curl localhost:3000/admin/root/edit-backoffice-user/1
```

With every value it reads:

```bash
curl localhost:3000/admin/root/edit-backoffice-user/1 \
  -H 'origin: my-origin' \
  -H 'host: my-host' \
  -H 'x-client-ip: my-x-client-ip' \
  -H 'x-forwarded-for: my-x-forwarded-for'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{Id:integer}` | whole number | `1` |  |

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `origin` | header | text | no | `my-origin` | the origin of the page that sent the request, sent by the browser — read by [`same-origin`](same_origin.md), which runs first |
| `host` | header | text | no | `my-host` | the host the request was sent to — read by [`same-origin`](same_origin.md), which runs first |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client — read by [`client-ip`](client_ip.md), which runs first |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for — read by [`client-ip`](client_ip.md), which runs first |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |
| `404` | `{Id:integer}` is not a whole number, so this route does not answer the address. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`authentication`](authentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`root-guard`](root_guard.md) | always |
| [`same-origin`](same_origin.md) | always |
| [`client-ip`](client_ip.md) | always |
| [`security-headers`](security_headers.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/edit_backoffice_user_page/` · Backoffice Users · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
