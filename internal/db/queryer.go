package db

import (
	nslog "github.com/13k/night-stalker/internal/logger"
)

type Queryer interface {
	QueryBuilder
	QueryExecutor
}

func NewQueryer(sb SQLBuilder, sx SQLExecutor, log *nslog.Logger) Queryer {
	return &queryer{
		QueryBuilder:  NewQueryBuilder(sb),
		QueryExecutor: NewQueryExecutor(sx, log),
	}
}

var _ Queryer = (*queryer)(nil)

type queryer struct {
	QueryBuilder
	QueryExecutor
}
