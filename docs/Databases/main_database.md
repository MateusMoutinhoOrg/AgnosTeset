# `MainDatabase`

`sandbox/internal/databases/main_database/`, keys under `main-database`. Build one with
`main_database.New(sandbox)` — it touches no key, so building one is free.

## `admin-users`

| Field | Type | Required | Target |
| --- | --- | --- | --- |

## Methods

| Method | What it does |
| --- | --- |
| `AddAdminUsers(props AdminUsersNew) (AdminUsersItem, error)` | inserts one admin-users record |
| `FindAdminUsersById(id int64) (AdminUsersItem, bool)` | reads one admin-users record by its permanent id |
| `ListAdminUsers(filtrage AdminUsersFiltrage) ([]AdminUsersItem, error)` | reads every admin-users record the filtrage keeps |
| `PageAdminUsers(position int, chunk int) ([]AdminUsersItem, error)` | reads one page of admin-users records, counted from 1 |
| `CountAdminUsers() (int, error)` | is how many admin-users records are live |
| `RemoveAdminUsers(id int64) error` | deletes one admin-users record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
