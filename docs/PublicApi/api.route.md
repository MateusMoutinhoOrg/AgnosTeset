# `sandbox/api/route.go`

| Constant | Value | Description |
| --- | --- | --- |
| `EqualTrigger` | `iota` | EqualTrigger matches a text that is exactly the trigger's Value. |
| `PrefixTrigger` |  | PrefixTrigger matches a text that is the Value or continues it with a new segment: "/admin" matches "/admin" and "/admin/users", never "/administrator". A Value of "/" matches every path. |
| `TextPrefixTrigger` |  | TextPrefixTrigger matches a text that begins with the Value, whatever follows it: "/admin" matches "/administrator" too. |
| `SuffixTrigger` |  | SuffixTrigger matches a text that ends with the Value. |
| `RegexTrigger` |  | RegexTrigger matches a text the Value, a regular expression, matches. |
| `StringPath` | `iota` | StringPath takes any slice, bound as a string. |
| `IntegerPath` |  | IntegerPath takes one segment reading as a whole number, bound as an int. |
| `NumberPath` |  | NumberPath takes one segment reading as a number, bound as a float64. |
| `UuidPath` |  | UuidPath takes one segment reading as a canonical uuid, bound as a string. |
| `HeaderParam` | `iota` | HeaderParam reads a request header, matched without regard to case. |
| `QueryParam` |  | QueryParam reads a query-string parameter. |
| `CookieParam` |  | CookieParam reads a request cookie. |
| `StringType` | `iota` | StringType is bound as a string. |
| `NumberType` |  | NumberType is bound as a float64. |
| `BooleanType` |  | BooleanType is bound as a bool: true/1 or false/0. |
| `DateTimeType` |  | DateTimeType is bound as a string that has to read as RFC 3339. |
| `StringArrayType` |  | StringArrayType is bound as a []string: every occurrence of a query key, or a header's comma-separated values. |
| `IntegerType` |  | IntegerType is bound as an int. |
| `IntegerArrayType` |  | IntegerArrayType is bound as a []int, read the way a StringArrayType is. |
| `AnyMethod` | `"ANY"` | AnyMethod is the one entry of AcceptMethods that accepts every http method. |

## `TriggerType`

TriggerType is how a Trigger compares the text it is handed.

`type TriggerType int`

## `Trigger`

Trigger is the condition a path slice or a parameter value has to meet for a route to run at all — the parsed form of one `trigger:` of route.yaml. Failing it is a non-match, never a 400: the request is for some other route.

| Field | Type | Description |
| --- | --- | --- |
| `Exist` | `bool` | Exist tells a declared trigger from none at all; an entry with none matches whatever the request brought. |
| `Type` | `TriggerType` | Type is how Value is compared. |
| `Value` | `string` | Value is what the text is compared against. |
| `Negate` | `bool` | Negate inverts the comparison: the trigger holds when the text does not match. |
| `IgnoreCase` | `bool` | IgnoreCase compares without regard to case. |

## `PathType`

PathType is what one segment a Path reads has to convert to. A segment that will not is a non-match: the url is for some other route.

`type PathType int`

## `Path`

Path is one entry of `paths` in route.yaml: the slice of request segments from Start to End, both inclusive, End -1 standing for the last segment. The slice reads as "/" followed by its segments joined by "/", and is bound to the Entries field tagged with its Id.

| Field | Type | Description |
| --- | --- | --- |
| `Id` | `string` | Id is the Entries field the slice is bound to. |
| `Start` | `int` | Start is the index of the first segment of the slice. |
| `End` | `int` | End is the index of the last segment of the slice, -1 for the last segment of the request. |
| `Type` | `PathType` | Type is what the slice converts to; anything but StringPath reads one segment alone. |
| `Description` | `string` | Description is the one-line help text. |
| `Trigger` | `Trigger` | Trigger is what the slice has to match for the route to run. |

## `ParameterFont`

ParameterFont is one place of the request a Parameter is read from.

`type ParameterFont int`

## `ParameterType`

ParameterType is the type a Parameter is converted to before it reaches Entries.

`type ParameterType int`

## `Parameter`

Parameter is one entry of `parameters` in route.yaml: one value read off the request under Key, from the first of Fonts that carries it, and bound to the Entries field tagged with its Id.

| Field | Type | Description |
| --- | --- | --- |
| `Id` | `string` | Id is the Entries field the value is bound to. |
| `Key` | `string` | Key is the query key or the header name the value is read under. |
| `Fonts` | `[]ParameterFont` | Fonts are the places the value is read from, in order: the first one that brings a value wins. |
| `Required` | `bool` | Required reports that the request is answered 400 without it. |
| `Type` | `ParameterType` | Type is what the value is converted to. |
| `Default` | `string` | Default is the value bound when the request brings none, spelled as route.yaml writes it; HasDefault tells an empty default from none. |
| `HasDefault` | `bool` |  |
| `Trigger` | `Trigger` | Trigger is what the value has to match for the route to run. |
| `Description` | `string` | Description is the one-line help text. |

## `RouteBody`

RouteBody is the request body a route declares — the parsed form of `body:` in its route.yaml. The dispatch enforces ContentType and MaxBytes before the handler runs; reading the body itself is the handler's to ask for, through the ReadBody its own package generates.

| Field | Type | Description |
| --- | --- | --- |
| `Type` | `string` | Type is the declared type: "none", "raw", "text", "json" or "form". |
| `Required` | `bool` | Required reports that an absent or empty body is rejected. |
| `MaxBytes` | `int` | MaxBytes is the largest body the route reads; a longer one is 413. |
| `ContentType` | `string` | ContentType is the media type the route accepts; a divergent one is 415. It is empty when the route accepts any. |
| `Schema` | `string` | Schema is the declared json-schema in canonical form, "" when the route declares none. |

## `RouteFailure`

RouteFailure is one way a request did not get answered: what the dispatch would have written, and why. Every failure the server layer raises — a header that will not bind, a body the schema rejected, a handler that returned an error, a path nothing matched — arrives at one of the project's own Handle* files as this, read off the Failure of the route it is handed.

| Field | Type | Description |
| --- | --- | --- |
| `Status` | `int` | Status is the http status the failure carries: one of the Status* constants of server.go. |
| `Field` | `string` | Field is the header, query parameter, path segment or json path that failed, "" when the failure names no single value. |
| `Message` | `string` | Message is the one-line reason, written for whoever called the route. |
| `Cause` | `string` | Cause is what went wrong underneath — an error's text, or the value a handler panicked with — "" when there is nothing below the message. It is for the log, not for the caller. |

## `Route`

Route is one http route of the project, as the sandbox offers it: the whole of what its route.yaml declares, plus the handler behind it. Server.Routes holds one per sandbox/internal/routeslist/<name>/, each built by that package's generated NewRoute, in run order — lowest Priority first. What Server.Routes holds is the declaration alone: nothing is ever bound onto it. The dispatch copies it with BindRoute, puts the request and the response on the copy, and hands the copy to IsActionable and RequestHandler.

| Field | Type | Description |
| --- | --- | --- |
| `Name` | `string` | Name is the package directory of the route, snake_case. |
| `AcceptMethods` | `[]string` | AcceptMethods are the http methods it answers to ("GET", "POST"), or AnyMethod alone for every one. |
| `Priority` | `int` | Priority is the rung this route runs on when several match one request: the dispatch runs them from the lowest upwards and stops at the first handler that sets a status. |
| `ResponseType` | `string` | ResponseType is the Content-Type set on the response before the handler runs; the handler may set another. |
| `Segments` | `int` | Segments is how many segments the request path has to have for the route to run, 0 for any count. |
| `After` | `bool` | After reports a route of the `after` phase: it runs once the chain has answered, whatever answered it, and never answers itself. |
| `Pattern` | `string` | Pattern is its path as it reads in docs and messages. |
| `Category` | `string` | Category groups it on the generated Routes page. |
| `Help` | `string` | Help is the one-line description. |
| `LongDescription` | `string` | LongDescription is the paragraph the Routes page prints. |
| `Examples` | `[]string` | Examples are whole requests the Routes page prints. |
| `Hidden` | `bool` | Hidden keeps it off the Routes page without disabling it. |
| `Paths` | `[]Path` | Paths are the slices of the request path it reads, in declaration order. |
| `Parameters` | `[]Parameter` | Parameters are the headers and query parameters it reads, in declaration order. |
| `Body` | `RouteBody` | Body is the request body declaration. |
| `InternalPurehandler` | `any` | InternalPurehandler is the route package's own InternalPureHandler, closed over the sandbox: a func(route *Route, entries *Entries, response *serverdeps.Response) error whose Entries is that package's generated struct. It is held as any because every route's Entries is a type of its own; RequestHandler builds and fills one through Deps.Reflectdeps and calls it. |
| `Request` | `any` | Request is the http request this copy was bound from and Response the one being written. Both are handed over as any: sandbox/api may name no type of sandbox/deps, so the server layer reads them back through routeio.RequestOf and routeio.ResponseOf. |
| `Response` | `any` |  |
| `Locals` | `map[string]any` | Locals is one request's scratch space, shared by every route of the chain that runs for it: what a middleware stores there, the routes after it read. Read and write it through routeio.SetLocal and routeio.GetLocal. |
| `Failure` | `*RouteFailure` | Failure is why this route is being handed to one of the project's Handle* files, nil on a normal run. It is set by routeio.Fail, which is the one way any part of the server layer raises a failure. |
| `IsActionable` | `func(bound *Route) bool` | IsActionable reports whether one bound copy — its Request set — is for this route: the method is accepted, and every path slice and every parameter declaring a trigger matches it. |
| `MatchesPath` | `func(bound *Route) bool` | MatchesPath reports whether one bound copy's request path is for this route whatever its method — what tells a 405 from a 404. |
| `RequestHandler` | `func(bound *Route) error` | RequestHandler binds one bound copy's request onto a fresh Entries and runs InternalPurehandler with it. It returns the failure the handler did not answer itself, nil otherwise; what it answered with is the status it wrote, and a handler writing none hands the request to the next route of the chain. |

| Function | Description |
| --- | --- |
| `NewRoute() *Route` | NewRoute returns an empty Route with every slice open. The generic sandbox/internal/server/route.NewRoute fills the matcher and the handler on top of it, and a route's generated NewRoute its declaration. |
| `BindRoute(route *Route) *Route` | BindRoute copies one declaration into the route a single request runs on: the same declared fields — the slices are read-only and shared — with no request, response or failure yet. The dispatch calls it once per request, so two requests in flight never share a bound value. |

[every contract](doc.md)
