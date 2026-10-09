# `Backup`

`sandbox/internal/databases/backup/`, keys under `data/backup`. Build one with
`backup.New(sandbox)` — it touches no key, so building one is free.

## `blob`

| Field | Type | Required | Target |
| --- | --- | --- | --- |
| `sha` | `key` | yes |  |
| `value` | `bytes` | yes |  |

## Methods

| Method | What it does |
| --- | --- |
| `AddBlob(props BlobInput) (BlobRecord, error)` | inserts one blob record |
| `FindBlobById(id int64) (BlobRecord, bool)` | reads one blob record by its permanent id |
| `FindBlobBySha(value string) (BlobRecord, bool)` | reads one blob record by its indexed sha |
| `ListBlobs(filter BlobFilter) ([]BlobRecord, error)` | reads every blob record the filter keeps |
| `ListBlobsPage(offset int, limit int) ([]BlobRecord, error)` | reads up to limit blob records after the first offset, every one past them when limit is 0 |
| `CountBlob() (int, error)` | is how many blob records are live |
| `SetBlobSha(id int64, value string) error` | writes a new sha on one blob record |
| `SetBlobValue(id int64, value []byte) error` | writes a new value on one blob record |
| `RemoveBlob(id int64) error` | deletes one blob record and everything nested under it |

A query this page does not list goes in `methods_custom.go`, hand-written beside these and
rewritten by no build.

[every database](doc.md)
