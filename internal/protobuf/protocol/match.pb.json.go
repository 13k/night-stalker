// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: protocol/match.proto

package protocol

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
)

// MarshalJSON implements json.Marshaler
func (msg *Match) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := (&jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		OrigName:     true,
	}).Marshal(&buf, msg)
	return buf.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Match) UnmarshalJSON(b []byte) error {
	return (&jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}).Unmarshal(bytes.NewReader(b), msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Match_Player) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	err := (&jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: true,
		OrigName:     true,
	}).Marshal(&buf, msg)
	return buf.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Match_Player) UnmarshalJSON(b []byte) error {
	return (&jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}).Unmarshal(bytes.NewReader(b), msg)
}