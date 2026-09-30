package jwtdeps

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	jwtdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/jwtdeps"

	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

// claimsSet is what a token carries on the wire: the registered claims and
// the private `host` claim.
type claimsSet struct {
	jwt.RegisteredClaims
	Host string `json:"host,omitempty"`
}

// sign fills jwtdeps.Sandbox.Sign: the claims become a claimsSet and are
// signed with HMAC-SHA256 over secret.
func sign(claims jwtdeps.Claims, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("jwt: empty secret")
	}
	set := claimsSet{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   claims.Subject,
			IssuedAt:  jwt.NewNumericDate(time.Unix(claims.IssuedAt, 0)),
			ExpiresAt: jwt.NewNumericDate(time.Unix(claims.ExpiresAt, 0)),
		},
		Host: claims.Host,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, set).SignedString([]byte(secret))
}

// parse fills jwtdeps.Sandbox.Parse: only HS256 is accepted, and a token
// without an `exp` claim is refused along with an expired one.
func parse(token string, secret string) (jwtdeps.Claims, error) {
	if secret == "" {
		return jwtdeps.Claims{}, errors.New("jwt: empty secret")
	}
	set := claimsSet{}
	_, err := jwt.ParseWithClaims(token, &set, func(*jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithExpirationRequired())
	if err != nil {
		return jwtdeps.Claims{}, err
	}
	claims := jwtdeps.Claims{Subject: set.Subject, Host: set.Host}
	if set.IssuedAt != nil {
		claims.IssuedAt = set.IssuedAt.Unix()
	}
	if set.ExpiresAt != nil {
		claims.ExpiresAt = set.ExpiresAt.Unix()
	}
	return claims, nil
}

// Bind fills deps.Deps.Jwtdeps with github.com/golang-jwt/jwt/v5.
func Bind(deps *deps.Deps) {
	deps.Jwtdeps = jwtdeps.Sandbox{
		Sign: func(claims jwtdeps.Claims, secret string) (string, error) {
			return sign(claims, secret)
		},
		Parse: func(token string, secret string) (jwtdeps.Claims, error) {
			return parse(token, secret)
		},
	}
}
