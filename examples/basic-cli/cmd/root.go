package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/scottkgregory/mamba"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	mamba.MustBind(&AppConfig{}, rootCmd, &mamba.Options{PrefixEmbedded: false})
}

func initConfig() {
	cfgFile := viper.GetString("configfile")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}
