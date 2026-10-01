# `GET /admin/list-backoffice-users`

Lists backoffice users, filtered and paginated

## Try it

Only what is required:

```bash
curl localhost:3000/admin/list-backoffice-users
```

With every value it reads:

```bash
curl 'localhost:3000/admin/list-backoffice-users?search=my-search&role=my-role&page=1&limit=20&notice=my-notice'
```

## Query string, headers and cookies

| Name | Sent in | What goes there | Required | Example | Description |
| --- | --- | --- | --- | --- | --- |
| `search` | query string | text | no | `my-search` | matches the username or the email, case-insensitive |
| `role` | query string | text | no | `my-role` | root, viewer, or empty for every role |
| `page` | query string | whole number | no — `1` when left out | `1` | the page to show, counted from 1 |
| `limit` | query string | whole number | no — `20` when left out | `20` | how many users a page shows, from 1 to 100 |
| `notice` | query string | text | no | `my-notice` | the outcome code of the last add, edit or remove |

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |
| `400` | Something you sent is missing or has the wrong type or format. The answer's `field` names it. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/list_backoffice_users/` · Admin · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
