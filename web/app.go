package web

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsrds "github.com/13k/night-stalker/internal/redis"
	nswebhdl "github.com/13k/night-stalker/web/internal/handlers"
	nswebmw "github.com/13k/night-stalker/web/internal/middleware"
	nswebrds "github.com/13k/night-stalker/web/internal/redis"
)

const (
	busBufSize = 10
)

type AppOptions struct {
	Log             *nslog.Logger
	DB              *nsdb.DB
	Redis           *nsrds.Redis
	StaticFS        http.FileSystem
	Address         string
	CertFile        string
	CertKeyFile     string
	CertHosts       []string
	CertCacheDir    string
	ShutdownTimeout time.Duration
}

type App struct {
	options AppOptions
	log     *nslog.Logger
	db      *nsdb.DB
	dbl     *nsdbda.Loader
	bus     *nsbus.Bus
	rds     *nsrds.Redis
	pubsub  *nswebrds.PubSub
	engine  *echo.Echo
	sv      *http.Server
	ctx     context.Context
	cancel  context.CancelFunc
}

func New(options AppOptions) (*App, error) {
	bus := nsbus.New(nsbus.Options{
		Cap: busBufSize,
		Log: options.Log,
	})

	app := &App{
		options: options,
		engine:  echo.New(),
		sv:      &http.Server{},
		log:     options.Log,
		bus:     bus,
		db:      options.DB,
		dbl:     nsdbda.NewLoader(options.DB),
		rds:     options.Redis,
	}

	app.configurePubSub()
	app.configureEngine()

	if err := app.configureServer(); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) configurePubSub() {
	app.pubsub = nswebrds.NewPubSub(&nswebrds.PubSubOptions{
		Log:        app.log,
		Bus:        app.bus,
		Redis:      app.rds,
		DataLoader: app.dbl,
	})
}

func (app *App) configureEngine() {
	app.engine.Logger = app.log.EchoLogger()
	app.engine.StdLogger = app.log.StdLogger()
	app.engine.Debug = app.log.IsLevelEnabled(nslog.LevelDebug)
	app.engine.HTTPErrorHandler = app.handleError

	app.engine.Use(nswebmw.Context())
	app.engine.Use(nswebmw.Logger(app.log))
	app.engine.Use(nswebmw.MediaType())
	app.engine.Use(nswebmw.ErrorHandler())
	app.engine.Use(nswebmw.Recover())

	root := app.engine

	api := root.Group("/api")

	apiV1 := api.Group("/v1")
	apiV1.GET("/search", app.serveSearch)
	apiV1.GET("/leagues", app.serveLeagues)
	apiV1.GET("/heroes", app.serveHeroes)
	apiV1.GET("/heroes/:id/matches", app.serveHeroMatches)
	apiV1.GET("/live_matches", app.serveLiveMatches)
	apiV1.GET("/players/:account_id/matches", app.servePlayerMatches)

	root.GET("/ws", app.serveWS)

	assetHandler := nswebhdl.AssetHandler(app.options.StaticFS)

	root.GET("/", assetHandler)
	root.GET("/*", assetHandler)
}

func (app *App) configureServer() error {
	var err error

	if app.options.CertFile != "" {
		app.sv.TLSConfig = &tls.Config{}
		app.sv.TLSConfig.Certificates = make([]tls.Certificate, 1)
		app.sv.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(app.options.CertFile, app.options.CertKeyFile)

		if err != nil {
			app.log.WithError(err).Error("error loading certificate")
			return err
		}
	} else if len(app.options.CertHosts) > 0 {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(app.options.CertHosts...),
			Cache:      autocert.DirCache(app.options.CertCacheDir),
		}

		app.sv.TLSConfig = certManager.TLSConfig()
	}

	app.sv.Addr = app.options.Address

	return nil
}

func (app *App) Start() error {
	app.ctx, app.cancel = context.WithCancel(context.Background())

	if err := app.pubsub.Start(app.ctx); err != nil {
		return err
	}

	return app.engine.StartServer(app.sv)
}

func (app *App) Stop() {
	app.cancel()

	ctx, cancel := context.WithTimeout(context.Background(), app.options.ShutdownTimeout)

	defer cancel()

	if err := app.sv.Shutdown(ctx); err != nil {
		app.log.WithError(err).Error("server shutdown error")
	}

	app.log.Warn("stop")
}

func (app *App) Routes() []*echo.Route {
	return app.engine.Routes()
}
