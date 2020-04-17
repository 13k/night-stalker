package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"

	nslog "github.com/13k/night-stalker/internal/logger"
)

type SQLExecutor interface {
	// Begin can simulate nested transactions using savepoints
	Begin(ctx context.Context, options *sql.TxOptions) (SQLTx, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	ScanStruct(ctx context.Context, dst interface{}, query string, args ...interface{}) (exists bool, err error)
	ScanStructs(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	ScanVal(ctx context.Context, dst interface{}, query string, args ...interface{}) (exists bool, err error)
	ScanVals(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Trace(*TraceEntry)
}

var _ SQLExecutor = (*sqlExec)(nil)

type sqlExec struct {
	gqdb *goqu.Database
	log  *nslog.Logger
}

func (sx *sqlExec) Begin(ctx context.Context, options *sql.TxOptions) (SQLTx, error) {
	start := time.Now()
	gqtx, err := sx.gqdb.BeginTx(ctx, options)

	sx.Trace(&TraceEntry{
		start: start,
		end:   time.Now(),
		query: "BEGIN",
		err:   err,
	})

	if err != nil {
		return nil, err
	}

	return newSQLTx(ctx, gqtx, sx.Trace), nil
}

func (sx *sqlExec) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return sx.gqdb.ExecContext(ctx, query, args...)
}

func (sx *sqlExec) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return sx.gqdb.QueryContext(ctx, query, args...)
}

func (sx *sqlExec) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return sx.gqdb.QueryRowContext(ctx, query, args...)
}

func (sx *sqlExec) ScanStruct(ctx context.Context, dst interface{}, query string, args ...interface{}) (bool, error) {
	return sx.gqdb.ScanStructContext(ctx, dst, query, args...)
}

func (sx *sqlExec) ScanStructs(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return sx.gqdb.ScanStructsContext(ctx, dst, query, args...)
}

func (sx *sqlExec) ScanVal(ctx context.Context, dst interface{}, query string, args ...interface{}) (bool, error) {
	return sx.gqdb.ScanValContext(ctx, dst, query, args...)
}

func (sx *sqlExec) ScanVals(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return sx.gqdb.ScanValsContext(ctx, dst, query, args...)
}

func (sx *sqlExec) Trace(e *TraceEntry) {
	if sx.log == nil {
		return
	}

	msg, keyvals := e.Logfmt()
	l := sx.log.WithOFields(keyvals...)

	if e.err != nil {
		l = l.WithError(e.err)
	}

	if e.err != nil {
		l.Errorf(msg)
	} else {
		l.Tracef(msg)
	}
}
