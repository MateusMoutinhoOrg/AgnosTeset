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

// IssueToken signs a session token for user on host, valid for
// SessionSeconds, and makes sure host has a hosts record under user, so the
// autentication middleware finds the mincreation it checks the token against.
func IssueToken(sandbox *api.Sandbox, user maindatabase.BackofficeuserItem, host string) (string, error) {
	hosts, err := hostsOf(sandbox, user.Id, host)
	if err != nil {
		return "", err
	}
	if len(hosts) == 0 {
		_, err = maindatabase.New(sandbox).AddBackofficeuserHosts(user.Id, maindatabase.HostsNew{Host: host})
		if err != nil {
			return "", err
		}
	}

	now := nowSeconds(sandbox)
	return sandbox.Deps.Jwtdeps.Sign(jwtdeps.Claims{
		Subject:   sandbox.Deps.Stringsdeps.FormatInt(user.Id, 10),
		IssuedAt:  now,
		ExpiresAt: now + SessionSeconds,
		Host:      host,
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

// hostsOf is every hosts record of the user with id userId that names host.
func hostsOf(sandbox *api.Sandbox, userId int64, host string) ([]maindatabase.HostsItem, error) {
	all, err := maindatabase.New(sandbox).ListBackofficeuserHosts(userId)
	if err != nil {
		return nil, err
	}
	matching := []maindatabase.HostsItem{}
	for _, item := range all {
		if item.Host == host {
			matching = append(matching, item)
		}
	}
	return matching, nil
}

// UserOfToken answers the user a session token was issued for, or false when
// the token is invalid, expired, issued on another host than host, issued no
// later than the mincreation of host (a logout on it), or its user no longer
// exists.
func UserOfToken(sandbox *api.Sandbox, token string, host string) (maindatabase.BackofficeuserItem, bool) {
	if token == "" || host == "" {
		return maindatabase.BackofficeuserItem{}, false
	}
	claims, err := sandbox.Deps.Jwtdeps.Parse(token, sandbox.Config.Secret)
	if err != nil || claims.Host != host {
		return maindatabase.BackofficeuserItem{}, false
	}
	id, err := sandbox.Deps.Stringsdeps.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return maindatabase.BackofficeuserItem{}, false
	}
	user, ok := maindatabase.New(sandbox).FindBackofficeuserById(id)
	if !ok {
		return maindatabase.BackofficeuserItem{}, false
	}

	hosts, err := hostsOf(sandbox, user.Id, host)
	if err != nil || len(hosts) == 0 {
		return maindatabase.BackofficeuserItem{}, false
	}
	for _, item := range hosts {
		// Strictly after: a token issued in the very second of a logout is
		// one the logout meant to end.
		if claims.IssuedAt <= item.Mincreation {
			return maindatabase.BackofficeuserItem{}, false
		}
	}
	return user, true
}

// Logout ends every session of user on host: the mincreation of host becomes
// now, so every token issued on it until now is refused from here on.
func Logout(sandbox *api.Sandbox, user maindatabase.BackofficeuserItem, host string) error {
	db := maindatabase.New(sandbox)
	now := nowSeconds(sandbox)

	hosts, err := hostsOf(sandbox, user.Id, host)
	if err != nil {
		return err
	}
	if len(hosts) == 0 {
		_, err = db.AddBackofficeuserHosts(user.Id, maindatabase.HostsNew{Host: host, Mincreation: now})
		return err
	}
	for _, item := range hosts {
		err = maindatabase.UpdateBackofficeuserHostsMincreation(sandbox, db, user.Id, item.Id, now)
		if err != nil {
			return err
		}
	}
	return nil
}
