package standard

import (
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/argvdeps"
	embeddeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/embeddeps"
	hashdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/hashdeps"
	keep "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/keep"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serverdeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/sortdeps"
	std "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/std"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/stringsdeps"
	templatedeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/templatedeps"
	deps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	argvdeps.Bind(&deps)
	embeddeps.Bind(&deps)
	hashdeps.Bind(&deps)
	keep.Bind(&deps)
	serializables.Bind(&deps)
	serverdeps.Bind(&deps)
	sortdeps.Bind(&deps)
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	templatedeps.Bind(&deps)
	return deps
}
