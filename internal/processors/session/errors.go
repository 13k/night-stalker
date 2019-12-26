package session

import (
	"errors"
)

var (
	ErrInvalidServerAddress = errors.New("Invalid server address")
	ErrSteamLogOnFailed     = errors.New("Steam login failed")
	ErrSteamLoggedOff       = errors.New("Steam logged off")
	ErrSteamDisconnected    = errors.New("Steam disconnected")
	ErrDotaGCWelcomeTimeout = errors.New("Dota2 GC welcome timeout")
	ErrDotaClientSuspended  = errors.New("Dota2 client suspended")
)
