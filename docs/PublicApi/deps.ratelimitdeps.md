# `deps.Ratelimitdeps`

`sandbox/deps/ratelimitdeps`

## `Sandbox`

Sandbox is the rate limiter injected whole as the Deps.Ratelimitdeps field. Every function is safe to call from requests served concurrently.

| Field | Type | Description |
| --- | --- | --- |
| `Hit` | `func(key string, windowSeconds int64) int` | Hit records one hit under key, in a window of windowSeconds, and returns how many hits the window holds with it included. |
| `Count` | `func(key string, windowSeconds int64) int` | Count returns how many hits key holds in its open window, 0 when it has none or its window of windowSeconds closed. It records nothing. |
| `Reset` | `func(key string)` | Reset forgets every hit of key. |

[every contract](doc.md)
