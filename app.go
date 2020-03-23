package ns

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nssess "github.com/13k/night-stalker/internal/processors/session"
)

const (
	busBufSize             = uint(100)
	tvGamesInterval        = 60 * time.Second
	rtStatsPoolSize        = 10
	rtStatsInterval        = 30 * time.Second
	matchInfoPoolSize      = 10
	matchInfoInterval      = 45 * time.Second
	defaultShutdownTimeout = 10 * time.Second
)

type AppOptions struct {
	Log             *nslog.Logger
	DB              *gorm.DB
	Redis           *redis.Client
	Credentials     *SteamCredentials
	ShutdownTimeout time.Duration
}

type App struct {
	options    AppOptions
	log        *nslog.Logger
	bus        *nsbus.Bus
	db         *gorm.DB
	rds        *redis.Client
	api        *geyser.Client
	apiDota    *geyserd2.Client
	supervisor *oversight.Tree
	ctx        context.Context
	cancel     context.CancelFunc
}

func New(options AppOptions) (*App, error) {
	bus := nsbus.New(nsbus.Options{
		Cap: busBufSize,
		Log: options.Log,
	})

	app := &App{
		options: options,
		log:     options.Log,
		db:      options.DB,
		rds:     options.Redis,
		bus:     bus,
	}

	if app.options.ShutdownTimeout == 0 {
		app.options.ShutdownTimeout = defaultShutdownTimeout
	}

	if err := app.setupAPI(); err != nil {
		return nil, err
	}

	app.setupSupervisor()
	app.setupContext()

	return app, nil
}

func (app *App) setupAPI() error {
	var options []geyser.ClientOption
	var err error

	if app.options.Credentials.APIKey != "" {
		options = append(options, geyser.WithKey(app.options.Credentials.APIKey))
	}

	if app.api, err = geyser.New(options...); err != nil {
		app.log.WithError(err).Error("error creating API client")
		return err
	}

	if app.apiDota, err = geyserd2.New(options...); err != nil {
		app.log.WithError(err).Error("error creating Dota2 API client")
		return err
	}

	return nil
}

func (app *App) setupSupervisor() {
	sessOptions := nssess.ManagerOptions{
		Log:                   app.log,
		Bus:                   app.bus,
		Credentials:           app.options.Credentials.sessionCredentials(),
		ShutdownTimeout:       app.options.ShutdownTimeout,
		TVGamesInterval:       tvGamesInterval,
		RealtimeStatsPoolSize: rtStatsPoolSize,
		RealtimeStatsInterval: rtStatsInterval,
		MatchInfoPoolSize:     matchInfoPoolSize,
		MatchInfoInterval:     matchInfoInterval,
	}

	session := nssess.NewManager(sessOptions)

	app.supervisor = oversight.New(
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(app.log.WithPackage("supervisor").OversightLogger()),
		oversight.Process(session.ChildSpec()),
	)
}

func (app *App) setupContext() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = nsctx.WithLogger(ctx, app.log)
	ctx = nsctx.WithBus(ctx, app.bus)
	ctx = nsctx.WithDB(ctx, app.db)
	ctx = nsctx.WithRedis(ctx, app.rds)
	ctx = nsctx.WithAPI(ctx, app.api)
	ctx = nsctx.WithDotaAPI(ctx, app.apiDota)

	app.ctx = ctx
	app.cancel = cancel
}

func (app *App) Start() error {
	defer func() {
		app.cancel()
		app.bus.Shutdown()
	}()

	return app.supervisor.Start(app.ctx)
}

func (app *App) Stop() {
	app.cancel()
}
