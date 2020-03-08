package middleware

import (
	"time"

	"github.com/labstack/echo/v4"

	nslog "github.com/13k/night-stalker/internal/logger"
)

func Logger(logger *nslog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			l := logger
			req := c.Request()
			res := c.Response()
			start := time.Now()

			if err := next(c); err != nil {
				c.Error(err)
				l = l.WithError(err)
			}

			stop := time.Now()
			latency := stop.Sub(start)

			var id string
			if id = req.Header.Get(echo.HeaderXRequestID); id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			var bytesIn string
			if bytesIn = req.Header.Get(echo.HeaderContentLength); bytesIn == "" {
				bytesIn = "0"
			}

			if id != "" {
				l = l.WithField("id", id)
			}

			l = l.WithOFields(
				"remote_ip", c.RealIP(),
				"host", req.Host,
				"method", req.Method,
				"uri", req.RequestURI,
				"user_agent", req.UserAgent(),
				"status", res.Status,
				"latency", int64(latency),
				"latency_human", latency.String(),
				"bytes_in", bytesIn,
				"bytes_out", res.Size,
			)

			switch {
			case res.Status >= 500:
				l.Error("request")
			case res.Status >= 400:
				l.Warn("request")
			default:
				l.Info("request")
			}

			return nil
		}
	}
}
