package protocol

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
)

var (
	jsonMarshaler = &jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		OrigName:     true,
	}
)

func Marshal(w io.Writer, message proto.Message) error {
	return jsonMarshaler.Marshal(w, message)
}

func MarshalBytes(message proto.Message) ([]byte, error) {
	var buf bytes.Buffer

	if err := Marshal(&buf, message); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *LiveMatch) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Hero) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}
