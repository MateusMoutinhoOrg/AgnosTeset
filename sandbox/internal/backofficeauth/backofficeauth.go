package backofficeauth

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/jwtdeps"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/internal/databases/maindatabase"
)

// CookieName is the cookie the session token travels in. The autentication
// route declares a cookie parameter under the same key.
const CookieName = "admin_token"

// SessionSeconds is how long a session token is valid after login.
const SessionSeconds = 30 * 60

// Role is what the role column of a backofficeuser stands for.
type Role int64

const (
	// RoleRoot is the root role.
	RoleRoot Role = iota
	// RoleViewer is the viewer role.
	RoleViewer
)

// Roles is every role a backoffice user may hold, in the order they are offered.
func Roles(sandbox *api.Sandbox) []Role {
	return []Role{RoleRoot, RoleViewer}
}

// RoleName is the display name of a role, "unknown" for one Roles does not hold.
func RoleName(sandbox *api.Sandbox, role Role) string {
	switch role {
	case RoleRoot:
		return "root"
	case RoleViewer:
		return "viewer"
	}
	return "unknown"
}

// ParseRole is the role RoleName spells as name, or false when none does.
func ParseRole(sandbox *api.Sandbox, name string) (Role, bool) {
	for _, role := range Roles(sandbox) {
		if RoleName(sandbox, role) == name {
			return role, true
		}
	}
	return RoleRoot, false
}

// ValidRole tells whether role is one of Roles.
func ValidRole(sandbox *api.Sandbox, role Role) bool {
	for _, known := range Roles(sandbox) {
		if known == role {
			return true
		}
	}
	return false
}

// nowSeconds is the current time in seconds since the Unix epoch.
func nowSeconds(sandbox *api.Sandbox) int64 {
	return sandbox.Deps.Std.Now() / 1_000_000_000
}

// PasswordSha is how add-backoffice-user stores a password: SHA-256 of the
// server secret followed by the password, lower-case hex.
func PasswordSha(sandbox *api.Sandbox, password string) string {
	return sandbox.Deps.Hashdeps.Sha256Hex([]byte(sandbox.Config.Secret + password))
}

// FindByLogin looks a backoffice user up by username, then by email.
func FindByLogin(sandbox *api.Sandbox, login string) (maindatabase.BackofficeuserItem, bool, error) {
	db := maindatabase.New(sandbox)
	for _, filtrage := range []maindatabase.BackofficeuserFiltrage{
		{UsernameEquals: login},
		{EmailEquals: login},
	} {
		found, err := db.ListBackofficeuser(filtrage)
		if err != nil {
			return maindatabase.BackofficeuserItem{}, false, err
		}
		if len(found) > 0 {
			return found[0], true, nil
		}
	}
	return maindatabase.BackofficeuserItem{}, false, nil
}

// Authenticate answers the user whose login (username or email) and password
// match, or false when either does not.
func Authenticate(sandbox *api.Sandbox, login string, password string) (maindatabase.BackofficeuserItem, bool, error) {
	user, ok, err := FindByLogin(sandbox, login)
	if err != nil || !ok {
		return user, false, err
	}
	if user.Passwordsha != PasswordSha(sandbox, password) {
		return maindatabase.BackofficeuserItem{}, false, nil
	}
	return user, true, nil
}

// IssueToken opens a session for user, whose request came from the client ip
// ip, and signs its token, valid for SessionSeconds and bound to ip. The session is a sessions record under user, living as long
// as the token; its id travels as the token's `jti`, so the autentication
// middleware can tell whether that one session is still open. The user's
// expired sessions are dropped first, so they never pile up.
func IssueToken(sandbox *api.Sandbox, user maindatabase.BackofficeuserItem, ip string) (string, error) {
	now := nowSeconds(sandbox)
	err := dropExpired(sandbox, user.Id, now)
	if err != nil {
		return "", err
	}

	session, err := maindatabase.New(sandbox).AddBackofficeuserSessions(user.Id, maindatabase.SessionsNew{
		Expiresat: now + SessionSeconds,
	})
	if err != nil {
		return "", err
	}

	return sandbox.Deps.Jwtdeps.Sign(jwtdeps.Claims{
		Id:        sandbox.Deps.Stringsdeps.FormatInt(session.Id, 10),
		Subject:   sandbox.Deps.Stringsdeps.FormatInt(user.Id, 10),
		IssuedAt:  now,
		ExpiresAt: now + SessionSeconds,
		Ip:        ip,
	}, sandbox.Config.Secret)
}

// SessionCookie is the Set-Cookie value carrying token: HttpOnly,
// SameSite=Strict, on every path, expiring with the token.
func SessionCookie(sandbox *api.Sandbox, token string) string {
	return sandbox.Deps.Std.Sprintf("%s=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=Strict", CookieName, token, SessionSeconds)
}

// ClearedCookie is the Set-Cookie value that removes the session cookie.
func ClearedCookie(sandbox *api.Sandbox) string {
	return sandbox.Deps.Std.Sprintf("%s=; Path=/; Max-Age=0; HttpOnly; SameSite=Strict", CookieName)
}

// BearerToken is the token an Authorization header carries in the Bearer
// scheme, the scheme matched regardless of case, or "" when it carries none.
// The api-autentication middleware reads the API token from there, where the
// autentication one reads the session token from the session cookie.
func BearerToken(sandbox *api.Sandbox, authorization string) string {
	fields := sandbox.Deps.Stringsdeps.Fields(authorization)
	if len(fields) != 2 || sandbox.Deps.Stringsdeps.ToLower(fields[0]) != "bearer" {
		return ""
	}
	return fields[1]
}

// dropExpired deletes every session of the user with id userId that expired
// by now: its token is refused anyway, so the record is only garbage.
func dropExpired(sandbox *api.Sandbox, userId int64, now int64) error {
	db := maindatabase.New(sandbox)
	sessions, err := db.ListBackofficeuserSessions(userId)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		if session.Expiresat <= now {
			err = maindatabase.RemoveBackofficeuserSessions(sandbox, db, userId, session.Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// findSession is the session with id sessionId of the user with id userId.
func findSession(sandbox *api.Sandbox, userId int64, sessionId int64) (maindatabase.SessionsItem, bool) {
	sessions, err := maindatabase.New(sandbox).ListBackofficeuserSessions(userId)
	if err != nil {
		return maindatabase.SessionsItem{}, false
	}
	for _, session := range sessions {
		if session.Id == sessionId {
			return session, true
		}
	}
	return maindatabase.SessionsItem{}, false
}

// SessionOfToken answers the user a session token was issued for and the
// session it names, or false when the token is invalid or expired, was issued
// to another client ip than ip, names a session that was closed by a logout,
// or its user no longer exists.
func SessionOfToken(sandbox *api.Sandbox, token string, ip string) (maindatabase.BackofficeuserItem, maindatabase.SessionsItem, bool) {
	none := func() (maindatabase.BackofficeuserItem, maindatabase.SessionsItem, bool) {
		return maindatabase.BackofficeuserItem{}, maindatabase.SessionsItem{}, false
	}
	if token == "" || ip == "" {
		return none()
	}
	claims, err := sandbox.Deps.Jwtdeps.Parse(token, sandbox.Config.Secret)
	if err != nil || claims.Ip != ip {
		return none()
	}
	userId, err := sandbox.Deps.Stringsdeps.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return none()
	}
	sessionId, err := sandbox.Deps.Stringsdeps.ParseInt(claims.Id, 10, 64)
	if err != nil {
		return none()
	}
	user, ok := maindatabase.New(sandbox).FindBackofficeuserById(userId)
	if !ok {
		return none()
	}
	session, ok := findSession(sandbox, user.Id, sessionId)
	if !ok || session.Expiresat <= nowSeconds(sandbox) {
		return none()
	}
	return user, session, true
}

// Logout closes session of user: its record is deleted, so its token is
// refused from here on, along with every other session of user that expired.
func Logout(sandbox *api.Sandbox, user maindatabase.BackofficeuserItem, session maindatabase.SessionsItem) error {
	err := maindatabase.RemoveBackofficeuserSessions(sandbox, maindatabase.New(sandbox), user.Id, session.Id)
	if err != nil {
		return err
	}
	return dropExpired(sandbox, user.Id, nowSeconds(sandbox))
}
