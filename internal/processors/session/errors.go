package session

import (
	"time"
)

type ErrInvalidServerAddress struct {
	Address string
}

func (*ErrInvalidServerAddress) Error() string {
	return "invalid server address"
}

type ErrSteamLogOnFailed struct {
	Reason string
}

func (*ErrSteamLogOnFailed) Error() string {
	return "steam login failed"
}

type ErrSteamLoggedOff struct {
	Reason string
}

func (*ErrSteamLoggedOff) Error() string {
	return "steam logged off"
}

type ErrSteamDisconnected struct{}

func (*ErrSteamDisconnected) Error() string {
	return "steam disconnected"
}

type ErrDotaGCWelcomeTimeout struct {
	RetryCount    int
	RetryInterval time.Duration
}

func (*ErrDotaGCWelcomeTimeout) Error() string {
	return "dota GC welcome timeout"
}

type ErrDotaClientSuspended struct {
	Until *time.Time
}

func (*ErrDotaClientSuspended) Error() string {
	return "dota client suspended"
}
