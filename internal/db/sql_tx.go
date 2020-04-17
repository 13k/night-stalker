package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type SQLTx interface {
	SQLExecutor

	Commit() error
	Rollback() error
}

var _ SQLTx = (*sqlTx)(nil)

type sqlTx struct {
	ctx       context.Context
	gqtx      *goqu.TxDatabase
	savepoint string
	tracer    TracerFunc
}

func newSQLTx(ctx context.Context, gqtx *goqu.TxDatabase, tracer TracerFunc) *sqlTx {
	return &sqlTx{ctx: ctx, gqtx: gqtx, tracer: tracer}
}

func (tx *sqlTx) Begin(ctx context.Context, _ *sql.TxOptions) (SQLTx, error) {
	savepoint := genSavepointName()

	_, err := tx.ExecTrace(ctx, fmt.Sprintf("SAVEPOINT %s", savepoint))

	if err != nil {
		return nil, err
	}

	tx2 := newSQLTx(ctx, tx.gqtx, tx.tracer)
	tx2.savepoint = savepoint

	return tx2, nil
}

func (tx *sqlTx) Commit() error {
	if tx.savepoint != "" {
		_, err := tx.ExecTrace(tx.ctx, fmt.Sprintf("RELEASE SAVEPOINT %s", tx.savepoint))
		return err
	}

	start := time.Now()
	err := tx.gqtx.Commit()

	tx.Trace(&TraceEntry{
		start: start,
		end:   time.Now(),
		query: "COMMIT",
		err:   err,
	})

	return err
}

func (tx *sqlTx) Rollback() error {
	if tx.savepoint != "" {
		_, err := tx.ExecTrace(tx.ctx, fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", tx.savepoint))
		return err
	}

	start := time.Now()
	err := tx.gqtx.Rollback()

	tx.Trace(&TraceEntry{
		start: start,
		end:   time.Now(),
		query: "ROLLBACK",
		err:   err,
	})

	return err
}

func (tx *sqlTx) ExecTrace(ctx context.Context, query string, args ...interface{}) (r sql.Result, err error) {
	var (
		start time.Time = time.Now()
		rows  int64
	)

	defer func() {
		tx.Trace(&TraceEntry{
			start: start,
			end:   time.Now(),
			query: query,
			args:  args,
			rows:  rows,
			err:   err,
		})
	}()

	r, err = tx.Exec(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	rows, err = r.RowsAffected()

	if err != nil {
		return nil, err
	}

	return
}

func (tx *sqlTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.gqtx.ExecContext(ctx, query, args...)
}

func (tx *sqlTx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.gqtx.QueryContext(ctx, query, args...)
}

func (tx *sqlTx) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return tx.gqtx.QueryRowContext(ctx, query, args...)
}

func (tx *sqlTx) ScanStruct(ctx context.Context, dst interface{}, query string, args ...interface{}) (bool, error) {
	return tx.gqtx.ScanStructContext(ctx, dst, query, args...)
}

func (tx *sqlTx) ScanStructs(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return tx.gqtx.ScanStructsContext(ctx, dst, query, args...)
}

func (tx *sqlTx) ScanVal(ctx context.Context, dst interface{}, query string, args ...interface{}) (bool, error) {
	return tx.gqtx.ScanValContext(ctx, dst, query, args...)
}

func (tx *sqlTx) ScanVals(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return tx.gqtx.ScanValsContext(ctx, dst, query, args...)
}

func (tx *sqlTx) Trace(e *TraceEntry) {
	if tx.tracer != nil {
		tx.tracer(e)
	}
}
