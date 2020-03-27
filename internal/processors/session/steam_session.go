package session

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	"github.com/faceit/go-steam/protocol/steamlang"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nserr "github.com/13k/night-stalker/internal/errors"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nssteam "github.com/13k/night-stalker/internal/steam"
)

const (
	steamSessionProcessorName = "session/steam"
)

type steamSessionOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	Addr            *netutil.PortAddr
	Credentials     *Credentials
	MachineHash     steam.SentryHash
	LoginKey        string
	ShutdownTimeout time.Duration
}

type steamSessionState struct {
	closed    bool
	waitClose bool
	err       error
	t         *nssteam.StateTransition
}

func (s *steamSessionState) reset() {
	s.closed = false
	s.waitClose = false
	s.err = nil
	s.t = nil
}

var _ nsproc.Processor = (*steamSession)(nil)

type steamSession struct {
	options           steamSessionOptions
	log               *nslog.Logger
	bus               *nsbus.Bus
	steam             *nssteam.Client
	ctx               context.Context
	busSubSteamEvents *nsbus.Subscription
	state             steamSessionState
}

func newSteamSession(options steamSessionOptions) *steamSession {
	return &steamSession{
		options: options,
		log:     options.Log.WithPackage("steam"),
		bus:     options.Bus,
	}
}

func (s *steamSession) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if s.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(s.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     steamSessionProcessorName,
		Start:    s.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (s *steamSession) Start(ctx context.Context) error {
	s.state.reset()

	if err := s.setupContext(ctx); err != nil {
		return err
	}

	s.busSubscribe()

	if err := s.connect(); err != nil {
		return xerrors.Errorf("connection error: %w", err)
	}

	return s.loop()
}

func (s *steamSession) setupContext(ctx context.Context) error {
	if s.steam = nsctx.GetSteam(ctx); s.steam == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextSteamClient)
	}

	s.ctx = ctx

	return nil
}

func (s *steamSession) busSubscribe() {
	if s.busSubSteamEvents == nil {
		s.busSubSteamEvents = s.bus.Sub(nsbus.TopicSteamEvents)
	}
}

func (s *steamSession) busUnsubscribe() {
	if s.busSubSteamEvents != nil {
		s.bus.Unsub(s.busSubSteamEvents)
		s.busSubSteamEvents = nil
	}
}

func (s *steamSession) connect() error {
	if s.ctx.Err() != nil {
		return xerrors.Errorf("connection error: %w", s.ctx.Err())
	}

	if s.options.Addr != nil {
		s.log.WithField("server", s.options.Addr.String()).Trace("connecting")
		s.steam.ConnectTo(s.options.Addr)
	} else {
		s.log.Trace("initializing steam directory")

		if err := steam.InitializeSteamDirectory(); err != nil {
			return xerrors.Errorf("error initializing steam directory: %w", err)
		}

		s.log.Trace("connecting to random server")
		s.steam.Connect()
	}

	s.state.closed = false
	return nil
}

func (s *steamSession) disconnect() {
	s.log.Trace("disconnecting")
	s.steam.Disconnect()
	s.state.closed = true
}

func (s *steamSession) pub(ev interface{}) error {
	msg := nsbus.Message{
		Topic:   nsbus.TopicSessionSteam,
		Payload: ev,
	}

	if err := s.bus.Pub(msg); err != nil {
		return xerrors.Errorf("error publishing event %T: %w", ev, err)
	}

	return nil
}

func (s *steamSession) cancel(err error) {
	s.state.err = err
	s.disconnect()
}

func (s *steamSession) stop() {
	s.busUnsubscribe()
	s.log.Trace("stop")
}

func (s *steamSession) loop() error {
	defer s.stop()

	s.log.Trace("start")

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case busmsg, ok := <-s.busSubSteamEvents.C:
			if !ok {
				s.state.err = xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: s.busSubSteamEvents,
				})

				break
			}

			if steammsg, ok := busmsg.Payload.(*nsbus.SteamEventMessage); ok {
				s.handleSteamEvent(steammsg.Event)
			}
		}

		if t := s.state.t; t != nil {
			busmsg := &nsbus.SteamSessionChangeMessage{
				StateTransition: t,
				IsReady:         t.Next.Ready,
				Err:             s.state.err,
			}

			if err := s.pub(busmsg); err != nil {
				return err
			}

			s.state.t = nil
		}

		if s.state.err != nil {
			if s.state.waitClose {
				s.state.waitClose = false
				continue
			}

			if !s.state.closed {
				s.disconnect()
				continue
			}

			return s.state.err
		}
	}
}

func (s *steamSession) handleSteamEvent(ev interface{}) {
	switch ev := ev.(type) {
	case *steam.ConnectedEvent:
		s.onSteamConnect(ev)
	case *steam.LogOnFailedEvent:
		s.onSteamLogOnFail(ev)
	case *steam.LoggedOnEvent:
		s.onSteamLogOn(ev)
	case *steam.WebSessionIdEvent:
		s.onSteamWebSession(ev)
	case *steam.LoggedOffEvent:
		s.onSteamLogOff(ev)
	case *steam.DisconnectedEvent:
		s.onSteamDisconnect(ev)
	case *steam.SteamFailureEvent:
		s.onSteamFailure(ev)
	case steam.FatalErrorEvent:
		s.onFatalError(ev)
	case *nssteam.ClientStateChanged:
		s.onClientStateChange(ev)
	}
}

func (s *steamSession) onSteamConnect(_ *steam.ConnectedEvent) {
	s.log.
		WithField("username", s.options.Credentials.Username).
		Info("connected, logging in")

	logOnDetails := &steam.LogOnDetails{
		Username:               s.options.Credentials.Username,
		Password:               s.options.Credentials.Password,
		AuthCode:               s.options.Credentials.AuthCode,
		TwoFactorCode:          s.options.Credentials.TwoFactorCode,
		SentryFileHash:         s.options.MachineHash,
		ShouldRememberPassword: s.options.Credentials.RememberPassword,
	}

	if logOnDetails.Password == "" {
		logOnDetails.LoginKey = s.options.LoginKey
	}

	s.steam.Auth.LogOn(logOnDetails)
}

func (s *steamSession) onSteamLogOnFail(ev *steam.LogOnFailedEvent) {
	s.cancel(&nssteam.ErrLogOnFailed{
		Reason: ev.Result.String(),
	})
}

func (s *steamSession) onSteamLogOn(_ *steam.LoggedOnEvent) {
	s.steam.Social.SetPersonaName(s.options.Credentials.Username)
	s.steam.Social.SetPersonaState(steamlang.EPersonaState_Online)

	t := nssteam.NewStateTransition(
		nssteam.State{Ready: false},
		nssteam.State{Ready: true},
	)

	s.steam.Session.StateChange(t)
}

func (s *steamSession) onSteamWebSession(_ *steam.WebSessionIdEvent) {
	// s.steam.Web.LogOn()
}

func (s *steamSession) onSteamLogOff(ev *steam.LoggedOffEvent) {
	s.log.Warn("logged off")

	s.state.err = &nssteam.ErrLoggedOff{
		Reason: ev.Result.String(),
	}

	t := nssteam.NewStateTransition(
		nssteam.State{Ready: true},
		nssteam.State{Ready: false},
	)

	s.steam.Session.StateChange(t)
}

func (s *steamSession) onSteamFailure(ev *steam.SteamFailureEvent) {
	s.log.WithField("reason", ev.Result.String()).Error("steam failure")

	s.state.waitClose = true
	s.state.err = &nssteam.ErrFailure{
		Reason: ev.Result.String(),
	}
}

func (s *steamSession) onFatalError(ev steam.FatalErrorEvent) {
	s.log.WithError(ev).Error("steam fatal error")

	s.state.waitClose = true
	s.state.err = ev
}

func (s *steamSession) onSteamDisconnect(_ *steam.DisconnectedEvent) {
	s.log.Error("steam disconnected")

	if s.state.err == nil {
		s.state.err = &nssteam.ErrDisconnected{}
	}

	s.state.waitClose = true

	t := nssteam.NewStateTransition(
		nssteam.State{Ready: true},
		nssteam.State{Ready: false},
	)

	s.steam.Session.StateChange(t)
}

func (s *steamSession) onClientStateChange(ev *nssteam.ClientStateChanged) {
	s.state.t = ev.Transition

	if s.state.t.ReadyToUnready {
		err := s.state.err

		if err == nil {
			err = &nssteam.ErrLostSession{}
		}

		s.state.closed = true
		s.state.err = &nssteam.ErrNoSession{
			Err: nserr.Wrap("session error", err),
		}
	}
}
