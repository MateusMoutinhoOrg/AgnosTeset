package backofficeusers

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficeauth"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/backofficetokens"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
)

// ListPath is the page every add, edit and remove sends the browser back to.
const ListPath = "/admin/list-backoffice-users"

// DefaultLimit is how many users a page of the list shows when none is asked.
const DefaultLimit = 20

// MaxLimit is the most users one page of the list shows.
const MaxLimit = 100

// MinPasswordLength is the fewest characters a new password may have.
const MinPasswordLength = 8

// emailPattern is the shape an email has to have: something, an @, a domain
// with a dot. The mailbox itself is never checked.
const emailPattern = `^[^@\s]+@[^@\s]+\.[^@\s]+$`

// The notices an add, an edit or a remove hands the list page through
// ListPath?notice=<code>. Each one is a fixed word the page words itself, so
// nothing the client sent is ever shown back.
const (
	// NoticeAdded follows a user added.
	NoticeAdded = "added"
	// NoticeUpdated follows a user edited.
	NoticeUpdated = "updated"
	// NoticeRemoved follows a user removed.
	NoticeRemoved = "removed"
	// NoticeNotFound follows an edit or a remove of a user that does not exist.
	NoticeNotFound = "not-found"
	// NoticeSelf follows a root trying to remove its own account.
	NoticeSelf = "self"
)

// Query is what the list page is asked for.
type Query struct {
	// Search keeps the users whose username or email holds it, regardless of
	// case; "" keeps every user.
	Search string
	// Role keeps the users of one role, by its backofficeauth.RoleName; ""
	// or a name no role has keeps every role.
	Role string
	// Page is the page to show, counted from 1.
	Page int
	// Limit is how many users a page shows.
	Limit int
}

// Listing is one page of the users a Query keeps.
type Listing struct {
	// Users is the page itself.
	Users []maindatabase.BackofficeuserItem
	// Search is the Query's Search, trimmed.
	Search string
	// Role is the Query's Role, "" when it names no role.
	Role string
	// Total is how many users the Query keeps, across every page.
	Total int
	// Page is the page shown, within 1 and Pages.
	Page int
	// Pages is how many pages Total fills, at least 1.
	Pages int
	// Limit is how many users a page shows, within 1 and MaxLimit.
	Limit int
}

// Fields is what the add and edit forms send for one user.
type Fields struct {
	Username string
	Email    string
	// Password is the plain password; "" on an edit keeps the current one.
	Password string
	Role     int64
}

// ListLocation is ListPath carrying notice for the list page to show.
func ListLocation(sandbox *api.Sandbox, notice string) string {
	return ListPath + "?notice=" + notice
}

// Find is the user with id id, or false when there is none.
func Find(sandbox *api.Sandbox, id int64) (maindatabase.BackofficeuserItem, bool) {
	return maindatabase.New(sandbox).FindBackofficeuserById(id)
}

// List answers the page of users query asks for. The filters run here rather
// than through the generated filtrage: the search matches anywhere in either
// field, and a zero RoleMin/RoleMax turns its filter off, so the root role (0)
// could never be asked for there.
func List(sandbox *api.Sandbox, query Query) (Listing, error) {
	users, err := maindatabase.New(sandbox).ListBackofficeuser(maindatabase.BackofficeuserFiltrage{})
	if err != nil {
		return Listing{}, err
	}

	listing := Listing{Search: sandbox.Deps.Stringsdeps.TrimSpace(query.Search)}
	role, byRole := backofficeauth.ParseRole(sandbox, query.Role)
	if byRole {
		listing.Role = query.Role
	}
	search := sandbox.Deps.Stringsdeps.ToLower(listing.Search)

	kept := []maindatabase.BackofficeuserItem{}
	for _, user := range users {
		if byRole && backofficeauth.Role(user.Role) != role {
			continue
		}
		if search != "" && !holds(sandbox, user.Username, search) && !holds(sandbox, user.Email, search) {
			continue
		}
		kept = append(kept, user)
	}

	listing.Total = len(kept)
	listing.Limit = query.Limit
	if listing.Limit < 1 {
		listing.Limit = DefaultLimit
	}
	listing.Limit = min(listing.Limit, MaxLimit)
	listing.Pages = max(1, (listing.Total+listing.Limit-1)/listing.Limit)
	listing.Page = min(max(query.Page, 1), listing.Pages)

	from := (listing.Page - 1) * listing.Limit
	listing.Users = kept[from:min(from+listing.Limit, listing.Total)]
	return listing, nil
}

// Add inserts the user fields describe. It answers a message for the form when
// fields are refused, and the user added with "" once it is.
func Add(sandbox *api.Sandbox, fields Fields) (maindatabase.BackofficeuserItem, string, error) {
	fields = trimmed(sandbox, fields)
	message, err := validate(sandbox, fields, true, 0)
	if err != nil || message != "" {
		return maindatabase.BackofficeuserItem{}, message, err
	}

	user, err := maindatabase.New(sandbox).AddBackofficeuser(maindatabase.BackofficeuserNew{
		Username:    fields.Username,
		Email:       fields.Email,
		Passwordsha: backofficeauth.PasswordSha(sandbox, fields.Password),
		Role:        fields.Role,
	})
	return user, "", err
}

// Update writes fields over the user with id id, the password only when one
// was given. It answers a message for the form when fields are refused —
// demoting the last root among them — and "" once the user is written. The
// user's sessions stay open, and a new role holds from their next request,
// because the autentication middleware reads the user afresh on each one.
func Update(sandbox *api.Sandbox, id int64, fields Fields) (string, error) {
	fields = trimmed(sandbox, fields)
	message, err := validate(sandbox, fields, false, id)
	if err != nil || message != "" {
		return message, err
	}

	user, ok := Find(sandbox, id)
	if !ok {
		return "That user no longer exists.", nil
	}
	if backofficeauth.Role(user.Role) == backofficeauth.RoleRoot && backofficeauth.Role(fields.Role) != backofficeauth.RoleRoot {
		roots, err := countRoots(sandbox)
		if err != nil {
			return "", err
		}
		if roots <= 1 {
			return "This is the last root user, so it has to stay root.", nil
		}
	}

	db := maindatabase.New(sandbox)
	err = db.UpdateBackofficeuserUsername(id, fields.Username)
	if err != nil {
		return "", err
	}
	err = db.UpdateBackofficeuserEmail(id, fields.Email)
	if err != nil {
		return "", err
	}
	err = db.UpdateBackofficeuserRole(id, fields.Role)
	if err != nil {
		return "", err
	}
	if fields.Password != "" {
		err = db.UpdateBackofficeuserPasswordsha(id, backofficeauth.PasswordSha(sandbox, fields.Password))
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

// Remove deletes the user with id id, and every session and API token of it
// with it, on behalf of actor. It answers the notice the list page shows next:
// a root may not remove its own account, and since actor is a root, another
// root always remains.
func Remove(sandbox *api.Sandbox, actor maindatabase.BackofficeuserItem, id int64) (string, error) {
	if id == actor.Id {
		return NoticeSelf, nil
	}
	_, ok := Find(sandbox, id)
	if !ok {
		return NoticeNotFound, nil
	}
	err := backofficetokens.RemoveOfOwner(sandbox, id)
	if err != nil {
		return "", err
	}
	err = maindatabase.New(sandbox).RemoveBackofficeuser(id)
	if err != nil {
		return "", err
	}
	return NoticeRemoved, nil
}

// holds tells whether text holds search, which is already lower case,
// regardless of case.
func holds(sandbox *api.Sandbox, text string, search string) bool {
	return sandbox.Deps.Stringsdeps.Contains(sandbox.Deps.Stringsdeps.ToLower(text), search)
}

// trimmed is fields with the spaces around the username and the email cut.
func trimmed(sandbox *api.Sandbox, fields Fields) Fields {
	fields.Username = sandbox.Deps.Stringsdeps.TrimSpace(fields.Username)
	fields.Email = sandbox.Deps.Stringsdeps.TrimSpace(fields.Email)
	return fields
}

// validate answers why fields may not be stored, "" when they may. A new user
// (creating) needs a password; an edited one, with id id, keeps its own
// username and email without them counting as taken. Login reads its field as
// a username or an email, so neither may be any other user's username or
// email, regardless of case.
func validate(sandbox *api.Sandbox, fields Fields, creating bool, id int64) (string, error) {
	strings := sandbox.Deps.Stringsdeps
	if fields.Username == "" {
		return "The username is required.", nil
	}
	if strings.ContainsAny(fields.Username, " \t\r\n") {
		return "The username cannot contain spaces.", nil
	}
	email, err := strings.MatchPattern(emailPattern, fields.Email)
	if err != nil {
		return "", err
	}
	if !email {
		return "Enter a valid email address.", nil
	}
	if !backofficeauth.ValidRole(sandbox, backofficeauth.Role(fields.Role)) {
		return "Choose a valid role.", nil
	}
	if creating && fields.Password == "" {
		return "The password is required.", nil
	}
	if fields.Password != "" && len([]rune(fields.Password)) < MinPasswordLength {
		return sandbox.Deps.Std.Sprintf("The password needs at least %d characters.", MinPasswordLength), nil
	}

	users, err := maindatabase.New(sandbox).ListBackofficeuser(maindatabase.BackofficeuserFiltrage{})
	if err != nil {
		return "", err
	}
	username := strings.ToLower(fields.Username)
	address := strings.ToLower(fields.Email)
	for _, user := range users {
		if !creating && user.Id == id {
			continue
		}
		taken := []string{strings.ToLower(user.Username), strings.ToLower(user.Email)}
		for _, login := range taken {
			if login == username {
				return "That username is already in use.", nil
			}
			if login == address {
				return "That email is already in use.", nil
			}
		}
	}
	return "", nil
}

// countRoots is how many users hold the root role.
func countRoots(sandbox *api.Sandbox) (int, error) {
	users, err := maindatabase.New(sandbox).ListBackofficeuser(maindatabase.BackofficeuserFiltrage{})
	if err != nil {
		return 0, err
	}
	roots := 0
	for _, user := range users {
		if backofficeauth.Role(user.Role) == backofficeauth.RoleRoot {
			roots++
		}
	}
	return roots, nil
}
