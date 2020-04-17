package websocket

import (
	"context"

	"github.com/docker/go-units"
	ws "github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

type ConnOptions struct {
	ReqCtx *nswebctx.Context
	Log    *nslog.Logger
	Bus    *nsbus.Bus
}

type Conn struct {
	*ws.Conn

	options ConnOptions
	ctx     context.Context
	cancel  context.CancelFunc
	c       *nswebctx.Context
	log     *nslog.Logger
	bus     *nsbus.Bus
}

func NewConn(conn *ws.Conn, options ConnOptions) *Conn {
	return &Conn{
		Conn:    conn,
		options: options,
		c:       options.ReqCtx,
		log:     options.Log.WithPackage("ws").WithField("id", options.ReqCtx.ID()),
		bus:     options.Bus,
	}
}

func (c *Conn) Serve(ctx context.Context) {
	c.ctx, c.cancel = context.WithCancel(ctx)

	go c.rx()
	go c.tx()
}

func (c *Conn) rx() {
	defer c.cancel()

	for {
		if _, _, err := c.NextReader(); err != nil {
			return
		}
	}
}

func (c *Conn) tx() {
	busSubLiveMatches := c.bus.Sub(nsbus.TopicWebPatternLiveMatchesAll)

	defer func() {
		c.bus.Unsub(busSubLiveMatches)
		c.Close()
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		case busmsg, ok := <-busSubLiveMatches.C:
			if !ok {
				c.log.Warn("live matches channel closed")
				return
			}

			if msg, ok := busmsg.Payload.(*nspb.LiveMatchesChange); ok {
				c.handleLiveMatchesChange(msg)
			}
		}
	}
}

func (c *Conn) handleLiveMatchesChange(msg *nspb.LiveMatchesChange) {
	l := c.log.WithOFields(
		"op", msg.Op.String(),
		"count", len(msg.Change.Matches),
	)

	wsmsg, err := proto.Marshal(msg)

	if err != nil {
		l.WithError(err).Error("error serializing message")
		return
	}

	if err = c.WriteMessage(ws.BinaryMessage, wsmsg); err != nil {
		l.WithError(err).Error("error sending message")
		return
	}

	bytesOut := proto.Size(msg)

	l.WithOFields(
		"tx", bytesOut,
		"tx_h", units.BytesSize(float64(bytesOut)),
	).Debug("sent message")
}
