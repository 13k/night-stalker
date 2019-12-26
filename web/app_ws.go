package web

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

var wsUpgrader = ws.Upgrader{}

func (app *App) serveWS(c echo.Context) error {
	if !c.IsWebSocket() {
		return c.NoContent(http.StatusNotFound)
	}

	conn, err := wsUpgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		app.wslog.WithError(err).Error("error accepting connection")
		return err
	}

	go app.serveWSConn(conn)

	return nil
}

func (app *App) serveWSConn(conn *ws.Conn) {
	id := uuid.New()
	l := app.wslog.WithField("id", id.String())

	liveMatchesSub := app.bus.Sub(nsbus.TopicLiveMatches)
	connClosed := make(chan bool)

	defer func() {
		app.bus.Unsub(liveMatchesSub)
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
		case busmsg, ok := <-liveMatchesSub:
			if !ok {
				l.Debug("live matches channel closed")
				return
			}

			if msg, ok := busmsg.(*nsbus.LiveMatchesProtoMessage); ok {
				wsmsg, err := json.Marshal(msg.Matches)

				if err != nil {
					l.WithError(err).Error("error serializing message")
					continue
				}

				if err = conn.WriteMessage(ws.TextMessage, wsmsg); err != nil {
					l.WithError(err).Error("error sending message")
					continue
				}

				l.Debug("sent message")
			}
		}
	}
}
