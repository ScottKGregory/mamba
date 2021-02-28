package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/mamba"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Messages  Messages `config:","`
	Snakes    []string `config:"[\"adder\"],A list of snakes. Hsssss!"`
	*Embedded `config:","`
}

type Messages struct {
	Greeting string `config:"Hello there!,The greating to use"`
}

type Embedded struct {
	Farewell string `config:"Goodbye!,The farewell to use"`
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

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	mamba.Bind(AppConfig{}, rootCmd, &mamba.Options{LogLevel: zerolog.TraceLevel})
}
