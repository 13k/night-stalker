package processors

import (
	"errors"
)

var (
	ErrProcessorContextLogger      = errors.New("context missing Logger")
	ErrProcessorContextDatabase    = errors.New("context missing Database")
	ErrProcessorContextRedis       = errors.New("context missing Redis")
	ErrProcessorContextSteamClient = errors.New("context missing Steam client")
	ErrProcessorContextDotaClient  = errors.New("context missing Dota2 client")
	ErrProcessorContextBus         = errors.New("context missing Bus")
	ErrProcessorContextAPI         = errors.New("context missing Steam API")
	ErrProcessorContextDotaAPI     = errors.New("context missing Dota2 API")
)
