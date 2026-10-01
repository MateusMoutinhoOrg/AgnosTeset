# `POST /api/admin/get-backoffice-user`

Answers one backoffice user by id

## Try it

```bash
curl -X POST localhost:3000/api/admin/get-backoffice-user \
  -H 'Content-Type: application/json' \
  -d '{"id":1}'
```

More examples:

```bash
curl -X POST localhost:3000/api/admin/get-backoffice-user -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"id":2}'
```

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
| [`api-autentication`](api_autentication.md) | depends on the address — `explain-route` gives the exact answer |

---

For developers: `sandbox/internal/routeslist/api_get_backoffice_user/` · Backoffice Users API · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
