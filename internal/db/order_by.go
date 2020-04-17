package db

import (
	"github.com/doug-martin/goqu/v9/exp"
)

var (
	OrderFields OrderByFields
)

type OrderByField struct {
	Field     string
	Ascending bool
}

func NewOrderByField(f string, asc bool) *OrderByField {
	return &OrderByField{
		Field:     f,
		Ascending: asc,
	}
}

type OrderByFields []*OrderByField

func (s OrderByFields) Asc(f string) OrderByFields {
	return append(s, NewOrderByField(f, true))
}

func (s OrderByFields) Desc(f string) OrderByFields {
	return append(s, NewOrderByField(f, false))
}

type OrderByFieldMapper func(string) exp.Orderable

func (s OrderByFields) MapOrderBy(fn OrderByFieldMapper) []exp.OrderedExpression {
	if len(s) == 0 {
		return nil
	}

	orderBy := make([]exp.OrderedExpression, len(s))

	for i, of := range s {
		orderable := fn(of.Field)

		if of.Ascending {
			orderBy[i] = orderable.Asc()
		} else {
			orderBy[i] = orderable.Desc()
		}
	}

	return orderBy
}
