# `PublicUser`

`sandbox/internal/databases/public_user/`, keys under `public-user`. Build one with
`public_user.New(sandbox)` — it touches no key, so building one is free.

## `users`

| Field | Type | Required | Target |
| --- | --- | --- | --- |

## Methods

| Method | What it does |
| --- | --- |
| `AddUsers(props UsersInput) (UsersRecord, error)` | inserts one users record |
| `FindUsersById(id int64) (UsersRecord, bool)` | reads one users record by its permanent id |
| `ListUserses(filter UsersFilter) ([]UsersRecord, error)` | reads every users record the filter keeps |
| `ListUsersesPage(offset int, limit int) ([]UsersRecord, error)` | reads up to limit users records after the first offset, every one past them when limit is 0 |
| `CountUsers() (int, error)` | is how many users records are live |
| `RemoveUsers(id int64) error` | deletes one users record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
