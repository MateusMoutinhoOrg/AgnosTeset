package api

// Sandbox is the whole library: one field per contract declared in
// sandbox/api/, each built by the New<Contract> of its own package under
// sandbox/internal/. sandbox.New returns it, and nothing callable lives outside
// of it.
type Sandbox struct {
	Config Config
}
