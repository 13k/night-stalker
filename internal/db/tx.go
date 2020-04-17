package db

import (
	"context"
	"database/sql"

	nslog "github.com/13k/night-stalker/internal/logger"
)

var _ ModelQueryer = (*Tx)(nil)

type Tx struct {
	sqltx SQLTx
	sb    SQLBuilder
	log   *nslog.Logger
	q     Queryer
	m     ModelPersistence
}

func NewTx(sb SQLBuilder, sqltx SQLTx, log *nslog.Logger) *Tx {
	q := NewQueryer(sb, sqltx, log)
	mq := NewModelPersistence(q, q, log)

	return &Tx{
		sb:    sb,
		sqltx: sqltx,
		log:   log,
		q:     q,
		m:     mq,
	}
}

func (tx *Tx) Q() Queryer {
	return tx.q
}

func (tx *Tx) M() ModelPersistence {
	return tx.m
}

func (tx *Tx) Begin(ctx context.Context, options *sql.TxOptions) (*Tx, error) {
	sqltx, err := tx.sqltx.Begin(ctx, options)

	if err != nil {
		return nil, err
	}

	return NewTx(tx.sb, sqltx, tx.log), nil
}

func (tx *Tx) Commit() error {
	return tx.sqltx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.sqltx.Rollback()
}
