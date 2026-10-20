# `deps.OpinionatedAgnosDatabase`

`sandbox/deps/OpinionatedAgnosDatabase`

| Constant | Value | Description |
| --- | --- | --- |
| `OpRenameField` | `"rename-field"` | OpRenameField moves every stored value of Field to the name To. A nested collection (Type object) is moved whole, ids and all. |
| `OpRetypeField` | `"retype-field"` | OpRetypeField converts every stored value of Field from Type to the type To; a value that does not convert fails the migration. |
| `OpRequireField` | `"require-field"` | OpRequireField gives Default to every record holding no value for Field, when Default is set. A record still holding none fails the check that follows the last step. |
| `OpRemoveField` | `"remove-field"` | OpRemoveField deletes every stored value of Field, or every record of the nested collection it names. |
| `OpRemoveTable` | `"remove-table"` | OpRemoveTable deletes every record of Table. |
| `TypeKey` | `"key"` | TypeKey is a unique, indexed text field. |
| `TypeString` | `"string"` | TypeString is a plain text field. |
| `TypeInteger` | `"integer"` | TypeInteger is a whole-number field. |
| `TypeNumber` | `"number"` | TypeNumber is a floating-point field. |
| `TypeBytes` | `"bytes"` | TypeBytes is a binary field. |
| `TypeLink` | `"link"` | TypeLink is the id of a record of another table. |
| `TypeObject` | `"object"` | TypeObject is a collection nested under each record. |
| `VersionTable` | `"agnos-version"` | VersionTable is the table every database records the version of its data in, under its own key-prefix, so the version travels with the data. |
| `MigratingDir` | `"agnos-migrating"` | MigratingDir is the folder a migration works in: the copy each step runs on, swapped in only once every step and check passed. |
| `PreMigrationDir` | `"agnos-pre-migration"` | PreMigrationDir is the folder each migrated database's previous data is kept in, as <key-prefix>-v<version>: undoing a migration is moving it back. |

## `Op`

Op is one operation of a version, as migrations/v<N>.ops.yaml declares it. Parent names the object field the field sits in, empty for a field of the table itself.

| Field | Type | Description |
| --- | --- | --- |
| `Kind` | `string` | Kind is one of the Op* constants. |
| `Table` | `string` | Table is the table the operation reaches. |
| `Parent` | `string` | Parent is the nested collection holding Field, "" for the table's own. |
| `Field` | `string` | Field is the field the operation reaches, "" for OpRemoveTable. |
| `Type` | `string` | Type is the field's type when the operation runs: the old one, for a rename or a retype. |
| `To` | `string` | To is the new name of OpRenameField and the new type of OpRetypeField. |
| `Default` | `string` | Default is the value OpRequireField fills, as text; converted to Type. |

## `Version`

Version is one version of a database: the schema its data has once it reached it, and what brings the data of the version before to it.

| Field | Type | Description |
| --- | --- | --- |
| `Version` | `int` | Version is the number, 1 for the first. |
| `Schemas` | `[]databasedeps.Schema` | Schemas is the schema of the data at this version. |
| `Ops` | `[]Op` | Ops is what brings the data of the version before to this one, run in order; empty on version 1. |
| `Custom` | `func(from databasedeps.Database, to databasedeps.Database) error` | Custom is the hand-written step of this version, run after Ops, with from opened on the schema of the version before and to on this one — the same folder read two ways. Nil when the version has none. |

## `Migrations`

Migrations is one database as the engine reads it, built by its generated.migrations.go on every call so Dir follows --database.

| Field | Type | Description |
| --- | --- | --- |
| `Name` | `string` | Name is the database's own name, for the messages. |
| `Dir` | `string` | Dir is the folder every database lives under (sandbox.Config.DatabaseDir). |
| `KeyPrefix` | `string` | KeyPrefix is the folder inside Dir the database's records live in. |
| `Versions` | `[]Version` | Versions is every version, the first one first; the last is the one the code expects. |

## `MigrateProps`

MigrateProps is everything Migrate and CheckVersions read. The generated code builds it per call, so nothing the lib holds goes stale.

| Field | Type | Description |
| --- | --- | --- |
| `Databases` | `databasedeps.Databases` | Databases opens every database the engine reads or writes. |
| `IoDeps` | `iodeps.Contract` | IoDeps copies, renames and removes the folders of a migration. |
| `Logf` | `func(format string, a ...any) (n int, err error)` | Logf reports each step taken. |
| `Command` | `string` | Command is the command line that migrates, named by CheckVersions' failures ("<project> migrate-database"). |
| `Migrations` | `[]Migrations` | Migrations is every database to check or migrate. |

## `Contract`

Contract is the database lib injected whole as the Deps.OpinionatedAgnosDatabase field.

| Field | Type | Description |
| --- | --- | --- |
| `Fail` | `func(failure *databasedeps.Error) error` | Fail turns one failure the database reported into an error the sandbox carries. A nil *databasedeps.Error is success and answers nil. |
| `Collection` | `func(database databasedeps.Database, name string) (databasedeps.Collection, error)` | Collection resolves one collection of a database by name. A name the database does not declare, or a database Databases.New refused, is an error rather than a zero collection. |
| `ReadString` | `func(record databasedeps.Record, field string) (string, error)` | ReadString reads one Key or String field of a record. |
| `ReadInt` | `func(record databasedeps.Record, field string) (int64, error)` | ReadInt reads one Int or Link field of a record. |
| `ReadFloat` | `func(record databasedeps.Record, field string) (float64, error)` | ReadFloat reads one Float field of a record. |
| `ReadBytes` | `func(record databasedeps.Record, field string) ([]byte, error)` | ReadBytes reads one Bytes field of a record, byte for byte. |
| `TextMatches` | `func(value string, startsWith string, equals string) bool` | TextMatches is the filter a generated <T>Filter applies to one text field: an empty needle passes everything, so a zero value turns the filter off. |
| `BytesMatches` | `func(value []byte, startsWith []byte, equals []byte) bool` | BytesMatches is TextMatches for a binary field, compared byte for byte: an empty needle passes everything. |
| `IntInRange` | `func(value int64, min int64, max int64) bool` | IntInRange is TextMatches for a whole-number field: a zero bound is no bound. |
| `FloatInRange` | `func(value float64, min float64, max float64) bool` | FloatInRange is TextMatches for a floating-point field. |
| `Migrate` | `func(props MigrateProps) error` | Migrate brings the data of every database of props up to its last version. Each one is copied into MigratingDir, every step from the version its data records runs on the copy, the copy is repaired and checked against the last schema, and only then is the live folder moved into PreMigrationDir and the copy moved in its place. A failure before that move leaves the live folder as it was; a crash during it is finished by the next call. |
| `CheckVersions` | `func(props MigrateProps) error` | CheckVersions refuses a database whose data is at another version than its code expects, or that a migration left half done, naming props.Command. A database holding no data at all is stamped with the last version instead; one holding data but no version is at version 1, written before versions existed. |

[every contract](doc.md)
