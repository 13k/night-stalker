package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"cirello.io/oversight"
	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	pkgKey   = "pkg"
	pkgSep   = "/"
	errorKey = "error"
)

var (
	ErrChildClose = errors.New("cannot close child Logger")
)

type Logger struct {
	logger      *log15Logger
	stdLogger   *log.Logger
	level       Level
	outputs     []io.Writer
	multiOutput io.Writer
	handler     log15.Handler
	pkgPath     []string
	pkg         string
	parent      *Logger
}

func New(level Level, outputs ...io.Writer) *Logger {
	if len(outputs) == 0 {
		outputs = append(outputs, os.Stdout)
	}

	multiOutput := io.MultiWriter(outputs...)

	l := &Logger{
		level:       level,
		outputs:     outputs,
		multiOutput: multiOutput,
		handler:     createHandler(level, outputs...),
	}

	l.logger = newLog15Logger()
	l.logger.SetHandler(l)

	return l
}

func (l *Logger) Close() error {
	if l.parent != nil {
		return ErrChildClose
	}

	for _, output := range l.outputs {
		if closer, ok := output.(io.Closer); ok && !isStdio(output) {
			if err := closer.Close(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *Logger) child(keyvalues ...interface{}) *Logger {
	if len(keyvalues)%2 != 0 {
		panic("len(keyvalues) should be an even number")
	}

	child := &Logger{
		logger:      l.logger.child(keyvalues...),
		level:       l.level,
		outputs:     l.outputs,
		multiOutput: l.multiOutput,
		handler:     l.handler,
		pkgPath:     l.pkgPath,
		pkg:         l.pkg,
		parent:      l,
	}

	child.logger.SetHandler(child)

	return child
}

func (l *Logger) IsLevelEnabled(lvl Level) bool {
	return l.level.Enables(lvl)
}

func (l *Logger) StdLogger() *log.Logger {
	if l.stdLogger == nil {
		l.stdLogger = log.New(l.multiOutput, l.pkg, 0)
	}

	return l.stdLogger
}

func (l *Logger) MigrateLogger() migrate.Logger {
	return newMigrateLogger(l)
}

func (l *Logger) LogrusLogger() logrus.FieldLogger {
	return newLogrusLogger(l)
}

func (l *Logger) EchoLogger() echo.Logger {
	return newEchoLogger(l)
}

func (l *Logger) OversightLogger() oversight.Logger {
	return newOversightLogger(l)
}

func (l *Logger) Log(r *log15.Record) error {
	if l.pkg != "" {
		r.Ctx = append([]interface{}{pkgKey, l.pkg}, r.Ctx...)
	}

	return l.handler.Log(r)
}

func (l *Logger) WithPackage(pkg string) *Logger {
	child := l.child()
	child.pkgPath = append(child.pkgPath, pkg)
	child.pkg = strings.Join(child.pkgPath, pkgSep)
	return child
}

func (l *Logger) WithFields(fieldSet ...Fields) *Logger {
	fields := FieldSet(fieldSet).Merge()

	if len(fields) == 0 {
		return l
	}

	return l.child(log15.Ctx(fields))
}

func (l *Logger) WithOFields(fields ...interface{}) *Logger {
	if len(fields) == 0 {
		return l
	}

	return l.child(fields...)
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return l.child(key, value)
}

func (l *Logger) WithError(err error) *Logger {
	return l.WithField(errorKey, err.Error())
}

func (l *Logger) Panic(msg string) {
	l.logger.Crit(msg)
	panic(msg)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicln(args ...interface{}) {
	l.Panic(fmt.Sprintln(args...))
}

func (l *Logger) Fatal(msg string) {
	l.logger.Crit(msg)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.Fatal(fmt.Sprintln(args...))
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorln(args ...interface{}) {
	l.Error(fmt.Sprintln(args...))
}

func (l *Logger) Errorx(err error) {
	if _, ok := err.(xerrors.Formatter); ok {
		l.Errorf("%+v", err)
	}
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnln(args ...interface{}) {
	l.Warn(fmt.Sprintln(args...))
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Infoln(args ...interface{}) {
	l.Info(fmt.Sprintln(args...))
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugln(args ...interface{}) {
	l.Debug(fmt.Sprintln(args...))
}

func (l *Logger) Trace(msg string) {
	l.logger.Trace(msg)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Trace(fmt.Sprintf(format, args...))
}

func (l *Logger) Traceln(args ...interface{}) {
	l.Trace(fmt.Sprintln(args...))
}

func (l *Logger) Print(msg string) {
	l.Info(msg)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.Infoln(args...)
}
