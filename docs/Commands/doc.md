# Commands

`FinanceApp <command> [flags] [args]`. `FinanceApp help <command>` prints
the same for one command; an empty command line prints the general help and exits 2.

One page per command, each rendered from that command's `entries.yaml`
([EntriesYaml](../EntriesYaml/doc.md)) on each build — open the one you need rather than this
whole page. Hidden commands are not listed. Flags may appear anywhere on the command line;
positionals bind in order after them. A `repeatable` field is given once per value.

## Info

| Command | Does |
| --- | --- |
| [`help`](help.md) | Display help for a command |
| [`version`](version.md) | Print the installed version |

## Server

| Command | Does |
| --- | --- |
| [`start-server`](start-server.md) | Starts the http server |

Output channels and exit codes are in [Rules](../Rules/doc.md#output-channels).
