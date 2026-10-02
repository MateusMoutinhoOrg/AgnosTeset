# `sandbox/api/backofficeconfig.go`

## `BackofficeConfig`

BackofficeConfig is the part of the Config the backoffice reads, embedded in api.Config so each field is read as sandbox.Config.<Field>. The backoffice-server middleware fills it in front of start-server, before the first request is served. Written once by `agnos backoffice-init` and the project's from then on; `agnos backoffice-purge` removes it.

| Field | Type | Description |
| --- | --- | --- |
| `Secret` | `string` | Secret signs the backoffice session tokens. It is read from the <NAME>_SECRET environment variable, never from the command line. |
| `AllowXForwardedFor` | `bool` | AllowXForwardedFor trusts the last entry of X-Forwarded-For as the client ip, the one a reverse proxy in front of the server appended. Off, the client ip is the connection's own. |
| `InsecureHttp` | `bool` | InsecureHttp serves the backoffice over plain http, for local development: the session cookie drops Secure and no Strict-Transport-Security is sent. |

[every contract](doc.md)
