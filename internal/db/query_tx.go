package db

import (
	nslog "github.com/13k/night-stalker/internal/logger"
)

type QueryTx interface {
	QueryExecutor

	Commit() error
	Rollback() error
}

var _ QueryTx = (*queryTx)(nil)

type queryTx struct {
	QueryExecutor

	sqltx SQLTx
}

func NewQueryTx(sqltx SQLTx, log *nslog.Logger) QueryTx {
	return &queryTx{
		QueryExecutor: NewQueryExecutor(sqltx, log),
		sqltx:         sqltx,
	}
}

func (tx *queryTx) Commit() error {
	return tx.sqltx.Commit()
}

func (tx *queryTx) Rollback() error {
	return tx.sqltx.Rollback()
}
