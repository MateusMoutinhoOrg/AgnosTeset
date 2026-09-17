# `sandbox/api/route.go`

## `RouteField`

RouteField is one value a route binds off a request — the parsed form of one captured segment of `paths`, or of one entry under `headers` or `params` in that route's route.yaml. It is matched by Id: the header name, the query key or the name the captured segment is declared under, all read back with the Get* readers of Route.

| Field | Type | Description |
| --- | --- | --- |
| `Type` | `string` | Type is the declared type: "string", "boolean", "int" or "float". |
| `Id` | `string` | Id is the name the field is declared and read back under. |
| `Required` | `bool` | Required reports that the request is rejected without it. |
| `Array` | `bool` | Array reports that every occurrence is kept, not just the first; on the last captured segment of `paths` it takes every segment left in the path. |
| `Description` | `string` | Description is the one-line help text. |
| `Examples` | `[]string` | Examples are whole requests the docs print. |
| `Default` | `string` | Default is the value bound when the field is absent, spelled as it is written in route.yaml. |
| `HasDefault` | `bool` | HasDefault tells a declared empty default from no default at all. |
| `Min` | `float64` | Min and Max bound a numeric value; HasMin and HasMax tell a bound of zero from no bound. |
| `HasMin` | `bool` |  |
| `Max` | `float64` |  |
| `HasMax` | `bool` |  |

## `RoutePath`

RoutePath is one segment of the route's url, in declaration order. Exactly one of the two is filled: a literal the request path has to spell, or a captured field that matches whatever sits in its place.

| Field | Type | Description |
| --- | --- | --- |
| `Identifier` | `string` | Identifier is the literal segment, leading slash included ("/users"); "/" alone fixes the empty path and matches no segment of its own. It is empty on a capture. |
| `Field` | `*RouteField` | Field is the captured segment's declaration, nil on a literal. Only the last entry of Paths may declare an Array, and it then takes every segment left in the path. |

## `RouteBody`

RouteBody is the request body a route declares — the parsed form of `body:` in its route.yaml. The dispatch enforces ContentType and MaxBytes before the handler runs; reading the body itself is the handler's to ask for, through the ReadBody its own package generates.

| Field | Type | Description |
| --- | --- | --- |
| `Type` | `string` | Type is the declared type: "none", "raw", "text" or "json". |
| `Required` | `bool` | Required reports that an absent or empty body is rejected. |
| `MaxBytes` | `int` | MaxBytes is the largest body the route reads; a longer one is 413. |
| `ContentType` | `string` | ContentType is the media type the route accepts; a divergent one is 415. It is empty when the route accepts any. |
| `Schema` | `string` | Schema is the declared json-schema in canonical form, "" when the route declares none. |

## `Route`

Route is one http route of the project, as the sandbox offers it: the whole of what its route.yaml declares, plus the handler behind it. Server.Routes holds one per sandbox/internal/routes/<name>/, each built by that package's generated NewRoute, in match order. What Server.Routes holds is the declaration alone: nothing is ever bound onto it. The server dispatch copies it with BindRoute, fills that copy's Items from the path, the headers and the query string, and hands it to Handler; a caller holding the sandbox can do the same.

| Field | Type | Description |
| --- | --- | --- |
| `Name` | `string` | Name is the package directory of the route, snake_case. |
| `Method` | `string` | Method is the http method it answers to ("GET", "POST"). |
| `Pattern` | `string` | Pattern is its path as it reads in docs and messages ("/users/{tenant}/create"). |
| `Paths` | `[]RoutePath` | Paths are the url's segments, in declaration order. |
| `Headers` | `[]RouteField` | Headers are the request headers it binds, in declaration order. |
| `Params` | `[]RouteField` | Params are the query parameters it binds, in declaration order. |
| `Body` | `RouteBody` | Body is the request body declaration. |
| `Category` | `string` | Category groups it on the generated Routes page. |
| `Help` | `string` | Help is the one-line description. |
| `LongDescription` | `string` | LongDescription is the paragraph the Routes page prints. |
| `Examples` | `[]string` | Examples are whole requests the Routes page prints. |
| `Hidden` | `bool` | Hidden keeps it off the Routes page without disabling it. |
| `Items` | `map[string][]any` | Items holds the values bound to the declaration: one entry per field Id, in request order for an array, one element long for a scalar, absent when nothing was bound. The Get* fields read it. |
| `Request` | `any` | Request is the http request this instance was bound from and Response the one being written. Both are handed over as any: sandbox/api may name no type of sandbox/deps, so the server layer reads them back through routeio.RequestOf and routeio.ResponseOf. |
| `Response` | `any` |  |
| `GetItem` | `func(id string) []any` | GetItem returns every value bound under one Id, nil when none was. |
| `GetString` | `func(id string) string` | GetString returns the first string bound under one Id, "" when none was. |
| `GetBool` | `func(id string) bool` | GetBool returns the first boolean bound under one Id, false when none was. |
| `GetInt` | `func(id string) int` | GetInt returns the first int bound under one Id, 0 when none was. |
| `GetFloat` | `func(id string) float64` | GetFloat returns the first float bound under one Id, 0 when none was. |
| `GetStrings` | `func(id string) []string` | GetStrings returns every string bound under one Id, in order. |
| `GetInts` | `func(id string) []int` | GetInts returns every int bound under one Id, in order. |
| `GetFloats` | `func(id string) []float64` | GetFloats returns every float bound under one Id, in order. |
| `Handler` | `func(route *Route) int` | Handler runs the route against one bound copy — the values in its Items and the response it carries — and returns the status it answered with. It takes that copy rather than closing over one, so the declaration on Server.Routes is shared by every request while nothing bound ever is. It is the route package's own RouteHandler, closed over the sandbox. |

| Function | Description |
| --- | --- |
| `NewRoute() *Route` | NewRoute returns an empty Route with Items open and every Get* reader bound to it. A generated NewRoute fills the declaration and the handler on top of what this returns, so every route reads its values the same way. |
| `BindRoute(route *Route) *Route` | BindRoute copies one declaration into the route a single request runs on: the same declared fields — the slices are read-only and shared — with Items empty, the readers pointed at the copy and the handler carried over. The dispatch calls it once per request, so two requests in flight never share a bound value. |

[every contract](doc.md)
