package standard

import (
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/argvdeps"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serverdeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/sortdeps"
	std "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/std"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/stringsdeps"
	deps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	argvdeps.Bind(&deps)
	serializables.Bind(&deps)
	serverdeps.Bind(&deps)
	sortdeps.Bind(&deps)
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	return deps
}
