# `POST /admin/root/edit-user/{Id:integer}`

Edits a backoffice user; a blank password keeps the current one

## Try it

Only what is required:

```bash
curl -X POST localhost:3000/admin/root/edit-user/1 \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'email=text&role=1&username=text'
```

With every value it reads:

```bash
curl -X POST localhost:3000/admin/root/edit-user/1 \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'email=text&password=text&role=1&username=text'
```

## In the address

| Part | What goes there | Example | Description |
| --- | --- | --- | --- |
| `{Id:integer}` | whole number | `1` |  |

## Body

Send form fields, like `name=value&other=value` with the header `Content-Type: application/x-www-form-urlencoded`, up to 1 MB. The body is required.

| Field | What goes there | Required | Rules |
| --- | --- | --- | --- |
| `email` | text | yes |  |
| `password` | text | no |  |
| `role` | whole number | yes |  |
| `username` | text | yes |  |

Example:

```text
email=text&password=text&role=1&username=text
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `text/html`. |
| `400` | Something you sent is missing or has the wrong type or format. The answer's `field` names it. |
| `404` | `{Id:integer}` is not a whole number, so this route does not answer the address. |
| `413` | The body is larger than 1 MB. |
| `415` | The body was not sent with `Content-Type: application/x-www-form-urlencoded`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | depends on the address — `explain-route` gives the exact answer |
| [`root-guard`](root_guard.md) | always |

---

For developers: `sandbox/internal/routeslist/edit_user/` · Admin · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
