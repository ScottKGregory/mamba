
# 🐍 Mamba
*Because who doesn't need more snakes?*

[![Build Status](https://travis-ci.com/ScottKGregory/mamba.svg?branch=main)](https://travis-ci.com/ScottKGregory/mamba)
[![Go Report Card](https://goreportcard.com/badge/github.com/ScottKGregory/mamba)](https://goreportcard.com/report/github.com/ScottKGregory/mamba)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/scottkgregory/mamba)
![GitHub](https://img.shields.io/github/license/scottkgregory/mamba)

## What is Mamba?

Mamba is a utility designed to make working with the combination of [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) require less boilerplate setup.

## Quick start

1. Install Mamba
```
go get github.com/scottkgregory/mamba
```

2. Annotation your config struct using the config tag in the form `config:"default,Description"`. Arrays, slices, and maps allow for setting default values via json.
```go
type Config struct {
	Root   string   `config:"defaultRoot,The root directory to do a thing with"`
	Number int      `config:"12,A number to use for a thing"`
	Snakes []string `config:"[\"adder\"],A list of snakes. Hsssss!"`
}
```

3. Call Mamba in the init of your Cobra command
```go
  mamba.Bind(Config{}, rootCmd)
```

4. Run your program with the `--help` flag to view your bound flags

5. Go forth and use your config! From this point on your config values will all be available via Viper

## Examples

- [basic-cli](./examples/basic-cli)

## Configuration

Options can be supplied to `mamba.Bind` to modify the way in which Mamba operates.
```go
mamba.Bind(AppConfig{}, rootCmd, &mamba.Options{ LogLevel: zerolog.TraceLevel })
```