package db

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/13k/night-stalker/models"
)

func WhereBinaryOp(db *gorm.DB, model models.Model, col string, op string, value interface{}) *gorm.DB {
	op = strings.ToUpper(op)
	quote := db.Dialect().Quote
	placeholder := "?"

	switch op {
	case "IN":
		placeholder = "(?)"
	}

	return db.Where(fmt.Sprintf("%s.%s %s %s", quote(model.TableName()), quote(col), op, placeholder), value)
}

func Eq(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, "=", value)
}

func Neq(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, "!=", value)
}

func In(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, "IN", value)
}

func Gt(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, ">", value)
}

func GtEq(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, ">=", value)
}

func Lt(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, "<", value)
}

func LtEq(db *gorm.DB, model models.Model, col string, value interface{}) *gorm.DB {
	return WhereBinaryOp(db, model, col, "<=", value)
}

func Group(db *gorm.DB, model models.Model, col string) *gorm.DB {
	quote := db.Dialect().Quote
	return db.Group(fmt.Sprintf("%s.%s", quote(model.TableName()), quote(col)))
}
