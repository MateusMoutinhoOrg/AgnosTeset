package standard

import (
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/argvdeps"
	embeddeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/embeddeps"
	hashdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/hashdeps"
	reflectdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/reflectdeps"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/serverdeps"
	signaldeps "github.com/MateusMoutinhoOrg/AgnosTeset/adapters/libs/signaldeps"
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
	reflectdeps.Bind(&deps)
	serializables.Bind(&deps)
	serverdeps.Bind(&deps)
	signaldeps.Bind(&deps)
	sortdeps.Bind(&deps)
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	templatedeps.Bind(&deps)
	return deps
}
