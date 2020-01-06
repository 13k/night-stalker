package web

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/cskr/pubsub"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	busBufSize = 10
)

type AppOptions struct {
	Log             *nslog.Logger
	DB              *gorm.DB
	Redis           *redis.Client
	StaticFS        http.FileSystem
	Address         string
	CertFile        string
	CertKeyFile     string
	CertHosts       []string
	CertCacheDir    string
	ShutdownTimeout time.Duration
}

type App struct {
	options                 *AppOptions
	log                     *nslog.Logger
	wslog                   *nslog.Logger
	db                      *gorm.DB
	bus                     *pubsub.PubSub
	engine                  *echo.Echo
	sv                      *http.Server
	ctx                     context.Context
	cancel                  context.CancelFunc
	stimeout                time.Duration
	rds                     *redis.Client
	rdsSubLiveMatchesUpdate *redis.PubSub
}

func New(options *AppOptions) (*App, error) {
	app := &App{
		options:  options,
		engine:   echo.New(),
		sv:       &http.Server{},
		log:      options.Log,
		wslog:    options.Log.WithPackage("ws"),
		db:       options.DB,
		rds:      options.Redis,
		bus:      pubsub.New(busBufSize),
		stimeout: options.ShutdownTimeout,
	}

	if err := app.configureEngine(); err != nil {
		return nil, err
	}

	if err := app.configureServer(); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) configureEngine() error { //nolint: unparam
	app.engine.Logger = app.log.EchoLogger()
	app.engine.Debug = app.log.Debugging()

	app.engine.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Output: app.log.Output(),
	}))

	app.engine.Use(mw.RecoverWithConfig(mw.RecoverConfig{
		DisableStackAll: true,
	}))

	root := app.engine

	api := root.Group("/api")

	apiV1 := api.Group("/v1")
	apiV1.GET("/heroes", app.serveHeroes)
	apiV1.GET("/live_matches", app.serveLiveMatches)
	apiV1.GET("/players/:account_id", app.servePlayer)

	root.GET("/ws", app.serveWS)

	assetHandler := http.FileServer(app.options.StaticFS)

	root.GET("/", echo.WrapHandler(assetHandler))
	root.GET("/*", echo.WrapHandler(assetHandler))

	return nil
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

	if err := app.subscribeLiveMatchesUpdate(); err != nil {
		return err
	}

	return app.engine.StartServer(app.sv)
}

func (app *App) Stop() {
	app.cancel()

	ctx, cancel := context.WithTimeout(context.Background(), app.stimeout)

	defer cancel()

	if err := app.sv.Shutdown(ctx); err != nil {
		app.log.WithError(err).Error("server shutdown error")
	}

	app.log.Warn("stop")
}

func (app *App) Routes() []*echo.Route {
	return app.engine.Routes()
}
