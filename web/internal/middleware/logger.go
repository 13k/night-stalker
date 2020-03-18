package middleware

import (
	"time"

	"github.com/docker/go-units"
	"github.com/labstack/echo/v4"

	nslog "github.com/13k/night-stalker/internal/logger"
	nsstrconv "github.com/13k/night-stalker/internal/strconv"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func Logger(logger *nslog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*nswebctx.Context)
			req := c.Request()
			res := c.Response()
			start := time.Now()

			err := next(cc)

			if cc.SkipLogging() {
				return err
			}

			stop := time.Now()
			latency := stop.Sub(start)

			var requestID string
			if requestID = req.Header.Get(echo.HeaderXRequestID); requestID == "" {
				requestID = res.Header().Get(echo.HeaderXRequestID)
			}

			var bytesIn uint64
			if h := req.Header.Get(echo.HeaderContentLength); h != "" {
				bytesIn = nsstrconv.SafeParseUint(h)
			}

			l := logger.WithOFields(
				"id", cc.ID(),
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

			if requestID != "" {
				l = l.WithField("request_id", requestID)
			}

			if err != nil {
				l = l.WithError(err)
			}

			switch {
			case res.Status >= 500:
				l.Errorz()
			case res.Status >= 400:
				l.Warnz()
			default:
				l.Infoz()
			}

			return err
		}
	}
}
