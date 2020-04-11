package models

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

const (
	columnPK        = "id"
	columnCreatedAt = "created_at"
	columnUpdatedAt = "updated_at"
	columnDeletedAt = "deleted_at"
)

type Table interface {
	exp.IdentifierExpression

	PK() Column
	CreatedAt() Column
	UpdatedAt() Column
	DeletedAt() Column
}

type Column interface {
	exp.IdentifierExpression
}

var _ Table = (*table)(nil)

type table struct {
	exp.IdentifierExpression
}

func NewTable(name string) Table {
	return &table{IdentifierExpression: goqu.T(name)}
}

func (t *table) PK() Column {
	return t.Col(columnPK)
}

func (t *table) CreatedAt() Column {
	return t.Col(columnCreatedAt)
}

func (t *table) UpdatedAt() Column {
	return t.Col(columnUpdatedAt)
}

func (t *table) DeletedAt() Column {
	return t.Col(columnDeletedAt)
}
