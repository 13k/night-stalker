package session

import (
	"context"
	"fmt"
	"time"

	"cirello.io/oversight"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	nssteam "github.com/13k/night-stalker/internal/steam"
	nsm "github.com/13k/night-stalker/models"
)

const (
	processorName = "session"
)

type ManagerOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	Credentials     *Credentials
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Manager)(nil)

type Manager struct {
	options            ManagerOptions
	login              *nsm.SteamLogin
	log                *nslog.Logger
	bus                *nsbus.Bus
	db                 *nsdb.DB
	steam              *nssteam.Client
	dota               *nsdota2.Client
	ctx                context.Context
	cancel             context.CancelFunc
	supervisor         *supervisor
	busSubSteamEvents  *nsbus.Subscription
	busSubSessionSteam *nsbus.Subscription
	busSubSessionDota  *nsbus.Subscription
	err                error
}

func NewManager(options ManagerOptions) *Manager {
	return &Manager{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
		login:   &nsm.SteamLogin{},
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
		Restart:  oversight.Transient(),
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
	p.err = nil

	if err := p.setup(ctx); err != nil {
		return err
	}

	go p.loop()

	p.err = p.supervisor.Start(p.ctx)

	<-p.ctx.Done()
	p.ctx = nil

	return p.err
}

func (p *Manager) stop() {
	p.busUnsubscribe()
	p.cancel()
	p.log.Warn("stop")
}

func (p *Manager) loop() {
	defer p.stop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return
		case busmsg, ok := <-p.busSubSessionSteam.C:
			if !ok {
				p.err = xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubSessionSteam,
				})

				return
			}

			if sessmsg, ok := busmsg.Payload.(*nsbus.SteamSessionChangeMessage); ok {
				p.handleSteamSessionChange(sessmsg)
			}
		case busmsg, ok := <-p.busSubSessionDota.C:
			if !ok {
				p.err = xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubSessionDota,
				})

				return
			}

			if sessmsg, ok := busmsg.Payload.(*nsbus.DotaSessionChangeMessage); ok {
				p.handleDotaSessionChange(sessmsg)
			}
		case busmsg, ok := <-p.busSubSteamEvents.C:
			if !ok {
				p.err = xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubSteamEvents,
				})

				return
			}

			if steammsg, ok := busmsg.Payload.(*nsbus.SteamEventMessage); ok {
				if err := p.handleSteamEvent(steammsg.Event); err != nil {
					p.err = xerrors.Errorf("error handling event %T: %w", steammsg.Event, err)
					return
				}
			}
		}
	}
}

func (p *Manager) handleError(err error) {
	msg := fmt.Sprintf("%s error", processorName)
	l := p.log

	if e := (&ErrInvalidServerAddress{}); xerrors.As(err, &e) {
		msg = "invalid server address"
		l = l.WithField("address", e.Address)
	} else if e := (&nssteam.ErrLogOnFailed{}); xerrors.As(err, &e) {
		msg = "steam logon failed"
		l = l.WithField("reason", e.Reason)
	} else if e := (&nsdota2.ErrClientSuspended{}); xerrors.As(err, &e) {
		msg = "dota client suspended"
		l = l.WithField("until", e.Until)
	} else if e := (&nsdota2.ErrWelcomeTimeout{}); xerrors.As(err, &e) {
		msg = "dota welcome timeout"
		l = l.WithOFields(
			"retry_count", e.RetryCount,
			"retry_interval", e.RetryInterval,
		)
	}

	l.WithError(err).Error(msg)
}
