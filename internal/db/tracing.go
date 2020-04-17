package db

import (
	"fmt"
	"time"
)

const (
	fmtTraceLogMessage = "%s %v"
)

type TracerFunc func(*TraceEntry)

type TraceEntry struct {
	query string
	args  []interface{}
	start time.Time
	end   time.Time
	rows  int64
	err   error
}

// Logfmt returns a log message and log keyvalue fields (except the error) for the logfmt format.
func (e *TraceEntry) Logfmt() (string, []interface{}) {
	duration := e.end.Sub(e.start)

	message := fmt.Sprintf(fmtTraceLogMessage, e.query, e.args)

	keyvals := []interface{}{
		"rows", e.rows,
		"duration", int64(duration),
		"duration_human", duration.String(),
	}

	return message, keyvals
}
