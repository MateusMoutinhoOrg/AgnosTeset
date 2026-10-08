package osenv

import (
	"os"

	envdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/envdeps"

	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

// Bind fills deps.Deps.EnvDeps with the standard library's os.
func Bind(deps *deps.Deps) {
	deps.EnvDeps = envdeps.Contract{
		Getenv: func(key string) string {
			return os.Getenv(key)
		},
	}
}
