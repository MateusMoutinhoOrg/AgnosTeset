# `GET /{*Rest}`

Serves any file of the embedded assets/frontend tree

## Try it

Only what is required:

```bash
curl localhost:3000/my-rest
```

With every value it reads:

```bash
curl localhost:3000/my-rest \
  -H 'x-client-ip: my-x-client-ip' \
  -H 'x-forwarded-for: my-x-forwarded-for'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{*Rest}` | the rest of the address — one part or more, like `a/b.png` | `my-rest` | the file under assets/frontend; none is its index.html |

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `x-client-ip` | header | text | no | `my-x-client-ip` | the ip of the connection, set by the server and never by the client — read by [`client-ip`](client_ip.md), which runs first |
| `x-forwarded-for` | header | text | no | `my-x-forwarded-for` | the client chain a reverse proxy appended to, read only with --allow-x-forwarded-for — read by [`client-ip`](client_ip.md), which runs first |

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
| [`authentication`](authentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`create-backoffice-api-token-page`](create_backoffice_api_token_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`home`](home.md) | depends on the address — `explain-route` gives the exact answer |
| [`list-backoffice-api-tokens`](list_backoffice_api_tokens.md) | depends on the address — `explain-route` gives the exact answer |
| [`list-backoffice-users`](list_backoffice_users.md) | depends on the address — `explain-route` gives the exact answer |
| [`add-backoffice-user-page`](add_backoffice_user_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`edit-backoffice-user-page`](edit_backoffice_user_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`root-guard`](root_guard.md) | depends on the address — `explain-route` gives the exact answer |
| [`same-origin`](same_origin.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-authentication`](api_authentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-me`](api_me.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-root-guard`](api_root_guard.md) | depends on the address — `explain-route` gives the exact answer |
| [`client-ip`](client_ip.md) | always |
| [`health`](health.md) | depends on the address — `explain-route` gives the exact answer |
| [`security-headers`](security_headers.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/frontend/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
