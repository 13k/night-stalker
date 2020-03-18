package json

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	ProtoMarshalOptions = protojson.MarshalOptions{
		Multiline:       false,
		UseProtoNames:   true,
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}

	ProtoUnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

// ProtoMarshal encodes a protobuf Message to JSON using the default ProtoMarshalOptions.
func ProtoMarshal(m proto.Message) ([]byte, error) {
	return ProtoMarshalOptions.Marshal(m)
}

// ProtoUnmarshal decodes JSON data into a protobuf Message using the default ProtoUnmarshalOptions.
func ProtoUnmarshal(b []byte, m proto.Message) error {
	return ProtoUnmarshalOptions.Unmarshal(b, m)
}
