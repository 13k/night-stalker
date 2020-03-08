package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var _ logrus.FieldLogger = (*logrusLogger)(nil)

type logrusLogger struct {
	logger *Logger
	logrus *logrus.Logger
}

func newLogrusLogger(l *Logger) *logrusLogger {
	lr := logrus.New()
	lr.Out = l.multiOutput
	lr.Level = l.level.toLogrus()

	return &logrusLogger{
		logger: l,
		logrus: lr,
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) *logrus.Entry {
	lr := newLogrusLogger(l.logger.WithField(key, value))
	return &logrus.Entry{Logger: lr.logrus}
}

func (l *logrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	lr := newLogrusLogger(l.logger.WithFields(Fields(fields)))
	return &logrus.Entry{Logger: lr.logrus}
}

func (l *logrusLogger) WithError(err error) *logrus.Entry {
	lr := newLogrusLogger(l.logger.WithError(err))
	return &logrus.Entry{Logger: lr.logrus}
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

func (l *logrusLogger) Warning(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(fmt.Sprint(args...))
}

func (l *logrusLogger) Debugln(args ...interface{}) {
	l.Debug(args...)
}

func (l *logrusLogger) Infoln(args ...interface{}) {
	l.Info(args...)
}

func (l *logrusLogger) Println(args ...interface{}) {
	l.Print(args...)
}

func (l *logrusLogger) Warnln(args ...interface{}) {
	l.Warn(args...)
}

func (l *logrusLogger) Warningln(args ...interface{}) {
	l.Warn(args...)
}

func (l *logrusLogger) Errorln(args ...interface{}) {
	l.Error(args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
	l.Fatal(args...)
}

func (l *logrusLogger) Panicln(args ...interface{}) {
	l.Panic(args...)
}
