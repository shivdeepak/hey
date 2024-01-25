package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rs/zerolog/log"
)

func initViper() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configType := "yaml"
		configName := ".hey"

		viper.SetConfigFile(filepath.Join(home, ".config", configName+"."+configType))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	log.Debug().Msg("Viper Initialized!")
}

func initViperConfigFile() {

}

func loadViperConfigFile() {

}
