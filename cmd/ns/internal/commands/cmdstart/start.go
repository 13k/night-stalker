package cmdstart

import (
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	ns "github.com/13k/night-stalker"
	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdrds "github.com/13k/night-stalker/cmd/ns/internal/redis"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
)

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "Start stalking",
	RunE:  run,
}

const (
	defaultShutdownTimeout = 10 * time.Second
)

var (
	flagShutdownTimeout time.Duration
)

func init() {
	Cmd.Flags().StringP("redis", "r", "", "redis URL")
	Cmd.Flags().DurationVar(&flagShutdownTimeout, "stimeout", defaultShutdownTimeout, "shutdown timeout")

	v.MustBindFlagLookup(v.KeyRedisURL, Cmd, "redis")
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	rds, err := nscmdrds.Connect()

	if err != nil {
		return xerrors.Errorf("error connecting to redis: %w", err)
	}

	defer rds.Close()

	credentials := &ns.SteamCredentials{
		Username:         v.GetString(v.KeySteamUser),
		Password:         v.GetString(v.KeySteamPassword),
		RememberPassword: true,
	}

	app, err := ns.New(ns.AppOptions{
		Log:             log,
		DB:              db,
		Redis:           rds,
		Credentials:     credentials,
		ShutdownTimeout: flagShutdownTimeout,
	})

	if err != nil {
		return xerrors.Errorf("error initializing application: %w", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		log.Warn("interrupted, stopping...")
		app.Stop()
	}()

	if err := app.Start(); err != nil {
		return xerrors.Errorf("application error: %w", err)
	}

	return nil
}
