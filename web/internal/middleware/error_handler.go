package middleware

import (
	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*nswebctx.Context)

			err := next(cc)

			if err != nil {
				cc.Error(err)
			}

			return err
		}
	}
}
