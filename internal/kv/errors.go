package kv

import (
	"fmt"
	"strings"

	kv "github.com/galaco/KeyValues"
)

func IsKeyMissingError(err error) bool {
	return strings.HasPrefix(err.Error(), "could not find key")
}

func IsInvalidTypeError(err error) bool {
	_, ok := err.(*InvalidTypeError)
	return ok
}

type InvalidTypeError struct {
	Key          string
	Type         kv.ValueType
	ExpectedType kv.ValueType

	message string
}

func (err *InvalidTypeError) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("expected value type '%s' for key '%s' but got '%s'", err.ExpectedType, err.Key, err.Type)
	}

	return err.message
}
