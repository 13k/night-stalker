package sql

import (
	"database/sql/driver"

	"github.com/lib/pq"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type HeroRoles []nspb.HeroRole

func NewHeroRolesFromInt64s(s []int64) HeroRoles {
	if s == nil {
		return nil
	}

	r := make(HeroRoles, len(s))

	for i, n := range s {
		r[i] = nspb.HeroRole(n)
	}

	return r
}

func (a HeroRoles) ToInt64s() []int64 {
	if a == nil {
		return nil
	}

	s := make([]int64, len(a))

	for i, r := range a {
		s[i] = int64(r)
	}

	return s
}

// Scan implements the sql.Scanner interface.
func (a *HeroRoles) Scan(src interface{}) error {
	var s pq.Int64Array

	if err := s.Scan(src); err != nil {
		return err
	}

	*a = NewHeroRolesFromInt64s(s)

	return nil
}

// Value implements the driver.Valuer interface.
func (a HeroRoles) Value() (driver.Value, error) {
	return pq.Int64Array(a.ToInt64s()).Value()
}
