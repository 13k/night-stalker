package cmdroot

import (
	"fmt"
	"strings"

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
	cfgDefaultLogFile = "-"
)

var Cmd = &cobra.Command{
	Use:              "ns <command>",
	Short:            "Stalk dota2 players",
	Run:              run,
	PersistentPreRun: preRun,
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default automatic detection)")
	Cmd.PersistentFlags().StringP("db", "d", "", "database URL")
	Cmd.PersistentFlags().StringP("log", "l", cfgDefaultLogFile, `log file. "-" logs to stdout`)
	Cmd.PersistentFlags().BoolP("debug", "D", false, "enable debug logging")
	Cmd.PersistentFlags().BoolP("tee", "t", false, "when logging to a file, also log to stdout")
	Cmd.PersistentFlags().BoolP("trace", "T", false, "enable tracing logging")

	v.MustBindPersistentFlagLookup(v.KeyLogFile, Cmd, "log")
	v.MustBindPersistentFlagLookup(v.KeyLogDebug, Cmd, "debug")
	v.MustBindPersistentFlagLookup(v.KeyLogDebug, Cmd, "trace")
	v.MustBindPersistentFlagLookup(v.KeyLogTee, Cmd, "tee")
	v.MustBindPersistentFlagLookup(v.KeyDbURL, Cmd, "db")

	Cmd.AddCommand(cmddebug.Cmd)
	Cmd.AddCommand(cmdfollow.Cmd)
	Cmd.AddCommand(cmdimport.Cmd)
	Cmd.AddCommand(cmdmigrate.Cmd)
	Cmd.AddCommand(cmdstart.Cmd)
	Cmd.AddCommand(cmdweb.Cmd)
}

func initConfig() {
	v.AutoConfig(cfgBaseName, cfgEnvPrefix, cfgFile)
}

func run(cmd *cobra.Command, args []string) {
	if err := cmd.Usage(); err != nil {
		panic(err)
	}
}

func preRun(cmd *cobra.Command, args []string) {
	flagLog := cmd.Flags().Lookup("log")

	if flagLog == nil {
		return
	}

	var cmdNames []string
	c := cmd

	for c != nil {
		cmdNames = append([]string{c.Name()}, cmdNames...)
		c = c.Parent()
	}

	cmdName := strings.Join(cmdNames, "-")
	lpath := nscmdlog.ParseLogfilePath(flagLog.Value.String(), cmdName)

	if err := flagLog.Value.Set(lpath); err != nil {
		panic(err)
	}
}

func Execute(meta *nscmdmeta.Meta) error {
	Cmd.Version = fmt.Sprintf("%s (rev %s)", meta.Version, meta.Revision)
	return Cmd.Execute()
}
