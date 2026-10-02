# Routes

Every address this server answers. Open one to see what to send, a request you can run as it is,
and what comes back.

The requests call `localhost:3000`, where `teste start-server` listens when that port is free —
it prints the address it took. Change it to wherever your server runs.

## How to read an address

| In the address | Means | For example |
| --- | --- | --- |
| `GET`, `POST`, … | the method to send it with; `ANY` takes every one | `curl -X POST …` |
| `/users` | exactly that text | `/users` |
| `{name}` | a value you choose | `/users/{tenant}` -> `/users/acme` |
| `{name:integer}` | a value of that type: `integer`, `number` or `uuid` | `/articles/{id:integer}` -> `/articles/42` |
| `{*name}` | the rest of the address, one part or more | `/files/{*file}` -> `/files/a/b.png` |
| `*` | anything else, or nothing | `/admin/*` -> `/admin`, `/admin/users` |
| `(a\|b)` | one of these words | `/(en\|pt)` -> `/en` |

## Middleware

| Route | What it does |
| --- | --- |
| [`ANY /admin/* !(/admin/login)`](autentication.md) | Requires a valid backoffice session on /admin, except /admin/login |
| [`ANY /admin/root/*`](root_guard.md) | Lets only root users reach /admin/root |
| [`ANY /admin/*`](same_origin.md) | Refuses a request to /admin another site's page sent |
| [`ANY /api/admin/*`](api_autentication.md) | Requires a valid API token on every /api/admin route |
| [`ANY /api/admin/root/*`](api_root_guard.md) | Lets only root users reach /api/admin/root |
| [`ANY /*`](client_ip.md) | Works out the client ip every route after it reads, from the connection or the reverse proxy in front |
| [`ANY /~(^/(api/)?admin(/|$))`](security_headers.md) | Sends the security headers on every /admin and /api/admin response |

## Backoffice API Tokens

| Route | What it does |
| --- | --- |
| [`POST /admin/create-backoffice-api-token`](create_backoffice_api_token.md) | Creates an API token and shows it once |
| [`GET /admin/create-backoffice-api-token`](create_backoffice_api_token_page.md) | Shows the form that creates an API token |
| [`GET /admin/list-backoffice-api-tokens`](list_backoffice_api_tokens.md) | Lists your API tokens, or every user's for a root |
| [`POST /admin/revoke-backoffice-api-token/{Id:integer}`](revoke_backoffice_api_token.md) | Revokes an API token: your own, or anyone's for a root |

## Backoffice

| Route | What it does |
| --- | --- |
| [`GET /admin/home`](home.md) | Shows the backoffice home page to the signed-in user |
| [`POST /admin/login`](login.md) | Signs a backoffice user in and sets the session cookie |
| [`POST /admin/logout`](logout.md) | Ends the current session |

## Backoffice Users

| Route | What it does |
| --- | --- |
| [`GET /admin/list-backoffice-users`](list_backoffice_users.md) | Lists backoffice users, filtered and paginated |
| [`POST /admin/root/add-backoffice-user`](add_backoffice_user.md) | Adds a backoffice user |
| [`GET /admin/root/add-backoffice-user`](add_backoffice_user_page.md) | Shows the form that adds a backoffice user |
| [`POST /admin/root/edit-backoffice-user/{Id:integer}`](edit_backoffice_user.md) | Edits a backoffice user; a blank password keeps the current one |
| [`GET /admin/root/edit-backoffice-user/{Id:integer}`](edit_backoffice_user_page.md) | Shows the form that edits a backoffice user |
| [`POST /admin/root/remove-backoffice-user/{Id:integer}`](remove_backoffice_user.md) | Removes a backoffice user and every session of it |

## Backoffice Users API

| Route | What it does |
| --- | --- |
| [`POST /api/admin/get-backoffice-user`](api_get_backoffice_user.md) | Answers one backoffice user by id |
| [`POST /api/admin/list-backoffice-users`](api_list_backoffice_users.md) | Lists backoffice users, filtered and paginated |
| [`POST /api/admin/root/add-backoffice-user`](api_add_backoffice_user.md) | Adds a backoffice user |
| [`POST /api/admin/root/edit-backoffice-user`](api_edit_backoffice_user.md) | Edits a backoffice user; a missing or blank password keeps the current one |
| [`POST /api/admin/root/remove-backoffice-user`](api_remove_backoffice_user.md) | Removes a backoffice user and every session of it |

## Backoffice API

| Route | What it does |
| --- | --- |
| [`GET /api/admin/me`](api_me.md) | Answers the backoffice user of the Bearer token |

## Assets

| Route | What it does |
| --- | --- |
| [`GET /{*Rest}`](frontend.md) | Serves any file of the embedded assets/frontend tree |

## Server

| Route | What it does |
| --- | --- |
| [`GET /health`](health.md) | Reports that the server is up |

## When something goes wrong

| Status | Means |
| --- | --- |
| `400` | Something you sent is missing or has the wrong type or format |
| `401` | You have to identify yourself first — a token, for example |
| `403` | You are identified, but not allowed to do this |
| `404` | No route answers this address |
| `405` | The address exists, but not for this method — a `GET` where it takes a `POST`, for example |
| `413` | The body is too large |
| `415` | The body is not in the format the route reads — check `Content-Type` |
| `500` | The server failed while answering |

Unless the project changed it, the answer to an error is JSON naming what went wrong and, when
it is one value, which one:

```json
{"error": "required parameter 'authorization' is missing", "field": "authorization"}
```

For developers: each page is generated on every build from
`sandbox/internal/routeslist/<name>/route.yaml` ([RouteYaml](../RouteYaml/doc.md)); hidden routes
are left out. `agnos list-routes` prints the routes in the order they run, and
`agnos explain-route <METHOD> <path>` which ones a request reaches. The error answers are
the eight files of `sandbox/internal/server/errors/` ([RouteYaml](../RouteYaml/doc.md#failures)).
