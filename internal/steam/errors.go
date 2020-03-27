package steam

import (
	"fmt"

	nserr "github.com/13k/night-stalker/internal/errors"
)

type ErrLogOnFailed struct {
	Reason string
}

func (err *ErrLogOnFailed) Error() string {
	return fmt.Sprintf("steam login failed with reason %s", err.Reason)
}

type ErrLoggedOff struct {
	Reason string
}

func (err *ErrLoggedOff) Error() string {
	return fmt.Sprintf("steam logged off with reason %s", err.Reason)
}

type ErrFailure struct {
	Reason string
}

func (err *ErrFailure) Error() string {
	return fmt.Sprintf("steam failed with reason %s", err.Reason)
}

type ErrDisconnected struct{}

func (*ErrDisconnected) Error() string {
	return "steam disconnected"
}

type ErrLostSession struct{}

func (*ErrLostSession) Error() string {
	return "steam session lost"
}

type ErrNoSession struct {
	*nserr.Err
}
