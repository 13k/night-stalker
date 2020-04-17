package db

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"

	nslog "github.com/13k/night-stalker/internal/logger"
)

var _ ModelQueryer = (*DB)(nil)

type DB struct {
	sqldb *sql.DB
	sx    SQLExecutor
	sb    SQLBuilder
	log   *nslog.Logger
	q     Queryer
	m     ModelPersistence
}

func New(sqldb *sql.DB, dialect string, log *nslog.Logger) *DB {
	gqdb := goqu.New(dialect, sqldb)
	sx := &sqlExec{gqdb: gqdb, log: log}
	q := NewQueryer(gqdb, sx, log)
	mq := NewModelPersistence(q, q, log)

	return &DB{
		sqldb: sqldb,
		sx:    sx,
		sb:    gqdb,
		log:   log,
		q:     q,
		m:     mq,
	}
}

func (db *DB) SQLDB() *sql.DB {
	return db.sqldb
}

func (db *DB) Close() error {
	return db.sqldb.Close()
}

func (db *DB) Q() Queryer {
	return db.q
}

func (db *DB) M() ModelPersistence {
	return db.m
}

func (db *DB) Begin(ctx context.Context, options *sql.TxOptions) (*Tx, error) {
	sqltx, err := db.sx.Begin(ctx, options)

	if err != nil {
		return nil, err
	}

	return NewTx(db.sb, sqltx, db.log), nil
}
