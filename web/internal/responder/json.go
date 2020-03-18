package responder

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"

	nsjson "github.com/13k/night-stalker/internal/json"
	nswebmime "github.com/13k/night-stalker/web/internal/mime"
)

func JSON() Responder {
	return NewResponder(nswebmime.MediaTypeJSON, encodeJSON)
}

func encodeJSON(v interface{}) ([]byte, error) {
	switch body := v.(type) {
	case proto.Message:
		return nsjson.ProtoMarshal(body)
	// case []proto.Message:
	// 	return nsjson.ProtoMarshal(body)
	default:
		// TODO: use a fast encoder
		return json.Marshal(body)
	}
}
