package logger

import (
	"strings"

	"github.com/golang-migrate/migrate/v4"
)

var _ migrate.Logger = (*migrateLogger)(nil)

type migrateLogger struct {
	logger *Logger
}

func newMigrateLogger(l *Logger) *migrateLogger {
	return &migrateLogger{logger: l}
}

func (l *migrateLogger) Verbose() bool {
	return false
}

func (l *migrateLogger) Printf(format string, v ...interface{}) {
	format = strings.TrimSpace(format)
	l.logger.Infof(format, v...)
}
