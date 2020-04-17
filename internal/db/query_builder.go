package db

import (
	"github.com/doug-martin/goqu/v9"
)

type QueryBuilder interface {
	Select(cols ...interface{}) *SelectQuery
	Insert(table goqu.Expression) *InsertQuery
	Update(table goqu.Expression) *UpdateQuery
	Delete(table goqu.Expression) *DeleteQuery
}

type queryBuilder struct {
	sb SQLBuilder
}

func NewQueryBuilder(sb SQLBuilder) QueryBuilder {
	return &queryBuilder{sb: sb}
}

func (qb *queryBuilder) Select(cols ...interface{}) *SelectQuery {
	return NewSelectQuery(qb.sb, cols...)
}

func (qb *queryBuilder) Insert(table goqu.Expression) *InsertQuery {
	return NewInsertQuery(qb.sb, table)
}

func (qb *queryBuilder) Update(table goqu.Expression) *UpdateQuery {
	return NewUpdateQuery(qb.sb, table)
}

func (qb *queryBuilder) Delete(table goqu.Expression) *DeleteQuery {
	return NewDeleteQuery(qb.sb, table)
}
