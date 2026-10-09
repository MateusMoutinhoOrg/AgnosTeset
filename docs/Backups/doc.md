# Backups

Snapshots of every database of `data/`, taken, downloaded, uploaded and restored from the
[Backoffice](../Backoffice/doc.md) by a root. Kept in the `backup` database
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

## Where it lives

| Path | What |
|---|---|
| `sandbox/internal/snapshots/` | take (`StartCreate`), restore (`StartRestore`), `Export`, `Import`, the one-job lock |
| `sandbox/internal/server/backoffice/backofficesnapshots/` | list path, notices, the zip download writer |
| `sandbox/internal/server/backoffice/backofficerender/backup_snapshots.go`, `assets/backoffice/backup_snapshots.html` | the list page |
| `sandbox/internal/server/backoffice/backofficeapi/snapshots.go` | the JSON documents |
| `sandbox/internal/routes/backoffice/root/*backup_snapshot*`, `sandbox/internal/routes/api/backoffice/root/api_*backup_snapshot*` | the routes |
| `sandbox/deps/archivedeps/`, `adapters/impls/ziparchive/` | zip and unzip, over `archive/zip` |
| `sandbox/deps/iodeps/`, `adapters/impls/osio/` | the filesystem `data/` is read and written through |

## Routes

Root only: `backoffice-root-guard` and `backoffice-api-root-guard` answer `403` to anyone else.

| Page | Does |
|---|---|
| GET `/admin/root/list-backup-snapshots` | every snapshot, newest first; reloads every 5 s while a job runs or one is `creating` |
| POST `/admin/root/create-backup-snapshot` | starts one, `303` to the list at once |
| POST `/admin/root/restore-backup-snapshot/{id}` | starts restoring one, `303` to the list at once |
| GET `/admin/root/download-backup-snapshot/{id}` | the zip, `attachment; filename="<name>.zip"` |
| POST `/admin/root/upload-backup-snapshot` | the zip as the whole body (`application/zip`), sent by `backoffice.js`: `201` or `400` |

| JSON (`Authorization: Bearer <token>`) | Answers |
|---|---|
| GET `/api/admin/root/list-backup-snapshots` | `200 {snapshots: [{id, name, data, status}], busy}` |
| POST `/api/admin/root/create-backup-snapshot` | `202 {status: "creating", snapshot}`; `409` while a job runs |
| POST `/api/admin/root/restore-backup-snapshot` `{id}` | `202 {status: "restoring"}`; `404`, `400` not ready, `409` |
| GET `/api/admin/root/download-backup-snapshot/{id}` | the zip; `404`, `400` not ready |
| POST `/api/admin/root/upload-backup-snapshot` | `201 {snapshot}`; `400` with why |

Every request with its curl is in [Routes](../Routes/doc.md).

## Rules

- A snapshot is every file of every directory of `data/` but `data/backup`; the store's `.keep-tmp-*`
  and `*.keeplock` are left out.
- Create and restore answer at once and run on a goroutine after the route returns, one job at a
  time: a second one is refused (`busy`, `409`). Upload runs in the request and beside a job.
- Restore: first a `pre-restore-<ts>` snapshot of the current data, then every blob of the target is
  checked, and only then is every directory of `data/` but `data/backup` replaced. Sessions and
  API tokens are restored too, so the root restoring may be signed out.
- Archive: `snapshot.json` (`{name, data, files}`) + `data/<db>/...`; unzipped at the project root it
  restores by hand.
- Upload: at most 256 MiB, 1 GiB unpacked; every entry under `data/<db>/`, never `data/backup`, an
  empty, `.` or `..` segment, a `\` or a `:`, nor the same path twice. The name comes from
  `snapshot.json` when it matches `^[A-Za-z0-9][A-Za-z0-9._-]{0,99}$`, else `upload-<ts>`; a name
  taken gets `-2`, `-3`….
- A large snapshot may need `start-server --read-timeout-ms` / `--write-timeout-ms` above 10 s.
- Not done: removing a snapshot (and its orphan blobs), scheduled snapshots, a cli command.
