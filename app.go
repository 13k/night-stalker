package ns

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/faceit/go-steam"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nssess "github.com/13k/night-stalker/internal/processors/session"
)

const (
	busBufSize               = 10
	liveMatchesQueryInterval = 30 * time.Second
	rtStatsPoolSize          = 5
	defaultShutdownTimeout   = 10 * time.Second
)

type AppOptions struct {
	Log             *nslog.Logger
	DB              *gorm.DB
	Redis           *redis.Client
	Credentials     *SteamCredentials
	ShutdownTimeout time.Duration
}

type App struct {
	options    *AppOptions
	log        *nslog.Logger
	db         *gorm.DB
	rds        *redis.Client
	steam      *steam.Client
	dota       *dota2.Dota2
	api        *geyser.Client
	apiDota    *geyserd2.Client
	supervisor *oversight.Tree
	ctx        context.Context
	cancel     context.CancelFunc
	bus        *nsbus.Bus
}

func New(options *AppOptions) (*App, error) {
	ns := &App{
		options: options,
		log:     options.Log,
		db:      options.DB,
		rds:     options.Redis,
		bus:     nsbus.New(busBufSize),
	}

	if ns.options.ShutdownTimeout == 0 {
		ns.options.ShutdownTimeout = defaultShutdownTimeout
	}

	if err := ns.setupAPI(); err != nil {
		return nil, err
	}

	ns.setupSteam()
	ns.setupDota()
	ns.setupSupervisor()
	ns.setupContext()

	return ns, nil
}

func (ns *App) setupAPI() error {
	var options []geyser.ClientOption
	var err error

	if ns.options.Credentials.APIKey != "" {
		options = append(options, geyser.WithKey(ns.options.Credentials.APIKey))
	}

	if ns.api, err = geyser.New(options...); err != nil {
		ns.log.WithError(err).Error("error creating API client")
		return err
	}

	if ns.apiDota, err = geyserd2.New(options...); err != nil {
		ns.log.WithError(err).Error("error creating Dota2 API client")
		return err
	}

	return nil
}

func (ns *App) setupSteam() {
	ns.steam = steam.NewClient()
}

func (ns *App) setupDota() {
	ns.dota = dota2.New(ns.steam, ns.log.Dota2Logger())
}

func (ns *App) setupSupervisor() {
	sessOptions := &nssess.ManagerOptions{
		Logger:                   ns.log,
		Credentials:              ns.options.Credentials.sessionCredentials(),
		ShutdownTimeout:          ns.options.ShutdownTimeout,
		LiveMatchesQueryInterval: liveMatchesQueryInterval,
		BusBufferSize:            busBufSize,
		RealtimeStatsPoolSize:    rtStatsPoolSize,
	}

	session := nssess.NewManager(sessOptions).ChildSpec()

	ns.supervisor = oversight.New(
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(ns.log.WithPackage("supervisor")),
		oversight.Process(session),
	)
}

func (ns *App) setupContext() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx = nsctx.WithLogger(ctx, ns.log)
	ctx = nsctx.WithDB(ctx, ns.db)
	ctx = nsctx.WithRedis(ctx, ns.rds)
	ctx = nsctx.WithBus(ctx, ns.bus)
	ctx = nsctx.WithSteam(ctx, ns.steam)
	ctx = nsctx.WithDota(ctx, ns.dota)
	ctx = nsctx.WithAPI(ctx, ns.api)
	ctx = nsctx.WithDotaAPI(ctx, ns.apiDota)

	ns.ctx = ctx
	ns.cancel = cancel
}

func (ns *App) Start() error {
	defer func() {
		ns.cancel()
		ns.bus.Shutdown()
	}()

	return ns.supervisor.Start(ns.ctx)
}

func (ns *App) Stop() {
	ns.cancel()
}
