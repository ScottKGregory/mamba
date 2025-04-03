package mamba

import (
	"github.com/scottkgregory/mamba/internal"
	"github.com/spf13/cobra"
)

// Expose types from internal package in one place
type Options = internal.Options

var InvalidTypeError = internal.InvalidTypeError
var BindError = internal.BindError
var ParseError = internal.ParseError
var TagParseError = internal.TagParseError

// MustBind calls the mamba.Bind method and panics if an error is returned.
func MustBind(obj any, cmd *cobra.Command, options ...*Options) {
	if err := Bind(obj, cmd, options...); err != nil {
		panic(err)
	}
}

// Bind recursively iterates over all properties of the given object, binding flags
// for each one that is tagged with the `config` tag.
//
// Nested objects will result in dot-notation flags, e.g. `server.port`. Other
// separators can be supplied via the Options param if full-stops are not desired.
func Bind(obj any, cmd *cobra.Command, options ...*Options) error {
	return internal.Bind(obj, cmd, options...)
}
