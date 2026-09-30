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

// IssueToken signs a session token for user, valid for SessionSeconds.
func IssueToken(sandbox *api.Sandbox, user maindatabase.BackofficeuserItem) (string, error) {
	now := nowSeconds(sandbox)
	return sandbox.Deps.Jwtdeps.Sign(jwtdeps.Claims{
		Subject:   sandbox.Deps.Stringsdeps.FormatInt(user.Id, 10),
		IssuedAt:  now,
		ExpiresAt: now + SessionSeconds,
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

// UserOf is the api.User a stored backoffice user stands for.
func UserOf(sandbox *api.Sandbox, item maindatabase.BackofficeuserItem) api.User {
	return api.User{
		Id:       sandbox.Deps.Stringsdeps.FormatInt(item.Id, 10),
		Email:    item.Email,
		Username: item.Username,
		Role:     api.UserRole(item.Role),
	}
}

// UserOfToken answers the user a session token was issued for, or false when
// the token is invalid, expired, or its user no longer exists.
func UserOfToken(sandbox *api.Sandbox, token string) (api.User, bool) {
	if token == "" {
		return api.User{}, false
	}
	claims, err := sandbox.Deps.Jwtdeps.Parse(token, sandbox.Config.Secret)
	if err != nil {
		return api.User{}, false
	}
	id, err := sandbox.Deps.Stringsdeps.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return api.User{}, false
	}
	item, ok := maindatabase.New(sandbox).FindBackofficeuserById(id)
	if !ok {
		return api.User{}, false
	}
	return UserOf(sandbox, item), true
}
