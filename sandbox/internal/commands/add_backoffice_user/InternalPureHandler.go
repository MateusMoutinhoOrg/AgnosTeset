package add_backoffice_user

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/generated/cliio"
)

// InternalPureHandler answers `add-backoffice-user`. It prepends the secret
// to the password, hashes the result with SHA-256, and inserts one
// backoffice user record into the database.
func InternalPureHandler(sandbox *api.Sandbox, props *api.CommandProps, entries *Entries, response *api.CommandResponse) error {
	// Hash: SHA-256(secret + password)
	combined := entries.Secret + entries.Password
	passwordSha := sandbox.Deps.Hashdeps.Sha256Hex([]byte(combined))

	db := maindatabase.New(sandbox)
	_, err := db.AddBackofficeuser(maindatabase.BackofficeuserNew{
		Username:    entries.Username,
		Email:       entries.Email,
		Passwordsha: passwordSha,
	})
	if err != nil {
		return cliio.Fail(sandbox, api.ExitFailure, "", "failed to add backoffice user: "+err.Error())
	}

	response.Printf("backoffice user %s <%s> created successfully\n", entries.Username, entries.Email)
	return nil
}

