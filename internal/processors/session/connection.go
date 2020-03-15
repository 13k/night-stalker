package session

import (
	"context"
	"fmt"
	"time"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/netutil"
	"github.com/faceit/go-steam/protocol/steamlang"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
)

type connectionReadyEvent struct{}
type connectionClosedEvent struct{}

type connOptions struct {
	Log         *nslog.Logger
	Address     *netutil.PortAddr
	Credentials *Credentials
	MachineHash steam.SentryHash
	LoginKey    string
}

type conn struct {
	options connOptions
	log     *nslog.Logger
	steam   *steam.Client
	ctx     context.Context
	cancel  context.CancelFunc
	bus     *nsbus.Bus
	ready   bool
}

func newConnection(ctx context.Context, options connOptions) (*conn, error) {
	c := &conn{
		options: options,
		log:     options.Log.WithPackage("conn"),
		steam:   steam.NewClient(),
		bus: nsbus.New(nsbus.Options{
			Cap:        1,
			Log:        options.Log,
			PubTimeout: 3 * time.Second,
		}),
	}

	c.ctx, c.cancel = context.WithCancel(ctx)

	if err := c.connect(); err != nil {
		return nil, xerrors.Errorf("connection error: %w", err)
	}

	return c, nil
}

func (c *conn) Client() *steam.Client {
	return c.steam
}

func (c *conn) Bus() *nsbus.Bus {
	return c.bus
}

func (c *conn) IsReady() bool {
	return c.ready
}

func (c *conn) Close(ctx context.Context) error {
	return c.disconnect(ctx)
}

func (c *conn) pub(ev interface{}) {
	l := c.log.WithField("event", fmt.Sprintf("%T", ev))

	msg := nsbus.Message{Topic: "events", Payload: ev}

	if err := c.bus.Pub(msg); err != nil {
		l.WithError(err).Error("error publishing event")
		return
	}

	l.Trace("published event")
}

func (c *conn) connect() error {
	if c.ready {
		c.log.Warn("called conn.connect() when connected")
		return nil
	}

	if c.ctx.Err() != nil {
		return xerrors.Errorf("connection error: %w", c.ctx.Err())
	}

	go c.loop()

	if c.options.Address != nil {
		c.steam.ConnectTo(c.options.Address)
	} else {
		if err := steam.InitializeSteamDirectory(); err != nil {
			return xerrors.Errorf("error initializing steam directory: %w", err)
		}

		c.steam.Connect()
	}

	return nil
}

func (c *conn) disconnect(ctx context.Context) error {
	if !c.steam.Connected() {
		c.log.Warn("called conn.disconnect() when disconnected")
		return nil
	}

	c.steam.Disconnect()
	c.ready = false

	select {
	case <-ctx.Done():
		return xerrors.Errorf("disconnect error: %w", ctx.Err())
	case <-c.ctx.Done():
		return nil
	}
}

func (c *conn) teardown() {
	c.ready = false
	c.pub(&connectionClosedEvent{})
	c.bus.Shutdown()
	c.log.Trace("stop")
}

func (c *conn) loop() {
	defer c.teardown()

	c.log.Trace("start")

	for {
		select {
		case <-c.ctx.Done():
			return
		case ev, ok := <-c.steam.Events():
			if !ok {
				c.log.Warn("steam events closed")
				return
			}

			c.handleEvent(ev)
		}
	}
}

func (c *conn) handleEvent(ev interface{}) {
	switch e := ev.(type) {
	case *steam.ConnectedEvent:
		c.onSteamConnect(e)
	case *steam.LogOnFailedEvent:
		c.onSteamLogOnFail(e)
	case *steam.LoggedOnEvent:
		c.onSteamLogOn(e)
	case *steam.WebSessionIdEvent:
		c.onSteamWebSession(e)
	case *steam.WebLoggedOnEvent:
		c.onSteamWebLogOn(e)
	case *steam.LoggedOffEvent:
		c.onSteamLogOff(e)
	case *steam.DisconnectedEvent:
		c.onSteamDisconnect(e)
	case steam.FatalErrorEvent:
		c.onFatalError(e)
	default:
		c.pub(ev)
	}
}

func (c *conn) onSteamConnect(_ *steam.ConnectedEvent) {
	c.log.
		WithField("username", c.options.Credentials.Username).
		Info("connected, logging in")

	logOnDetails := &steam.LogOnDetails{
		Username:               c.options.Credentials.Username,
		Password:               c.options.Credentials.Password,
		AuthCode:               c.options.Credentials.AuthCode,
		TwoFactorCode:          c.options.Credentials.TwoFactorCode,
		SentryFileHash:         c.options.MachineHash,
		ShouldRememberPassword: c.options.Credentials.RememberPassword,
	}

	if logOnDetails.Password == "" {
		logOnDetails.LoginKey = c.options.LoginKey
	}

	c.steam.Auth.LogOn(logOnDetails)
}

func (c *conn) onSteamLogOnFail(ev *steam.LogOnFailedEvent) {
	err := xerrors.Errorf("steam error: %w", &ErrSteamLogOnFailed{
		Reason: ev.Result.String(),
	})

	c.pub(err)
}

func (c *conn) onSteamLogOn(ev *steam.LoggedOnEvent) {
	c.steam.Social.SetPersonaName(c.options.Credentials.Username)
	c.steam.Social.SetPersonaState(steamlang.EPersonaState_Online)

	c.ready = true
	c.pub(ev)
	c.pub(&connectionReadyEvent{})
}

func (c *conn) onSteamWebSession(_ *steam.WebSessionIdEvent) {
	c.pub(&SteamWebSessionIDEvent{
		SessionID: c.steam.Web.SessionId,
	})

	c.steam.Web.LogOn()
}

func (c *conn) onSteamWebLogOn(_ *steam.WebLoggedOnEvent) {
	c.pub(&SteamWebLoggedOnEvent{
		AuthToken:  c.steam.Web.SteamLogin,
		AuthSecret: c.steam.Web.SteamLoginSecure,
	})
}

func (c *conn) onSteamLogOff(ev *steam.LoggedOffEvent) {
	c.pub(ev)

	if err := c.disconnect(context.Background()); err != nil {
		c.log.Errorx(xerrors.Errorf("error disconnecting after logoff: %w", err))
	}
}

func (c *conn) onSteamDisconnect(ev *steam.DisconnectedEvent) {
	c.pub(ev)
	c.cancel()
}

func (c *conn) onFatalError(ev steam.FatalErrorEvent) {
	c.pub(ev)
	c.cancel()
}
