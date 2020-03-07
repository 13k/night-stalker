package viper

import (
	"time"

	"github.com/spf13/viper"
)

type Key string

const (
	KeyLogFile                 Key = "log.file"
	KeyLogDebug                Key = "log.debug"
	KeyLogTee                  Key = "log.tee"
	KeyDbURL                   Key = "db.url"
	KeyDbDriver                Key = "db.driver"
	KeyDbConnMaxTotal          Key = "db.conn.max_total"
	KeyDbConnMaxIdle           Key = "db.conn.max_idle"
	KeyDbConnMaxLifetime       Key = "db.conn.max_lifetime"
	KeyRedisURL                Key = "redis.url"
	KeyRedisMaxRetries         Key = "redis.max_retries"
	KeyRedisMinRetryBackoff    Key = "redis.min_retry_backoff"
	KeyRedisMaxRetryBackoff    Key = "redis.max_retry_backoff"
	KeyRedisDialTimeout        Key = "redis.dial_timeout"
	KeyRedisReadTimeout        Key = "redis.read_timeout"
	KeyRedisWriteTimeout       Key = "redis.write_timeout"
	KeyRedisPoolSize           Key = "redis.pool_size"
	KeyRedisMinIdleConns       Key = "redis.min_idle_conns"
	KeyRedisMaxConnAge         Key = "redis.max_conn_age"
	KeyRedisPoolTimeout        Key = "redis.pool_timeout"
	KeyRedisIdleTimeout        Key = "redis.idle_timeout"
	KeyRedisIdleCheckFrequency Key = "redis.idle_check_freq"
	KeySteamUser               Key = "steam.user"
	KeySteamPassword           Key = "steam.password"
	KeyOpendotaAPIKey          Key = "opendota.api_key"
)

func Set(key Key, value interface{}) {
	viper.Set(string(key), value)
}

func GetString(key Key) string {
	return viper.GetString(string(key))
}

func GetBool(key Key) bool {
	return viper.GetBool(string(key))
}

func GetInt(key Key) int {
	return viper.GetInt(string(key))
}

func GetDuration(key Key) time.Duration {
	return viper.GetDuration(string(key))
}
