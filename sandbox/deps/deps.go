package deps

import (
	argvdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/argvdeps"
	reflectdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/reflectdeps"
	serializables "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serializables"
	serverdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
	signaldeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/signaldeps"
	sortdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/sortdeps"
	std "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/std"
	stringsdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/stringsdeps"
)

// Deps is every capability the sandbox needs from the outside world, one field
// per sub-contract directory of sandbox/deps/. An adapter fills the fields; the
// sandbox only calls them, which is what keeps it free of OS packages.
type Deps struct {
	Argvdeps      argvdeps.Sandbox
	Reflectdeps   reflectdeps.Sandbox
	Serializables serializables.Sandbox
	Serverdeps    serverdeps.Sandbox
	Signaldeps    signaldeps.Sandbox
	Sortdeps      sortdeps.Sandbox
	Std           std.Sandbox
	Stringsdeps   stringsdeps.Sandbox
}
