package web

import (
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
	nswebws "github.com/13k/night-stalker/web/internal/websocket"
)

var wsUpgrader = ws.Upgrader{}

func (app *App) serveWS(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	if !cc.IsWebSocket() {
		return cc.NoContent(http.StatusNotFound)
	}

	conn, err := wsUpgrader.Upgrade(cc.Response(), cc.Request(), nil)

	if err != nil {
		app.log.WithError(err).Error("error accepting connection")
		return err
	}

	cconn := nswebws.NewConn(conn, nswebws.ConnOptions{
		ReqCtx: cc,
		Log:    app.log,
		Bus:    app.bus,
	})

	cconn.Serve(app.ctx)

	return nil
}
