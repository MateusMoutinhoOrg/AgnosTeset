# `migrate-database`

Bring the stored data up to the version every database.yaml declares

```bash
testebackoffice migrate-database [Database] [--database <database>] [--help]
```

Runs every step from the version a database's data records to the one its database.yaml declares, on a copy of its folder under agnos-migrating/: the operations each version declares, then its hand-written step. The copy is repaired and checked against the last schema — every key unique, every required field held — and only then swapped in: the data it replaced is kept in agnos-pre-migration/<key-prefix>-v<version>. A step that fails leaves the data in use untouched, and a run cut short is settled by the next one. Stop every process using the data first: a running server is not checked again.

| Arg | Type | Default | Description |
| --- | --- | --- | --- |
| `Database` | string |  | the database to migrate, by the name its database.yaml declares (every database when left out) |

| Flag | Type | Default | Description | From |
| --- | --- | --- | --- | --- |
| `--database` | string | `data` | the folder every database lives under, relative to where the command runs, lower-case letters, digits, - and _ (defaults to data) | [database-dir](database-dir.md) |
| `--help`, `-h` | boolean |  | Print the help of the command this command line is for | [help-flag](help-flag.md) |

| Runs in front of it | When |
| --- | --- |
| [`database-dir`](database-dir.md) | always |
| [`database-version`](database-version.md) | always |
| [`help-flag`](help-flag.md) | always |

```bash
testebackoffice migrate-database
testebackoffice migrate-database app-database
```

Database · [every command](doc.md) · [CommandYaml](../CommandYaml/doc.md)
