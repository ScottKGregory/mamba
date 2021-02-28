package mamba

import "github.com/rs/zerolog"

// Options allows for configuring the mamba bindr.
type Options struct {
	// LogLevel sets the zerolog log level. By default it is set to none.
	LogLevel zerolog.Level
}
