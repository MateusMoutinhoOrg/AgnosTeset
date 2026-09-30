# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go project (binary name `teste`, module `github.com/MateusMoutinhoOrg/AgnosTeset`) scaffolded and maintained by the **`agnos`** code generator. Most of the tree is generated from YAML declarations; only a small set of files is hand-written. Before changing anything structural, read `docs/Rules/doc.md` (every rule, some enforced by `agnos verify`) and `docs/Workflow/doc.md` (the `agnos` command for each kind of change). The rest of `docs/` is also accurate, and much of it is generated.

## Commands

```bash
agnos build                 # verify + regenerate all generated files + go mod tidy + compile. Run after every hand edit.
agnos verify                # schema/layer check only, writes nothing
go build -o release/teste ./cmd/main
go run ./cmd/main start-server --secret <s> [--addr 3000:4000]
go run ./cmd/main add-backoffice-user --username u --email e --password p --secret <s>
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
- **Per-request state goes on `props *api.RouteProps`** (fields declared in `sandbox/api/routeprops.go`), which is fresh for each request and shared along its chain. `*api.Sandbox` is shared across all requests, so don't store per-request data such as the authenticated user on it.
- **Databases:** `sandbox/internal/databases/maindatabase/specs.yaml` generates `api.go`/`new.go`/`methods.go`. Edit the specs with `agnos add-table-field` and similar commands. Custom queries go in a hand-written `methods_custom.go`. Use `maindatabase.New(sandbox)` then `db.AddBackofficeuser(...)`. `Find<T>By<Field>` is generated only for fields of `type: key`. Other fields are reached through `List<T>` plus a filtrage.
- **Generated code** under `sandbox/internal/generated/`, `sandbox/new.go`, `sandbox/api/sandbox.go`, `README.md` and most of `docs/` is overwritten on every `build`. Never edit it. Change the source declaration instead (see `docs/GeneratedFiles/doc.md`). Files marked "written once" (constructors, `routeprops.go`, `handle_*.go`, route/command stubs) are yours after creation.
- **Assets:** everything under `assets/` is embedded via `assets/asset.go` and read through `sandbox.Deps.Embeddeps` (e.g. `templates/login.html`). `assets/frontend/**` is served as-is by the `frontend` route.
- **Output:** use `response.Printf` / `deps.Std` channels, never `fmt.Printf`. Exit codes: `api.ExitOk`=0, `api.ExitFailure`=1, `api.ExitUsage`=2.

## Backoffice auth (in progress on `backoffice-auth`)

- `start-server --secret` copies the secret into `sandbox.Config.Secret` (`UserConfig`) for route handlers.
- `add-backoffice-user` stores `SHA-256(secret + password)` in the `backofficeuser` table. `role` is an int that maps to `api.UserRole` (`UserRoleRoot`=0, `UserRoleViewer`=1).
- Routes under `routeslist/admin/`:
  - `autentication` (spelled that way): an `ANY /admin/...` middleware at priority 10.
  - `login`: `POST /admin/login`, form body.
  - `home`: `GET /admin/home`.
- The login and middleware handlers are still stubs. Comments in each describe the intended flow: a JWT cookie (HttpOnly, SameSite=Strict, 30 min), rendering `templates/login.html` on failure, and redirecting to `/admin/home`.
