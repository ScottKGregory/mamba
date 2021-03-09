package internal

// Options allows for configuring the mamba binder.
type Options struct {
	// Setting Persistent to true will cause all flags to be bound as persistent.
	// Persistence can also be set individually in the config struct tag.
	Persistent bool
}
