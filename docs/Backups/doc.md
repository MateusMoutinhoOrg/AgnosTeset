# Backups

Snapshots of every database of `data/`, taken, downloaded, uploaded, restored and deleted from the
[Backoffice](../Backoffice/doc.md) by a root — or built by hand over the JSON api, one file per
request, for a backup too long for one job or one archive. Kept in the `backup` database
([Databases](../Databases/backup.md)), under `./data/backup`, gitignored.

## Model

| Table | Holds |
|---|---|
| `blob` | one file content, `sha` its SHA-256 (indexed), `value` its bytes; stored once whatever the number of snapshots holding it |
| `snapshot` | `name` (indexed), `data` (Unix seconds it was taken at), `status`, and `content`: one `(path, sha)` per file, `path` relative to `data/` |

| `status` | Means |
|---|---|
| `creating` | its files are being stored |
| `ready` | holds every file; the only status that downloads or restores |
| `failed` | its job failed or the server stopped during it: `start-server` marks those in the background, create and restore answering `busy` until it is done (up to a minute after a crash, while the store's write lease expires) |
| `open` | built by hand: takes files until it is closed, then `ready`; a server that stops leaves it `open` |

## Where it lives

| Path | What |
|---|---|
| `sandbox/internal/snapshots/` | take (`StartCreate`), restore (`StartRestore`), `Export`, `Import`, `Remove`, `StartOptimize`, the one-job lock |
| `sandbox/internal/snapshots/manual.go` | by hand: `CreateOpen`, `AddFile`, `AddBlob`, `AddReference`, `Close`, `ListFiles`, `ReadFile` |
| `sandbox/internal/databases/backup/methods_custom.go` | `ListBlobShas`: every blob's id and sha, without reading its value; `HasBlob`: whether a sha is stored, the same way; `RemoveSnapshotContent`: one `(path, sha)` of a snapshot |
| `sandbox/internal/server/backoffice/backofficesnapshots/` | list path, notices, the zip and the file download writers |
| `sandbox/internal/server/backoffice/backofficerender/backup_snapshots.go`, `assets/backoffice/backup_snapshots.html` | the list page |
| `sandbox/internal/server/backoffice/backofficeapi/snapshots.go` | the JSON documents |
| `sandbox/internal/routes/backoffice/root/*backup_snapshot*`, `sandbox/internal/routes/api/backoffice/root/api_*backup_*` | the routes |
| `sandbox/deps/archivedeps/`, `adapters/impls/ziparchive/` | zip and unzip, over `archive/zip` |
| `sandbox/deps/iodeps/`, `adapters/impls/osio/` | the filesystem `data/` is read and written through |

## Routes

Root only: `backoffice-root-guard` and `backoffice-api-root-guard` answer `403` to anyone else.

| Page | Does |
|---|---|
| GET `/admin/root/list-backup-snapshots?prefix=` | every snapshot, newest first, or the ones whose name starts with `prefix`; reloads every 5 s while a job runs or one is `creating` |
| POST `/admin/root/create-backup-snapshot` | form field `name` (optional): starts one, `303` to the list at once |
| POST `/admin/root/restore-backup-snapshot/{id}` | starts restoring one, `303` to the list at once |
| GET `/admin/root/download-backup-snapshot/{id}` | the zip, `attachment; filename="<name>.zip"` |
| POST `/admin/root/upload-backup-snapshot` | the zip as the whole body (`application/zip`), sent by `backoffice.js`: `201`, `400` or `409` |
| POST `/admin/root/remove-backup-snapshot/{id}` | deletes one, `303` to the list |
| POST `/admin/root/optimize-backup-storage` | starts removing every blob no snapshot holds, `303` to the list at once |

| JSON (`Authorization: Bearer <token>`) | Answers |
|---|---|
| GET `/api/admin/root/list-backup-snapshots?prefix=` | `200 {snapshots: [{id, name, data, status}], busy}`, narrowed to the names starting with `prefix` when given |
| POST `/api/admin/root/create-backup-snapshot` `{name?}` | `202 {status: "creating", snapshot}`; `400` invalid name, `409` name taken or while a job runs |
| POST `/api/admin/root/restore-backup-snapshot` `{id}` | `202 {status: "restoring"}`; `404`, `400` not ready, `409` |
| GET `/api/admin/root/download-backup-snapshot/{id}` | the zip; `404`, `400` not ready |
| POST `/api/admin/root/upload-backup-snapshot` | `201 {snapshot}`; `400` with why, `409` |
| POST `/api/admin/root/remove-backup-snapshot` `{id}` | `200 {status: "ok"}`; `404`, `409` |
| POST `/api/admin/root/optimize-backup-storage` | `202 {status: "optimizing"}`; `409` |

By hand — `path` is relative to `data/`, as in `backofficedb/backoffice-user/1/values/username`:

| JSON (`Authorization: Bearer <token>`) | Answers |
|---|---|
| POST `/api/admin/root/create-empty-backup-snapshot` `{name?}` | `201 {snapshot}`, `open`; `400` invalid name, `409` name taken or while a job runs |
| POST `/api/admin/root/add-backup-snapshot-file/{id}/{path}`, the file as the body (`application/octet-stream`) | `201 {file: {path, sha}}`; `400` invalid path or not `open`, `404`, `409` |
| POST `/api/admin/root/add-backup-blob`, a content as the body (`application/octet-stream`) | `201 {blob: {sha, size}}`; `409` |
| POST `/api/admin/root/add-backup-snapshot-reference` `{id, path, sha}` | `201 {file: {path, sha}}`; `400` invalid path or sha or not `open`, `404` snapshot or sha not stored (`field: sha`), `409` |
| POST `/api/admin/root/close-backup-snapshot` `{id}` | `200 {snapshot}`, `ready`; `400` not `open`, holding no file, or naming a content not stored (its path in the message), `404`, `409` |
| GET `/api/admin/root/list-backup-snapshot-files/{id}?prefix=` | `200 {snapshot, files: [{path, sha}]}`, in path order, narrowed to the paths starting with `prefix` when given; any status; `404` |
| GET `/api/admin/root/download-backup-snapshot-file/{id}/{path}` | the file, `attachment; filename="<last segment>"`; any status; `404` snapshot, path or content |

Every request with its curl is in [Routes](../Routes/doc.md).

## Rules

- Names: a create may name its snapshot; the name is trimmed and must match
  `^[A-Za-z0-9][A-Za-z0-9._-]{0,99}$` (`400`, or the `invalid-name` notice), and one another snapshot
  holds is refused (`409`, or `name-taken`), never renamed. Without one it is `snapshot-<ts>`; a
  restore's safety copy is `pre-restore-<ts>`, an upload without a usable name `upload-<ts>`, `<ts>`
  being `YYYYMMDD-hhmmss`.
- Search: `prefix` keeps the snapshots whose name starts with it, case-sensitive, its surrounding
  spaces trimmed; empty keeps every one. Naming snapshots by a scheme (`daily-…`, `before-…`) is
  what makes a prefix find a family of them.
- A snapshot is every file of every directory of `data/` but `data/backup`; the store's `.keep-tmp-*`
  and `*.keeplock` are left out.
- Create, restore and optimize answer at once and run on a goroutine after the route returns. Every
  job — those three, an upload, a delete and every write by hand — runs one at a time: a second one
  is refused (`busy`, `409`). An upload is checked before, so a bad archive is a `400` even then.
- Delete removes the snapshot and its `(path, sha)` list, never a blob: another snapshot may hold
  the same content. Optimize first runs the store's `Repair` over `blob` and `snapshot` (holes,
  indexes, what a non-live record left), then removes every blob no snapshot names — a `failed`
  one's included — and writes how many it removed on stderr. A nested record a killed process was
  inserting under a snapshot is not reached by `Repair`: its few bytes of `(path, sha)` stay.
- Restore: first a `pre-restore-<ts>` snapshot of the current data, then every blob of the target is
  checked, and only then is every directory of `data/` but `data/backup` replaced. Sessions and
  API tokens are restored too, so the root restoring may be signed out.
- Archive: `snapshot.json` (`{name, data, files}`) + `data/<db>/...`; unzipped at the project root it
  restores by hand.
- Upload: at most 256 MiB, 1 GiB unpacked; every entry under `data/<db>/`, never `data/backup`, an
  empty, `.` or `..` segment, a `\` or a `:`, nor the same path twice. The name comes from
  `snapshot.json` when it matches `^[A-Za-z0-9][A-Za-z0-9._-]{0,99}$`, else `upload-<ts>`; a name
  taken gets `-2`, `-3`….
- By hand: `create-empty-backup-snapshot` records an `open` snapshot (`manual-<ts>` without a
  name); `add-backup-snapshot-file` stores one file in it, or `add-backup-blob` stores a content and
  `add-backup-snapshot-reference` names it at a path; `close-backup-snapshot` turns it `ready` once
  every sha it names is stored. Only an `open` snapshot takes files, and only a `ready` one
  downloads whole or restores; list and download one file read any status, so an interrupted
  backup lists what it already holds and resumes from there. Every write is a job (`409` while
  another runs); list and download one file are not.
- A path added twice replaces the file: both `(path, sha)` stay until the close, the one added last
  is the one listed and downloaded, and the close removes the others.
- A close refuses a snapshot holding no file — restoring it would empty every database — and one
  naming a sha not stored, which stays `open`.
- A sha is the 64 hex digits of the SHA-256, read in lowercase. A blob no snapshot names is removed
  by the next optimize: one sent to `add-backup-blob` and not referenced yet answers `404` on its
  reference, sent again. That `404` is also how to skip a content already stored: reference it
  first, send it only when it is missing.
- One file is at most 256 MiB; paths follow the upload's rules.
- A large snapshot may need `start-server --read-timeout-ms` / `--write-timeout-ms` above 10 s.
- Not done: scheduled snapshots, a cli command.
