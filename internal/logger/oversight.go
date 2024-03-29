package logger

import (
	"fmt"

	"cirello.io/oversight"
)

var _ oversight.Logger = (*oversightLogger)(nil)

type oversightLogger struct {
	logger *Logger
}

func newOversightLogger(l *Logger) *oversightLogger {
	return &oversightLogger{logger: l}
}

func (l *oversightLogger) Printf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *oversightLogger) Println(args ...interface{}) {
	l.logger.Debug(fmt.Sprintln(args...))
}
