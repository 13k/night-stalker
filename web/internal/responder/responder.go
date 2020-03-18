package responder

import (
	"io"

	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

type EncoderFunc func(interface{}) ([]byte, error)

type Responder interface {
	MediaType() *nswebmime.MediaType
	Encode(interface{}) ([]byte, error)
}

func NewResponder(mt *nswebmime.MediaType, enc EncoderFunc) Responder {
	return &responder{
		mt:  mt,
		enc: enc,
	}
}

type Streamer interface {
	Stream(interface{}) (io.Reader, error)
}

func IsStreamer(r Responder) bool {
	_, ok := r.(Streamer)
	return ok
}

func AsStreamer(r Responder) (Streamer, bool) {
	s, ok := r.(Streamer)
	return s, ok
}

type responder struct {
	mt  *nswebmime.MediaType
	enc EncoderFunc
}

func (r *responder) MediaType() *nswebmime.MediaType {
	return r.mt
}

func (r *responder) Encode(v interface{}) ([]byte, error) {
	return r.enc(v)
}
