# Commands

`Url-Shortner <command> [flags] [args]`. `Url-Shortner help <command>` prints
the same for one command; an empty command line prints the general help and exits 2.

Hidden commands are not listed. Flags may appear anywhere on the command line; positionals
bind in order after them. A `repeatable` field is given once per value. Every section below is
rendered from that command's `entries.yaml` ([EntriesYaml](../EntriesYaml/doc.md)) on each
build.

## Info

### `help` — `--help`

Display help for a command

```bash
Url-Shortner help [<command>]
```

When called without arguments, lists every available command grouped by category. When called with a command name, shows detailed usage, arguments, flags, and examples for that command.

| Argument | Type | Default | Description |
| --- | --- | --- | --- |
| `command` | string |  | The command to describe; omit it to list every command |

```bash
Url-Shortner help
Url-Shortner help start
```

### `version` — `--version`

Print the installed version

```bash
Url-Shortner version
```

Prints the current version of the installed binary and exits.

```bash
Url-Shortner version
```

## Server

### `start-server`

Starts the http server

```bash
Url-Shortner start-server [--addr <addr>] [--read-timeout-ms <read_timeout_ms>] [--write-timeout-ms <write_timeout_ms>]
```

Opens the port and serves every route declared under sandbox/internal/routes, until the process is stopped.

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--addr` | string | `:8080` | the address the server listens on |
| `--read-timeout-ms` | int | `10000` | how long a request has to arrive, in milliseconds |
| `--write-timeout-ms` | int | `10000` | how long a response has to be written, in milliseconds |

```bash
Url-Shortner start-server
Url-Shortner start-server --addr :3000
```

Output channels and exit codes are in [Rules](../Rules/doc.md#output-channels).
