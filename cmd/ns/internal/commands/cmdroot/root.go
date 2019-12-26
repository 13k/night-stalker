package cmdroot

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmddebug"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdfollow"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdmigrate"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdstart"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdweb"
)

const (
	cfgEnvPrefix      = "ns"
	cfgBaseName       = "config"
	cfgDefaultLogFile = "-"
)

var Cmd = &cobra.Command{
	Use:   "ns <command>",
	Short: "Stalk dota2 players",
	Run:   run,
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default automatic detection).")
	Cmd.PersistentFlags().StringP("db", "d", "", "database URL")
	Cmd.PersistentFlags().StringP("log", "l", cfgDefaultLogFile, `log file. "-" logs to stdout`)
	Cmd.PersistentFlags().BoolP("tee", "t", false, "when logging to a file, also log to stdout")
	Cmd.PersistentFlags().BoolP("debug", "D", false, "enable debug logging")

	if err := viper.BindPFlag("log.file", Cmd.PersistentFlags().Lookup("log")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("log.debug", Cmd.PersistentFlags().Lookup("debug")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("log.stdout", Cmd.PersistentFlags().Lookup("tee")); err != nil {
		panic(err)
	}

	if err := viper.BindPFlag("db.url", Cmd.PersistentFlags().Lookup("db")); err != nil {
		panic(err)
	}

	Cmd.AddCommand(cmddebug.Cmd)
	Cmd.AddCommand(cmdfollow.Cmd)
	Cmd.AddCommand(cmdimport.Cmd)
	Cmd.AddCommand(cmdmigrate.Cmd)
	Cmd.AddCommand(cmdstart.Cmd)
	Cmd.AddCommand(cmdweb.Cmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, err := os.Getwd()

		if err != nil {
			fmt.Printf("Error determining current directory: %s\n", err)
			os.Exit(1)
		}

		viper.AddConfigPath(cwd)
		viper.SetConfigName(cfgBaseName)
	}

	viper.SetEnvPrefix(cfgEnvPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration: %s\n", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	if err := cmd.Usage(); err != nil {
		panic(err)
	}
}

func Execute() error {
	return Cmd.Execute()
}
