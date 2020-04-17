package db

import (
	"github.com/doug-martin/goqu/v9"
)

type SQLBuilder interface {
	Insert(table interface{}) *goqu.InsertDataset
	Select(cols ...interface{}) *goqu.SelectDataset
	Update(table interface{}) *goqu.UpdateDataset
	Delete(table interface{}) *goqu.DeleteDataset
}
