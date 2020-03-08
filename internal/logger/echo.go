package logger

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	elog "github.com/labstack/gommon/log"
)

var _ echo.Logger = (*echoLogger)(nil)

type echoLogger struct {
	logger *Logger
}

func newEchoLogger(l *Logger) *echoLogger {
	return &echoLogger{logger: l}
}

func (l *echoLogger) Output() io.Writer     { return l.logger.multiOutput }
func (l *echoLogger) SetOutput(w io.Writer) {}
func (l *echoLogger) Prefix() string        { return l.logger.pkg }
func (l *echoLogger) SetPrefix(p string)    {}
func (l *echoLogger) SetLevel(v elog.Lvl)   {}
func (l *echoLogger) SetHeader(h string)    {}

func (l *echoLogger) Level() elog.Lvl {
	return l.logger.level.toEcho()
}

func (l *echoLogger) Print(i ...interface{}) {
	l.logger.Info(fmt.Sprint(i...))
}

func (l *echoLogger) Printf(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *echoLogger) Printj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Info("")
}

func (l *echoLogger) Debug(i ...interface{}) {
	l.logger.Debug(fmt.Sprint(i...))
}

func (l *echoLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *echoLogger) Debugj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Debug("")
}

func (l *echoLogger) Info(i ...interface{}) {
	l.logger.Info(fmt.Sprint(i...))
}

func (l *echoLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *echoLogger) Infoj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Info("")
}

func (l *echoLogger) Warn(i ...interface{}) {
	l.logger.Warn(fmt.Sprint(i...))
}

func (l *echoLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *echoLogger) Warnj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Warn("")
}

func (l *echoLogger) Error(i ...interface{}) {
	l.logger.Error(fmt.Sprint(i...))
}

func (l *echoLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *echoLogger) Errorj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Error("")
}

func (l *echoLogger) Fatal(i ...interface{}) {
	l.logger.Fatal(fmt.Sprint(i...))
}

func (l *echoLogger) Fatalj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Fatal("")
}

func (l *echoLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *echoLogger) Panic(i ...interface{}) {
	l.logger.Panic(fmt.Sprint(i...))
}

func (l *echoLogger) Panicj(j elog.JSON) {
	l.logger.WithFields(Fields(j)).Panic("")
}

func (l *echoLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
