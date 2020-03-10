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

const (
	levelLog15Trace = log15.LvlDebug + 1
)

const (
	levelUnknownName = "unknown"
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
		LevelTrace: levelLog15Trace,
	}

	log15LevelsReverseMapping = map[log15.Lvl]Level{
		log15.LvlCrit:   LevelFatal,
		log15.LvlError:  LevelError,
		log15.LvlWarn:   LevelWarn,
		log15.LvlInfo:   LevelInfo,
		log15.LvlDebug:  LevelDebug,
		levelLog15Trace: LevelTrace,
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

	logrusLevelsReverseMapping = map[logrus.Level]Level{
		logrus.PanicLevel: LevelPanic,
		logrus.FatalLevel: LevelFatal,
		logrus.ErrorLevel: LevelError,
		logrus.WarnLevel:  LevelWarn,
		logrus.InfoLevel:  LevelInfo,
		logrus.DebugLevel: LevelDebug,
		logrus.TraceLevel: LevelTrace,
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

func levelFromLog15(lvl log15.Lvl) Level {
	if l, ok := log15LevelsReverseMapping[lvl]; ok {
		return l
	}

	return LevelInfo
}

func levelFromLogrus(lvl logrus.Level) Level {
	if l, ok := logrusLevelsReverseMapping[lvl]; ok {
		return l
	}

	return LevelInfo
}

func (l Level) String() string {
	if s, ok := levelNames[l]; ok {
		return s
	}

	return levelUnknownName
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
