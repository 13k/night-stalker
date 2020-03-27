package models

import (
	"github.com/lib/pq"
)

// TruncateUint truncates an uint64 to the maximum positive int64 value.
//
// Returns zero if it overflows, otherwise returns the unmodified value.
func TruncateUint(i uint64) uint64 {
	isig := int64(i)

	if isig < 0 {
		return 0
	}

	return i
}

// Uint32Array converts a slice of uint32 values to a `pq.Int64Array`.
//
// Returns nil if the given slice is nil.
func Uint32Array(s []uint32) pq.Int64Array {
	if s == nil {
		return nil
	}

	arr := make([]int64, len(s))

	for i, n := range s {
		arr[i] = int64(n)
	}

	return arr
}

// Int32Array converts a slice of int32 values to a `pq.Int64Array`.
//
// Returns nil if the given slice is nil.
func Int32Array(s []int32) pq.Int64Array {
	if s == nil {
		return nil
	}

	arr := make([]int64, len(s))

	for i, n := range s {
		arr[i] = int64(n)
	}

	return arr
}
