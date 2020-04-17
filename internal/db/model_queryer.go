package db

import (
	"context"
	"database/sql"
)

type ModelQueryer interface {
	Q() Queryer
	M() ModelPersistence
	Begin(context.Context, *sql.TxOptions) (*Tx, error)
}
