package protocol

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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

func (m *Search) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Search_Player) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Team) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Hero) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *HeroMatches) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *LiveMatch) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Player) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *LiveMatches) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *LiveMatch_Player) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *PlayerMatches) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *LiveMatchesChange) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Match) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}

func (m *Match_Player) MarshalJSON() ([]byte, error) {
	return MarshalBytes(m)
}
