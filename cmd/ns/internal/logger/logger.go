package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	nsio "github.com/13k/night-stalker/internal/io"
	nslog "github.com/13k/night-stalker/internal/logger"
)

func New() (*nslog.Logger, error) {
	var out io.Writer = os.Stdout

	logFile := viper.GetString("log.file")

	if logFile != "" && logFile != "-" {
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return nil, err
		}

		if viper.GetBool("log.stdout") {
			out = nsio.MultiWriteCloser(f, os.Stdout)
		} else {
			out = f
		}
	}

	level := logrus.InfoLevel

	if viper.GetBool("log.debug") {
		level = logrus.DebugLevel
	}

	return nslog.New(out, level)
}
