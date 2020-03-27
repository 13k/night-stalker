package dota2

import (
	"fmt"
	"time"

	"github.com/paralin/go-dota2/protocol"

	nserr "github.com/13k/night-stalker/internal/errors"
)

type ErrWelcomeTimeout struct {
	RetryCount    int
	RetryInterval time.Duration
}

func (err *ErrWelcomeTimeout) Error() string {
	return fmt.Sprintf("dota welcome timeout after %d tries with interval %s", err.RetryCount, err.RetryInterval)
}

type ErrClientSuspended struct {
	Until time.Time
}

func (err *ErrClientSuspended) Error() string {
	return fmt.Sprintf("dota client suspended until %s", err.Until)
}

type ErrLostSession struct {
	Status protocol.GCConnectionStatus
}

func (err *ErrLostSession) Error() string {
	return fmt.Sprintf("dota session lost with reason %s", err.Status.String())
}

type ErrNoSession struct {
	*nserr.Err
}
