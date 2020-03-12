package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	FilenameTimeFormat = "2006-01-02T15-04-05"
)

func ParseLogfilePath(lpath, name string) string {
	if lpath == "-" {
		return lpath
	}

	if fi, err := os.Stat(lpath); err == nil && fi.IsDir() {
		fname := fmt.Sprintf("%s.%s.log", name, time.Now().Format(FilenameTimeFormat))
		lpath = filepath.Join(lpath, fname)
	}

	return lpath
}

func New() (*nslog.Logger, error) {
	var outputs []io.Writer

	lpath := v.GetString(v.KeyLogFile)

	if lpath != "" && lpath != "-" {
		f, err := os.OpenFile(lpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			return nil, err
		}

		outputs = append(outputs, f)

		if v.GetBool(v.KeyLogTee) {
			outputs = append(outputs, os.Stdout)
		}
	}

	level := nslog.LevelInfo

	if v.GetBool(v.KeyLogDebug) {
		level = nslog.LevelDebug
	}

	if v.GetBool(v.KeyLogTrace) {
		level = nslog.LevelTrace
	}

	return nslog.New(level, outputs...), nil
}
