package cmdroot

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmddebug"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdfollow"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdmigrate"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdstart"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdweb"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdmeta "github.com/13k/night-stalker/cmd/ns/internal/meta"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
)

const (
	cfgEnvPrefix      = "ns"
	cfgBaseName       = "config"
	cfgDefaultLogPath = "-"

	flagHelpLogPath = `log file or directory
* with "-", logs to stdout
* with directory, it will create a log with the command name and timestamp
`
)

var Cmd = &cobra.Command{
	Use:               "ns <command>",
	Short:             "Stalk dota2 players",
	RunE:              run,
	PersistentPreRunE: preRun,
	SilenceErrors:     true,
	SilenceUsage:      true,
}

var (
	ErrCommandFailureLogged = errors.New("command error")

	flagConfigFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	Cmd.PersistentFlags().StringVarP(&flagConfigFile, "config", "c", "", "config file (default automatic detection)")
	Cmd.PersistentFlags().StringP("db", "d", "", "database URL")
	Cmd.PersistentFlags().StringP("log", "l", cfgDefaultLogPath, flagHelpLogPath)
	Cmd.PersistentFlags().BoolP("tee", "t", false, "when logging to a file, also log to stdout")
	Cmd.PersistentFlags().BoolP("debug", "D", false, "enable debug logging")
	Cmd.PersistentFlags().BoolP("trace", "T", false, "enable trace logging")

	v.MustBindPersistentFlagLookup(v.KeyLogPath, Cmd, "log")
	v.MustBindPersistentFlagLookup(v.KeyLogTee, Cmd, "tee")
	v.MustBindPersistentFlagLookup(v.KeyLogDebug, Cmd, "debug")
	v.MustBindPersistentFlagLookup(v.KeyLogTrace, Cmd, "trace")
	v.MustBindPersistentFlagLookup(v.KeyDbURL, Cmd, "db")

	Cmd.AddCommand(cmddebug.Cmd)
	Cmd.AddCommand(cmdfollow.Cmd)
	Cmd.AddCommand(cmdimport.Cmd)
	Cmd.AddCommand(cmdmigrate.Cmd)
	Cmd.AddCommand(cmdstart.Cmd)
	Cmd.AddCommand(cmdweb.Cmd)
}

func initConfig() {
	v.AutoConfig(cfgBaseName, cfgEnvPrefix, flagConfigFile)
}

func run(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func preRun(cmd *cobra.Command, args []string) error {
	if err := nscmdlog.Init(cmdNamePath(cmd)); err != nil {
		return err
	}

	return nil
}

func cmdNamePath(cmd *cobra.Command) []string {
	var path []string

	for c := cmd; c != nil; c = c.Parent() {
		path = append([]string{c.Name()}, path...)
	}

	return path
}

func Execute(meta *nscmdmeta.Meta) error {
	Cmd.Version = fmt.Sprintf("%s (rev %s)", meta.Version, meta.Revision)

	if err := Cmd.Execute(); err != nil {
		if log := nscmdlog.Instance(); log != nil {
			log.WithError(err).Error("command error")
			return ErrCommandFailureLogged
		}

		return err
	}

	return nil
}
