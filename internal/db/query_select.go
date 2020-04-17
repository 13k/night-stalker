package db

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type SelectQueryFilter interface {
	Empty() bool
	Validate() error
	Filter(*SelectQuery) *SelectQuery
}

var _ Query = (*SelectQuery)(nil)

type SelectQuery struct {
	*query

	ds *goqu.SelectDataset
}

func NewSelectQuery(b SQLBuilder, cols ...interface{}) *SelectQuery {
	return &SelectQuery{
		query: &query{b: b},
		ds:    b.Select(cols...),
	}
}

func (q *SelectQuery) clone() *SelectQuery {
	return &SelectQuery{
		query: q.query.clone(),
		ds:    q.ds,
	}
}

func (q *SelectQuery) withDS(ds *goqu.SelectDataset) *SelectQuery {
	newQ := q.clone()
	newQ.ds = ds
	return newQ
}

func (q *SelectQuery) Trace() *SelectQuery {
	newQ := q.clone()
	newQ.trace = true
	return newQ
}

func (q *SelectQuery) Expr() exp.Expression {
	return q.ds.Expression()
}

func (q *SelectQuery) ToSQL() (string, []interface{}, error) {
	return q.ds.ToSQL()
}

func (q *SelectQuery) Filter(f SelectQueryFilter) *SelectQuery {
	return f.Filter(q)
}

//#region goqu.SelectDataset API {{{

func (q *SelectQuery) As(alias string) *SelectQuery {
	return q.withDS(q.ds.As(alias))
}

func (q *SelectQuery) ClearLimit() *SelectQuery {
	return q.withDS(q.ds.ClearLimit())
}

func (q *SelectQuery) ClearOffset() *SelectQuery {
	return q.withDS(q.ds.ClearOffset())
}

func (q *SelectQuery) ClearOrder() *SelectQuery {
	return q.withDS(q.ds.ClearOrder())
}

func (q *SelectQuery) ClearSelect() *SelectQuery {
	return q.withDS(q.ds.ClearSelect())
}

func (q *SelectQuery) ClearWhere() *SelectQuery {
	return q.withDS(q.ds.ClearWhere())
}

func (q *SelectQuery) ClearWindow() *SelectQuery {
	return q.withDS(q.ds.ClearWindow())
}

func (q *SelectQuery) CompoundFromSelf() *SelectQuery {
	return q.withDS(q.ds.CompoundFromSelf())
}

func (q *SelectQuery) CrossJoin(table exp.Expression) *SelectQuery {
	return q.withDS(q.ds.CrossJoin(table))
}

func (q *SelectQuery) Distinct(on ...interface{}) *SelectQuery {
	return q.withDS(q.ds.Distinct(on...))
}

func (q *SelectQuery) ForKeyShare(waitOption exp.WaitOption) *SelectQuery {
	return q.withDS(q.ds.ForKeyShare(waitOption))
}

func (q *SelectQuery) ForNoKeyUpdate(waitOption exp.WaitOption) *SelectQuery {
	return q.withDS(q.ds.ForNoKeyUpdate(waitOption))
}

func (q *SelectQuery) ForShare(waitOption exp.WaitOption) *SelectQuery {
	return q.withDS(q.ds.ForShare(waitOption))
}

func (q *SelectQuery) ForUpdate(waitOption exp.WaitOption) *SelectQuery {
	return q.withDS(q.ds.ForUpdate(waitOption))
}

func (q *SelectQuery) From(from ...interface{}) *SelectQuery {
	return q.withDS(q.ds.From(from...))
}

func (q *SelectQuery) FromSelf() *SelectQuery {
	return q.withDS(q.ds.FromSelf())
}

func (q *SelectQuery) FullJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.FullJoin(table, condition))
}

func (q *SelectQuery) FullOuterJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.FullOuterJoin(table, condition))
}

func (q *SelectQuery) GroupBy(groupBy ...interface{}) *SelectQuery {
	return q.withDS(q.ds.GroupBy(groupBy...))
}

func (q *SelectQuery) Having(expressions ...exp.Expression) *SelectQuery {
	return q.withDS(q.ds.Having(expressions...))
}

func (q *SelectQuery) InnerJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.InnerJoin(table, condition))
}

func (q *SelectQuery) Intersect(other *SelectQuery) *SelectQuery {
	return q.withDS(q.ds.Intersect(other.ds))
}

func (q *SelectQuery) IntersectAll(other *SelectQuery) *SelectQuery {
	return q.withDS(q.ds.IntersectAll(other.ds))
}

func (q *SelectQuery) Join(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.Join(table, condition))
}

func (q *SelectQuery) LeftJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.LeftJoin(table, condition))
}

func (q *SelectQuery) LeftOuterJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.LeftOuterJoin(table, condition))
}

func (q *SelectQuery) Limit(limit uint) *SelectQuery {
	return q.withDS(q.ds.Limit(limit))
}

func (q *SelectQuery) LimitAll() *SelectQuery {
	return q.withDS(q.ds.LimitAll())
}

func (q *SelectQuery) NaturalFullJoin(table exp.Expression) *SelectQuery {
	return q.withDS(q.ds.NaturalFullJoin(table))
}

func (q *SelectQuery) NaturalJoin(table exp.Expression) *SelectQuery {
	return q.withDS(q.ds.NaturalJoin(table))
}

func (q *SelectQuery) NaturalLeftJoin(table exp.Expression) *SelectQuery {
	return q.withDS(q.ds.NaturalLeftJoin(table))
}

func (q *SelectQuery) NaturalRightJoin(table exp.Expression) *SelectQuery {
	return q.withDS(q.ds.NaturalRightJoin(table))
}

func (q *SelectQuery) Offset(offset uint) *SelectQuery {
	return q.withDS(q.ds.Offset(offset))
}

func (q *SelectQuery) Order(order ...exp.OrderedExpression) *SelectQuery {
	return q.withDS(q.ds.Order(order...))
}

func (q *SelectQuery) OrderAppend(order ...exp.OrderedExpression) *SelectQuery {
	return q.withDS(q.ds.OrderAppend(order...))
}

func (q *SelectQuery) OrderPrepend(order ...exp.OrderedExpression) *SelectQuery {
	return q.withDS(q.ds.OrderPrepend(order...))
}

func (q *SelectQuery) Prepared(prepared bool) *SelectQuery {
	return q.withDS(q.ds.Prepared(prepared))
}

func (q *SelectQuery) RightJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.RightJoin(table, condition))
}

func (q *SelectQuery) RightOuterJoin(table exp.Expression, condition exp.JoinCondition) *SelectQuery {
	return q.withDS(q.ds.RightOuterJoin(table, condition))
}

func (q *SelectQuery) Select(selects ...interface{}) *SelectQuery {
	return q.withDS(q.ds.Select(selects...))
}

func (q *SelectQuery) SelectAppend(selects ...interface{}) *SelectQuery {
	return q.withDS(q.ds.SelectAppend(selects...))
}

func (q *SelectQuery) Union(other *SelectQuery) *SelectQuery {
	return q.withDS(q.ds.Union(other.ds))
}

func (q *SelectQuery) UnionAll(other *SelectQuery) *SelectQuery {
	return q.withDS(q.ds.UnionAll(other.ds))
}

func (q *SelectQuery) Where(expressions ...exp.Expression) *SelectQuery {
	return q.withDS(q.ds.Where(expressions...))
}

func (q *SelectQuery) Window(ws ...exp.WindowExpression) *SelectQuery {
	return q.withDS(q.ds.Window(ws...))
}

func (q *SelectQuery) WindowAppend(ws ...exp.WindowExpression) *SelectQuery {
	return q.withDS(q.ds.WindowAppend(ws...))
}

func (q *SelectQuery) With(name string, subquery exp.Expression) *SelectQuery {
	return q.withDS(q.ds.With(name, subquery))
}

func (q *SelectQuery) WithRecursive(name string, subquery exp.Expression) *SelectQuery {
	return q.withDS(q.ds.WithRecursive(name, subquery))
}

//#endregion }}}

func (q *SelectQuery) Eq(lhs exp.Comparable, v interface{}) *SelectQuery {
	return q.Where(lhs.Eq(v))
}

func (q *SelectQuery) GtEq(lhs exp.Comparable, v interface{}) *SelectQuery {
	return q.Where(lhs.Gte(v))
}

func (q *SelectQuery) LtEq(lhs exp.Comparable, v interface{}) *SelectQuery {
	return q.Where(lhs.Lte(v))
}

func (q *SelectQuery) In(lhs exp.Inable, v interface{}) *SelectQuery {
	return q.Where(lhs.In(v))
}

func (q *SelectQuery) ILike(lhs exp.Likeable, v interface{}) *SelectQuery {
	return q.Where(lhs.ILike(v))
}

func (q *SelectQuery) InnerJoinEq(lhs, rhs exp.IdentifierExpression) *SelectQuery {
	return q.InnerJoin(goqu.T(rhs.GetTable()), goqu.On(lhs.Eq(rhs)))
}
