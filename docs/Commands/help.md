# `help`

Display help for a command

```bash
teste help [Name…] [--help]
```

When called without arguments, lists every available command grouped by category. When called with a command name, shows detailed usage, arguments, flags, and examples for that command.

| Arg | Type | Default | Description |
| --- | --- | --- | --- |
| `Name` | string, repeatable |  | The command to describe; omit it to list every command |

| Flag | Type | Default | Description | From |
| --- | --- | --- | --- | --- |
| `--help` | boolean |  | Print the help of the command this command line is for | [help-flag](help-flag.md) |

| Runs in front of it | When |
| --- | --- |
| [`help-flag`](help-flag.md) | always |

```bash
teste help
teste help start
```

Info · [every command](doc.md) · [CommandYaml](../CommandYaml/doc.md)
