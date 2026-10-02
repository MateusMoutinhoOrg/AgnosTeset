package routeprops

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
)

// RouteProps is what one request carries from route to route of its chain: the
// dispatch builds one, empty, per request and hands the same one to every
// InternalPureHandler that runs for it, as its first argument. A middleware
// sets what it learned — the user it authenticated — and every route after it
// reads it, typed.
//
// Written once by `agnos build` and the project's from then on:
// declare here whatever the routes of this project hand each other.
type RouteProps struct {
	// User is the backoffice user the admin/autentication middleware
	// authenticated from the session cookie, or the one the
	// api/admin/api-autentication middleware authenticated from an API
	// token; nil when none.
	User *maindatabase.BackofficeuserItem
	// Session is the session of User the session cookie names, nil when
	// none — and always nil on /api/admin, which reads no cookie.
	Session *maindatabase.SessionsItem
	// ApiToken is the API token of User the Authorization header carried
	// on /api/admin, nil when none — and always nil on /admin.
	ApiToken *maindatabase.ApitokenItem
}
