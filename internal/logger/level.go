package logger

import (
	elog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"gopkg.in/inconshreveable/log15.v2"
)

type Level uint8

const (
	LevelPanic Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

var (
	levelNames = map[Level]string{
		LevelPanic: "panic",
		LevelFatal: "fatal",
		LevelError: "error",
		LevelWarn:  "warn",
		LevelInfo:  "info",
		LevelDebug: "debug",
		LevelTrace: "trace",
	}

	log15LevelsMapping = map[Level]log15.Lvl{
		LevelPanic: log15.LvlCrit,
		LevelFatal: log15.LvlCrit,
		LevelError: log15.LvlError,
		LevelWarn:  log15.LvlWarn,
		LevelInfo:  log15.LvlInfo,
		LevelDebug: log15.LvlDebug,
		LevelTrace: log15.LvlDebug,
	}

	logrusLevelsMapping = map[Level]logrus.Level{
		LevelPanic: logrus.PanicLevel,
		LevelFatal: logrus.FatalLevel,
		LevelError: logrus.ErrorLevel,
		LevelWarn:  logrus.WarnLevel,
		LevelInfo:  logrus.InfoLevel,
		LevelDebug: logrus.DebugLevel,
		LevelTrace: logrus.TraceLevel,
	}

	echoLevelsMapping = map[Level]elog.Lvl{
		LevelPanic: elog.ERROR,
		LevelFatal: elog.ERROR,
		LevelError: elog.ERROR,
		LevelWarn:  elog.WARN,
		LevelInfo:  elog.INFO,
		LevelDebug: elog.DEBUG,
		LevelTrace: elog.DEBUG,
	}
)

func (l Level) String() string {
	return levelNames[l]
}

func (l Level) Enables(other Level) bool {
	return other <= l
}

func (l Level) toLog15() log15.Lvl {
	if lvl, ok := log15LevelsMapping[l]; ok {
		return lvl
	}

	return log15.LvlInfo
}

func (l Level) toLogrus() logrus.Level {
	if lvl, ok := logrusLevelsMapping[l]; ok {
		return lvl
	}

	return logrus.InfoLevel
}

func (l Level) toEcho() elog.Lvl {
	if lvl, ok := echoLevelsMapping[l]; ok {
		return lvl
	}

	return elog.INFO
}
