# `start-server`

Starts the http server

```bash
Url-Shortner start-server [--addr <addr>] [--read-timeout-ms <read_timeout_ms>] [--write-timeout-ms <write_timeout_ms>] --root-password <root-password>
```

Opens the port and serves every route declared under sandbox/internal/routes, until the process is stopped.

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--addr` | string | `:8080` | the address the server listens on |
| `--read-timeout-ms` | int | `10000` | how long a request has to arrive, in milliseconds |
| `--write-timeout-ms` | int | `10000` | how long a response has to be written, in milliseconds |
| `--root-password` | string, required |  | root password required to list links |

```bash
Url-Shortner start-server
Url-Shortner start-server --addr :3000
```

Server · [every command](doc.md) · [EntriesYaml](../EntriesYaml/doc.md)
