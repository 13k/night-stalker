package db

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	dbDrivers = map[string]string{
		"postgresql": "postgres",
	}

	ErrMissingDatabaseURL = errors.New("empty database URL")
)

func ParseURL() error {
	dbURL := viper.GetString("db.url")

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

	viper.Set("db.driver", driver)

	return nil
}

func Connect() (*gorm.DB, error) {
	if err := ParseURL(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(viper.GetString("db.driver"), viper.GetString("db.url"))

	if err != nil {
		return nil, err
	}

	if n := viper.GetInt("db.max_total"); n > 0 {
		db.DB().SetMaxOpenConns(n)
	}

	if n := viper.GetInt("db.max_idle"); n > 0 {
		db.DB().SetMaxIdleConns(n)
	}

	if d := viper.GetDuration("db.max_lifetime"); d > 0 {
		db.DB().SetConnMaxLifetime(d)
	}

	return db, nil
}
