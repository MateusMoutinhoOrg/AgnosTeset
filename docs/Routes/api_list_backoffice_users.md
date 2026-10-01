# `POST /api/admin/list-backoffice-users`

Lists backoffice users, filtered and paginated

## Try it

Only what is required:

```bash
curl -X POST localhost:3000/api/admin/list-backoffice-users
```

With every value it reads:

```bash
curl -X POST localhost:3000/api/admin/list-backoffice-users \
  -H 'Content-Type: application/json' \
  -d '{"limit":1,"page":1,"role":"root","search":"text"}'
```

More examples:

```bash
curl -X POST localhost:3000/api/admin/list-backoffice-users -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"search":"ana","role":"viewer","page":1,"limit":20}'
```

## Body

Send JSON with the header `Content-Type: application/json`, up to 1 MB. The body is optional.

| Field | What goes there | Required | Rules |
| --- | --- | --- | --- |
| `limit` | whole number | no |  |
| `page` | whole number | no |  |
| `role` | text | no | one of `root`, `viewer` |
| `search` | text | no |  |

Example:

```json
{
  "limit": 1,
  "page": 1,
  "role": "root",
  "search": "text"
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
| [`api-autentication`](api_autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/api_list_backoffice_users/` · Backoffice Users API · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
