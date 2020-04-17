package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func AutoConfig(basename, envPrefix, file string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Printf("Error determining current directory: %s\n", err)
			os.Exit(1)
		}

		viper.AddConfigPath(cwd)
		viper.SetConfigName(basename)
	}

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading configuration: %s\n", err)
		os.Exit(1)
	}
}
