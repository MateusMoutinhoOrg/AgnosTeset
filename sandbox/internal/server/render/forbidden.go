package render

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
)

// ForbiddenPage is what templates/forbidden.html is rendered with.
type ForbiddenPage struct {
	Viewer Viewer
}

// Forbidden answers, under a 403, the page telling user that what they asked
// for needs the root role.
func Forbidden(sandbox *api.Sandbox, response *serverdeps.Response, user *maindatabase.BackofficeuserItem) error {
	return Html(sandbox, response, api.StatusForbidden, "templates/forbidden.html", ForbiddenPage{Viewer: viewerOf(sandbox, user)})
}
