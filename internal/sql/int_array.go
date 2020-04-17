package sql

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

type IntArrayScanner interface {
	SetInt64s([]int64)
}

type IntArrayValuer interface {
	ToInt64s() []int64
}

func IntArrayScan(src interface{}, dst IntArrayScanner) error {
	var arr pq.Int64Array

	if err := arr.Scan(src); err != nil {
		return err
	}

	dst.SetInt64s(arr)

	return nil
}

func IntArrayValue(s IntArrayValuer) (driver.Value, error) {
	return pq.Int64Array(s.ToInt64s()).Value()
}

func IntArrayEqual(left, right IntArrayValuer) bool {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	sLeft, sRight := left.ToInt64s(), right.ToInt64s()

	if len(sLeft) != len(sRight) {
		return false
	}

	for i, n := range sLeft {
		if sRight[i] != n {
			return false
		}
	}

	return true
}
