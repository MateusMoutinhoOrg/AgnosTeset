# `POST /admin/create-backoffice-api-token`

Creates an API token and shows it once

## Try it

Only what is required:

```bash
curl -X POST localhost:3000/admin/create-backoffice-api-token \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'expiration=7&name=text'
```

With every value it reads:

```bash
curl -X POST localhost:3000/admin/create-backoffice-api-token \
  -H 'origin: my-origin' \
  -H 'host: my-host' \
  -H 'x-client-ip: my-x-client-ip' \
  -H 'x-forwarded-for: my-x-forwarded-for' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'date=text&expiration=7&ips=text&name=text'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `origin` | header | text | no | `my-origin` | the origin of the page that sent the request, sent by the browser — read by [`same-origin`](same_origin.md), which runs first |
| `host` | header | text | no | `my-host` | the host the request was sent to — read by [`same-origin`](same_origin.md), which runs first |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client — read by [`client-ip`](client_ip.md), which runs first |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for — read by [`client-ip`](client_ip.md), which runs first |

## Body

Send form fields, like `name=value&other=value` with the header `Content-Type: application/x-www-form-urlencoded`, up to 1 MB. The body is required.

| Field | What goes there | Required | Rules |
| --- | --- | --- | --- |
| `date` | text | no |  |
| `expiration` | text | yes | one of `7`, `30`, `60`, `90`, `365`, `custom`, `never` |
| `ips` | text | no |  |
| `name` | text | yes |  |

Example:

```text
date=text&expiration=7&ips=text&name=text
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |
| `400` | Something you sent is missing or has the wrong type or format. The answer's `field` names it. |
| `413` | The body is larger than 1 MB. |
| `415` | The body was not sent with `Content-Type: application/x-www-form-urlencoded`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`authentication`](authentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`same-origin`](same_origin.md) | always |
| [`client-ip`](client_ip.md) | always |
| [`security-headers`](security_headers.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/create_backoffice_api_token/` · Backoffice API Tokens · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
