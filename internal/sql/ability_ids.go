package sql

import (
	"database/sql/driver"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type AbilityIDs []nspb.AbilityID

var _ IntArrayScanner = (*AbilityIDs)(nil)
var _ IntArrayValuer = (AbilityIDs)(nil)

func NewAbilityIDs(ids []uint32) AbilityIDs {
	if ids == nil {
		return nil
	}

	s := make(AbilityIDs, len(ids))

	for i, id := range ids {
		s[i] = nspb.AbilityID(id)
	}

	return s
}

func (s AbilityIDs) ToInt64s() []int64 {
	if s == nil {
		return nil
	}

	r := make([]int64, len(s))

	for i, id := range s {
		r[i] = int64(id)
	}

	return r
}

func (s *AbilityIDs) SetInt64s(arr []int64) {
	if arr == nil {
		*s = nil
		return
	}

	*s = make(AbilityIDs, len(arr))

	for i, n := range arr {
		(*s)[i] = nspb.AbilityID(n)
	}
}

// Scan implements the sql.Scanner interface.
func (s *AbilityIDs) Scan(src interface{}) error {
	return IntArrayScan(src, s)
}

// Value implements the driver.Valuer interface.
func (s AbilityIDs) Value() (driver.Value, error) {
	return IntArrayValue(s)
}
