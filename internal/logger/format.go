package logger

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-logfmt/logfmt"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	termTimeFormat          = "01-02|15:04:05"
	fmtTermTimeMessage      = "[%s] %-40s "
	fmtTermTimeMessageEmpty = "[%s] "
	fmtTermPkg              = "[%s] "

	recordKeyTime    = "t"
	recordKeyLevel   = "level"
	recordKeyPackage = "pkg"
	recordKeyMessage = "msg"
)

var (
	levelColors = map[Level]*color.Color{
		LevelPanic: color.New(color.FgHiMagenta),
		LevelFatal: color.New(color.FgMagenta),
		LevelError: color.New(color.FgRed),
		LevelWarn:  color.New(color.FgYellow),
		LevelInfo:  color.New(color.FgGreen),
		LevelDebug: color.New(color.FgCyan),
		LevelTrace: color.New(color.FgHiBlack),
	}

	pkgColors = map[Level]*color.Color{
		LevelPanic: color.New(color.FgHiMagenta),
		LevelFatal: color.New(color.FgMagenta),
		LevelError: color.New(color.FgRed),
		LevelWarn:  color.New(color.FgYellow),
		LevelInfo:  color.New(color.FgBlue),
		LevelDebug: color.New(color.FgHiBlack),
		LevelTrace: color.New(color.FgHiBlack),
	}

	levelPrinters = map[Level]termPrinter{
		LevelPanic: {color: levelColors[LevelPanic]},
		LevelFatal: {color: levelColors[LevelFatal]},
		LevelError: {color: levelColors[LevelError]},
		LevelWarn:  {color: levelColors[LevelWarn]},
		LevelInfo:  {color: levelColors[LevelInfo]},
		LevelDebug: {color: levelColors[LevelDebug]},
		LevelTrace: {color: levelColors[LevelTrace]},
	}

	pkgPrinters = map[Level]termPrinter{
		LevelPanic: {color: pkgColors[LevelPanic]},
		LevelFatal: {color: pkgColors[LevelFatal]},
		LevelError: {color: pkgColors[LevelError]},
		LevelWarn:  {color: pkgColors[LevelWarn]},
		LevelInfo:  {color: pkgColors[LevelInfo]},
		LevelDebug: {color: pkgColors[LevelDebug]},
		LevelTrace: {color: pkgColors[LevelTrace]},
	}
)

type logfmtRecord struct {
	time      time.Time
	lvl       Level
	pkg       string
	msg       string
	keyvalues []interface{}
}

func encodeLogfmt(w io.Writer, r *logfmtRecord, endRecord bool) error {
	header := []interface{}{
		recordKeyTime, r.time,
		recordKeyLevel, r.lvl,
	}

	if r.pkg != "" {
		header = append(header, recordKeyPackage, r.pkg)
	}

	header = append(header, recordKeyMessage, strings.TrimRight(r.msg, "\n"))
	keyvalues := append(header, r.keyvalues...)
	enc := logfmt.NewEncoder(w)

	if err := enc.EncodeKeyvals(keyvalues...); err != nil {
		return err
	}

	if endRecord {
		if err := enc.EndRecord(); err != nil {
			return err
		}
	}

	return nil
}

func extractPkgKeyvals(r *log15.Record) ([]interface{}, []interface{}) {
	var pkgkv []interface{}

	ctx := r.Ctx

	if len(ctx) > 1 {
		if key, ok := ctx[0].(string); ok && key == pkgKey {
			pkgkv, ctx = ctx[0:2], ctx[2:]
		}
	}

	return ctx, pkgkv
}

func LogfmtFormat() log15.Format {
	return log15.FormatFunc(func(r *log15.Record) []byte {
		ctx, pkgkv := extractPkgKeyvals(r)

		lfmtr := &logfmtRecord{
			time:      r.Time,
			lvl:       levelFromLog15(r.Lvl),
			msg:       r.Msg,
			keyvalues: ctx,
		}

		if len(pkgkv) > 1 {
			lfmtr.pkg = pkgkv[1].(string)
		}

		b := &bytes.Buffer{}

		if err := encodeLogfmt(b, lfmtr, true); err != nil {
			panic(err)
		}

		return b.Bytes()
	})
}

type termPrinter struct {
	color *color.Color
}

func (p termPrinter) Fprintf(w io.Writer, format string, values ...interface{}) (int, error) {
	if p.color != nil {
		return p.color.Fprintf(w, format, values...)
	}

	return fmt.Fprintf(w, format, values...)
}

func TerminalFormat() log15.Format {
	return log15.FormatFunc(func(r *log15.Record) []byte {
		timeStr := r.Time.Format(termTimeFormat)
		msg := strings.TrimRight(r.Msg, "\n")
		lvl := levelFromLog15(r.Lvl)
		lvlStr := strings.ToUpper(lvl.String())
		lvlPrinter := levelPrinters[lvl]
		pkgPrinter := pkgPrinters[lvl]
		ctx, pkgkv := extractPkgKeyvals(r)
		b := &bytes.Buffer{}

		lvlPrinter.Fprintf(b, "%5s", lvlStr)

		var timeMsgFmt string
		timeMsgArgs := []interface{}{timeStr}

		if len(msg) > 0 {
			timeMsgFmt = fmtTermTimeMessage
			timeMsgArgs = append(timeMsgArgs, msg)
		} else {
			timeMsgFmt = fmtTermTimeMessageEmpty
		}

		fmt.Fprintf(b, timeMsgFmt, timeMsgArgs...)

		if len(pkgkv) > 1 {
			pkgPrinter.Fprintf(b, fmtTermPkg, pkgkv[1])
		}

		enc := logfmt.NewEncoder(b)

		if err := enc.EncodeKeyvals(ctx...); err != nil {
			panic(err)
		}

		if err := enc.EndRecord(); err != nil {
			panic(err)
		}

		return b.Bytes()
	})
}
