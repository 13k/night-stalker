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
		app.wslog.WithError(err).Error("error accepting connection")
		return err
	}

	cconn := nswebws.NewConn(conn, nswebws.ConnOptions{
		Ctx:    app.ctx,
		ReqCtx: cc,
		Log:    app.wslog,
		Bus:    app.bus,
	})

	go cconn.Serve()

	return nil
}
