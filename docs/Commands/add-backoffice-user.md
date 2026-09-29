# `add-backoffice-user`

Creates an initial backoffice user in the database

```bash
teste add-backoffice-user --username <username> --email <email> --password <password> --secret <secret> [--help]
```

| Flag | Type | Default | Description | From |
| --- | --- | --- | --- | --- |
| `--username` | string, required |  | the username for the backoffice user | — |
| `--email` | string, required |  | the email for the backoffice user | — |
| `--password` | string, required |  | the plain-text password (hashed with secret before storing) | — |
| `--secret` | string, required |  | the secret prepended to the password before SHA-256 hashing | — |
| `--help`, `-h` | boolean |  | Print the help of the command this command line is for | [help-flag](help-flag.md) |

| Runs in front of it | When |
| --- | --- |
| [`help-flag`](help-flag.md) | always |

Backoffice · [every command](doc.md) · [CommandYaml](../CommandYaml/doc.md)
