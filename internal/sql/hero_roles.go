package sql

import (
	"database/sql/driver"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type HeroRoles []nspb.HeroRole

func (s HeroRoles) ToInt64s() []int64 {
	if s == nil {
		return nil
	}

	r := make([]int64, len(s))

	for i, id := range s {
		r[i] = int64(id)
	}

	return r
}

func (s *HeroRoles) SetInt64s(arr []int64) {
	if arr == nil {
		*s = nil
		return
	}

	*s = make(HeroRoles, len(arr))

	for i, n := range arr {
		(*s)[i] = nspb.HeroRole(n)
	}
}

// Scan implements the sql.Scanner interface.
func (s *HeroRoles) Scan(src interface{}) error {
	return IntArrayScan(src, s)
}

// Value implements the driver.Valuer interface.
func (s *HeroRoles) Value() (driver.Value, error) {
	return IntArrayValue(s)
}
