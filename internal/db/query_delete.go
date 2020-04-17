package db

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var _ Query = (*DeleteQuery)(nil)

type DeleteQuery struct {
	*query

	ds *goqu.DeleteDataset
}

func NewDeleteQuery(b SQLBuilder, t goqu.Expression) *DeleteQuery {
	return &DeleteQuery{
		query: &query{b: b},
		ds:    b.Delete(t),
	}
}

func (q *DeleteQuery) clone() *DeleteQuery {
	return &DeleteQuery{
		query: q.query.clone(),
		ds:    q.ds,
	}
}

func (q *DeleteQuery) withDS(ds *goqu.DeleteDataset) *DeleteQuery {
	newQ := q.clone()
	newQ.ds = ds
	return newQ
}

func (q *DeleteQuery) Trace() *DeleteQuery {
	newQ := q.clone()
	newQ.trace = true
	return newQ
}

func (q *DeleteQuery) ToSQL() (string, []interface{}, error) {
	return q.ds.ToSQL()
}

//#region goqu.DeleteDataset API {{{

func (q *DeleteQuery) ClearLimit() *DeleteQuery {
	return q.withDS(q.ds.ClearLimit())
}

func (q *DeleteQuery) ClearOrder() *DeleteQuery {
	return q.withDS(q.ds.ClearOrder())
}

func (q *DeleteQuery) ClearWhere() *DeleteQuery {
	return q.withDS(q.ds.ClearWhere())
}

func (q *DeleteQuery) From(table interface{}) *DeleteQuery {
	return q.withDS(q.ds.From(table))
}

func (q *DeleteQuery) Limit(limit uint) *DeleteQuery {
	return q.withDS(q.ds.Limit(limit))
}

func (q *DeleteQuery) LimitAll() *DeleteQuery {
	return q.withDS(q.ds.LimitAll())
}

func (q *DeleteQuery) Order(order ...exp.OrderedExpression) *DeleteQuery {
	return q.withDS(q.ds.Order(order...))
}

func (q *DeleteQuery) OrderAppend(order ...exp.OrderedExpression) *DeleteQuery {
	return q.withDS(q.ds.OrderAppend(order...))
}

func (q *DeleteQuery) OrderPrepend(order ...exp.OrderedExpression) *DeleteQuery {
	return q.withDS(q.ds.OrderPrepend(order...))
}

func (q *DeleteQuery) Prepared(prepared bool) *DeleteQuery {
	return q.withDS(q.ds.Prepared(prepared))
}

func (q *DeleteQuery) Returning(returning ...interface{}) *DeleteQuery {
	return q.withDS(q.ds.Returning(returning...))
}

func (q *DeleteQuery) Where(expressions ...exp.Expression) *DeleteQuery {
	return q.withDS(q.ds.Where(expressions...))
}

func (q *DeleteQuery) With(name string, subquery exp.Expression) *DeleteQuery {
	return q.withDS(q.ds.With(name, subquery))
}

func (q *DeleteQuery) WithRecursive(name string, subquery exp.Expression) *DeleteQuery {
	return q.withDS(q.ds.WithRecursive(name, subquery))
}

//#endregion }}}
