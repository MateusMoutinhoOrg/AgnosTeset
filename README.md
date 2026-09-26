# teste

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Agnos.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Agnos)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Agnos)](https://github.com/MateusMoutinhoOrg/Agnos/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue)](go.mod)

Place your description in AgnosConfig/docs/ReadmeHeader.md



## Documentation

### CliUsage

Driving the CLI from a terminal - install, commands, flags, exit codes

| Doc | Description |
| --- | --- |
| [CliInstall](docs/CliInstall/doc.md) | Install teste: a released binary per platform, or a build from source |
| [Commands](docs/Commands/doc.md) | Every command of teste, generated from the command declarations |
| [CliExamples](docs/CliExamples/doc.md) | Index of every runnable example of the teste cli |

### ServerUsage

Serving http - bring the server up, declare routes, read a body

| Doc | Description |
| --- | --- |
| [ServerUsage](docs/ServerUsage/doc.md) | Serve http from teste: bring the layer up, declare routes, read a body |
| [FrontUsage](docs/FrontUsage/doc.md) | Serve a website from teste: bring the front layer up, drop files in assets/frontend, add pages |
| [Routes](docs/Routes/doc.md) | Every route of teste, generated from the route declarations |

### LibUsage

Using the project as a Go module - wiring the deps, calling the sandbox

| Doc | Description |
| --- | --- |
| [LibUsage](docs/LibUsage/doc.md) | Use teste as a Go module: wire the deps, build the sandbox, call its API |
| [PublicApi](docs/PublicApi/doc.md) | Every exported symbol of teste, generated from the contract sources and their doc comments |
| [LibExamples](docs/LibExamples/doc.md) | Index of every runnable example of teste as a Go module |

### Architecture

How the project is put together - layers, boundaries, data flow

| Doc | Description |
| --- | --- |
| [Adapters](docs/Adapters/doc.md) | Contract, adapter and available: three units, one field of Deps, and who fills it |

### Development

Changing this repository - schema, build mechanics, recipes

| Doc | Description |
| --- | --- |
| [Requirements](docs/Requirements/doc.md) | The two tools this project needs — Go and agnos — installed per platform |
| [Workflow](docs/Workflow/doc.md) | Every change this project takes and the agnos command that makes it |
| [Rules](docs/Rules/doc.md) | Every rule the generators, `verify` and the hand-written files must hold to |
| [Structure](docs/Structure/doc.md) | The project schema: what lives where, what is generated, what verify enforces |

### Reference

Lookup tables - schemas, file formats, generated file listings

| Doc | Description |
| --- | --- |
| [EntriesYaml](docs/EntriesYaml/doc.md) | Every key of a command's entries.yaml and what the generated code does with it |
| [Extensions](docs/Extensions/doc.md) | The generation mechanics this project turns on, and what each one writes |
| [DepList](docs/DepList/doc.md) | Every dep `agnos add-dep` can add, the adapters that fill it, and what backs each one |
| [GeneratedFiles](docs/GeneratedFiles/doc.md) | Every file agnos writes into this project and whether build overwrites it |
| [LibExamples](docs/LibExamples/doc.md) | Index of every runnable example of teste as a Go module |
| [CliExamples](docs/CliExamples/doc.md) | Index of every runnable example of the teste cli |
| [RouteYaml](docs/RouteYaml/doc.md) | Every key of a route's route.yaml and what the generated code does with it |

## License

Place your license in this file. It is pasted verbatim into the License section of README.md.

