# `deps.OpinatedAgnosDatabase`

`sandbox/deps/OpinatedAgnosDatabase`

## `Sandbox`

Sandbox is the database lib injected whole as the Deps.OpinatedAgnosDatabase field.

| Field | Type | Description |
| --- | --- | --- |
| `Fail` | `func(failure *database.Error) error` | Fail turns one failure the database reported into an error the sandbox carries. A nil *database.Error is success and answers nil. |
| `Schema` | `func(handle database.DatabaseHandle, name string) (database.SchemaInstance, error)` | Schema resolves one collection of a handle by name. A name the handle does not declare is an error rather than a nil instance. |
| `ReadString` | `func(item database.SchemaItem, field string) (string, error)` | ReadString reads one Key or String field of a record. |
| `ReadInt` | `func(item database.SchemaItem, field string) (int64, error)` | ReadInt reads one Int or Link field of a record. |
| `ReadFloat` | `func(item database.SchemaItem, field string) (float64, error)` | ReadFloat reads one Float field of a record. |
| `TextMatches` | `func(value string, startsWith string, equals string) bool` | TextMatches is the filter a generated <T>Filtrage applies to one text field: an empty needle passes everything, so a zero value turns the filter off. |
| `IntInRange` | `func(value int64, min int64, max int64) bool` | IntInRange is TextMatches for a whole-number field: a zero bound is no bound. |
| `FloatInRange` | `func(value float64, min float64, max float64) bool` | FloatInRange is TextMatches for a floating-point field. |

[every contract](doc.md)
