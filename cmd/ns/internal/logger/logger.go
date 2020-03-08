package logger

import (
	"io"
	"os"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nslog "github.com/13k/night-stalker/internal/logger"
)

func New() (*nslog.Logger, error) {
	var outputs []io.Writer

	logFile := v.GetString(v.KeyLogFile)

	if logFile != "" && logFile != "-" {
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			return nil, err
		}

		outputs = append(outputs, f)

		if v.GetBool(v.KeyLogTee) {
			outputs = append(outputs, os.Stdout)
		}
	}

	level := nslog.LevelInfo

	if v.GetBool(v.KeyLogDebug) {
		level = nslog.LevelDebug
	}

	if v.GetBool(v.KeyLogTrace) {
		level = nslog.LevelTrace
	}

	return nslog.New(level, outputs...), nil
}
