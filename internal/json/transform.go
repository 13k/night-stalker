package json

import (
	"bytes"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

var jsoniterTypeNames = map[jsoniter.ValueType]string{
	jsoniter.ArrayValue:   "array",
	jsoniter.ObjectValue:  "object",
	jsoniter.StringValue:  "string",
	jsoniter.NumberValue:  "number",
	jsoniter.NilValue:     "null",
	jsoniter.BoolValue:    "bool",
	jsoniter.InvalidValue: "invalid",
}

type ErrUnexpectedType struct {
	Expected jsoniter.ValueType
	Actual   jsoniter.ValueType
}

func (err *ErrUnexpectedType) Error() string {
	return fmt.Sprintf(
		"expected %s, got %s",
		jsoniterTypeNames[err.Expected],
		jsoniterTypeNames[err.Actual],
	)
}

type Transform struct {
	*jsoniter.Iterator

	b []byte
}

func NewTransform(data []byte) (*Transform, error) {
	if !jsoniter.Valid(data) {
		return nil, xerrors.New("invalid response data")
	}

	it := jsoniter.ParseBytes(jsoniter.ConfigDefault, data)

	if it.Error != nil {
		return nil, xerrors.Errorf("error parsing response data: %w", it.Error)
	}

	t := &Transform{
		Iterator: it,
		b:        data,
	}

	return t, nil
}

func (t *Transform) ValidateNext(expectedType jsoniter.ValueType) error {
	if actualType := t.WhatIsNext(); actualType != expectedType {
		return xerrors.Errorf("transform: parse error: %w", &ErrUnexpectedType{
			Expected: expectedType,
			Actual:   actualType,
		})
	}

	return nil
}

func (t *Transform) Array(cb func() ([]byte, error)) ([]byte, error) {
	if err := t.ValidateNext(jsoniter.ArrayValue); err != nil {
		return nil, xerrors.Errorf("transform: array transform error: %w", err)
	}

	buf := &bytes.Buffer{}
	first := true
	skip := false

	buf.WriteByte('[')

	for {
		if !t.ReadArray() {
			break
		}

		if !first && !skip {
			buf.WriteByte(',')
		}

		first = false

		b, err := cb()

		if err != nil {
			return nil, xerrors.Errorf("transform: array transform callback error: %w", err)
		}

		skip = b == nil

		if skip {
			t.Skip()
		} else {
			buf.Write(b)
		}
	}

	buf.WriteByte(']')

	return buf.Bytes(), nil
}

func (t *Transform) Object(cb func(string) ([]byte, error)) ([]byte, error) {
	if err := t.ValidateNext(jsoniter.ObjectValue); err != nil {
		return nil, xerrors.Errorf("transform: object transform error: %w", err)
	}

	buf := &bytes.Buffer{}
	first := true
	skip := false

	buf.WriteByte('{')

	for {
		key := t.ReadObject()

		if key == "" {
			break
		}

		if !first && !skip {
			buf.WriteByte(',')
		}

		first = false

		b, err := cb(key)

		if err != nil {
			return nil, xerrors.Errorf("transform: object transform callback error: %w", err)
		}

		skip = b == nil

		if skip {
			t.Skip()
		} else {
			fmt.Fprintf(buf, `"%s":`, key)
			buf.Write(b)
		}
	}

	buf.WriteByte('}')

	return buf.Bytes(), nil
}
