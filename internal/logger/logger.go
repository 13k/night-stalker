package logger

import (
	"io"
	"os"
	"strings"

	elog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

const (
	pkgKey = "@pkg"
)

type Logger struct {
	logrus.FieldLogger

	logger  *logrus.Logger
	pkgPath []string
	pkg     string
}

func New(output io.Writer, level logrus.Level) (*Logger, error) {
	logger := logrus.New()

	logger.SetLevel(level)
	logger.SetOutput(output)
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}

	l := &Logger{
		FieldLogger: logger,
		logger:      logger,
	}

	return l, nil
}

func (l *Logger) isStdio() bool {
	return l.logger.Out != os.Stdout && l.logger.Out != os.Stderr
}

func (l *Logger) Output() io.Writer {
	return l.logger.Out
}

func (l *Logger) Debugging() bool {
	return l.logger.IsLevelEnabled(logrus.DebugLevel)
}

func (l *Logger) Close() error {
	if closer, ok := l.logger.Out.(io.Closer); ok && !l.isStdio() {
		return closer.Close()
	}

	return nil
}

func (l *Logger) Debugfn(fn func()) {
	if !l.logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}

	fn()
}

func (l *Logger) WithPackage(pkg string) *Logger {
	pkgPath := append(l.pkgPath, pkg)
	pkg = strings.Join(pkgPath, "/")

	return &Logger{
		FieldLogger: l.WithField(pkgKey, pkg),
		logger:      l.logger,
		pkgPath:     pkgPath,
		pkg:         pkg,
	}
}

func (l *Logger) MigrateLogger() *MigrateLogger {
	return &MigrateLogger{
		FieldLogger: l.WithPackage("migrate"),
		verbose:     l.Debugging(),
	}
}

func (l *Logger) Dota2Logger() *logrus.Entry {
	return l.WithPackage("dota2").WithFields(logrus.Fields{})
}

func (l *Logger) EchoLogger() *elog.Logger {
	logger := elog.New(l.pkg)
	logger.SetOutput(l.logger.Out)

	return logger
}
