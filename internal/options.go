package internal

// Options allows for configuring the mamba binder.
type Options struct {
	// Persistent (Default `false`) will cause all flags to be bound as persistent.
	// Persistence can also be set individually in the config struct tag.
	Persistent bool

	// Separator (Default `.`) will be used instead of dot notation when constructing flags.
	// For example a `-` could be supplied leading to flags like `server-port` insetad
	// of the standard `server.port`.
	Separator string

	// PrefixEmbedded (Defaults to True) allows you to control if properties of embedded structs will
	// have the struct name as a prefix. For example:
	//
	// type Config struct {
	//   *Server `config:""`
	// }
	//
	// type Server struct {
	//   Port int `config:"8080,The port to listen on"`
	// }
	//
	// With `PrefixEmbedded = true`: `server.port`
	// With `PrefixEmbedded` = false`: `port`
	//
	// Use with caution, properties with the same name in two embdedded structs will fail
	// to bind with this option set to false
	PrefixEmbedded bool
}
