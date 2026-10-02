package envdeps

import (
	"os"

	envdeps "github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/envdeps"

	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps"
)

// Bind fills deps.Deps.Envdeps with the standard library's os.
func Bind(deps *deps.Deps) {
	deps.Envdeps = envdeps.Sandbox{
		Getenv: func(key string) string {
			return os.Getenv(key)
		},
	}
}
