package sql

import (
	"database/sql/driver"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type AccountIDs []nspb.AccountID

var _ IntArrayScanner = (*AccountIDs)(nil)
var _ IntArrayValuer = (AccountIDs)(nil)

func NewAccountIDs(ids []uint32) AccountIDs {
	if ids == nil {
		return nil
	}

	s := make(AccountIDs, len(ids))

	for i, id := range ids {
		s[i] = nspb.AccountID(id)
	}

	return s
}

func (s AccountIDs) ToInt64s() []int64 {
	if s == nil {
		return nil
	}

	r := make([]int64, len(s))

	for i, id := range s {
		r[i] = int64(id)
	}

	return r
}

func (s *AccountIDs) SetInt64s(arr []int64) {
	if arr == nil {
		*s = nil
		return
	}

	*s = make(AccountIDs, len(arr))

	for i, n := range arr {
		(*s)[i] = nspb.AccountID(n)
	}
}

// Scan implements the sql.Scanner interface.
func (s *AccountIDs) Scan(src interface{}) error {
	return IntArrayScan(src, s)
}

// Value implements the driver.Valuer interface.
func (s AccountIDs) Value() (driver.Value, error) {
	return IntArrayValue(s)
}
