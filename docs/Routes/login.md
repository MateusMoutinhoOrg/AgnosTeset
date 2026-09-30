# `POST /admin/login`

## Try it

```bash
curl -X POST localhost:3000/admin/login \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'password=text&username=text'
```

## Body

Send form fields, like `name=value&other=value` with the header `Content-Type: application/x-www-form-urlencoded`, up to 1 MB. The body is required.

| Field | What goes there | Required | Rules |
| --- | --- | --- | --- |
| `password` | text | yes |  |
| `username` | text | yes |  |

Example:

```text
password=text&username=text
```

## What comes back

| Status | Means |
| --- | --- |
| `200` | It worked. The answer comes as `application/json`. |
| `400` | Something you sent is missing or has the wrong type or format. The answer's `field` names it. |
| `413` | The body is larger than 1 MB. |
| `415` | The body was not sent with `Content-Type: application/x-www-form-urlencoded`. |

Any route may also answer `404`, `405` or `500`: see [when something goes wrong](doc.md#when-something-goes-wrong).

## Runs first

These routes run before this one, on the same request. Any of them may refuse it — with `401`
or `403`, for example — or let it through to this route.

| Route | When |
| --- | --- |
| [`autentication`](autentication.md) | always |

---

For developers: `sandbox/internal/routeslist/login/` · frontend · [every route](doc.md) · [RouteYaml](../RouteYaml/doc.md)
