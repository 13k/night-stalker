package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	FilenameTimeFormat = "2006-01-02T15-04-05"
)

var (
	log *nslog.Logger

	baseName = "ns"
)

func Init(cmdPath []string) (err error) {
	baseName = strings.Join(cmdPath, "-")
	log, err = New()
	return err
}

func Instance() *nslog.Logger {
	return log
}

func GenerateFilename(name string) string {
	ts := time.Now().Format(FilenameTimeFormat)
	return fmt.Sprintf("%s.%s.log", name, ts)
}

func New() (*nslog.Logger, error) {
	var outputs []io.Writer

	lpath := v.GetString(v.KeyLogPath)

	if lpath != "" && lpath != "-" {
		if fi, err := os.Stat(lpath); err == nil && fi.IsDir() {
			lpath = filepath.Join(lpath, GenerateFilename(baseName))
		}

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
