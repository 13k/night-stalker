package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

type Err struct {
	msg   string
	err   error
	frame xerrors.Frame
}

// Wrap returns a new `Err` that wraps the given error `err`.
func Wrap(msg string, err error) *Err {
	return &Err{
		msg:   msg,
		err:   err,
		frame: xerrors.Caller(1),
	}
}

// Error implements `error`
func (err *Err) Error() string {
	return fmt.Sprint(err)
}

// Unwrap implements `xerrors.Wrapper`
func (err *Err) Unwrap() error {
	return err.err
}

// Format implements `fmt.Formatter`
func (err *Err) Format(f fmt.State, c rune) {
	xerrors.FormatError(err, f, c)
}

// FormatError implements `xerrors.Formatter`
func (err *Err) FormatError(p xerrors.Printer) error {
	p.Print(err.msg)

	if p.Detail() {
		err.frame.Format(p)
	}

	return err.err
}
