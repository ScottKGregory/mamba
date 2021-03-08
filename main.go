package mamba

import (
	"github.com/scottkgregory/mamba/internal"
	"github.com/spf13/cobra"
)

type Options = internal.Options

func MustBind(obj interface{}, cmd *cobra.Command, options ...*Options) {
	if err := Bind(obj, cmd, options...); err != nil {
		panic(err)
	}
}

func Bind(obj interface{}, cmd *cobra.Command, options ...*Options) error {
	return internal.Bind(obj, cmd, options...)
}
