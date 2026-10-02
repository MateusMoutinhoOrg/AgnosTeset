# `Backofficedb`

`sandbox/internal/databases/backofficedb/`, keys under `backofficedb`. Build one with
`backofficedb.New(sandbox)` — it touches no key, so building one is free.

## `backofficeuser`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `username` | `string` | yes |  |
| `email` | `string` | yes |  |
| `passwordhash` | `string` | yes |  |
| `role` | `int` |  |  |
| `sessions` | `database` |  |  |
| `sessions.expiresat` | `int` | yes |  |

## `apitoken`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `tokensha` | `key` | yes |  |
| `name` | `string` | yes |  |
| `prefix` | `string` | yes |  |
| `ownerid` | `int` |  |  |
| `createdat` | `int` |  |  |
| `expiresat` | `int` |  |  |
| `ips` | `string` |  |  |
| `lastusedat` | `int` |  |  |
| `lastusedip` | `string` |  |  |

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
| `UpdateBackofficeuserPasswordhash(id int64, value string) error` | writes a new passwordhash on one backofficeuser record |
| `UpdateBackofficeuserRole(id int64, value int64) error` | writes a new role on one backofficeuser record |
| `RemoveBackofficeuser(id int64) error` | deletes one backofficeuser record and everything nested under it |
| `AddBackofficeuserSessions(parent_id int64, props SessionsNew) (SessionsItem, error)` | inserts one sessions record under one backofficeuser record |
| `ListBackofficeuserSessions(parent_id int64) ([]SessionsItem, error)` | reads every sessions record of one backofficeuser record |
| `AddApitoken(props ApitokenNew) (ApitokenItem, error)` | inserts one apitoken record |
| `FindApitokenById(id int64) (ApitokenItem, bool)` | reads one apitoken record by its permanent id |
| `FindApitokenByTokensha(value string) (ApitokenItem, bool)` | reads one apitoken record by its indexed tokensha |
| `ListApitoken(filtrage ApitokenFiltrage) ([]ApitokenItem, error)` | reads every apitoken record the filtrage keeps |
| `PageApitoken(position int, chunk int) ([]ApitokenItem, error)` | reads one page of apitoken records, counted from 1 |
| `CountApitoken() (int, error)` | is how many apitoken records are live |
| `UpdateApitokenTokensha(id int64, value string) error` | writes a new tokensha on one apitoken record |
| `UpdateApitokenName(id int64, value string) error` | writes a new name on one apitoken record |
| `UpdateApitokenPrefix(id int64, value string) error` | writes a new prefix on one apitoken record |
| `UpdateApitokenOwnerid(id int64, value int64) error` | writes a new ownerid on one apitoken record |
| `UpdateApitokenCreatedat(id int64, value int64) error` | writes a new createdat on one apitoken record |
| `UpdateApitokenExpiresat(id int64, value int64) error` | writes a new expiresat on one apitoken record |
| `UpdateApitokenIps(id int64, value string) error` | writes a new ips on one apitoken record |
| `UpdateApitokenLastusedat(id int64, value int64) error` | writes a new lastusedat on one apitoken record |
| `UpdateApitokenLastusedip(id int64, value string) error` | writes a new lastusedip on one apitoken record |
| `RemoveApitoken(id int64) error` | deletes one apitoken record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
