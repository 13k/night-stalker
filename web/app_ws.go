package web

import (
	"encoding/json"
	"net/http"

	"github.com/docker/go-units"
	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

var wsUpgrader = ws.Upgrader{}

func (app *App) serveWS(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	if !cc.IsWebSocket() {
		return cc.NoContent(http.StatusNotFound)
	}

	conn, err := wsUpgrader.Upgrade(cc.Response(), cc.Request(), nil)

	if err != nil {
		app.wslog.WithError(err).Error("error accepting connection")
		return err
	}

	go app.serveWSConn(cc, conn)

	return nil
}

func (app *App) serveWSConn(c *nswebctx.Context, conn *ws.Conn) {
	l := app.wslog.WithField("id", c.ID())

	busSubLiveMatches := app.bus.Sub(nsbus.TopicWebPatternLiveMatchesAll)
	connClosed := make(chan bool)

	defer func() {
		app.bus.Unsub(busSubLiveMatches)
		conn.Close()
		l.Debug("stop")
	}()

	l.Debug("client connected")

	go func() {
		for {
			if _, _, err := conn.NextReader(); err != nil {
				close(connClosed)
				return
			}
		}
	}()

	for {
		select {
		case <-connClosed:
			l.Debug("connection closed")
			return
		case <-app.ctx.Done():
			l.Debug("canceled")
			return
		case busmsg, ok := <-busSubLiveMatches.C:
			if !ok {
				l.Debug("live matches channel closed")
				return
			}

			if msg, ok := busmsg.Payload.(*nspb.LiveMatchesChange); ok {
				app.pushWSLiveMatches(c, conn, msg)
			}
		}
	}
}

func (app *App) pushWSLiveMatches(c *nswebctx.Context, conn *ws.Conn, msg *nspb.LiveMatchesChange) {
	l := app.wslog.WithOFields(
		"id", c.ID(),
		"op", msg.Op.String(),
		"count", len(msg.Change.Matches),
	)

	wsmsg, err := json.Marshal(msg)

	if err != nil {
		l.WithError(err).Error("error serializing message")
		return
	}

	if err = conn.WriteMessage(ws.TextMessage, wsmsg); err != nil {
		l.WithError(err).Error("error sending message")
		return
	}

	bytesOut := len(wsmsg)

	l.WithOFields(
		"tx", bytesOut,
		"tx_h", units.BytesSize(float64(bytesOut)),
	).Debug("sent message")
}
