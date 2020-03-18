package middleware

import (
	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func MediaType() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*nswebctx.Context)

			if err := cc.ValidateMediaType(); err != nil {
				return err
			}

			return next(c)
		}
	}
}
