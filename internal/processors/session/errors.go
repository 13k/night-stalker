package session

import (
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

type ErrInvalidServerAddress struct {
	Address string

	message string
}

func (err *ErrInvalidServerAddress) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("invalid server address '%s'", err.Address)
	}

	return err.message
}

func NewErrInvalidServerAddressX(addr string) error {
	return xerrors.Errorf("session error: %w", &ErrInvalidServerAddress{Address: addr})
}

type ErrSteamLogOnFailed struct {
	Reason string

	message string
}

func (err *ErrSteamLogOnFailed) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("steam login failed with reason '%s'", err.Reason)
	}

	return err.message
}

func NewErrSteamLogOnFailedX(reason string) error {
	return xerrors.Errorf("session error: %w", &ErrSteamLogOnFailed{Reason: reason})
}

type ErrSteamLoggedOff struct {
	Reason string

	message string
}

func (err *ErrSteamLoggedOff) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("steam logged off with reason '%s'", err.Reason)
	}

	return err.message
}

func NewErrSteamLoggedOffX(reason string) error {
	return xerrors.Errorf("session error: %w", &ErrSteamLoggedOff{Reason: reason})
}

type ErrSteamDisconnected struct{}

func (err *ErrSteamDisconnected) Error() string {
	return "steam disconnected"
}

func NewErrSteamDisconnectedX() error {
	return xerrors.Errorf("session error: %w", &ErrSteamDisconnected{})
}

type ErrDotaGCWelcomeTimeout struct {
	RetryCount    int
	RetryInterval time.Duration

	message string
}

func (err *ErrDotaGCWelcomeTimeout) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("dota GC hello failed after %d tries", err.RetryCount)
	}

	return err.message
}

func NewErrDotaGCWelcomeTimeoutX(count int, interval time.Duration) error {
	return xerrors.Errorf("session error: %w", &ErrDotaGCWelcomeTimeout{
		RetryCount:    count,
		RetryInterval: interval,
	})
}

type ErrDotaClientSuspended struct {
	Until *time.Time

	message string
}

func (err *ErrDotaClientSuspended) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("dota client suspended until %s", err.Until)
	}

	return err.message
}

func NewErrDotaClientSuspendedX(until *time.Time) error {
	return xerrors.Errorf("session error: %w", &ErrDotaClientSuspended{Until: until})
}
