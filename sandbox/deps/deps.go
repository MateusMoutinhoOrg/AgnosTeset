package deps

import (
	OpinionatedAgnosCli "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinionatedAgnosCli"
	OpinionatedAgnosDatabase "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinionatedAgnosDatabase"
	OpinionatedAgnosFront "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinionatedAgnosFront"
	OpinionatedAgnosServer "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinionatedAgnosServer"
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/argvdeps"
	databasedeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/databasedeps"
	embeddeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/embeddeps"
	envdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/envdeps"
	hashdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/hashdeps"
	jwtdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/jwtdeps"
	passworddeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/passworddeps"
	randdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/randdeps"
	ratelimitdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/ratelimitdeps"
	serializabledeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serializabledeps"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	signaldeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/signaldeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/sortdeps"
	stddeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/stddeps"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/stringsdeps"
	timedeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/timedeps"
)

// Deps is every capability the sandbox needs from the outside world, one field
// per sub-contract directory of sandbox/deps/. An adapter fills the fields; the
// sandbox only calls them, which is what keeps it free of OS packages.
type Deps struct {
	OpinionatedAgnosCli      OpinionatedAgnosCli.Contract
	OpinionatedAgnosDatabase OpinionatedAgnosDatabase.Contract
	OpinionatedAgnosFront    OpinionatedAgnosFront.Contract
	OpinionatedAgnosServer   OpinionatedAgnosServer.Contract
	ArgvDeps                 argvdeps.Contract
	DatabaseDeps             databasedeps.Sandbox
	EmbedDeps                embeddeps.Contract
	EnvDeps                  envdeps.Contract
	HashDeps                 hashdeps.Contract
	JwtDeps                  jwtdeps.Contract
	PasswordDeps             passworddeps.Contract
	RandDeps                 randdeps.Contract
	RatelimitDeps            ratelimitdeps.Contract
	SerializableDeps         serializabledeps.Contract
	ServerDeps               serverdeps.Contract
	SignalDeps               signaldeps.Contract
	SortDeps                 sortdeps.Contract
	StdDeps                  stddeps.Contract
	StringsDeps              stringsdeps.Contract
	TimeDeps                 timedeps.Contract
}
