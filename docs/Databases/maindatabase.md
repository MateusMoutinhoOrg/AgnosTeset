# `Maindatabase`

`sandbox/internal/databases/maindatabase/`, keys under `maindatabase`. Build one with
`maindatabase.New(sandbox)` — it touches no key, so building one is free.

## `backofficeuser`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `username` | `string` | yes |  |
| `email` | `string` | yes |  |
| `passwordsha` | `string` | yes |  |
| `role` | `int` |  |  |

## Methods

| Method | What it does |
| --- | --- |
| `AddBackofficeuser(props BackofficeuserNew) (BackofficeuserItem, error)` | inserts one backofficeuser record |
| `FindBackofficeuserById(id int64) (BackofficeuserItem, bool)` | reads one backofficeuser record by its permanent id |
| `ListBackofficeuser(filtrage BackofficeuserFiltrage) ([]BackofficeuserItem, error)` | reads every backofficeuser record the filtrage keeps |
| `PageBackofficeuser(position int, chunk int) ([]BackofficeuserItem, error)` | reads one page of backofficeuser records, counted from 1 |
| `CountBackofficeuser() (int, error)` | is how many backofficeuser records are live |
| `UpdateBackofficeuserUsername(id int64, value string) error` | writes a new username on one backofficeuser record |
| `UpdateBackofficeuserEmail(id int64, value string) error` | writes a new email on one backofficeuser record |
| `UpdateBackofficeuserPasswordsha(id int64, value string) error` | writes a new passwordsha on one backofficeuser record |
| `UpdateBackofficeuserRole(id int64, value int64) error` | writes a new role on one backofficeuser record |
| `RemoveBackofficeuser(id int64) error` | deletes one backofficeuser record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
