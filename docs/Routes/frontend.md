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
| [`create-backoffice-api-token-page`](create_backoffice_api_token_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`home`](home.md) | depends on the address — `explain-route` gives the exact answer |
| [`list-backoffice-api-tokens`](list_backoffice_api_tokens.md) | depends on the address — `explain-route` gives the exact answer |
| [`list-backoffice-users`](list_backoffice_users.md) | depends on the address — `explain-route` gives the exact answer |
| [`add-backoffice-user-page`](add_backoffice_user_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`edit-backoffice-user-page`](edit_backoffice_user_page.md) | depends on the address — `explain-route` gives the exact answer |
| [`root-guard`](root_guard.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-autentication`](api_autentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-me`](api_me.md) | depends on the address — `explain-route` gives the exact answer |
| [`api-root-guard`](api_root_guard.md) | depends on the address — `explain-route` gives the exact answer |
| [`health`](health.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/frontend/` · Assets · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
