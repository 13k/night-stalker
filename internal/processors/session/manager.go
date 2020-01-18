package session

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	"github.com/faceit/go-steam/protocol/steamlang"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsgc "github.com/13k/night-stalker/internal/processors/gc"
	nslm "github.com/13k/night-stalker/internal/processors/livematches"
	nsrts "github.com/13k/night-stalker/internal/processors/rtstats"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "session"

	helloRetryCount    = 360
	helloRetryInterval = 10 * time.Second
)

type ManagerOptions struct {
	Logger                   *nslog.Logger
	Credentials              *Credentials
	ShutdownTimeout          time.Duration
	LiveMatchesQueryInterval time.Duration
	BusBufferSize            int
	RealtimeStatsPoolSize    int
}

var _ nsproc.Processor = (*Manager)(nil)

type Manager struct {
	options       *ManagerOptions
	ctx           context.Context
	sessionCtx    context.Context
	sessionCancel context.CancelFunc
	login         *models.SteamLogin
	log           *nslog.Logger
	db            *gorm.DB
	steam         *steam.Client
	dota          *dota2.Dota2
	bus           *nsbus.Bus
	supervisor    *oversight.Tree
	ready         bool
	helloTicker   *time.Ticker
	helloCount    int
}

func NewManager(options *ManagerOptions) *Manager {
	p := &Manager{
		options: options,
		log:     options.Logger.WithPackage(processorName),
		login:   &models.SteamLogin{},
	}

	p.setupSupervisor()

	return p
}

func (p *Manager) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if p.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(p.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Permanent(),
		Start:    p.Start,
		Shutdown: shutdown,
	}
}

func (p *Manager) setupSupervisor() {
	dispatcherSpec := nsgc.NewDispatcher(&nsgc.DispatcherOptions{
		Logger:          p.log,
		BufferSize:      p.options.BusBufferSize,
		ShutdownTimeout: p.options.ShutdownTimeout,
	}).ChildSpec()

	liveMatchesSpec := nslm.NewWatcher(&nslm.WatcherOptions{
		Logger:          p.log,
		Interval:        p.options.LiveMatchesQueryInterval,
		ShutdownTimeout: p.options.ShutdownTimeout,
	}).ChildSpec()

	rtStatsSpec := nsrts.NewMonitor(&nsrts.MonitorOptions{
		Logger:          p.log,
		PoolSize:        p.options.RealtimeStatsPoolSize,
		ShutdownTimeout: p.options.ShutdownTimeout,
	}).ChildSpec()

	p.supervisor = oversight.New(
		oversight.NeverHalt(),
		oversight.WithRestartStrategy(oversight.OneForOne()),
		oversight.WithLogger(p.log.WithPackage("supervisor")),
		oversight.Process(dispatcherSpec, liveMatchesSpec, rtStatsSpec),
	)
}

func (p *Manager) Start(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return nsproc.ErrProcessorContextDotaClient
	}

	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	if err := p.loadLogin(); err != nil {
		p.log.WithError(err).Error("error loading login info")
		return err
	}

	if p.isSuspended() {
		p.log.WithError(ErrDotaClientSuspended).Error("client suspended")
		return ErrDotaClientSuspended
	}

	if err := p.connectSteam(); err != nil {
		p.log.WithError(err).Error("failed to connect to steam")
		return err
	}

	p.ctx = ctx

	return p.loop()
}

func (p *Manager) stop() {
	p.log.Warn("stopping...")
	p.cancelSession()
	p.dota.Close()
	p.steam.Disconnect()
}

func (p *Manager) publishEvent(ev interface{}) {
	p.bus.TryPub(&nsbus.SteamEventMessage{Event: ev}, nsbus.TopicSteamEvents)
}

func (p *Manager) startSession() {
	p.log.Debug("startSession()")

	p.sessionCtx, p.sessionCancel = context.WithCancel(p.ctx)

	go func() {
		if err := p.supervisor.Start(p.sessionCtx); err != nil {
			p.log.WithError(err).Error("supervisor error")
		}
	}()
}

func (p *Manager) cancelSession() {
	p.log.WithField("sessionCancel_nil", p.sessionCancel == nil).Debug("cancelSession()")

	if p.sessionCancel != nil {
		p.sessionCancel()
	}
}

func (p *Manager) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	defer func() {
		p.stop()
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case ev := <-p.steam.Events():
			if err := p.handleEvent(ev); err != nil {
				p.log.WithError(err).Error()
				return err
			}
		}
	}
}

func (p *Manager) handleEvent(ev interface{}) error {
	var err error

	switch e := ev.(type) {
	case *steam.ClientCMListEvent:
		err = p.onSteamServerList(e)
	case *steam.ConnectedEvent:
		err = p.onSteamConnect()
	case *steam.DisconnectedEvent:
		err = p.onSteamDisconnect()
	case *steam.LoggedOnEvent:
		err = p.onSteamLogOn(e)
	case *steam.LoginKeyEvent:
		err = p.onSteamLoginKey(e)
	case *steam.MachineAuthUpdateEvent:
		err = p.onSteamMachineAuth(e)
	case *steam.WebSessionIdEvent:
		err = p.onSteamWebSession(e)
	case *steam.WebLoggedOnEvent:
		err = p.onSteamWebLogOn(e)
	case *steam.LoggedOffEvent:
		err = p.onSteamLogOff(e)
	case *steam.LogOnFailedEvent:
		err = p.onSteamLogOnFail(e)
	case *events.ClientSuspended:
		err = p.onDotaClientSuspended(e)
	case events.ClientStateChanged:
		err = p.onDotaGCStateChange(e)
	case *events.ClientWelcomed:
		err = p.onDotaWelcome(e)
	case steam.FatalErrorEvent:
		p.log.WithError(e).Error("steam error")
		err = e
	default:
		p.publishEvent(ev)
	}

	return err
}

func (p *Manager) isSuspended() bool {
	return p.login.SuspendedUntil != nil && time.Now().Before(*p.login.SuspendedUntil)
}

func (p *Manager) randomServer() (*models.SteamServer, error) {
	server := &models.SteamServer{}
	result := p.db.Order("random()").Take(&server)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return server, nil
}

func (p *Manager) saveServers(addresses []*netutil.PortAddr) error {
	var err error

	tx := p.db.Begin()

	for _, addr := range addresses {
		server := &models.SteamServer{}
		err = p.db.
			Where(&models.SteamServer{Address: addr.String()}).
			FirstOrCreate(server).
			Error

		if err != nil {
			break
		}
	}

	if err != nil {
		p.log.WithError(err).Error("error saving servers")
		return tx.Rollback().Error
	}

	return tx.Commit().Error
}

func (p *Manager) loadLogin() error {
	err := p.db.
		Where(&models.SteamLogin{Username: p.options.Credentials.Username}).
		FirstOrInit(p.login).
		Error

	if err != nil {
		p.log.WithError(err).Error("error loading login")
	}

	return err
}

func (p *Manager) updateLogin(update *models.SteamLogin) error {
	p.log.Debug("update login")

	err := p.db.
		Where(&models.SteamLogin{Username: p.login.Username}).
		Assign(update).
		FirstOrCreate(p.login).
		Error

	if err != nil {
		p.log.WithError(err).Error("error updating login")
	}

	return err
}

func (p *Manager) saveDotaWelcome(welcome *protocol.CMsgClientWelcome) error {
	update := &models.SteamLogin{
		GameVersion:       welcome.GetVersion(),
		LocationCountry:   welcome.GetLocation().GetCountry(),
		LocationLatitude:  welcome.GetLocation().GetLatitude(),
		LocationLongitude: welcome.GetLocation().GetLongitude(),
	}

	return p.updateLogin(update)
}

func (p *Manager) connectSteam() error {
	server, err := p.randomServer()

	if err != nil {
		p.log.WithError(err).Error("error loading steam server")
		return err
	}

	if server != nil {
		addr := netutil.ParsePortAddr(server.Address)

		if addr == nil {
			p.log.WithError(err).WithField("address", server.Address).Error("error parsing server address")
			return ErrInvalidServerAddress
		}

		p.steam.ConnectTo(addr)
	} else {
		if err := steam.InitializeSteamDirectory(); err != nil {
			return err
		}

		p.steam.Connect()
	}

	return nil
}

func (p *Manager) onSteamConnect() error { //nolint: unparam
	p.log.WithField("username", p.options.Credentials.Username).Info("connected, logging in")

	logOnDetails := &steam.LogOnDetails{
		Username:               p.options.Credentials.Username,
		Password:               p.options.Credentials.Password,
		AuthCode:               p.options.Credentials.AuthCode,
		TwoFactorCode:          p.options.Credentials.TwoFactorCode,
		SentryFileHash:         steam.SentryHash(p.login.MachineHash),
		ShouldRememberPassword: p.options.Credentials.RememberPassword,
	}

	if logOnDetails.Password == "" {
		logOnDetails.LoginKey = p.login.LoginKey
	}

	p.steam.Auth.LogOn(logOnDetails)

	return nil
}

func (p *Manager) onSteamDisconnect() error {
	p.log.Info("disconnected")
	return ErrSteamDisconnected
}

func (p *Manager) onSteamServerList(e *steam.ClientCMListEvent) error { //nolint: unparam
	p.log.WithField("count", len(e.Addresses)).Debug("received server list")

	if err := p.saveServers(e.Addresses); err != nil {
		return nil
	}

	return nil
}

func (p *Manager) onSteamLoginKey(e *steam.LoginKeyEvent) error { //nolint: unparam
	p.log.Debug("received login key")

	update := &models.SteamLogin{
		UniqueID: e.UniqueId,
		LoginKey: e.LoginKey,
	}

	if err := p.updateLogin(update); err != nil {
		return nil
	}

	return nil
}

func (p *Manager) onSteamMachineAuth(e *steam.MachineAuthUpdateEvent) error { //nolint: unparam
	p.log.Debug("received machine hash")

	update := &models.SteamLogin{MachineHash: e.Hash}

	if err := p.updateLogin(update); err != nil {
		return nil
	}

	return nil
}

func (p *Manager) onSteamWebSession(_ *steam.WebSessionIdEvent) error { //nolint: unparam
	p.log.Debug("received web session id")

	update := &models.SteamLogin{
		WebSessionID: p.steam.Web.SessionId,
	}

	if err := p.updateLogin(update); err != nil {
		return nil
	}

	// p.steam.Web.LogOn()

	return nil
}

func (p *Manager) onSteamWebLogOn(_ *steam.WebLoggedOnEvent) error { //nolint: unparam
	p.log.Debug("web logged on")

	update := &models.SteamLogin{
		WebAuthToken:       p.steam.Web.SteamLogin,
		WebAuthTokenSecure: p.steam.Web.SteamLoginSecure,
	}

	if err := p.updateLogin(update); err != nil {
		return nil
	}

	return nil
}

func (p *Manager) onSteamLogOnFail(e *steam.LogOnFailedEvent) error {
	p.log.
		WithError(ErrSteamLogOnFailed).
		WithField("reason", e.Result.String()).
		Error("steam error")

	return ErrSteamLogOnFailed
}

func (p *Manager) onSteamLogOff(e *steam.LoggedOffEvent) error {
	p.log.
		WithError(ErrSteamLoggedOff).
		WithField("reason", e.Result.String()).
		Error("steam error")

	return ErrSteamLoggedOff
}

func (p *Manager) onSteamLogOn(e *steam.LoggedOnEvent) error {
	p.log.Info("logged in")
	p.steam.Social.SetPersonaName(p.login.Username)
	p.steam.Social.SetPersonaState(steamlang.EPersonaState_Online)

	loginUpdate := &models.SteamLogin{
		SteamID:                   e.ClientSteamId,
		AccountFlags:              uint32(e.AccountFlags),
		WebAuthNonce:              e.Body.GetWebapiAuthenticateUserNonce(),
		CellID:                    e.Body.GetCellId(),
		CellIDPingThreshold:       e.Body.GetCellIdPingThreshold(),
		EmailDomain:               e.Body.GetEmailDomain(),
		VanityURL:                 e.Body.GetVanityUrl(),
		OutOfGameHeartbeatSeconds: e.Body.GetOutOfGameHeartbeatSeconds(),
		InGameHeartbeatSeconds:    e.Body.GetInGameHeartbeatSeconds(),
		PublicIP:                  e.Body.GetPublicIp(),
		ServerTime:                e.Body.GetRtime32ServerTime(),
		SteamTicket:               e.Body.GetSteam2Ticket(),
		UsePics:                   e.Body.GetUsePics(),
		CountryCode:               e.Body.GetIpCountryCode(),
		ParentalSettings:          e.Body.GetParentalSettings(),
		ParentalSettingSignature:  e.Body.GetParentalSettingSignature(),
		LoginFailuresToMigrate:    e.Body.GetCountLoginfailuresToMigrate(),
		DisconnectsToMigrate:      e.Body.GetCountDisconnectsToMigrate(),
		OgsDataReportTimeWindow:   e.Body.GetOgsDataReportTimeWindow(),
		ClientInstanceID:          e.Body.GetClientInstanceId(),
		ForceClientUpdateCheck:    e.Body.GetForceClientUpdateCheck(),
	}

	if err := p.updateLogin(loginUpdate); err != nil {
		return err
	}

	return p.connectDota()
}

func (p *Manager) connectDota() error {
	p.dota.SetPlaying(true)
	p.log.Info("playing Dota 2")

	go p.dotaHello()

	return nil
}

func (p *Manager) dotaHello() {
	if p.ready {
		p.log.Warn("tried to send GC hello while connected")
		return
	}

	p.helloCount = 1
	p.helloTicker = time.NewTicker(helloRetryInterval)
	sessionSub := p.bus.Sub(nsbus.TopicSession)

	defer func() {
		p.helloTicker.Stop()
		p.bus.Unsub(sessionSub)
	}()

	p.dota.SayHello()

	for {
		select {
		case msg := <-sessionSub:
			if msg == nil {
				return
			}

			if s, ok := msg.(*nsbus.SessionChangeMessage); ok && s.IsReady {
				return
			}
		case <-p.helloTicker.C:
			if p.helloCount >= helloRetryCount {
				p.log.WithError(ErrDotaGCWelcomeTimeout).Error("dota GC error")
				p.stop()
				return
			}

			p.helloCount++
			p.dota.SayHello()
		}
	}
}

func (p *Manager) onDotaWelcome(e *events.ClientWelcomed) error {
	if p.helloTicker != nil {
		p.helloTicker.Stop()
	}

	if err := p.saveDotaWelcome(e.Welcome); err != nil {
		return err
	}

	return nil
}

func (p *Manager) onDotaClientSuspended(e *events.ClientSuspended) error {
	until := time.Unix(int64(e.GetTimeEnd()), 0)

	p.log.
		WithError(ErrDotaClientSuspended).
		WithField("until", until).
		Error("dota error")

	p.cancelSession()

	update := &models.SteamLogin{SuspendedUntil: &until}

	if err := p.updateLogin(update); err != nil {
		return err
	}

	return ErrDotaClientSuspended
}

func (p *Manager) onDotaGCStateChange(e events.ClientStateChanged) error { //nolint: unparam
	p.ready = e.NewState.IsReady()

	if !e.OldState.IsReady() && e.NewState.IsReady() {
		p.log.Info("dota connected")
		p.startSession()
	} else if e.OldState.IsReady() && !e.NewState.IsReady() {
		p.log.Warn("dota disconnected")
		p.cancelSession()
		go p.dotaHello()
	}

	return nil
}
