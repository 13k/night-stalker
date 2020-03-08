package logger

import (
	"github.com/golang-migrate/migrate/v4"
)

var _ migrate.Logger = (*migrateLogger)(nil)

type migrateLogger struct {
	logger  *Logger
	verbose bool
}

func newMigrateLogger(l *Logger) *migrateLogger {
	return &migrateLogger{
		logger:  l,
		verbose: l.IsLevelEnabled(LevelDebug),
	}
}

func (l *migrateLogger) Verbose() bool {
	return l.verbose
}

func (l *migrateLogger) Printf(format string, v ...interface{}) {
	var fn func(string, ...interface{})

	if l.verbose {
		fn = l.logger.Debugf
	} else {
		fn = l.logger.Infof
	}

	fn(format, v...)
}
