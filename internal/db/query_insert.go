package db

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var _ Query = (*InsertQuery)(nil)

type InsertQuery struct {
	*query

	ds *goqu.InsertDataset
}

func NewInsertQuery(b SQLBuilder, t goqu.Expression) *InsertQuery {
	return &InsertQuery{
		query: &query{b: b},
		ds:    b.Insert(t),
	}
}

func (q *InsertQuery) clone() *InsertQuery {
	return &InsertQuery{
		query: q.query.clone(),
		ds:    q.ds,
	}
}

func (q *InsertQuery) withDS(ds *goqu.InsertDataset) *InsertQuery {
	newQ := q.clone()
	newQ.ds = ds
	return newQ
}

func (q *InsertQuery) Trace() *InsertQuery {
	newQ := q.clone()
	newQ.trace = true
	return newQ
}

func (q *InsertQuery) ToSQL() (string, []interface{}, error) {
	return q.ds.ToSQL()
}

//#region goqu.InsertDataset API {{{

func (q *InsertQuery) ClearCols() *InsertQuery {
	return q.withDS(q.ds.ClearCols())
}

func (q *InsertQuery) ClearOnConflict() *InsertQuery {
	return q.withDS(q.ds.ClearOnConflict())
}

func (q *InsertQuery) ClearRows() *InsertQuery {
	return q.withDS(q.ds.ClearRows())
}

func (q *InsertQuery) ClearVals() *InsertQuery {
	return q.withDS(q.ds.ClearVals())
}

func (q *InsertQuery) Cols(cols ...interface{}) *InsertQuery {
	return q.withDS(q.ds.Cols(cols...))
}

func (q *InsertQuery) ColsAppend(cols ...interface{}) *InsertQuery {
	return q.withDS(q.ds.ColsAppend(cols...))
}

func (q *InsertQuery) FromQuery(from exp.AppendableExpression) *InsertQuery {
	return q.withDS(q.ds.FromQuery(from))
}

func (q *InsertQuery) Into(table interface{}) *InsertQuery {
	return q.withDS(q.ds.Into(table))
}

func (q *InsertQuery) OnConflict(conflict exp.ConflictExpression) *InsertQuery {
	return q.withDS(q.ds.OnConflict(conflict))
}

func (q *InsertQuery) Prepared(prepared bool) *InsertQuery {
	return q.withDS(q.ds.Prepared(prepared))
}

func (q *InsertQuery) Returning(returning ...interface{}) *InsertQuery {
	return q.withDS(q.ds.Returning(returning...))
}

func (q *InsertQuery) Rows(rows ...interface{}) *InsertQuery {
	return q.withDS(q.ds.Rows(rows...))
}

func (q *InsertQuery) Vals(vals ...[]interface{}) *InsertQuery {
	return q.withDS(q.ds.Vals(vals...))
}

func (q *InsertQuery) With(name string, subquery exp.Expression) *InsertQuery {
	return q.withDS(q.ds.With(name, subquery))
}

func (q *InsertQuery) WithRecursive(name string, subquery exp.Expression) *InsertQuery {
	return q.withDS(q.ds.WithRecursive(name, subquery))
}

//#endregion }}}
