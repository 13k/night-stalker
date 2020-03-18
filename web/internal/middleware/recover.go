package middleware

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

var (
	recoverConfig = mw.RecoverConfig{
		DisableStackAll: true,
	}
)

func Recover() echo.MiddlewareFunc {
	return mw.RecoverWithConfig(recoverConfig)
}
