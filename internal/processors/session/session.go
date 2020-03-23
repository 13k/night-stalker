package session

import (
	"context"
	"fmt"
	"time"

	"github.com/faceit/go-steam"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/state"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	helloRetryCount    = 360
	helloRetryInterval = 10 * time.Second
)

type sessionClosedEvent struct{}

type sessionStateChangeEvent struct {
	Previous       state.Dota2State
	Current        state.Dota2State
	UnreadyToReady bool
	ReadyToUnready bool
}

type sessionOptions struct {
	Log *nslog.Logger
}

type session struct {
	options sessionOptions
	log     *nslog.Logger
	conn    *conn
	steam   *steam.Client
	dota    *dota2.Dota2
	ctx     context.Context
	cancel  context.CancelFunc
	bus     *nsbus.Bus
	ready   bool
}

func newSession(ctx context.Context, conn *conn, options sessionOptions) (*session, error) {
	s := &session{
		options: options,
		log:     options.Log.WithPackage("sess"),
		conn:    conn,
		steam:   conn.Client(),
		dota:    dota2.New(conn.Client(), options.Log.WithPackage("dota2").LogrusLogger()),
		bus: nsbus.New(nsbus.Options{
			Cap:        1,
			Log:        options.Log,
			PubTimeout: 3 * time.Second,
		}),
	}

	s.ctx, s.cancel = context.WithCancel(ctx)

	if err := s.start(); err != nil {
		return nil, xerrors.Errorf("session error: %w", err)
	}

	return s, nil
}

func (s *session) Client() *dota2.Dota2 {
	return s.dota
}

func (s *session) Bus() *nsbus.Bus {
	return s.bus
}

func (s *session) Context() context.Context {
	return s.ctx
}

func (s *session) IsReady() bool {
	return s.ready
}

func (s *session) Close() {
	s.dota.Close()
	s.cancel()
}

func (s *session) pub(ev interface{}) {
	l := s.log.WithField("event", fmt.Sprintf("%T", ev))

	msg := nsbus.Message{Topic: "events", Payload: ev}

	if err := s.bus.Pub(msg); err != nil {
		l.WithError(err).Error("error publishing event")
		return
	}

	l.Trace("published event")
}

func (s *session) start() error {
	if s.ready {
		s.log.Warn("called session.start() when live")
		return nil
	}

	if s.ctx.Err() != nil {
		return xerrors.Errorf("session error: %w", s.ctx.Err())
	}

	go s.loop()

	if s.conn.IsReady() {
		go s.greet()
	}

	return nil
}

func (s *session) teardown() {
	s.ready = false
	s.pub(&sessionClosedEvent{})
	s.bus.Shutdown()
	s.log.Trace("stop")
}

func (s *session) loop() {
	subConn := s.conn.Bus().Sub("events")

	defer func() {
		s.conn.Bus().Unsub(subConn)
		s.teardown()
	}()

	s.log.Trace("start")

	for {
		select {
		case <-s.ctx.Done():
			return
		case busMsg, ok := <-subConn.C:
			if !ok {
				s.log.Warn("connection events subscription closed")
				return
			}

			s.handleEvent(busMsg.Payload)
		}
	}
}

func (s *session) greet() {
	s.dota.SetPlaying(true)
	s.log.Debug("playing Dota 2")

	if err := s.hello(); err != nil {
		s.pub(xerrors.Errorf("dota hello error: %w", err))
		return
	}
}

func (s *session) hello() error {
	retryCount := 0
	t := time.NewTicker(helloRetryInterval)
	subSession := s.bus.Sub("events")

	defer func() {
		t.Stop()
		s.bus.Unsub(subSession)
		s.log.Trace("hello() stop")
	}()

	s.log.Trace("hello() start")

	for {
		s.dota.SayHello()
		retryCount++

		select {
		case <-s.ctx.Done():
			return xerrors.Errorf("session error: %w", s.ctx.Err())
		case <-t.C:
			if retryCount >= helloRetryCount {
				return xerrors.Errorf("session error: %w", &ErrDotaGCWelcomeTimeout{
					RetryCount:    retryCount,
					RetryInterval: helloRetryInterval,
				})
			}
		case busMsg, ok := <-subSession.C:
			if !ok {
				return xerrors.Errorf("session bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: subSession,
				})
			}

			if stateMsg, ok := busMsg.Payload.(*sessionStateChangeEvent); ok && stateMsg.Current.IsReady() {
				return nil
			}
		}
	}
}

func (s *session) handleEvent(ev interface{}) {
	switch e := ev.(type) {
	case *connectionReadyEvent:
		go s.greet()
	case *connectionClosedEvent:
		s.cancel()
	case events.ClientStateChanged:
		s.onDotaGCStateChange(e)
	}
}

func (s *session) onDotaGCStateChange(ev events.ClientStateChanged) {
	busMsg := &sessionStateChangeEvent{
		Previous:       ev.OldState,
		Current:        ev.NewState,
		UnreadyToReady: !ev.OldState.IsReady() && ev.NewState.IsReady(),
		ReadyToUnready: ev.OldState.IsReady() && !ev.NewState.IsReady(),
	}

	s.ready = busMsg.Current.IsReady()
	s.pub(busMsg)

	if busMsg.ReadyToUnready {
		s.cancel()
	}
}
