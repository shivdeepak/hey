package cmd

import (
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

		viper.SetConfigFile(filepath.Join(home, ".config", "hey", "config.yaml"))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}

	log.Debug().Msg("Viper Initialized!")
}

func initViperConfigFile() {
	configFilePath := viper.ConfigFileUsed()
	log.Debug().Msgf("Checking if config file %s exists", configFilePath)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		log.Debug().Msg("Config file doesn't exist. Creating now.")

		mkdir_err := os.MkdirAll(filepath.Dir(configFilePath), os.ModePerm)
		if mkdir_err != nil {
			log.Error().Err(mkdir_err).Msgf("Error creating directory for config file: %s", configFilePath)
			os.Exit(1)
		}

		file_write_err := viper.SafeWriteConfigAs(configFilePath)
		if file_write_err != nil {
			log.Error().Err(file_write_err).Msgf("Error creating configuration file: %s", configFilePath)
			os.Exit(1)
		} else {
			log.Debug().Msgf("Successfully created the configuration file: %s", configFilePath)
		}

		read_config_err := viper.ReadInConfig()
		if read_config_err != nil {
			log.Error().Err(read_config_err).Msg("Error reading configuration file")
		} else {
			log.Debug().Msg("Config file read!")
		}
	}

}

func saveCurrentViperConfig() {
	configFilePath := viper.ConfigFileUsed()
	file_write_err := viper.WriteConfigAs(configFilePath)
	if file_write_err != nil {
		log.Error().Err(file_write_err).Msgf("Error writing to configuration file: %s", configFilePath)
		os.Exit(1)
	} else {
		log.Debug().Msgf("Successfully updated configuration file: %s", configFilePath)
	}
}
