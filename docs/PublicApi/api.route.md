# `sandbox/api/route.go`

| Constant | Value | Description |
| --- | --- | --- |
| `StringPath` | `opinatedagnosserver.StringPath` | StringPath takes any slice, bound as a string. |
| `IntegerPath` | `opinatedagnosserver.IntegerPath` | IntegerPath takes one segment reading as a whole number, bound as an int. |
| `NumberPath` | `opinatedagnosserver.NumberPath` | NumberPath takes one segment reading as a number, bound as a float64. |
| `UuidPath` | `opinatedagnosserver.UuidPath` | UuidPath takes one segment reading as a canonical uuid, bound as a string. |
| `HeaderParam` | `opinatedagnosserver.HeaderParam` | HeaderParam reads a request header, matched without regard to case. |
| `QueryParam` | `opinatedagnosserver.QueryParam` | QueryParam reads a query-string parameter. |
| `CookieParam` | `opinatedagnosserver.CookieParam` | CookieParam reads a request cookie. |
| `StringType` | `opinatedagnosserver.StringType` | StringType is bound as a string. |
| `NumberType` | `opinatedagnosserver.NumberType` | NumberType is bound as a float64. |
| `BooleanType` | `opinatedagnosserver.BooleanType` | BooleanType is bound as a bool: true/1 or false/0. |
| `DateTimeType` | `opinatedagnosserver.DateTimeType` | DateTimeType is bound as a string that has to read as RFC 3339. |
| `StringArrayType` | `opinatedagnosserver.StringArrayType` | StringArrayType is bound as a []string. |
| `IntegerType` | `opinatedagnosserver.IntegerType` | IntegerType is bound as an int. |
| `IntegerArrayType` | `opinatedagnosserver.IntegerArrayType` | IntegerArrayType is bound as a []int. |
| `AnyMethod` | `opinatedagnosserver.AnyMethod` | AnyMethod is the one entry of AcceptMethods that accepts every http method. |

## `PathType`

PathType is what one segment a Path reads has to convert to.

`type PathType = opinatedagnosserver.PathType`

## `Path`

Path is one entry of `paths` in route.yaml.

`type Path = opinatedagnosserver.Path`

## `ParameterFont`

ParameterFont is one place of the request a Parameter is read from.

`type ParameterFont = opinatedagnosserver.ParameterFont`

## `ParameterType`

ParameterType is the type a Parameter is converted to before it reaches Entries.

`type ParameterType = opinatedagnosserver.ParameterType`

## `Parameter`

Parameter is one entry of `parameters` in route.yaml.

`type Parameter = opinatedagnosserver.Parameter`

## `RouteBody`

RouteBody is the request body a route declares — the parsed form of `body:` in its route.yaml.

`type RouteBody = opinatedagnosserver.RouteBody`

## `RouteFailure`

RouteFailure is one way a request did not get answered; an InternalPureHandler refuses a request by returning one, built by Deps.OpinatedAgnosServer.Fail.

`type RouteFailure = opinatedagnosserver.RouteFailure`

## `Route`

Route is one http route of the project: the whole of what its route.yaml declares, plus the handler behind it.

`type Route = opinatedagnosserver.Route`

[every contract](doc.md)
