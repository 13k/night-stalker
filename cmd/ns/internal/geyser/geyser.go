package geyser

import (
	"github.com/13k/geyser"
	gsdota2 "github.com/13k/geyser/dota2"

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

func NewDota2() (*gsdota2.Client, error) {
	return gsdota2.New(getOptions()...)
}
