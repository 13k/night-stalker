package processors

import (
	"errors"
)

var (
	ErrProcessorContextLogger      = errors.New("ErrProcessorContextLogger")
	ErrProcessorContextDatabase    = errors.New("ErrProcessorContextDatabase")
	ErrProcessorContextRedis       = errors.New("ErrProcessorContextRedis")
	ErrProcessorContextSteamClient = errors.New("ErrProcessorContextSteamClient")
	ErrProcessorContextDotaClient  = errors.New("ErrProcessorContextDotaClient")
	ErrProcessorContextPubSub      = errors.New("ErrProcessorContextPubSub")
	ErrProcessorContextAPI         = errors.New("ErrProcessorContextAPI")
	ErrProcessorContextDotaAPI     = errors.New("ErrProcessorContextDotaAPI")
)
