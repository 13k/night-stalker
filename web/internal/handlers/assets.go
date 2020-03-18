package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func AssetHandler(fs http.FileSystem) echo.HandlerFunc {
	httpHandler := http.FileServer(fs)
	echoHandler := echo.WrapHandler(httpHandler)

	return func(c echo.Context) error {
		cc := c.(*nswebctx.Context)
		cc.SkipLogging(true)

		return echoHandler(cc)
	}
}
