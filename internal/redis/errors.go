package redis

import (
	"github.com/go-redis/redis/v7"

	nserr "github.com/13k/night-stalker/internal/errors"
)

type ErrCommandFailure struct {
	*nserr.Err

	Cmd redis.Cmder
	Key string
}

type ErrPubsubFailure struct {
	*nserr.Err

	Cmd   redis.Cmder
	Topic string
}
