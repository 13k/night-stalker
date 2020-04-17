package db

type Query interface {
	ToSQL() (string, []interface{}, error)
	IsTraceEnabled() bool
}

type query struct {
	b     SQLBuilder
	trace bool
}

func (q *query) clone() *query {
	return &query{
		b:     q.b,
		trace: q.trace,
	}
}

func (q *query) Builder() SQLBuilder {
	return q.b
}

func (q *query) IsTraceEnabled() bool {
	return q.trace
}
