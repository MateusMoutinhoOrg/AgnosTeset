package api

// UserConfig is the user-visible subset of api.Config.
//
// This type is embedded in sandbox/api/config.go. The only thing that writes it
// is sandbox/start.go, and you can add fields to it with impunity: if you add
// a method, just make sure sandbox/api/config.go embeds the updated UserConfig,
// and nothing breaks. If you add a field called Cli, Config, Actions or Deps,
// the verify tool will complain, and you should rename it.
//
// This is the only API contract you need to worry about. The implementation
// lives under sandbox/internal/, and can be refactored or replaced wholesale.
type UserConfig struct {
	Secret string
}
