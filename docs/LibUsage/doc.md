# LibUsage

`NewRoutes` is a Go module before it is anything else: every feature lives in `sandbox/`
and is reachable from any Go program that imports it.

```bash
go get github.com/MateusMoutinhoOrg/AgnosTeset@latest
```

## Wiring

`sandbox/` performs no OS effects of its own — filesystem, clock, stdout, processes all
arrive through a `deps.Deps` struct. `adapters/availables/standard` builds the ready-made
assembly, and `sandbox.New` turns it into the API object, which carries the deps on
`Sandbox.Deps` — so everything inside reaches them through the api it was handed.

```go
package main

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/adapters/availables/standard"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox"
)

func main() {
	deps := standard.New()    // every adapter lib bound
	lib := sandbox.New(&deps) // *api.Sandbox

	_ = lib
}
```

## What the sandbox exposes

`*api.Sandbox` is a flat struct, one field per contract declared in `sandbox/api/`.
Everything callable from Go is behind one of them.

| Field | Type |
| --- | --- |
| `lib.Cli` | `api.Cli` |
| `lib.Config` | `api.Config` |
| `lib.Server` | `api.Server` |

`lib.Cli.Commands` (`[]api.Command`) is the command surface itself: every command
the project declares, each carrying its flags, its args and the `Handler` that runs it.
`api.BindCommand(&command)` copies one into the command a single run binds to, so a caller
drives a command without a command line — bind the values into the copy's `Items` and call
`copy.Handler(copy)`.

`lib.Server.Routes` (`[]api.Route`) is the http surface the same way:
every route the project declares, in run order — lowest `Priority` first — each carrying its
`paths`, its headers, its params, its body and the `Handler` that answers it.
`api.BindRoute(&route)` copies one into the route a single request runs on, so a caller drives a
route without a socket — bind the values into the copy's `Items` and call `copy.Handler(copy)`,
which returns the failure it did not answer itself and `nil` otherwise. What it answered with is
the status it wrote on the response, never what it returned.

[PublicApi](../PublicApi/doc.md) lists every one of them — signatures, props structs and
dependency contracts — generated from `sandbox/api/` itself on every build.

## Custom deps

Every sub-contract is a struct of function fields, so any of them can be swapped for a
test double, an in-memory implementation or an instrumented wrapper. Patch fields **before**
`sandbox.New(&deps)`: the constructors capture the pointer.

```go
deps := standard.New()

var out bytes.Buffer
deps.Std.Printf = func(f string, a ...any) (int, error) {
	return fmt.Fprintf(&out, f, a...)
}

lib := sandbox.New(&deps)
```

The contracts available to patch:

| Field | Contract package |
| --- | --- |
| `deps.Argvdeps` | `sandbox/deps/argvdeps` |
| `deps.Serializables` | `sandbox/deps/serializables` |
| `deps.Serverdeps` | `sandbox/deps/serverdeps` |
| `deps.Sortdeps` | `sandbox/deps/sortdeps` |
| `deps.Std` | `sandbox/deps/std` |
| `deps.Stringsdeps` | `sandbox/deps/stringsdeps` |

Each one is filled by a matching implementation under `adapters/libs/`, every package
exposing the same `Bind(deps *deps.Deps)` entry point:

| Adapter lib | Binder |
| --- | --- |
| `adapters/libs/argvdeps` | `argvdeps.Bind(&deps)` |
| `adapters/libs/serializables` | `serializables.Bind(&deps)` |
| `adapters/libs/serverdeps` | `serverdeps.Bind(&deps)` |
| `adapters/libs/sortdeps` | `sortdeps.Bind(&deps)` |
| `adapters/libs/std` | `std.Bind(&deps)` |
| `adapters/libs/stringsdeps` | `stringsdeps.Bind(&deps)` |

Starting from `standard.New()` is the safe default: an unfilled field is a nil func that
panics on first call. For a permanent mix, write your own
`adapters/availables/<name>/new.go` binding only the libs you want — `standard/new.go` is
regenerated on every build, while other directories under `availables/` are left alone.

`sandbox/api` is pure contract and `sandbox/` never touches the OS, so both are safe to import
anywhere; the rest of the rules a caller can count on are in [Rules](../Rules/doc.md#layers),
and [DepList](../DepList/doc.md) lists every contract that can be added.
