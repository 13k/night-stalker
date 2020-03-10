package logger

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	logrusFieldKeyPackage = "@pkg"
)

var _ logrus.FieldLogger = (*logrusLogger)(nil)

type logrusLogger struct {
	*Logger
	logrus *logrus.Logger
}

func newLogrusLogger(l *Logger) *logrusLogger {
	lr := logrus.New()
	lr.Out = l.multiOutput
	lr.Level = l.level.toLogrus()
	lr.Formatter = logrusFuncFormatter(logrusLogfmtFormatter)

	return &logrusLogger{
		Logger: l,
		logrus: lr,
	}
}

func keyvaluesToLogrusFields(l *Logger) logrus.Fields {
	fields := logrus.Fields{logrusFieldKeyPackage: l.pkg}
	ctxLen := len(l.logger.ctx)

	for i := 0; i < ctxLen; i += 2 {
		if k, ok := l.logger.ctx[i].(string); ok && i+1 < ctxLen {
			fields[k] = l.logger.ctx[i+1]
		}
	}

	return fields
}

func (l *logrusLogger) newEntry() *logrus.Entry {
	return &logrus.Entry{
		Logger: l.logrus,
		Level:  l.logrus.Level,
		Data:   keyvaluesToLogrusFields(l.Logger),
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) *logrus.Entry {
	lr := newLogrusLogger(l.Logger.WithField(key, value))
	return lr.newEntry()
}

func (l *logrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	lr := newLogrusLogger(l.Logger.WithFields(Fields(fields)))
	return lr.newEntry()
}

func (l *logrusLogger) WithError(err error) *logrus.Entry {
	lr := newLogrusLogger(l.Logger.WithError(err))
	return lr.newEntry()
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.Logger.Panic(fmt.Sprint(args...))
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(fmt.Sprint(args...))
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.Logger.Error(fmt.Sprint(args...))
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...))
}

func (l *logrusLogger) Warning(args ...interface{}) {
	l.Logger.Warn(fmt.Sprint(args...))
}

func (l *logrusLogger) Warningf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *logrusLogger) Warningln(args ...interface{}) {
	l.Logger.Warnln(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.Logger.Info(fmt.Sprint(args...))
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(fmt.Sprint(args...))
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.Logger.Print(fmt.Sprint(args...))
}

var _ logrus.Formatter = (logrusFuncFormatter)(nil)

type logrusFuncFormatter func(*logrus.Entry) ([]byte, error)

func (fn logrusFuncFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return fn(entry)
}

func logrusLogfmtFormatter(entry *logrus.Entry) ([]byte, error) {
	r := &logfmtRecord{
		time: entry.Time,
		lvl:  levelFromLogrus(entry.Level),
		msg:  entry.Message,
		pkg:  entry.Data[logrusFieldKeyPackage].(string),
	}

	r.keyvalues = make([]interface{}, 0, len(entry.Data))

	for k, v := range entry.Data {
		if k != logrusFieldKeyPackage {
			r.keyvalues = append(r.keyvalues, k, v)
		}
	}

	b := &bytes.Buffer{}

	if err := encodeLogfmt(b, r, true); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
