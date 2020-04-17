package db

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsdb "github.com/13k/night-stalker/internal/db"
	nslog "github.com/13k/night-stalker/internal/logger"
)

var (
	dbDrivers = map[string]string{
		"postgresql": "postgres",
	}

	ErrMissingDatabaseURL = errors.New("empty database URL")
)

func ParseURL() error {
	dbURL := v.GetString(v.KeyDbURL)

	if dbURL == "" {
		return ErrMissingDatabaseURL
	}

	uri, err := url.Parse(dbURL)

	if err != nil {
		return fmt.Errorf("invalid database URL %q: %s", dbURL, err)
	}

	driver, ok := dbDrivers[uri.Scheme]

	if !ok {
		return fmt.Errorf("unknown database scheme %q", uri.Scheme)
	}

	v.Set(v.KeyDbDriver, driver)

	return nil
}

func Connect(l *nslog.Logger) (*nsdb.DB, error) {
	if err := ParseURL(); err != nil {
		return nil, err
	}

	driver := v.GetString(v.KeyDbDriver)
	sqldb, err := sql.Open(driver, v.GetString(v.KeyDbURL))

	if err != nil {
		return nil, err
	}

	if n := v.GetInt(v.KeyDbConnMaxTotal); n > 0 {
		sqldb.SetMaxOpenConns(n)
	}

	if n := v.GetInt(v.KeyDbConnMaxIdle); n > 0 {
		sqldb.SetMaxIdleConns(n)
	}

	if d := v.GetDuration(v.KeyDbConnMaxLifetime); d > 0 {
		sqldb.SetConnMaxLifetime(d)
	}

	db := nsdb.New(sqldb, driver, l)

	return db, nil
}
