package redis

import (
	"errors"

	"github.com/go-redis/redis/v7"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsrds "github.com/13k/night-stalker/internal/redis"
)

var (
	ErrMissingRedisURL = errors.New("empty redis URL")
)

func Connect() (*nsrds.Redis, error) {
	redisURL := v.GetString(v.KeyRedisURL)

	if redisURL == "" {
		return nil, ErrMissingRedisURL
	}

	options, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}

	options.MaxRetries = v.GetInt(v.KeyRedisMaxRetries)
	options.MinRetryBackoff = v.GetDuration(v.KeyRedisMinRetryBackoff)
	options.MaxRetryBackoff = v.GetDuration(v.KeyRedisMaxRetryBackoff)
	options.DialTimeout = v.GetDuration(v.KeyRedisDialTimeout)
	options.ReadTimeout = v.GetDuration(v.KeyRedisReadTimeout)
	options.WriteTimeout = v.GetDuration(v.KeyRedisWriteTimeout)
	options.PoolSize = v.GetInt(v.KeyRedisPoolSize)
	options.MinIdleConns = v.GetInt(v.KeyRedisMinIdleConns)
	options.MaxConnAge = v.GetDuration(v.KeyRedisMaxConnAge)
	options.PoolTimeout = v.GetDuration(v.KeyRedisPoolTimeout)
	options.IdleTimeout = v.GetDuration(v.KeyRedisIdleTimeout)
	options.IdleCheckFrequency = v.GetDuration(v.KeyRedisIdleCheckFrequency)

	client := redis.NewClient(options)

	if _, err := client.Ping().Result(); err != nil {
		client.Close()
		return nil, err
	}

	return nsrds.NewRedis(client), nil
}
