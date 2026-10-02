# `sandbox/api/userconfig.go`

## `UserConfig`

UserConfig is the user-visible subset of api.Config. This type is embedded in sandbox/api/config.go. The only thing that writes it is sandbox/start.go, and you can add fields to it with impunity: if you add a method, just make sure sandbox/api/config.go embeds the updated UserConfig, and nothing breaks. If you add a field called Cli, Config, Actions or Deps, the verify tool will complain, and you should rename it. This is the only API contract you need to worry about. The implementation lives under sandbox/internal/, and can be refactored or replaced wholesale.

| Field | Type | Description |
| --- | --- | --- |
| `Secret` | `string` | Secret signs the backoffice session tokens. start-server reads it from the TESTE_SECRET environment variable, never from the command line. |
| `AllowXForwardedFor` | `bool` | AllowXForwardedFor trusts the last entry of X-Forwarded-For as the client ip, the one a reverse proxy in front of the server appended. Off, the client ip is the connection's own. |
| `InsecureHttp` | `bool` | InsecureHttp serves the backoffice over plain http, for local development: the session cookie drops Secure and no Strict-Transport-Security is sent. |

[every contract](doc.md)
