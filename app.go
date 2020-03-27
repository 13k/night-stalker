package ns

import (
	"context"
	"fmt"
	"time"

	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/faceit/go-steam"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsrds "github.com/13k/night-stalker/internal/redis"
	nssteam "github.com/13k/night-stalker/internal/steam"
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
	Redis           *nsrds.Redis
	Credentials     *SteamCredentials
	ShutdownTimeout time.Duration
}

type App struct {
	options    AppOptions
	log        *nslog.Logger
	bus        *nsbus.Bus
	db         *gorm.DB
	rds        *nsrds.Redis
	steam      *nssteam.Client
	dota       *nsdota2.Client
	api        *geyser.Client
	apiDota    *geyserd2.Client
	supervisor *supervisor
	ctx        context.Context
	cancel     context.CancelFunc
}

func New(options AppOptions) (*App, error) {
	log := options.Log.WithPackage("app")

	bus := nsbus.New(nsbus.Options{
		Cap: busBufSize,
		Log: log,
	})

	app := &App{
		options: options,
		log:     log,
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

	app.setupSteam()
	app.setupDota()
	app.setupContext()
	app.setupSupervisor()

	return app, nil
}

func (app *App) setupSteam() {
	app.steam = nssteam.NewClient(steam.NewClient())
}

func (app *App) setupDota() {
	dota2 := dota2.New(app.steam.Client, app.log.WithPackage("dota2").LogrusLogger())
	app.dota = nsdota2.NewClient(dota2)
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

func (app *App) setupContext() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = nsctx.WithLogger(ctx, app.options.Log)
	ctx = nsctx.WithBus(ctx, app.bus)
	ctx = nsctx.WithDB(ctx, app.db)
	ctx = nsctx.WithRedis(ctx, app.rds)
	ctx = nsctx.WithSteam(ctx, app.steam)
	ctx = nsctx.WithDota(ctx, app.dota)
	ctx = nsctx.WithAPI(ctx, app.api)
	ctx = nsctx.WithDotaAPI(ctx, app.apiDota)

	app.ctx = ctx
	app.cancel = cancel
}

func (app *App) setupSupervisor() {
	app.supervisor = newSupervisor(supervisorOptions{
		Log:                   app.options.Log,
		Bus:                   app.bus,
		ShutdownTimeout:       app.options.ShutdownTimeout,
		Credentials:           app.options.Credentials,
		TVGamesInterval:       tvGamesInterval,
		RealtimeStatsPoolSize: rtStatsPoolSize,
		RealtimeStatsInterval: rtStatsInterval,
		MatchInfoPoolSize:     matchInfoPoolSize,
		MatchInfoInterval:     matchInfoInterval,
	})
}

func (app *App) Start() error {
	defer func() {
		app.cancel()
		app.bus.Shutdown()
		app.log.Warn("stopped")
	}()

	go app.eventsLoop()

	app.log.Info("starting")

	return app.supervisor.Start(app.ctx)
}

func (app *App) Stop() {
	app.cancel()
}

func (app *App) eventsLoop() {
	for {
		select {
		case <-app.ctx.Done():
			return
		case ev, ok := <-app.steam.Events():
			if !ok {
				app.log.Warn("steam events channel closed")
				return
			}

			err := app.bus.Pub(nsbus.Message{
				Topic:   nsbus.TopicSteamEvents,
				Payload: &nsbus.SteamEventMessage{Event: ev},
			})

			if err != nil {
				app.log.WithOFields(
					"topic", nsbus.TopicSteamEvents,
					"event", fmt.Sprintf("%T", ev),
				).WithError(err).Error("error publishing event")
			}
		}
	}
}
