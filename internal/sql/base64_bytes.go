package sql

import (
	"database/sql/driver"
	"encoding/base64"
	"errors"
)

var (
	ErrInvalidBase64BytesValue = errors.New("Invalid value for Base64Bytes")

	b64Enc = base64.StdEncoding
)

type Base64Bytes []byte

func (b *Base64Bytes) Scan(v interface{}) error {
	if v == nil {
		*b = nil
		return nil
	}

	var encoded []byte

	switch value := v.(type) {
	case []byte:
		encoded = value
	case string:
		encoded = []byte(value)
	default:
		return ErrInvalidBase64BytesValue
	}

	*b = make(Base64Bytes, b64Enc.DecodedLen(len(encoded)))

	if _, err := b64Enc.Decode(*b, encoded); err != nil {
		return err
	}

	return nil
}

func (b Base64Bytes) Value() (driver.Value, error) {
	encoded := make([]byte, b64Enc.EncodedLen(len(b)))
	b64Enc.Encode(encoded, b)
	return encoded, nil
}
