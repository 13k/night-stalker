package logger

import (
	"github.com/sirupsen/logrus"
)

type MigrateLogger struct {
	logrus.FieldLogger
	verbose bool
}

func (l *MigrateLogger) Verbose() bool {
	return l.verbose
}
