package standard

import (
	OpinionatedAgnosCli "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/OpinionatedAgnosCli"
	OpinionatedAgnosDatabase "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/OpinionatedAgnosDatabase"
	OpinionatedAgnosFront "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/OpinionatedAgnosFront"
	OpinionatedAgnosServer "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/OpinionatedAgnosServer"
	cryptorand "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/cryptorand"
	databasedeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/databasedeps"
	goembed "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/goembed"
	golangjwt "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/golangjwt"
	memoryratelimit "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/memoryratelimit"
	nethttpserver "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/nethttpserver"
	osenv "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/osenv"
	ossignal "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/ossignal"
	osstd "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/osstd"
	pbkdf2password "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/pbkdf2password"
	sha256hash "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/sha256hash"
	stdargv "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/stdargv"
	stdserializable "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/stdserializable"
	stdsort "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/stdsort"
	stdstrings "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/stdstrings"
	stdtime "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/impls/stdtime"
	deps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	OpinionatedAgnosCli.Bind(&deps)
	OpinionatedAgnosDatabase.Bind(&deps)
	OpinionatedAgnosFront.Bind(&deps)
	OpinionatedAgnosServer.Bind(&deps)
	cryptorand.Bind(&deps)
	databasedeps.Bind(&deps)
	goembed.Bind(&deps)
	golangjwt.Bind(&deps)
	memoryratelimit.Bind(&deps)
	nethttpserver.Bind(&deps)
	osenv.Bind(&deps)
	ossignal.Bind(&deps)
	osstd.Bind(&deps)
	pbkdf2password.Bind(&deps)
	sha256hash.Bind(&deps)
	stdargv.Bind(&deps)
	stdserializable.Bind(&deps)
	stdsort.Bind(&deps)
	stdstrings.Bind(&deps)
	stdtime.Bind(&deps)
	return deps
}
