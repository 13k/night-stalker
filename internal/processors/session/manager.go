package session

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "session"
)

type ManagerOptions struct {
	Log                   *nslog.Logger
	Bus                   *nsbus.Bus
	Credentials           *Credentials
	ShutdownTimeout       time.Duration
	TVGamesInterval       time.Duration
	RealtimeStatsPoolSize int
	RealtimeStatsInterval time.Duration
	MatchInfoPoolSize     int
	MatchInfoInterval     time.Duration
}

var _ nsproc.Processor = (*Manager)(nil)

type Manager struct {
	options    ManagerOptions
	login      *models.SteamLogin
	ctx        context.Context
	log        *nslog.Logger
	bus        *nsbus.Bus
	db         *gorm.DB
	conn       *conn
	session    *session
	supervisor *supervisor
}

func NewManager(options ManagerOptions) *Manager {
	return &Manager{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
		login:   &models.SteamLogin{},
	}
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
		Start:    p.Start,
		Restart:  oversight.Permanent(),
		Shutdown: shutdown,
	}
}

func (p *Manager) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Manager) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	if err := p.loadLogin(); err != nil {
		return xerrors.Errorf("error loading login info: %w", err)
	}

	if p.isSuspended() {
		return xerrors.Errorf("fatal error: %w", &ErrDotaClientSuspended{
			Until: p.login.SuspendedUntil,
		})
	}

	return p.loop()
}

func (p *Manager) stop() {
	p.log.Warn("stopping...")
	p.disconnect()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Manager) loop() error {
	defer p.stop()

	var subConn *nsbus.Subscription
	var subSession *nsbus.Subscription

	defer func() {
		if p.conn != nil && subConn != nil {
			p.conn.Bus().Unsub(subConn)
		}

		if p.session != nil && subSession != nil {
			p.session.Bus().Unsub(subSession)
		}
	}()

	p.log.Info("start")

	// FIXME: connect/startSession must run concurrently with loop

	for {
		if p.conn == nil {
			if err := p.connect(); err != nil {
				return xerrors.Errorf("error creating connection: %w", err)
			}

			subConn = nil

			if p.session != nil {
				p.closeSession()
			}
		}

		if p.session == nil {
			if err := p.startSession(); err != nil {
				return xerrors.Errorf("error creating session: %w", err)
			}

			subSession = nil
		}

		if subConn == nil {
			subConn = p.conn.Bus().Sub("events")
		}

		if subSession == nil {
			subSession = p.session.Bus().Sub("events")
		}

		select {
		case <-p.ctx.Done():
			return nil
		case busMsg, ok := <-subConn.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: subConn,
				})
			}

			if _, ok := busMsg.Payload.(*connectionClosedEvent); ok {
				p.log.Warn("connection closed")
				p.disconnect()
				continue
			}

			if err := p.handleEvent(busMsg.Payload); err != nil {
				return err
			}
		case busMsg, ok := <-subSession.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: subSession,
				})
			}

			if _, ok := busMsg.Payload.(*sessionClosedEvent); ok {
				p.log.Warn("session closed")
				p.closeSession()
				continue
			}

			if stateMsg, ok := busMsg.Payload.(*sessionStateChangeEvent); ok && stateMsg.UnreadyToReady {
				p.startSupervisor()
			}
		}
	}
}

func (p *Manager) busPubEvent(ev interface{}) error {
	return p.bus.Pub(nsbus.Message{
		Topic:   nsbus.TopicSteamEvents,
		Payload: &nsbus.SteamEventMessage{Event: ev},
	})
}

func (p *Manager) handleError(err error) {
	l := p.log

	if e := (&ErrInvalidServerAddress{}); xerrors.As(err, &e) {
		l = l.WithField("address", e.Address)
	} else if e := (&ErrSteamDisconnected{}); xerrors.As(err, &e) {
	} else if e := (&ErrSteamLogOnFailed{}); xerrors.As(err, &e) {
		l = l.WithField("reason", e.Reason)
	} else if e := (&ErrSteamLoggedOff{}); xerrors.As(err, &e) {
		l = l.WithField("reason", e.Reason)
	} else if e := (&ErrDotaClientSuspended{}); xerrors.As(err, &e) {
		l = l.WithField("until", e.Until)
	} else if e := (&ErrDotaGCWelcomeTimeout{}); xerrors.As(err, &e) {
		l = l.WithOFields(
			"retry_count", e.RetryCount,
			"retry_interval", e.RetryInterval,
		)
	}

	l.WithError(err).Error("session error")
	p.log.Errorx(err)
}
