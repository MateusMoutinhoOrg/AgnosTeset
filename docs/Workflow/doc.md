# Workflow

Every change this project takes and the command that makes it. `agnos` owns every generated
file; what stays hand-written is listed in [GeneratedFiles](../GeneratedFiles/doc.md), and the
rules each recipe holds to are in [Rules](../Rules/doc.md).

## The loop

```bash
agnos build      # verify + regenerate every generated file + go mod tidy + compile
agnos verify     # the schema check alone, writes nothing
```

`build` is the only thing that regenerates the dispatch, the wiring,
`README.md` and `docs/`, so run it after every hand edit. It is idempotent: a second run leaves
the tree unchanged. Every command below takes `--path <dir>` (default `.`) and `-q`, and runs
`build` for you.

No recipe below asks for a Go file to be created by hand except the cases listed under
[Hand-written code](#hand-written-code).

## Choose what agnos generates

```bash
agnos list-extensions             # every generation mechanic and whether it is on
agnos enable-extension readme     # start generating README.md again
agnos disable-extension doc       # stop generating docs/, keep what is there
```

`AgnosConfig/extensions.yaml` is what `build` reads to decide what to render. Turning a
mechanic off stops the generation and removes nothing: the files stay, and they are yours to
edit. Every key is in [Extensions](../Extensions/doc.md).

## Change the command surface

```bash
agnos add-command <name> --help "one line" --category "Core"
agnos add-flag <name> --command <cmd> --type string --description "..." [--default . | --required]
agnos add-arg  <name> --command <cmd> --type int --min 1 --description "..."
agnos set-command <cmd> --long-description "..." --example "<cmd> --flag v" --identifier <alias>
agnos remove-flag <name> --command <cmd>
agnos remove-arg  <name> --command <cmd>
agnos remove-command <cmd>
```

`add-command` writes `sandbox/internal/commands/<name>/entries.yaml` (the declaration) and a
stub `handler.go` (yours), then generates `new.go` — the `api.Command` that joins
`Cli.Commands`. Every key these editors write is in
[EntriesYaml](../EntriesYaml/doc.md); never edit `entries.yaml` by hand.

Then write `handler.go` — the whole hand-written half of a command:

```go
func CommandHandler(sandbox *api.Sandbox, command *api.Command) int {
	if err := something(sandbox, command.GetString("path")); err != nil {
		sandbox.Deps.Std.Error("%s\n", err.Error())
		return api.ExitFailure
	}
	sandbox.Deps.Std.Printf("%s\n", result)
	return api.ExitOk
}
```

Every value arrives typed, defaulted and range-checked: bad input already exited 2 before the
handler ran. [Commands](../Commands/doc.md) documents the command on the next build.


## Change the route surface

```bash
agnos add-route <name> --trigger /<path> --method POST --help "one line" --category "Users"
agnos set-route <route> --method PUT --example "curl localhost:8080/users"
agnos add-segment <name> --route <route>            # a capture; --identifier /users for a literal
agnos add-segment <name> --route <route> --array    # the last one, taking the rest of the path
agnos add-header <name> --route <route> --required
agnos add-param <name> --route <route> --type int --default 1 --min 1
agnos set-body <route> --type json --required --max-bytes 2097152
agnos add-body-field <dotted.name> --route <route> --format email --required
agnos import-body <route> --file payload.json --required --infer-format
agnos set-segment <name> --route <route> --type int  # and set-header / set-param
agnos set-body-field <dotted.name> --route <route> --max 130 --clear format
agnos show-route <route>                            # the whole declaration as a tree
agnos remove-segment <name> --route <route>         # and remove-header / remove-param /
agnos remove-body-field <dotted.name> --route <route>
agnos remove-route <route>
```

`add-route` writes `sandbox/internal/routes/<name>/route.yaml` (the declaration) and a stub
`handler.go` (yours), then generates `new.go` — the `api.Route` that lands in
`Server.Routes`, which the dispatch reads every request against.
One editor per place the declaration holds something, so every key of
[RouteYaml](../RouteYaml/doc.md) is reachable from the command line and `route.yaml` is never
edited by hand. `add-body-field` takes a dotted path (`address.city`) and creates the objects
it passes through; `set-body` covers the envelope around the schema — how the body is read,
whether it is required, its size limit and its content-type.

Each `add-` has a `set-` beside it, so a bound that was forgotten is added to the declaration
that is there instead of removing it and declaring it again: the keys given are written over
the ones already declared, `--clear <key>` takes one off, and the result goes through the same
constructor the `add-` side calls. `import-body` is `add-body-field` run once per key of an
example payload — a document pasted with `--json` or read with `--file`, inferring a type per
key, the objects and lists around them and, with `--infer-format`, the four formats a string
may spell; it never writes over a property already declared, and `--replace` starts the schema
over. `show-route` prints the whole declaration as a tree — the path, the headers, the
parameters and the body schema property by property — and is the one of them that writes
nothing.

Then write `handler.go` — the whole hand-written half of a route:

```go
func RouteHandler(sandbox *api.Sandbox, route *api.Route, response serverdeps.Response) int {
	body, status := ReadBody(sandbox, route, response)
	if status != api.StatusOk {
		return status
	}
	return writeJson(sandbox, response, api.StatusCreated, create(sandbox, route.GetString("tenant"), body))
}
```

`route` arrives bound, converted and range-checked — every value read back by the name its
declaration gives it (`GetString`, `GetInt`, `GetBool`, `GetStrings`) — so a bad request was
already answered `400` before the handler ran. The body is the exception — it is read only when `ReadBody` asks for
it. [Routes](../Routes/doc.md) documents the route on the next build, and
[ServerUsage](../ServerUsage/doc.md) is the whole recipe.


## Change the page surface

```bash
agnos add-page <name> --trigger /<path> --title "One Line"
agnos remove-page <name>                             # the route and the html both
```

`add-page` writes the route that answers the page (`route.yaml` + a `handler.go` rendering
through `pageio`) and `assets/frontend/pages/<name>.html`, then generates its `new.go` like
any other route's. A page **is** a route, so every editor of
[RouteYaml](../RouteYaml/doc.md) works on its declaration and [Routes](../Routes/doc.md)
documents it.

Then write the html, naming assets rather than hardcoding links:

```html
{{ dirref "styles" }}
<h1>{{ .Title }}</h1>
```

Every `{{ .Field }}` of the page is one exported field of the `pageVars` struct in its
handler, so a new variable is a compile error until it is declared.
[FrontUsage](../FrontUsage/doc.md) is the whole recipe.

## Change the database surface

```bash
agnos add-database app-database --prefix app
agnos add-table url --database app-database
agnos add-table-field alias --database app-database --table url --type key --required
agnos add-table-field visits --database app-database --table url --type database
agnos add-table-field agent --database app-database --table url --parent visits
agnos show-database app-database                  # read the declaration back
```

`set-table-field` and the `remove-` half of each pair are the inverses. Every command rewrites
`sandbox/internal/databases/<db>/specs.yaml` and runs `build`, which regenerates `api.go`,
`new.go` and `methods.go` from it — the records, the insert structs, the filtrage and the body
of every method.

Then call it from wherever needs it:

```go
db := app_database.New(sandbox)
url, err := db.AddUrl(app_database.UrlNew{Alias: "gh", Link: "https://github.com"})
found, ok := db.FindUrlByAlias("gh")
```

A query the declaration cannot describe goes in `methods_custom.go` beside them, hand-written
and rewritten by no build. [Databases](../Databases/doc.md) is the whole recipe.
## Add reusable logic

`sandbox/internal/<pkg>/`, one directory per concern, imported by whatever needs it. No
declaration, no generated counterpart — write the package and run `build`.

## Add a surface to the sandbox api

The api is what a Go caller gets back from `sandbox.New` (see [LibUsage](../LibUsage/doc.md)).
Two hand-written places, then `build` regenerates `sandbox/api/sandbox.go` and
`sandbox/new.go` around them:

1. `sandbox/api/<x>.go` — the contract: `type <X> struct { ... }` of function fields, named
   after the file, every declaration doc-commented (those comments render
   [PublicApi](../PublicApi/doc.md)). It becomes the `api.Sandbox` field `<X>`.
2. `sandbox/internal/<x>/new.go` — `func New<X>(sandbox *api.Sandbox) api.<X>`, assigning each
   field of the contract, with the implementation beside it.

`build` then writes `sandbox/constructors/<x>/constructor.go` — `sandbox.<X> =
<x>.New<X>(sandbox)` — **once**, and `sandbox/new.go` calls it. From there the constructor is
yours: wrap the implementation, decorate the contract, or build a different one entirely.

## Construct a field yourself

`sandbox/new.go` is one `<x>.Constructor(&self)` per directory of `sandbox/constructors/`, so
adding a directory is adding a call. Write `sandbox/constructors/<x>/constructor.go` with
`func Constructor(sandbox *api.Sandbox)` in `package <x>`, run `build`, and it is wired — the
same way a generated one is, and with no generated file to fight over. Editing a constructor
`build` wrote earlier works for the same reason: nothing rewrites it.

## Add a dependency

Everything the sandbox is not allowed to do itself — filesystem, clock, network, subprocess —
arrives through `sandbox.Deps`. Install a ready-made one:

```bash
agnos list-deps                 # every installable contract
agnos add-dep <dep>             # sandbox/deps/<dep>/ + its default adapter + the go.mod require
agnos add-dep <dep> --adapter <adapter>
agnos remove-dep <dep> [--with-adapters]
```

[DepList](../DepList/doc.md) is the catalogue. For one of your own, write the two halves:

1. `sandbox/deps/<x>/<x>.go` — `type Sandbox struct { ... }` of function fields, no import at all.
2. `adapters/libs/<x>/<x>.go` — `func Bind(deps *deps.Deps) { deps.<X> = <x>.Sandbox{...} }`, any
   import allowed, beside an `adapter.yaml` saying `dep: <x>`.

Then bind it: add `<x>` to `adapters/availables/standard/available.yaml`, or let
`agnos add-dep` do both for a dep of the catalogue. Reach it as `sandbox.Deps.<X>`
from anywhere inside `sandbox/`.

One contract may have several adapters — see [Adapters](../Adapters/doc.md).

## Add a doc

```bash
agnos add-doc <Name> --theme <id> --description "one line"    # themes: AgnosConfig/themes.yaml
agnos add-doc <Name>/<Sub> --description "one line"           # sub-doc, no theme
agnos remove-doc <Name>
```

Write `docs/<Name>/doc.md`; `README.md`'s index, and the parent `Index.md` of a sub-doc, are
regenerated. Describe any new path worth naming in `AgnosConfig/structure.yaml` — that
file is what renders [Structure](../Structure/doc.md).

## Add an example

```bash
agnos add-cli-example <name>       # examples/cli/<name>/example.sh
agnos add-lib-example <name>       # examples/lib/<name>/example.go
agnos exec-test                    # run them all, check each against its golden
agnos exec-test --only <name>      # one example, both sides
agnos update-test <name>           # rewrite that one golden with what it produces now
agnos exec-test --update           # rewrite every golden at once
agnos remove-cli-example <name>
agnos remove-lib-example <name>
```

Write the example itself, ending with the copy out of `TestDir` into `AssertDir` that says what
it asserts: `result.yaml` records `AssertDir`, and an example that copies nothing out fails.
The golden is written by the first `exec-test` and refreshed with `update-test <name>`, which
prints what it changes before writing. Details in [LibExamples](../LibExamples/doc.md) and
[CliExamples](../CliExamples/doc.md).

## Hand-written code

| File | Written when |
| --- | --- |
| `sandbox/internal/commands/<name>/handler.go` | a command does something |
| `sandbox/internal/routes/<name>/handler.go` | a route answers something |
| `assets/frontend/pages/<page>.html`, `assets/frontend/static/**` | a page looks like something |
| `sandbox/internal/<pkg>/*.go` | logic worth reusing |
| `sandbox/api/<x>.go` + `sandbox/internal/<x>/new.go` | a new api surface |
| `sandbox/constructors/<x>/constructor.go` | how a field of the `Sandbox` is built |
| `sandbox/deps/<x>/<x>.go` + `adapters/libs/<x>/<x>.go` + its `adapter.yaml` | a new dependency |

Everything else is regenerated over. Two more files are yours: `AgnosConfig/docs/ReadmeHeader.md`
is the whole of `README.md` above the documentation index, and `LICENSE` is pasted verbatim into
its License section — put whatever license you want there.

## Ship

```bash
agnos compile --target all   # cross-compile ./cmd/main into release/
agnos publish                # build, compile, then a gh release
```

`go build -o release/FinanceApp ./cmd/main` is the plain local binary.
`publish` names the release after `version` in `AgnosConfig/project.yaml`; bump it there
first. `compile` targets: `linux86`, `linuxarm64`, `linuxi32`, `mac86`, `macarm64`,
`windows86`, `windowsi32`, or `all`.
