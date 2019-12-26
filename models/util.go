package models

import (
	"math"
	"time"

	pbt "github.com/golang/protobuf/ptypes"
	pbts "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lib/pq"
)

func TruncateUint(i uint64) uint64 {
	isig := int64(i)

	if isig < 0 {
		return 0
	}

	return i
}

func NullUnixTimestamp(sec int64) *time.Time {
	if sec == 0 {
		return nil
	}

	t := time.Unix(sec, 0)

	return &t
}

func NullUnixTimestampFrac(fSec float64) *time.Time {
	if fSec == 0 {
		return nil
	}

	sec, frac := math.Modf(fSec)
	t := time.Unix(int64(sec), int64(frac*float64(time.Second)))

	return &t
}

func NullTimestampProto(t *time.Time) (*pbts.Timestamp, error) {
	if t == nil {
		return nil, nil
	}

	return pbt.TimestampProto(*t)
}

func Uint32Array(s []uint32) pq.Int64Array {
	arr := make([]int64, len(s))

	for i, n := range s {
		arr[i] = int64(n)
	}

	return arr
}

func Int32Array(s []int32) pq.Int64Array {
	arr := make([]int64, len(s))

	for i, n := range s {
		arr[i] = int64(n)
	}

	return arr
}
