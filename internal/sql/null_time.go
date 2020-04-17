package sql

import (
	"database/sql"
	"time"

	pbt "github.com/golang/protobuf/ptypes"
	pbts "github.com/golang/protobuf/ptypes/timestamp"

	nsjson "github.com/13k/night-stalker/internal/json"
)

var zeroNullTime = sql.NullTime{}

func NullTimeUnix(sec int64) sql.NullTime {
	if sec == 0 {
		return zeroNullTime
	}

	return sql.NullTime{
		Time:  time.Unix(sec, 0),
		Valid: true,
	}
}

func NullTimeIsZero(t sql.NullTime) bool {
	if !t.Valid {
		return true
	}

	return t.Time.IsZero()
}

func NullTimeEqual(left, right sql.NullTime) bool {
	if !left.Valid && !right.Valid {
		return true
	}

	if left.Valid && right.Valid {
		return left.Time.Equal(right.Time)
	}

	return false
}

// NullTimeProto converts a NullTime into a protobuf Timestamp.
//
// Returns nil with nil error if the given NullTime is invalid.
func NullTimeProto(t sql.NullTime) (*pbts.Timestamp, error) {
	if !t.Valid {
		return nil, nil
	}

	return pbt.TimestampProto(t.Time)
}

func NullTimeFromUnixJSON(jt *nsjson.UnixTime) sql.NullTime {
	if jt == nil {
		return zeroNullTime
	}

	return sql.NullTime{
		Time:  jt.Time,
		Valid: !jt.IsZero(),
	}
}
