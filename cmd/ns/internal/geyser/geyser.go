package geyser

import (
	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
)

func getOptions() []geyser.ClientOption {
	var options []geyser.ClientOption

	if apiKey := v.GetString(v.KeySteamAPIKey); apiKey != "" {
		options = append(options, geyser.WithKey(apiKey))
	}

	return options
}

func New() (*geyser.Client, error) {
	return geyser.New(getOptions()...)
}

func NewDota2() (*geyserd2.Client, error) {
	return geyserd2.New(getOptions()...)
}
