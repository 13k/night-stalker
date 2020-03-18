package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func Context() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(nswebctx.NewContext(uuid.New(), c))
		}
	}
}
