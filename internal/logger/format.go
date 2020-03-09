package logger

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-logfmt/logfmt"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	// termTimeFormat = "02-01|15:04:05"
	termTimeFormat = "0102T150405"
	pkgColor       = 34
)

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
		common := []interface{}{
			r.KeyNames.Time, r.Time,
			r.KeyNames.Lvl, LevelFromLog15(r.Lvl),
		}

		if len(pkgkv) > 0 {
			common = append(common, pkgkv...)
		}

		common = append(common, r.KeyNames.Msg, r.Msg)
		data, err := logfmt.MarshalKeyvals(append(common, ctx...)...)

		if err != nil {
			panic(err)
		}

		return append(data, '\n')
	})
}

func TerminalFormat() log15.Format {
	return log15.FormatFunc(func(r *log15.Record) []byte {
		lvl := LevelFromLog15(r.Lvl)
		lvlstr := strings.ToUpper(lvl.String())
		color := lvl.Color()
		ctx, pkgkv := extractPkgKeyvals(r)
		values := []interface{}{lvlstr, r.Time.Format(termTimeFormat), r.Msg}

		if color > 0 {
			values = append([]interface{}{color}, values...)
		}

		var format string

		if len(pkgkv) > 1 {
			if color > 0 {
				format = "\x1b[%dm%5s\x1b[0m[%s] %-40s \x1b[%dm[%s]\x1b[0m "
				values = append(values, pkgColor, pkgkv[1])
			} else {
				format = "%5s[%s] %-40s [%s] "
				values = append(values, pkgkv[1])
			}
		} else {
			if color > 0 {
				format = "\x1b[%dm%5s\x1b[0m[%s] %-40s "
			} else {
				format = "%5s[%s] %-40s "
			}
		}

		b := &bytes.Buffer{}

		if _, err := fmt.Fprintf(b, format, values...); err != nil {
			panic(err)
		}

		/*
			if color > 0 {
				fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] %s ", color, lvlstr, r.Time.Format(termTimeFormat), r.Msg)
			} else {
				fmt.Fprintf(b, "[%s] [%s] %s ", lvlstr, r.Time.Format(termTimeFormat), r.Msg)
			}

			// try to justify the log output for short messages
			if len(ctx) > 0 && len(r.Msg) < termMsgJust {
				b.Write(bytes.Repeat([]byte{' '}, termMsgJust-len(r.Msg)))
			}
		*/

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
