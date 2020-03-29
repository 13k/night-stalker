package responder

import (
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	nsjson "github.com/13k/night-stalker/internal/json"
	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

func JSON() Responder {
	return NewResponder(nswebmime.MediaTypeJSON, encodeJSON)
}

func encodeJSON(v interface{}) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return nsjson.ProtoMarshal(msg)
	}

	return jsoniter.Marshal(v)
}
