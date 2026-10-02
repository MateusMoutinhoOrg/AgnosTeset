# `start-server`

Starts the http server

```bash
teste start-server [--addr <addr>] [--read-timeout-ms <read-timeout-ms>] [--write-timeout-ms <write-timeout-ms>] [--shutdown-timeout-ms <shutdown-timeout-ms>] [--allow-x-forwarded-for] [--insecure-http] [--help]
```

Opens the port and serves every route declared under sandbox/internal/routeslist, until the process is stopped. An interrupt (Ctrl+C) or a termination request stops it gracefully: no new request is taken, and the ones in flight get --shutdown-timeout-ms to finish.

The secret that signs the backoffice sessions is read from the TESTE_SECRET environment variable, at least 32 characters (openssl rand -hex 32); without it the server does not start. It is never a flag, which every user of the machine reads, nor a file, which can end up committed with the code.

The session cookie is Secure and Strict-Transport-Security is sent, so the backoffice is served over https, by a reverse proxy that terminates TLS; --insecure-http turns both off, for local development over plain http.

Behind one reverse proxy, bind the server to an address only the proxy reaches (--addr 127.0.0.1:3000) and pass --allow-x-forwarded-for: the client ip is then the last entry of X-Forwarded-For, the one the proxy appended. In nginx: proxy_set_header Host $host; proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for. With more layers in front (a load balancer, a CDN), have nginx resolve the client ip with its real_ip module and send proxy_set_header X-Forwarded-For $remote_addr.

| Flag | Type | Default | Description | From |
| --- | --- | --- | --- | --- |
| `--addr` | string | `3000:4000` | the port the server listens on, or a range of ports it takes the first free one of, with or without a host (8080, 4000:5000, 127.0.0.1:4000:5000, :8080) | — |
| `--read-timeout-ms` | integer | `10000` | how long a request has to arrive, in milliseconds | — |
| `--write-timeout-ms` | integer | `10000` | how long a response has to be written, in milliseconds | — |
| `--shutdown-timeout-ms` | integer | `10000` | how long the requests in flight have to finish once the server is asked to stop, in milliseconds (0 waits for them) | — |
| `--allow-x-forwarded-for` | boolean |  | trust the last entry of X-Forwarded-For as the client ip: turn it on only behind one reverse proxy that appends it (nginx), with the server bound to an address only that proxy reaches (--addr 127.0.0.1:3000) | — |
| `--insecure-http` | boolean |  | serve the backoffice over plain http, for local development only: the session cookie drops Secure and no Strict-Transport-Security is sent | — |
| `--help`, `-h` | boolean |  | Print the help of the command this command line is for | [help-flag](help-flag.md) |

| Runs in front of it | When |
| --- | --- |
| [`help-flag`](help-flag.md) | always |

```bash
teste start-server
teste start-server --addr 4000:5000
teste start-server --addr 127.0.0.1:3000 --allow-x-forwarded-for
```

Server · [every command](doc.md) · [CommandYaml](../CommandYaml/doc.md)
