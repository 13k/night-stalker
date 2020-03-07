package db

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
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

func Connect() (*gorm.DB, error) {
	if err := ParseURL(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(v.GetString(v.KeyDbDriver), v.GetString(v.KeyDbURL))

	if err != nil {
		return nil, err
	}

	if n := v.GetInt(v.KeyDbConnMaxTotal); n > 0 {
		db.DB().SetMaxOpenConns(n)
	}

	if n := v.GetInt(v.KeyDbConnMaxIdle); n > 0 {
		db.DB().SetMaxIdleConns(n)
	}

	if d := v.GetDuration(v.KeyDbConnMaxLifetime); d > 0 {
		db.DB().SetConnMaxLifetime(d)
	}

	return db, nil
}
