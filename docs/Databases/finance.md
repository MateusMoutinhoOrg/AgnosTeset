# `Finance`

`sandbox/internal/databases/finance/`, keys under `finance`. Build one with
`finance.New(sandbox)` — it touches no key, so building one is free.

## `transactions`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `amount` | `float` |  |  |
| `description` | `string` |  |  |
| `type` | `string` |  |  |
| `date` | `string` |  |  |

## Methods

| Method | What it does |
| --- | --- |
| `AddTransactions(props TransactionsNew) (TransactionsItem, error)` | inserts one transactions record |
| `FindTransactionsById(id int64) (TransactionsItem, bool)` | reads one transactions record by its permanent id |
| `ListTransactions(filtrage TransactionsFiltrage) ([]TransactionsItem, error)` | reads every transactions record the filtrage keeps |
| `PageTransactions(position int, chunk int) ([]TransactionsItem, error)` | reads one page of transactions records, counted from 1 |
| `CountTransactions() (int, error)` | is how many transactions records are live |
| `UpdateTransactionsAmount(id int64, value float64) error` | writes a new amount on one transactions record |
| `UpdateTransactionsDescription(id int64, value string) error` | writes a new description on one transactions record |
| `UpdateTransactionsType(id int64, value string) error` | writes a new type on one transactions record |
| `UpdateTransactionsDate(id int64, value string) error` | writes a new date on one transactions record |
| `RemoveTransactions(id int64) error` | deletes one transactions record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
