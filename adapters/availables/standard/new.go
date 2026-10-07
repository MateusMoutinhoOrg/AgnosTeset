package standard

import (
	OpinatedAgnosCli "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/OpinatedAgnosCli"
	OpinatedAgnosDatabase "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/OpinatedAgnosDatabase"
	OpinatedAgnosFront "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/OpinatedAgnosFront"
	OpinatedAgnosServer "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/OpinatedAgnosServer"
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/argvdeps"
	database "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/database"
	embeddeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/embeddeps"
	envdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/envdeps"
	hashdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/hashdeps"
	jwtdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/jwtdeps"
	passworddeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/passworddeps"
	randdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/randdeps"
	ratelimitdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/ratelimitdeps"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serverdeps"
	signaldeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/signaldeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/sortdeps"
	std "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/std"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/stringsdeps"
	timedeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/timedeps"
	deps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	OpinatedAgnosCli.Bind(&deps)
	OpinatedAgnosDatabase.Bind(&deps)
	OpinatedAgnosFront.Bind(&deps)
	OpinatedAgnosServer.Bind(&deps)
	argvdeps.Bind(&deps)
	database.Bind(&deps)
	embeddeps.Bind(&deps)
	envdeps.Bind(&deps)
	hashdeps.Bind(&deps)
	jwtdeps.Bind(&deps)
	passworddeps.Bind(&deps)
	randdeps.Bind(&deps)
	ratelimitdeps.Bind(&deps)
	serializables.Bind(&deps)
	serverdeps.Bind(&deps)
	signaldeps.Bind(&deps)
	sortdeps.Bind(&deps)
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	timedeps.Bind(&deps)
	return deps
}
