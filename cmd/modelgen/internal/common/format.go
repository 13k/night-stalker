package common

import (
	"errors"
	"fmt"
	"go/format"
	"go/scanner"
)

func FormatCode(code []byte) ([]byte, error) {
	formatted, err := format.Source(code)

	if err != nil {
		if l := (scanner.ErrorList{}); errors.As(err, &l) {
			err = NewSyntaxError(string(code), l)
		}

		return nil, fmt.Errorf("error formatting generated code: %w", err)
	}

	return formatted, nil
}
