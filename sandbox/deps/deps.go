package deps

import (
	OpinatedAgnosCli "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinatedAgnosCli"
	OpinatedAgnosDatabase "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinatedAgnosDatabase"
	OpinatedAgnosFront "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinatedAgnosFront"
	OpinatedAgnosServer "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/OpinatedAgnosServer"
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/argvdeps"
	database "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/database"
	embeddeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/embeddeps"
	envdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/envdeps"
	hashdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/hashdeps"
	jwtdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/jwtdeps"
	passworddeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/passworddeps"
	randdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/randdeps"
	ratelimitdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/ratelimitdeps"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	signaldeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/signaldeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/sortdeps"
	std "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/std"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/stringsdeps"
	timedeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/timedeps"
)

// Deps is every capability the sandbox needs from the outside world, one field
// per sub-contract directory of sandbox/deps/. An adapter fills the fields; the
// sandbox only calls them, which is what keeps it free of OS packages.
type Deps struct {
	OpinatedAgnosCli      OpinatedAgnosCli.Sandbox
	OpinatedAgnosDatabase OpinatedAgnosDatabase.Sandbox
	OpinatedAgnosFront    OpinatedAgnosFront.Sandbox
	OpinatedAgnosServer   OpinatedAgnosServer.Sandbox
	Argvdeps              argvdeps.Sandbox
	Database              database.Sandbox
	Embeddeps             embeddeps.Sandbox
	Envdeps               envdeps.Sandbox
	Hashdeps              hashdeps.Sandbox
	Jwtdeps               jwtdeps.Sandbox
	Passworddeps          passworddeps.Sandbox
	Randdeps              randdeps.Sandbox
	Ratelimitdeps         ratelimitdeps.Sandbox
	Serializables         serializables.Sandbox
	Serverdeps            serverdeps.Sandbox
	Signaldeps            signaldeps.Sandbox
	Sortdeps              sortdeps.Sandbox
	Std                   std.Sandbox
	Stringsdeps           stringsdeps.Sandbox
	Timedeps              timedeps.Sandbox
}
