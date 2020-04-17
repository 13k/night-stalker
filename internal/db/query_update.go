package db

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var _ Query = (*UpdateQuery)(nil)

type UpdateQuery struct {
	*query

	ds *goqu.UpdateDataset
}

func NewUpdateQuery(b SQLBuilder, t goqu.Expression) *UpdateQuery {
	return &UpdateQuery{
		query: &query{b: b},
		ds:    b.Update(t),
	}
}

func (q *UpdateQuery) clone() *UpdateQuery {
	return &UpdateQuery{
		query: q.query.clone(),
		ds:    q.ds,
	}
}

func (q *UpdateQuery) withDS(ds *goqu.UpdateDataset) *UpdateQuery {
	newQ := q.clone()
	newQ.ds = ds
	return newQ
}

func (q *UpdateQuery) Trace() *UpdateQuery {
	newQ := q.clone()
	newQ.trace = true
	return newQ
}

func (q *UpdateQuery) ToSQL() (string, []interface{}, error) {
	return q.ds.ToSQL()
}

//#region goqu.UpdateDataset API {{{

func (q *UpdateQuery) ClearLimit() *UpdateQuery {
	return q.withDS(q.ds.ClearLimit())
}

func (q *UpdateQuery) ClearOrder() *UpdateQuery {
	return q.withDS(q.ds.ClearOrder())
}

func (q *UpdateQuery) ClearWhere() *UpdateQuery {
	return q.withDS(q.ds.ClearWhere())
}

func (q *UpdateQuery) From(tables ...interface{}) *UpdateQuery {
	return q.withDS(q.ds.From(tables...))
}

func (q *UpdateQuery) Limit(limit uint) *UpdateQuery {
	return q.withDS(q.ds.Limit(limit))
}

func (q *UpdateQuery) LimitAll() *UpdateQuery {
	return q.withDS(q.ds.LimitAll())
}

func (q *UpdateQuery) Order(order ...exp.OrderedExpression) *UpdateQuery {
	return q.withDS(q.ds.Order(order...))
}

func (q *UpdateQuery) OrderAppend(order ...exp.OrderedExpression) *UpdateQuery {
	return q.withDS(q.ds.OrderAppend(order...))
}

func (q *UpdateQuery) OrderPrepend(order ...exp.OrderedExpression) *UpdateQuery {
	return q.withDS(q.ds.OrderPrepend(order...))
}

func (q *UpdateQuery) Prepared(prepared bool) *UpdateQuery {
	return q.withDS(q.ds.Prepared(prepared))
}

func (q *UpdateQuery) Returning(returning ...interface{}) *UpdateQuery {
	return q.withDS(q.ds.Returning(returning...))
}

func (q *UpdateQuery) Set(values interface{}) *UpdateQuery {
	return q.withDS(q.ds.Set(values))
}

func (q *UpdateQuery) Table(table interface{}) *UpdateQuery {
	return q.withDS(q.ds.Table(table))
}

func (q *UpdateQuery) Where(expressions ...exp.Expression) *UpdateQuery {
	return q.withDS(q.ds.Where(expressions...))
}

func (q *UpdateQuery) With(name string, subquery exp.Expression) *UpdateQuery {
	return q.withDS(q.ds.With(name, subquery))
}

func (q *UpdateQuery) WithRecursive(name string, subquery exp.Expression) *UpdateQuery {
	return q.withDS(q.ds.WithRecursive(name, subquery))
}

//#endregion }}}
