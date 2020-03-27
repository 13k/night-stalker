package session

import (
	"context"
	"time"

	"cirello.io/oversight"
	d2events "github.com/paralin/go-dota2/events"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nserr "github.com/13k/night-stalker/internal/errors"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nssteam "github.com/13k/night-stalker/internal/steam"
)

const (
	dotaSessionProcessorName = "session/dota"
	helloRetryCount          = 360
	helloRetryInterval       = 10 * time.Second
)

type dotaSessionOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	ShutdownTimeout time.Duration
}

type dotaSessionState struct {
	greeting bool
	closed   bool
	err      error
	t        *nsdota2.StateTransition
}

func (s *dotaSessionState) reset() {
	s.closed = false
	s.err = nil
	s.t = nil
}

var _ nsproc.Processor = (*dotaSession)(nil)

type dotaSession struct {
	options            dotaSessionOptions
	log                *nslog.Logger
	steam              *nssteam.Client
	dota               *nsdota2.Client
	bus                *nsbus.Bus
	ctx                context.Context
	busSubSteamEvents  *nsbus.Subscription
	busSubSessionSteam *nsbus.Subscription
	state              dotaSessionState
}

func newDotaSession(options dotaSessionOptions) *dotaSession {
	return &dotaSession{
		options: options,
		log:     options.Log.WithPackage("dota"),
		bus:     options.Bus,
	}
}

func (s *dotaSession) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if s.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(s.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     dotaSessionProcessorName,
		Start:    s.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (s *dotaSession) Start(ctx context.Context) error {
	s.state.reset()

	if err := s.setupContext(ctx); err != nil {
		return err
	}

	s.busSubscribe()

	if s.steam.Session.IsReady() {
		go s.greet()
	}

	return s.loop()
}

func (s *dotaSession) setupContext(ctx context.Context) error {
	if s.steam = nsctx.GetSteam(ctx); s.steam == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextSteamClient)
	}

	if s.dota = nsctx.GetDota(ctx); s.dota == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaClient)
	}

	s.ctx = ctx

	return nil
}

func (s *dotaSession) busSubscribe() {
	if s.busSubSteamEvents == nil {
		s.busSubSteamEvents = s.bus.Sub(nsbus.TopicSteamEvents)
	}

	if s.busSubSessionSteam == nil {
		s.busSubSessionSteam = s.bus.Sub(nsbus.TopicSessionSteam)
	}
}

func (s *dotaSession) busUnsubscribe() {
	if s.busSubSteamEvents != nil {
		s.bus.Unsub(s.busSubSteamEvents)
		s.busSubSteamEvents = nil
	}

	if s.busSubSessionSteam != nil {
		s.bus.Unsub(s.busSubSessionSteam)
		s.busSubSessionSteam = nil
	}
}

func (s *dotaSession) pub(ev interface{}) error {
	msg := nsbus.Message{
		Topic:   nsbus.TopicSessionDota,
		Payload: ev,
	}

	if err := s.bus.Pub(msg); err != nil {
		return xerrors.Errorf("error publishing event %T: %w", ev, err)
	}

	return nil
}

func (s *dotaSession) greet() {
	if s.state.greeting {
		s.log.Warn("called greet() while greeting")
		return
	}

	s.state.greeting = true
	s.open()

	defer func() {
		s.state.greeting = false
	}()

	if err := s.hello(); err != nil {
		s.cancel(err)
	}
}

func (s *dotaSession) open() {
	s.dota.SetPlaying(true)
	s.log.Debug("playing Dota 2")
	s.state.closed = false
}

func (s *dotaSession) close() {
	s.log.Trace("closing")
	s.dota.Close()
	s.dota.SetPlaying(false)
	s.state.closed = true
}

func (s *dotaSession) cancel(err error) {
	s.state.err = err
	s.close()
}

func (s *dotaSession) stop() {
	s.busUnsubscribe()
	s.log.Trace("stop")
}

func (s *dotaSession) loop() error {
	defer s.stop()

	s.log.Trace("start")

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case busmsg, ok := <-s.busSubSessionSteam.C:
			if !ok {
				s.state.err = xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: s.busSubSessionSteam,
				})

				break
			}

			if steamsessmsg, ok := busmsg.Payload.(*nsbus.SteamSessionChangeMessage); ok {
				s.handleSteamSessionChange(steamsessmsg)
			}
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
			busmsg := &nsbus.DotaSessionChangeMessage{
				StateTransition: t,
				IsReady:         t.Next.IsReady(),
				Err:             s.state.err,
			}

			if err := s.pub(busmsg); err != nil {
				return err
			}

			s.state.t = nil
		}

		if s.state.err != nil {
			if !s.state.closed {
				s.close()
				continue
			}

			return s.state.err
		}
	}
}

func (s *dotaSession) hello() error {
	t := time.NewTicker(helloRetryInterval)
	subSteamSession := s.bus.Sub(nsbus.TopicSessionSteam)
	subDotaSession := s.bus.Sub(nsbus.TopicSessionDota)

	defer func() {
		t.Stop()
		s.bus.Unsub(subSteamSession)
		s.bus.Unsub(subDotaSession)
		s.log.Trace("hello stop")
	}()

	retryCount := 0
	tryHello := func() {
		s.dota.SayHello()
		retryCount++
	}

	s.log.Trace("hello start")

	tryHello()

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-t.C:
			if retryCount >= helloRetryCount {
				return xerrors.Errorf("session error: %w", &nsdota2.ErrWelcomeTimeout{
					RetryCount:    retryCount,
					RetryInterval: helloRetryInterval,
				})
			}

			tryHello()
		case busmsg, ok := <-subSteamSession.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: subSteamSession,
				})
			}

			if sessmsg, ok := busmsg.Payload.(*nsbus.SteamSessionChangeMessage); ok && !sessmsg.IsReady {
				s.log.Warn("steam disconnected while hello()ing")
				return &nsdota2.ErrNoSession{
					Err: nserr.Wrap("no steam session", sessmsg.Err),
				}
			}
		case busmsg, ok := <-subDotaSession.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: subDotaSession,
				})
			}

			if sessmsg, ok := busmsg.Payload.(*nsbus.DotaSessionChangeMessage); ok && sessmsg.IsReady {
				return nil
			}
		}
	}
}

func (s *dotaSession) handleSteamSessionChange(sessmsg *nsbus.SteamSessionChangeMessage) {
	if sessmsg.IsReady {
		go s.greet()
		return
	}

	s.cancel(&nsdota2.ErrNoSession{
		Err: nserr.Wrap("no steam session", sessmsg.Err),
	})
}

func (s *dotaSession) handleSteamEvent(ev interface{}) {
	switch ev := ev.(type) {
	case *d2events.ClientSuspended:
		s.onClientSuspended(ev)
	case d2events.ClientStateChanged:
		s.onClientStateChange(ev)
	}
}

func (s *dotaSession) onClientSuspended(ev *d2events.ClientSuspended) {
	s.cancel(&nsdota2.ErrClientSuspended{
		Until: time.Unix(int64(ev.GetTimeEnd()), 0),
	})
}

func (s *dotaSession) onClientStateChange(ev d2events.ClientStateChanged) {
	s.state.t = nsdota2.NewStateTransition(ev.OldState, ev.NewState)

	s.dota.Session.StateChange(s.state.t)

	if s.state.t.ReadyToUnready {
		err := s.state.err

		if err == nil {
			err = &nsdota2.ErrLostSession{Status: s.state.t.Next.ConnectionStatus}
		}

		s.state.err = &nsdota2.ErrNoSession{
			Err: nserr.Wrap("session error", err),
		}
	}
}
