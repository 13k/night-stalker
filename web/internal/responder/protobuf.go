package responder

import (
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"

	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

func init() {
	Register(Protobuf(), 1)
}

func Protobuf() Responder {
	return NewResponder(nswebmime.MediaTypeProtobuf, encodeProtobuf)
}

func encodeProtobuf(v interface{}) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	}

	return nil, xerrors.New("responder/protobuf: not a protobuf message")
}
