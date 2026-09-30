# `sandbox/api/usersandbox.go`

## `UserSandbox`

UserSandbox is the user-visible contract of the sandbox: a subset of api.Sandbox that does not include fields with names that would collide with top-level fields (Cli, Config, Actions, Deps). It is deliberately small and can be extended by hand without breaking changes because it is not rewritten by code generators. This type is embedded in sandbox/api/sandbox.go. The only thing that writes it is sandbox/start.go, and you can add fields to it with impunity: if you add a method, just make sure sandbox/api/sandbox.go embeds the updated UserSandbox, and nothing breaks. If you add a field called Cli, Config, Actions or Deps, the verify tool will complain, and you should rename it. This is the only API contract you need to worry about. The implementation lives under sandbox/internal/, and can be refactored or replaced wholesale.

| Field | Type |
| --- | --- |
| `User` | `User` |

[every contract](doc.md)
