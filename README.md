# 🐍 Mamba

_Because who doesn't need more snakes?_

![Build Status](https://github.com/OWNER/REPOSITORY/actions/workflows/go.yml/badge.svg)
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

1. Annotation your config struct using the config tag in the form `config:"default, description, persistent, shorthand"`. Arrays, slices, and maps allow for setting default values via json. Any of these values can be omitted.

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

```go
type AppConfig struct {
	ConfigFile string   `config:"./examples/basic-cli/config.yaml,The yaml config file to read, true, c"`
	Messages   Messages `config:""`
	Snakes     []string `config:"\"[\"\"Mamba\"\",\"\"Viper\"\"]\", A list of snakes. Hsssss!, true, s"`
	*Embedded  `config:""`
}

type Messages struct {
	Greeting string `config:"Hello there!, The greating to use, false, g"`
}

type Embedded struct {
	Farewell string `config:"Goodbye!, The farewell to use"`
}

var rootCmd = &cobra.Command{
	Use:   "basic-cli",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("messages.greeting"))
		fmt.Println("These are all the snakes that exist in the world:")

		for _, s := range viper.GetStringSlice("snakes") {
			fmt.Printf("  - %s\n", s)
		}
		fmt.Println(viper.GetString("embedded.farewell"))
	},
}
```

## Configuration

Options can be supplied to `mamba.Bind` to modify the way in which Mamba operates.

```go
mamba.Bind(AppConfig{}, rootCmd, &mamba.Options{ Persistent: true })
```
