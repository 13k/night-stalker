package errors

import (
	"golang.org/x/xerrors"
)

type Err struct {
	msg   string
	err   error
	frame xerrors.Frame
}

func Wrap(msg string, err error) *Err {
	return &Err{
		msg:   msg,
		err:   err,
		frame: xerrors.Caller(1),
	}
}

func (err *Err) Error() string {
	return err.msg
}

func (err *Err) Unwrap() error {
	return err.err
}

func (err *Err) FormatError(p xerrors.Printer) error {
	p.Print(err.msg)

	if p.Detail() {
		err.frame.Format(p)
	}

	return err.err
}
