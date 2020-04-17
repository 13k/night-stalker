package viper

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Key is a numeric type to ensure only known keys are used.
//
// All Set and Get* functions will panic if an unknown key is given.
type Key uint

const (
	keyInvalid Key = iota
	KeyLogPath
	KeyLogDebug
	KeyLogTrace
	KeyLogTee
	KeyDbURL
	KeyDbDriver
	KeyDbConnMaxTotal
	KeyDbConnMaxIdle
	KeyDbConnMaxLifetime
	KeyRedisURL
	KeyRedisMaxRetries
	KeyRedisMinRetryBackoff
	KeyRedisMaxRetryBackoff
	KeyRedisDialTimeout
	KeyRedisReadTimeout
	KeyRedisWriteTimeout
	KeyRedisPoolSize
	KeyRedisMinIdleConns
	KeyRedisMaxConnAge
	KeyRedisPoolTimeout
	KeyRedisIdleTimeout
	KeyRedisIdleCheckFrequency
	KeySteamUser
	KeySteamPassword
	KeySteamAPIKey
	KeyOpendotaAPIKey
)

var _ = keyInvalid

var keys = map[Key]string{
	KeyLogPath:                 "log.path",
	KeyLogDebug:                "log.debug",
	KeyLogTrace:                "log.trace",
	KeyLogTee:                  "log.tee",
	KeyDbURL:                   "db.url",
	KeyDbDriver:                "db.driver",
	KeyDbConnMaxTotal:          "db.conn.max_total",
	KeyDbConnMaxIdle:           "db.conn.max_idle",
	KeyDbConnMaxLifetime:       "db.conn.max_lifetime",
	KeyRedisURL:                "redis.url",
	KeyRedisMaxRetries:         "redis.max_retries",
	KeyRedisMinRetryBackoff:    "redis.min_retry_backoff",
	KeyRedisMaxRetryBackoff:    "redis.max_retry_backoff",
	KeyRedisDialTimeout:        "redis.dial_timeout",
	KeyRedisReadTimeout:        "redis.read_timeout",
	KeyRedisWriteTimeout:       "redis.write_timeout",
	KeyRedisPoolSize:           "redis.pool_size",
	KeyRedisMinIdleConns:       "redis.min_idle_conns",
	KeyRedisMaxConnAge:         "redis.max_conn_age",
	KeyRedisPoolTimeout:        "redis.pool_timeout",
	KeyRedisIdleTimeout:        "redis.idle_timeout",
	KeyRedisIdleCheckFrequency: "redis.idle_check_freq",
	KeySteamUser:               "steam.user",
	KeySteamPassword:           "steam.password",
	KeySteamAPIKey:             "steam.api_key",
	KeyOpendotaAPIKey:          "opendota.api_key",
}

func getkey(key Key) string {
	if s, ok := keys[key]; ok {
		return s
	}

	panic(fmt.Errorf("Invalid key %+v", key))
}

func Set(key Key, value interface{}) {
	viper.Set(getkey(key), value)
}

func GetBool(key Key) bool {
	return viper.GetBool(getkey(key))
}

func GetDuration(key Key) time.Duration {
	return viper.GetDuration(getkey(key))
}

func GetInt(key Key) int {
	return viper.GetInt(getkey(key))
}

func GetString(key Key) string {
	return viper.GetString(getkey(key))
}

func GetStringSlice(key Key) []string {
	return viper.GetStringSlice(getkey(key))
}
