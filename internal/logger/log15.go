package logger

import (
	"time"

	"github.com/go-stack/stack"
	"gopkg.in/inconshreveable/log15.v2"
)

var _ log15.Logger = (*log15Logger)(nil)

type log15Logger struct {
	log15.Logger

	ctx []interface{}
}

func newLog15Logger() *log15Logger {
	return &log15Logger{Logger: log15.New()}
}

func (l *log15Logger) child(ctx ...interface{}) *log15Logger {
	return &log15Logger{
		Logger: l.Logger.New(ctx...),
		ctx:    append(l.ctx, ctx...),
	}
}

func (l *log15Logger) Trace(msg string, ctx ...interface{}) {
	_ = l.GetHandler().Log(&log15.Record{
		Time: time.Now(),
		Lvl:  levelLog15Trace,
		Msg:  msg,
		Ctx:  append(l.ctx, ctx...),
		Call: stack.Caller(2),
		KeyNames: log15.RecordKeyNames{
			Time: recordKeyTime,
			Msg:  recordKeyMessage,
			Lvl:  recordKeyLevel,
		},
	})
}
