# `start-server`

Starts the http server

```bash
teste start-server [--addr <addr>] [--read-timeout-ms <read_timeout_ms>] [--write-timeout-ms <write_timeout_ms>] [--shutdown-timeout-ms <shutdown_timeout_ms>]
```

Opens the port and serves every route declared under sandbox/internal/routeslist, until the process is stopped. An interrupt (Ctrl+C) or a termination request stops it gracefully: no new request is taken, and the ones in flight get --shutdown-timeout-ms to finish.

| Flag | Type | Default | Description |
| --- | --- | --- | --- |
| `--addr` | string | `3000:4000` | the port the server listens on, or a range of ports it takes the first free one of, with or without a host (8080, 4000:5000, 127.0.0.1:4000:5000, :8080) |
| `--read-timeout-ms` | int | `10000` | how long a request has to arrive, in milliseconds |
| `--write-timeout-ms` | int | `10000` | how long a response has to be written, in milliseconds |
| `--shutdown-timeout-ms` | int | `10000` | how long the requests in flight have to finish once the server is asked to stop, in milliseconds (0 waits for them) |

```bash
teste start-server
teste start-server --addr 4000:5000
```

Server · [every command](doc.md) · [EntriesYaml](../EntriesYaml/doc.md)
