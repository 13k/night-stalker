package web

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var wsUpgrader = ws.Upgrader{}

type WSConn struct {
	*ws.Conn

	id uuid.UUID
}

func (app *App) serveWS(c echo.Context) error {
	if !c.IsWebSocket() {
		return c.NoContent(http.StatusNotFound)
	}

	conn, err := wsUpgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		app.wslog.WithError(err).Error("error accepting connection")
		return err
	}

	wsConn := &WSConn{
		Conn: conn,
		id:   uuid.New(),
	}

	go app.serveWSConn(wsConn)

	return nil
}

func (app *App) serveWSConn(conn *WSConn) {
	l := app.wslog.WithField("id", conn.id.String())

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
				app.pushWSLiveMatches(conn, msg)
			}
		}
	}
}

func (app *App) pushWSLiveMatches(conn *WSConn, msg *nspb.LiveMatchesChange) {
	l := app.wslog.WithOFields(
		"id", conn.id.String(),
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

	l.WithField("bytes_out", len(wsmsg)).Debug("sent message")
}
