package middleware

import (
	"time"

	"github.com/docker/go-units"
	"github.com/labstack/echo/v4"

	nslog "github.com/13k/night-stalker/internal/logger"
	nsstrconv "github.com/13k/night-stalker/internal/strconv"
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

			var bytesIn uint64
			if h := req.Header.Get(echo.HeaderContentLength); h != "" {
				bytesIn = nsstrconv.SafeParseUint(h)
			}

			if id != "" {
				l = l.WithField("id", id)
			}

			l = l.WithOFields(
				"status", res.Status,
				"remote_ip", c.RealIP(),
				"host", req.Host,
				"method", req.Method,
				"uri", req.RequestURI,
				"user_agent", req.UserAgent(),
				"latency", int64(latency),
				"latency_h", latency.String(),
				"rx", bytesIn,
				"rx_h", units.BytesSize(float64(bytesIn)),
				"tx", res.Size,
				"tx_h", units.BytesSize(float64(res.Size)),
			)

			switch {
			case res.Status >= 500:
				l.Errorz()
			case res.Status >= 400:
				l.Warnz()
			default:
				l.Infoz()
			}

			return nil
		}
	}
}
