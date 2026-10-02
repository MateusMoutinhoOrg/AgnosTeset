# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go project (binary name `teste`, module `github.com/MateusMoutinhoOrg/AgnosTeset`) scaffolded and maintained by the **`agnos`** code generator. Most of the tree is generated from YAML declarations; only a small set of files is hand-written. Before changing anything structural, read `docs/Rules/doc.md` (every rule, some enforced by `agnos verify`) and `docs/Workflow/doc.md` (the `agnos` command for each kind of change). The rest of `docs/` is also accurate, and much of it is generated.

## Commands

```bash
agnos build                 # verify + regenerate all generated files + go mod tidy + compile. Run after every hand edit.
agnos verify                # schema/layer check only, writes nothing
go build -o release/teste ./cmd/main
TESTE_SECRET=$(openssl rand -hex 32) go run ./cmd/main start-server [--addr 3000:4000] [--allow-x-forwarded-for] [--insecure-http]
go run ./cmd/main add-backoffice-user --username u --email e [--role root|viewer]   # prints a generated password once
```

- Compile with `./cmd/... ./sandbox/... ./adapters/...`, never `./...`.
- There are no `_test.go` files. Tests are golden-file examples under `examples/{cli,lib}/<name>/` (none exist yet): `agnos exec-test`, a single one with `agnos exec-test --only <name>`, refresh a golden with `agnos update-test <name>`.
- The database is written to `./maindatabase` in the working directory (gitignored).
- Use `agnos explain-route GET /admin/home` to see which routes a request reaches and why the others are skipped.

## Architecture

```
adapters/  -->  sandbox/  <--  cmd/main
(OS + 3rd party) (closed)      (wires adapters into sandbox, no logic)
```

- **`sandbox/` is closed.** Nothing in it imports anything outside `sandbox/`, **including the Go stdlib**. Every outside capability (hashing, strings, sorting, reflection, serialization, embedded files, HTTP, the database) is a contract struct of function fields in `sandbox/deps/<x>/`, reached as `sandbox.Deps.<X>` (e.g. `sandbox.Deps.Hashdeps.Sha256Hex`). Implementations live in `adapters/libs/<x>/` (`Bind(deps *deps.Deps)`), selected in `adapters/availables/standard/available.yaml`. To get a new capability, use `agnos add-dep` or write a contract and adapter pair. Don't import a stdlib package inside `sandbox/`.
- **`sandbox/api/`** holds the contracts that form the public API. `api.Sandbox` embeds `UserSandbox` (`usersandbox.go`) and `api.Config` embeds `UserConfig` (`userconfig.go`). Both files are hand-owned. Every exported symbol in `sandbox/api/` and `sandbox/deps/` needs a doc comment because `docs/PublicApi` is generated from them.
- **Every function in `sandbox/internal/`** takes `sandbox *api.Sandbox` as its first parameter.
- **Commands** are in `sandbox/internal/commands/<name>/`. **Routes** are in `sandbox/internal/routeslist/[<folder>/]<name>/`. Each has a `command.yaml`/`route.yaml` declaration plus generated `new.go`/`entries.go`. Only `InternalPureHandler.go` is hand-written. Change the YAML with `agnos add-flag`/`set-route`/`add-body-field`/etc., never by hand. Handler inputs arrive typed and validated on `entries`.
- **Route dispatch:** all matching routes run in `priority` order, lowest first (middleware on rung 10, routes on 100, `frontend` on 1000). A handler that neither sets a status nor writes has *declined*, and the next route runs. That's how middleware works. Refuse a request with `routeio.Fail(...)`. Unanswered and failed requests go to `sandbox/internal/server/errors/handle_*.go` (hand-owned).
- **Per-request state goes on `props *routeprops.RouteProps`** (fields declared in `sandbox/internal/routeprops/routeprops.go`), which is fresh for each request and shared along its chain. It lives under `sandbox/internal`, so a field may name a project type: `props.User` is the `*maindatabase.BackofficeuserItem` the auth middleware found (nil when none). `*api.Sandbox` is shared across all requests, so don't store per-request data such as the authenticated user on it. The cli mirror is `commandprops.CommandProps`.
- **Databases:** `sandbox/internal/databases/maindatabase/specs.yaml` generates `api.go`/`new.go`/`methods.go`. Edit the specs with `agnos add-table-field` and similar commands. Custom queries go in a hand-written `methods_custom.go`. Use `maindatabase.New(sandbox)` then `db.AddBackofficeuser(...)`. `Find<T>By<Field>` is generated only for fields of `type: key`. Other fields are reached through `List<T>` plus a filtrage.
- **Generated code** under `sandbox/internal/generated/`, `sandbox/new.go`, `sandbox/api/sandbox.go`, `README.md` and most of `docs/` is overwritten on every `build`. Never edit it. Change the source declaration instead (see `docs/GeneratedFiles/doc.md`). Files marked "written once" (constructors, `routeprops.go`, `handle_*.go`, route/command stubs) are yours after creation.
- **Assets:** everything under `assets/` is embedded via `assets/asset.go` and read through `sandbox.Deps.Embeddeps` (e.g. `templates/login.html`). `assets/frontend/**` is served as-is by the `frontend` route.
- **Output:** use `response.Printf` / `deps.Std` channels, never `fmt.Printf`. Exit codes: `api.ExitOk`=0, `api.ExitFailure`=1, `api.ExitUsage`=2.

## Backoffice auth (in progress on `backoffice-auth`)

- `start-server` reads the JWT secret from the `TESTE_SECRET` env var (`backofficeauth.ReadSecret`, at least 32 characters, never a flag or a file) into `sandbox.Config.Secret`, and refuses to start without it. The secret only signs session tokens.
- `start-server` flags, copied into `UserConfig`:
  - `--allow-x-forwarded-for` trusts the last `X-Forwarded-For` entry, for one reverse proxy in front. The server should be bound to an address only that proxy reaches, and it warns when it listens on every interface.
  - `--insecure-http` drops `Secure` from the cookie and stops HSTS, for local plain http.
- Passwords are stored in `passwordhash` as PBKDF2-HMAC-SHA256 (600k iterations, a salt per user) through `Deps.Passworddeps`, via `backofficeauth.HashPassword`.
  - `Authenticate` scans users once and hashes even for an unknown login, so timing doesn't tell the two apart.
- `add-backoffice-user` goes through `backofficeusers.Add`, with the same validation as the web. It generates the password and prints it once, so no password travels in argv. `--role` defaults to `viewer`.
- `role` is an int that maps to `backofficeauth.Role` (`RoleRoot`=0, `RoleViewer`=1).
- Middlewares in front of everything, before `autentication` and `api-autentication` (rung 10):
  - `client-ip` (rung 7, every path) puts the client ip on `props.ClientIp`, worked out by `httpguard.ClientIp`. Routes read that, never `X-Client-Ip`. The `serverdeps` adapter answers only the connection ip and joins every `X-Forwarded-For` line.
  - `security-headers` (rung 8, `/admin` and `/api/admin`) sets `httpguard.SecurityHeaders`: CSP, XFO, nosniff, Referrer-Policy, no-store and HSTS. The CSP runs no inline script, so page JS lives in `assets/frontend/admin/backoffice.js`. Use `data-confirm` on a form, never an `on*` attribute.
  - `same-origin` (rung 9, `/admin`) answers 403 to an `Origin` that isn't the request `Host`. It is `ANY` on purpose, because a method-specific middleware would turn 405s into 404s. A proxy must forward `Host`.
- Rate limiting is in `sandbox/internal/backofficethrottle/`, in memory through `Deps.Ratelimitdeps`, over a 15-minute window:
  - login allows 20 failures per ip and 10 per login, then answers 429 without checking the password;
  - `api-autentication` allows 20 invalid tokens per ip, then answers 429;
  - both 429s carry `Retry-After`.
- Routes under `routeslist/admin/`:
  - `autentication` (spelled that way): an `ANY /admin/...` middleware at priority 10.
  - `login`: `POST /admin/login`, form body.
  - `home`: `GET /admin/home`.
  - `list-backoffice-users`: `GET /admin/list-backoffice-users?search=&role=&page=&limit=&notice=`, open to any backoffice user.
  - the API token pages, open to any backoffice user (see below).
- Routes under `routeslist/admin/root/` (root only):
  - `root-guard`: an `ANY /admin/root/...` middleware at priority 11, right after `autentication`. It answers `templates/forbidden.html` (403) to non-roots.
  - `add-backoffice-user-page`/`add-backoffice-user`: `GET`/`POST /admin/root/add-backoffice-user`.
  - `edit-backoffice-user-page`/`edit-backoffice-user`: `GET`/`POST /admin/root/edit-backoffice-user/{id}`. A blank password keeps the current one.
  - `remove-backoffice-user`: `POST /admin/root/remove-backoffice-user/{id}`.
  - A handler never sees the method, so a form page and its action are two routes on one path.
- Session: a JWT cookie (HttpOnly, SameSite=Strict, Secure unless `--insecure-http`, 30 min) whose `jti` names a `sessions` record nested under the user. Logout or removing the user deletes that record. Removing a user also deletes their API tokens.
- A new password, through `backofficeusers.Update(sandbox, actor, session, id, fields)`, ends every session of the user (`backofficeauth.CloseSessions`) and revokes their API tokens. Only the session of a root editing their own account is spared. The list then shows the `password-changed` notice.
- User-management logic lives in `sandbox/internal/backofficeusers/`:
  - it lists, filters, paginates and validates users;
  - username and email are unique across both columns, ignoring case, and passwords need at least 8 characters;
  - a root can't remove their own account, and the last root can't be demoted.

  Pages are rendered by `sandbox/internal/render/` from `assets/templates/` (`render.BackofficeUsers` → `backoffice_users.html`, `render.Add/EditBackofficeUserForm` → `backoffice_user_form.html`). An action answers `303` to `/admin/list-backoffice-users?notice=<code>`, and `render.noticeOf` words each code.
- Everything here is named `backoffice*` because application users will come later. Don't give these routes, render helpers or templates generic `user` names.

## Backoffice API tokens

The only credential `/api/admin` accepts. Tokens are created and revoked on the HTML pages, and the API never issues or ends one.

- Token: `bo_` + 64 hex chars (`Deps.Randdeps.Hex(32)`), shown once.
  - The top-level `apitoken` table stores only `tokensha` = `SHA-256(token)`, a `key` field, so `FindApitokenByTokensha` is the lookup.
  - It also stores `prefix` (the first 11 chars, for display), `ownerid`, `createdat`, `expiresat` (0 = never), `ips` (comma-joined, "" = any ip), and `lastusedat`/`lastusedip`.
- A token acts as its owner, with the role read fresh on every request.
- Each user manages their own tokens. A root lists and revokes everyone's, and a revoke the actor may not make reads as `not-found`.
- Logic lives in `sandbox/internal/backofficetokens/`:
  - `Create` validates: the name is required, at most 100 characters and unique per owner ignoring case. The expiration is `7/30/60/90/365` days, `custom` (an `<input type="date">`, valid through that UTC day) or `never`. Each ip must be IPv4 or IPv6.
  - `List`, `Revoke`, `RemoveOfOwner` (the cascade from `backofficeusers.Remove`).
  - `Resolve` is what the API middleware calls. It refuses an unknown, expired or ip-disallowed token, and marks the token as last used.
- Routes under `routeslist/admin/`, all for any signed-in user:
  - `list-backoffice-api-tokens`: `GET /admin/list-backoffice-api-tokens?notice=`.
  - `create-backoffice-api-token-page`/`create-backoffice-api-token`: `GET`/`POST /admin/create-backoffice-api-token`. On success the POST answers the list page (`201`) with the token in full, instead of redirecting, so the token never travels in a url.
  - `revoke-backoffice-api-token`: `POST /admin/revoke-backoffice-api-token/{id}`. It answers `303` to the list with a notice.
- Rendering:
  - `render.BackofficeApiTokens` → `backoffice_api_tokens.html`;
  - `render.CreateBackofficeApiTokenForm` → `backoffice_api_token_form.html`;
  - dates are formatted with `Deps.Timedeps.FormatUnix`.
- These are local deps (`origin: local`), the same shape as `jwtdeps`:
  - `randdeps` (`crypto/rand`);
  - `timedeps` (`time`, UTC);
  - `envdeps` (`os.Getenv`);
  - `passworddeps` (`crypto/pbkdf2`);
  - `ratelimitdeps` (fixed-window counters behind a `sync.Mutex`).

## Backoffice JSON API (`/api/admin/`)

The JSON twin of the HTML routes above. It sits outside `/admin`, so the cookie `autentication` middleware never runs on it. Route names carry an `api-` prefix because a route name is unique across folders.

- Every parameter arrives in a JSON body (`Content-Type: application/json`). Every action is a `POST`, except `GET /api/admin/me`, which takes no parameters.
- Auth is `Authorization: Bearer <token>`, where the token is an API token from the section above. `api-autentication` puts the owner on `props.User` and the token on `props.ApiToken`. The API never reads the cookie and has no login or logout.
- `role` travels as its name (`"root"`/`"viewer"`), held by a schema `enum` and converted with `backofficeapi.Role`. A user is answered as `{id, username, email, role}`, never with `passwordhash`.
- Failures go through `routeio.Fail`, so they come back as the default `{"error","field"}` of the `handle_*.go`:
  - 401: a missing, unknown, revoked, expired or ip-disallowed token (with `WWW-Authenticate: Bearer`);
  - 429: too many invalid tokens from the ip (with `Retry-After`);
  - 403: not root, or a root removing their own account;
  - 404: unknown `id`;
  - 400: a schema violation or a `backofficeusers` message.
- `routeslist/api/admin/`:
  - `api-autentication` (`ANY /api/admin/...`, priority 10);
  - `api-me`;
  - `api-list-backoffice-users` (`{search?, role?, page?, limit?}`; an empty body lists everything);
  - `api-get-backoffice-user` (`{id}`).
- `routeslist/api/admin/root/`:
  - `api-root-guard` (priority 11);
  - `api-add-backoffice-user` (`201`);
  - `api-edit-backoffice-user` (`{id, username, email, role, password?}`);
  - `api-remove-backoffice-user` (`{id}`).
- The JSON documents are built in `sandbox/internal/backofficeapi/`, the JSON counterpart of `render`.
