package models

import (
	"math"
	"time"

	pbt "github.com/golang/protobuf/ptypes"
	pbts "github.com/golang/protobuf/ptypes/timestamp"
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

// NullUnixTimestamp converts a UNIX timestamp (in seconds) to a `*time.Time`.
//
// Returns nil if the timestamp is zero.
func NullUnixTimestamp(sec int64) *time.Time {
	if sec == 0 {
		return nil
	}

	t := time.Unix(sec, 0)

	return &t
}

// NullUnixTimestamp converts a fractional UNIX timestamp (in seconds) to a `*time.Time`.
//
// Returns nil if the timestamp is zero.
func NullUnixTimestampFrac(fSec float64) *time.Time {
	if fSec == 0 {
		return nil
	}

	sec, frac := math.Modf(fSec)
	t := time.Unix(int64(sec), int64(frac*float64(time.Second)))

	return &t
}

// NullTimestampProto converts a Time to a protobuf Timestamp.
//
// Returns nil with nil error if the given Time is nil.
func NullTimestampProto(t *time.Time) (*pbts.Timestamp, error) {
	if t == nil {
		return nil, nil
	}

	return pbt.TimestampProto(*t)
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
