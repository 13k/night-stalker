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
	Ctx    context.Context
	ReqCtx *nswebctx.Context
	Log    *nslog.Logger
	Bus    *nsbus.Bus
}

type Conn struct {
	*ws.Conn

	options ConnOptions
	ctx     context.Context
	c       *nswebctx.Context
	log     *nslog.Logger
	bus     *nsbus.Bus
}

func NewConn(conn *ws.Conn, options ConnOptions) *Conn {
	return &Conn{
		Conn:    conn,
		options: options,
		ctx:     options.Ctx,
		c:       options.ReqCtx,
		log:     options.Log.WithField("id", options.ReqCtx.ID()),
		bus:     options.Bus,
	}
}

func (c *Conn) Serve() {
	busSubLiveMatches := c.bus.Sub(nsbus.TopicWebPatternLiveMatchesAll)
	connClosed := make(chan bool)

	defer func() {
		c.bus.Unsub(busSubLiveMatches)
		c.Close()
		c.log.Debug("stop")
	}()

	c.log.Debug("client connected")

	go func() {
		for {
			if _, _, err := c.NextReader(); err != nil {
				close(connClosed)
				return
			}
		}
	}()

	for {
		select {
		case <-connClosed:
			c.log.Debug("connection closed")
			return
		case <-c.ctx.Done():
			c.log.Debug("canceled")
			return
		case busmsg, ok := <-busSubLiveMatches.C:
			if !ok {
				c.log.Debug("live matches channel closed")
				return
			}

			if msg, ok := busmsg.Payload.(*nspb.LiveMatchesChange); ok {
				c.pushWSLiveMatches(msg)
			}
		}
	}
}

func (c *Conn) pushWSLiveMatches(msg *nspb.LiveMatchesChange) {
	l := c.log.WithOFields(
		"op", msg.Op.String(),
		"count", len(msg.Change.Matches),
	)

	// wsmsg, err := json.Marshal(msg)
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
