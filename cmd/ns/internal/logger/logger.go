package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsio "github.com/13k/night-stalker/internal/io"
	nslog "github.com/13k/night-stalker/internal/logger"
)

func New() (*nslog.Logger, error) {
	var out io.Writer = os.Stdout

	logFile := v.GetString(v.KeyLogFile)

	if logFile != "" && logFile != "-" {
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return nil, err
		}

		if v.GetBool(v.KeyLogTee) {
			out = nsio.MultiWriteCloser(f, os.Stdout)
		} else {
			out = f
		}
	}

	level := logrus.InfoLevel

	if v.GetBool(v.KeyLogDebug) {
		level = logrus.DebugLevel
	}

	if v.GetBool(v.KeyLogTrace) {
		level = logrus.TraceLevel
	}

	return nslog.New(out, level)
}
