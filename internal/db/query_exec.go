package db

import (
	"context"
	"database/sql"
	"time"

	nslog "github.com/13k/night-stalker/internal/logger"
)

type QueryExecutor interface {
	Begin(ctx context.Context, options *sql.TxOptions) (QueryTx, error)
	Exec(ctx context.Context, q Query) (rows int64, err error)
	ScanStruct(ctx context.Context, q Query, dst interface{}) (exists bool, err error)
	ScanStructs(ctx context.Context, q Query, dst interface{}) error
	ScanVal(ctx context.Context, q Query, dst interface{}) (exists bool, err error)
	ScanVals(ctx context.Context, q Query, dst interface{}) error
}

func NewQueryExecutor(sx SQLExecutor, log *nslog.Logger) QueryExecutor {
	return &queryExec{sx: sx, log: log}
}

var _ QueryExecutor = (*queryExec)(nil)

type queryExec struct {
	sx  SQLExecutor
	log *nslog.Logger
}

func (qx *queryExec) Begin(ctx context.Context, options *sql.TxOptions) (QueryTx, error) {
	sqltx, err := qx.sx.Begin(ctx, options)

	if err != nil {
		return nil, err
	}

	return NewQueryTx(sqltx, qx.log), nil
}

func (qx *queryExec) Exec(ctx context.Context, q Query) (rows int64, err error) {
	return qx.exec(ctx, q)
}

func (qx *queryExec) ScanStruct(ctx context.Context, q Query, dst interface{}) (bool, error) {
	return qx.scanOne(ctx, q, qx.sx.ScanStruct, dst)
}

func (qx *queryExec) ScanStructs(ctx context.Context, q Query, dst interface{}) error {
	return qx.scanMany(ctx, q, qx.sx.ScanStructs, dst)
}

func (qx *queryExec) ScanVal(ctx context.Context, q Query, dst interface{}) (bool, error) {
	return qx.scanOne(ctx, q, qx.sx.ScanVal, dst)
}

func (qx *queryExec) ScanVals(ctx context.Context, q Query, dst interface{}) error {
	return qx.scanMany(ctx, q, qx.sx.ScanVals, dst)
}

func (qx *queryExec) exec(ctx context.Context, q Query) (rows int64, err error) {
	var (
		start time.Time = time.Now()
		query string
		args  []interface{}
		r     sql.Result
	)

	defer func() {
		if !qx.shouldTrace(q, err) {
			return
		}

		qx.sx.Trace(&TraceEntry{
			start: start,
			end:   time.Now(),
			query: query,
			args:  args,
			rows:  rows,
			err:   err,
		})
	}()

	query, args, err = q.ToSQL()

	if err != nil {
		return 0, err
	}

	r, err = qx.sx.Exec(ctx, query, args...)

	if err != nil {
		return 0, err
	}

	rows, err = r.RowsAffected()

	return
}

type scannerOne func(context.Context, interface{}, string, ...interface{}) (bool, error)

func (qx *queryExec) scanOne(
	ctx context.Context,
	q Query,
	scan scannerOne,
	dst interface{},
) (exists bool, err error) {
	var (
		start time.Time = time.Now()
		query string
		args  []interface{}
	)

	defer func() {
		if !qx.shouldTrace(q, err) {
			return
		}

		qx.traceScanOne(exists, &TraceEntry{
			start: start,
			end:   time.Now(),
			query: query,
			args:  args,
			err:   err,
		})
	}()

	query, args, err = q.ToSQL()

	if err != nil {
		return false, err
	}

	return scan(ctx, dst, query, args...)
}

type scannerMany func(context.Context, interface{}, string, ...interface{}) error

func (qx *queryExec) scanMany(ctx context.Context, q Query, scan scannerMany, dst interface{}) (err error) {
	var (
		start time.Time = time.Now()
		query string
		args  []interface{}
	)

	defer func() {
		if !qx.shouldTrace(q, err) {
			return
		}

		qx.traceScanMany(dst, &TraceEntry{
			start: start,
			end:   time.Now(),
			query: query,
			args:  args,
			err:   err,
		})
	}()

	query, args, err = q.ToSQL()

	if err != nil {
		return err
	}

	return scan(ctx, dst, query, args...)
}

func (qx *queryExec) shouldTrace(q Query, err error) bool {
	return qx.log != nil && (q.IsTraceEnabled() || err != nil)
}

func (qx *queryExec) traceScanOne(exists bool, e *TraceEntry) {
	if exists {
		e.rows = 1
	}

	qx.sx.Trace(e)
}

func (qx *queryExec) traceScanMany(dst interface{}, e *TraceEntry) {
	if e.err == nil && dst != nil {
		e.rows = int64(reflectSlicePtrLen(dst))
	}

	qx.sx.Trace(e)
}
