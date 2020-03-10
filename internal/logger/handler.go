package logger

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
	"gopkg.in/inconshreveable/log15.v2"
)

func isStdio(w io.Writer) bool {
	return w == os.Stdout || w == os.Stderr
}

func supportsColor(w io.Writer) bool {
	f, isFile := w.(*os.File)

	if !isFile {
		return false
	}

	return os.Getenv("TERM") != "dumb" && (isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd()))
}

func createHandler(level Level, outputs ...io.Writer) log15.Handler {
	handlers := make([]log15.Handler, len(outputs))

	for i, output := range outputs {
		var format log15.Format

		if supportsColor(output) {
			format = TerminalFormat()
		} else {
			format = LogfmtFormat()
		}

		handlers[i] = log15.StreamHandler(output, format)
	}

	multiHandler := log15.MultiHandler(handlers...)

	return log15.LvlFilterHandler(level.toLog15(), multiHandler)
}
