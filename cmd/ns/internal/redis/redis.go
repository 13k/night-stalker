package redis

import (
	"errors"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var (
	ErrMissingRedisURL = errors.New("empty redis URL")
)

func Connect() (*redis.Client, error) {
	redisURL := viper.GetString("redis.url")

	if redisURL == "" {
		return nil, ErrMissingRedisURL
	}

	options, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}

	options.MaxRetries = viper.GetInt("redis.max_retries")
	options.MinRetryBackoff = viper.GetDuration("redis.min_retry_backoff")
	options.MaxRetryBackoff = viper.GetDuration("redis.max_retry_backoff")
	options.DialTimeout = viper.GetDuration("redis.dial_timeout")
	options.ReadTimeout = viper.GetDuration("redis.read_timeout")
	options.WriteTimeout = viper.GetDuration("redis.write_timeout")
	options.PoolSize = viper.GetInt("redis.pool_size")
	options.MinIdleConns = viper.GetInt("redis.min_idle_conns")
	options.MaxConnAge = viper.GetDuration("redis.max_conn_age")
	options.PoolTimeout = viper.GetDuration("redis.pool_timeout")
	options.IdleTimeout = viper.GetDuration("redis.idle_timeout")
	options.IdleCheckFrequency = viper.GetDuration("redis.idle_check_freq")

	client := redis.NewClient(options)

	if _, err := client.Ping().Result(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}
