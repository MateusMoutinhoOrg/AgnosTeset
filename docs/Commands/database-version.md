# `database-version`

Refuse a command line whose databases are at another version

A middleware: it runs on rung 7, in front of every command line matching
`*`, and hands the line on unless it answers.

Runs in front of every command line but migrate-database, help and version, after --database is read, and answers none of them. It refuses one when a database's data is at an older version than its database.yaml declares — migrate-database is what brings it up — at a newer one, written by a newer build, or when a migration of it stopped half way. A database holding no data yet is stamped with the version it is declared at. It checks when the command starts, never again: stop a running server before migrating.

| Runs in front of | When |
| --- | --- |
| [`add-backoffice-user`](add-backoffice-user.md) | always |
| [`migrate-database`](migrate-database.md) | always |
| [`help`](help.md) | always |
| [`version`](version.md) | always |
| [`start-server`](start-server.md) | always |

Middlewares · [every command](doc.md) · [CommandYaml](../CommandYaml/doc.md)
