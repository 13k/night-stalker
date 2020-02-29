package cmdweb

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdwebroutes "github.com/13k/night-stalker/cmd/ns/internal/commands/cmdweb/routes"
	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdrds "github.com/13k/night-stalker/cmd/ns/internal/redis"
	nsweb "github.com/13k/night-stalker/web"
)

var Cmd = &cobra.Command{
	Use:   "web",
	Short: "Start HTTP server",
	Run:   run,
}

const (
	defaultAddress         = ":3000"
	defaultShutdownTimeout = 5 * time.Second
)

var (
	flagShutdownTimeout time.Duration
	flagAddress         string
	flagCertFile        string
	flagCertKeyFile     string
	flagCertHosts       []string
	flagCertCacheDir    string
)

func init() {
	Cmd.Flags().StringP("redis", "r", "", "redis URL")
	Cmd.Flags().StringVarP(&flagAddress, "listen", "L", defaultAddress, "listen address (host:port)")
	Cmd.Flags().StringVar(&flagCertFile, "crt", "", "certificate file")
	Cmd.Flags().StringVar(&flagCertKeyFile, "crtkey", "", "certificate key file")
	Cmd.Flags().StringVar(&flagCertCacheDir, "crtcache", "", "certificate cache directory")
	Cmd.Flags().StringSliceVar(&flagCertHosts, "crthost", nil, "certificate host(s)")
	Cmd.Flags().DurationVar(&flagShutdownTimeout, "stimeout", defaultShutdownTimeout, "shutdown timeout")

	if err := viper.BindPFlag("redis.url", Cmd.Flags().Lookup("redis")); err != nil {
		panic(err)
	}

	Cmd.AddCommand(cmdwebroutes.Cmd)
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	rds, err := nscmdrds.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to redis")
	}

	defer rds.Close()

	assets, err := assetsFS()

	if err != nil {
		log.WithError(err).Fatal("error loading assets")
	}

	app, err := nsweb.New(nsweb.AppOptions{
		Log:             log.WithPackage("web"),
		DB:              db,
		Redis:           rds,
		StaticFS:        assets,
		Address:         flagAddress,
		CertFile:        flagCertFile,
		CertKeyFile:     flagCertKeyFile,
		CertHosts:       flagCertHosts,
		CertCacheDir:    flagCertCacheDir,
		ShutdownTimeout: flagShutdownTimeout,
	})

	if err != nil {
		log.WithError(err).Fatal("error initializing application")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		log.Warn("interrupted, stopping...")
		app.Stop()
	}()

	if err := app.Start(); err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("web error")
	}
}

func assetsFS() (http.FileSystem, error) {
	return pkger.Open("/balanar/dist")
}
