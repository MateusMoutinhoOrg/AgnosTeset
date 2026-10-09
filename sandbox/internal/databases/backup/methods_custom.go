package backup

import (
	api "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
)

// BlobSha is one blob record without its value: its permanent id and its sha.
type BlobSha struct {
	Id  int64
	Sha string
}

// RepairBackup runs the store's Repair over the blob and the snapshot
// collections, the content nested under every snapshot included: what a
// process stopped in the middle of a write left behind — a blob's bytes among
// it — is deleted, and every invariant of both is restored. No generated
// method reaches Repair, so it is written here.
func RepairBackup(sandbox *api.Sandbox, db *Backup) error {
	for _, name := range []string{"blob", "snapshot"} {
		collection, err := sandbox.Deps.OpinionatedAgnosDatabase.Collection(db.database, name)
		if err != nil {
			return err
		}
		if failure := collection.Repair(); failure != nil {
			return sandbox.Deps.OpinionatedAgnosDatabase.Fail(failure)
		}
	}
	return nil
}

// ListBlobShas reads the id and the sha of every blob record, never its
// value. ListBlobs would load every stored content to answer the same, so a
// sweep over every blob — the one that removes the blobs no snapshot holds —
// is written here.
func ListBlobShas(sandbox *api.Sandbox, db *Backup) ([]BlobSha, error) {
	collection, err := sandbox.Deps.OpinionatedAgnosDatabase.Collection(db.database, "blob")
	if err != nil {
		return nil, err
	}
	records, failure := collection.ListAll()
	if failure != nil {
		return nil, sandbox.Deps.OpinionatedAgnosDatabase.Fail(failure)
	}
	listed := []BlobSha{}
	for _, record := range records {
		sha, err := sandbox.Deps.OpinionatedAgnosDatabase.ReadString(record, "sha")
		if err != nil {
			return nil, err
		}
		listed = append(listed, BlobSha{Id: record.ID, Sha: sha})
	}
	return listed, nil
}
