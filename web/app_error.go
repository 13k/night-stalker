package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func (app *App) handleError(err error, c echo.Context) {
	cc := c.(*nswebctx.Context)

	if cc.Response().Committed {
		app.log.Warn("response already committed")
		return
	}

	var he *echo.HTTPError

	switch e := err.(type) {
	case *echo.HTTPError:
		he = unwrapHTTPError(e)
	default:
		he = &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	if cc.Request().Method == http.MethodHead {
		err = cc.NoContent(he.Code)
	} else {
		useResponder := cc.Responder() != nil

		if useResponder {
			body := echo.Map{"error": he.Message}
			bodyStr := fmt.Sprintf("%+v", body)

			app.log.WithOFields(
				"status", he.Code,
				"body", bodyStr,
			).Debug("responding with")

			err = cc.RespondWith(he.Code, body)
		} else {
			message := fmt.Sprintf("%v", he.Message)
			err = cc.Blob(he.Code, echo.MIMETextPlain, []byte(message))
		}
	}

	if err != nil {
		app.log.WithError(err).Errorz()
	}
}

func unwrapHTTPError(err *echo.HTTPError) *echo.HTTPError {
	for e, ok := err.Message.(*echo.HTTPError); ok; err = e {
	}

	return err
}
