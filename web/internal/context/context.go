package context

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	nswebmime "github.com/13k/night-stalker/web/internal/mime"
	nswebres "github.com/13k/night-stalker/web/internal/responder"
)

type Context struct {
	echo.Context

	id               uuid.UUID
	mediaTypes       nswebmime.HTTPAcceptMediaTypes
	responder        nswebres.Responder
	responderVisited bool
	skipLogging      bool
}

func NewContext(id uuid.UUID, c echo.Context) *Context {
	return &Context{
		Context: c,
		id:      id,
	}
}

func (c *Context) ID() uuid.UUID {
	return c.id
}

func (c *Context) Error(err error) {
	c.Echo().HTTPErrorHandler(err, c)
}

func (c *Context) SkipLogging(v ...bool) bool {
	if len(v) > 0 {
		c.skipLogging = v[0]
	}

	return c.skipLogging
}

func (c *Context) MediaTypes() nswebmime.HTTPAcceptMediaTypes {
	if c.mediaTypes == nil {
		acceptHeader := c.Request().Header.Get(echo.HeaderAccept)
		c.mediaTypes = nswebmime.ParseHTTPAcceptMediaTypes(acceptHeader)
		c.mediaTypes.SortByQuality()
	}

	return c.mediaTypes
}

func (c *Context) ValidateMediaType() error {
	if c.Responder() == nil {
		return echo.ErrUnsupportedMediaType
	}

	return nil
}

func (c *Context) Responder() nswebres.Responder {
	if !c.responderVisited {
		for _, t := range c.MediaTypes() {
			if r := nswebres.GetByMatch(t); r != nil {
				c.responder = r
				break
			}
		}

		c.responderVisited = true
	}

	return c.responder
}

func (c *Context) RespondWith(status int, v interface{}) error {
	r := c.Responder()

	if r == nil {
		return echo.ErrUnsupportedMediaType
	}

	contentType := r.MediaType().Serialize()

	if s, ok := nswebres.AsStreamer(r); ok {
		stream, err := s.Stream(v)

		if err != nil {
			// TODO: wrap error
			return err
		}

		return c.Stream(status, contentType, stream)
	}

	data, err := r.Encode(v)

	if err != nil {
		// TODO: wrap error
		return err
	}

	c.Response().Header().Set(echo.HeaderContentLength, strconv.FormatInt(int64(len(data)), 10))

	return c.Blob(status, contentType, data)
}
