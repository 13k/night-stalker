package viper

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func MustBindFlag(key Key, cmd *cobra.Command, flag *pflag.Flag) {
	if err := viper.BindPFlag(string(key), flag); err != nil {
		panic(err)
	}
}

func MustBindFlagLookup(key Key, cmd *cobra.Command, flagName string) {
	flag := cmd.Flags().Lookup(flagName)

	if flag == nil {
		panic(fmt.Errorf("invalid flag name %q", flagName))
	}

	MustBindFlag(key, cmd, flag)
}

func MustBindPersistentFlagLookup(key Key, cmd *cobra.Command, flagName string) {
	flag := cmd.PersistentFlags().Lookup(flagName)

	if flag == nil {
		panic(fmt.Errorf("invalid flag name %q", flagName))
	}

	MustBindFlag(key, cmd, flag)
}
