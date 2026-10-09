package backofficeapi

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	serializabledeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serializabledeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/backup"
)

// SnapshotJSON is the JSON object a snapshot is answered as: its id, name,
// the Unix instant it was taken at as data, and its status. Its files never
// travel in JSON: they are the archive the download routes answer.
func SnapshotJSON(sandbox *api.Sandbox, snapshot backup.SnapshotRecord) *serializabledeps.SerializableObject {
	object := sandbox.Deps.SerializableDeps.CreateObject()
	object.AddItemToObject("id", snapshot.Id)
	object.AddItemToObject("name", snapshot.Name)
	object.AddItemToObject("data", snapshot.Data)
	object.AddItemToObject("status", snapshot.Status)
	return object
}

// SnapshotResponseJSON is {"snapshot": SnapshotJSON(snapshot)}.
func SnapshotResponseJSON(sandbox *api.Sandbox, snapshot backup.SnapshotRecord) *serializabledeps.SerializableObject {
	document := sandbox.Deps.SerializableDeps.CreateObject()
	document.AddItemToObject("snapshot", SnapshotJSON(sandbox, snapshot))
	return document
}

// SnapshotListJSON is every snapshot, newest first, and whether a snapshot or
// a restore is running.
func SnapshotListJSON(sandbox *api.Sandbox, listed []backup.SnapshotRecord, busy bool) *serializabledeps.SerializableObject {
	items := sandbox.Deps.SerializableDeps.CreateArray()
	for _, snapshot := range listed {
		items.AddItemToArray(SnapshotJSON(sandbox, snapshot))
	}

	document := sandbox.Deps.SerializableDeps.CreateObject()
	document.AddItemToObject("snapshots", items)
	document.AddItemToObject("busy", busy)
	return document
}

// SnapshotCreatingJSON is {"status": "creating", "snapshot": SnapshotJSON(snapshot)}:
// a snapshot recorded and still being filled.
func SnapshotCreatingJSON(sandbox *api.Sandbox, snapshot backup.SnapshotRecord) *serializabledeps.SerializableObject {
	document := sandbox.Deps.SerializableDeps.CreateObject()
	document.AddItemToObject("status", "creating")
	document.AddItemToObject("snapshot", SnapshotJSON(sandbox, snapshot))
	return document
}

// RestoringJSON is {"status": "restoring"}: a restore started and not waited
// for.
func RestoringJSON(sandbox *api.Sandbox) *serializabledeps.SerializableObject {
	document := sandbox.Deps.SerializableDeps.CreateObject()
	document.AddItemToObject("status", "restoring")
	return document
}
