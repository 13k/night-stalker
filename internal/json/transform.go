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
		return nil, xerrors.New("Transform: data is not valid JSON")
	}

	it := jsoniter.ParseBytes(jsoniter.ConfigDefault, data)

	if it.Error != nil {
		return nil, xerrors.Errorf("Transform: error parsing JSON: %w", it.Error)
	}

	t := &Transform{
		Iterator: it,
		b:        data,
	}

	return t, nil
}

func (t *Transform) ValidateNext(expectedType jsoniter.ValueType) error {
	if actualType := t.WhatIsNext(); actualType != expectedType {
		return xerrors.Errorf("Transform: parse error: %w", &ErrUnexpectedType{
			Expected: expectedType,
			Actual:   actualType,
		})
	}

	return nil
}

func (t *Transform) Array(cb func() ([]byte, error)) ([]byte, error) {
	if err := t.ValidateNext(jsoniter.ArrayValue); err != nil {
		return nil, xerrors.Errorf("Transform: Array transform error: %w", err)
	}

	var items [][]byte

	for {
		if !t.ReadArray() {
			break
		}

		b, err := cb()

		if err != nil {
			return nil, xerrors.Errorf("Transform: Array callback error: %w", err)
		}

		if b == nil {
			t.Skip()
		} else {
			items = append(items, b)
		}
	}

	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, item := range items {
		buf.Write(item)

		if i < len(items)-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteByte(']')

	return buf.Bytes(), nil
}

func (t *Transform) Object(cb func(string) ([]byte, error)) ([]byte, error) {
	if err := t.ValidateNext(jsoniter.ObjectValue); err != nil {
		return nil, xerrors.Errorf("Transform: Object transform error: %w", err)
	}

	type kv struct {
		key   string
		value []byte
	}

	var items []*kv

	for {
		key := t.ReadObject()

		if key == "" {
			break
		}

		b, err := cb(key)

		if err != nil {
			return nil, xerrors.Errorf("Transform: Object callback error: %w", err)
		}

		if b == nil {
			t.Skip()
		} else {
			items = append(items, &kv{
				key:   key,
				value: b,
			})
		}
	}

	var buf bytes.Buffer

	buf.WriteByte('{')

	for i, kv := range items {
		fmt.Fprintf(&buf, `"%s":`, kv.key)
		buf.Write(kv.value)

		if i < len(items)-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteByte('}')

	return buf.Bytes(), nil
}
