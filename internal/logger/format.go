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
	termTimeFormat = "01-02|15:04:05"

	recordKeyTime    = "t"
	recordKeyLevel   = "level"
	recordKeyPackage = "pkg"
	recordKeyMessage = "msg"
)

var (
	levelColorAttrs = map[Level][]color.Attribute{
		LevelPanic: {color.FgHiMagenta},
		LevelFatal: {color.FgMagenta},
		LevelError: {color.FgRed},
		LevelWarn:  {color.FgYellow},
		LevelInfo:  {color.FgGreen},
		LevelDebug: {color.FgCyan},
		LevelTrace: {color.FgHiBlack},
	}

	pkgColorAttrs = map[Level][]color.Attribute{
		LevelPanic: {color.FgHiMagenta},
		LevelFatal: {color.FgMagenta},
		LevelError: {color.FgRed},
		LevelWarn:  {color.FgYellow},
		LevelInfo:  {color.FgBlue},
		LevelDebug: {color.FgHiBlack},
		LevelTrace: {color.FgHiBlack},
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

func TerminalFormat() log15.Format {
	return log15.FormatFunc(func(r *log15.Record) []byte {
		lvl := levelFromLog15(r.Lvl)
		lvlColor := color.New(levelColorAttrs[lvl]...)
		ctx, pkgkv := extractPkgKeyvals(r)
		b := &bytes.Buffer{}

		lvlColor.Fprintf(b, "%5s", strings.ToUpper(lvl.String()))
		fmt.Fprintf(b, "[%s] %-40s ", r.Time.Format(termTimeFormat), strings.TrimRight(r.Msg, "\n"))

		if len(pkgkv) > 1 {
			pkgColor := color.New(pkgColorAttrs[lvl]...)
			pkgColor.Fprintf(b, "[%s] ", pkgkv[1])
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
