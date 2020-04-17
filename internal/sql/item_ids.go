package sql

import (
	"database/sql/driver"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type ItemIDs []nspb.ItemID

var _ IntArrayScanner = (*ItemIDs)(nil)
var _ IntArrayValuer = (ItemIDs)(nil)

func NewItemIDs(ids []uint32) ItemIDs {
	if ids == nil {
		return nil
	}

	s := make(ItemIDs, len(ids))

	for i, id := range ids {
		s[i] = nspb.ItemID(id)
	}

	return s
}

func (s ItemIDs) ToUint64s() []uint64 {
	if s == nil {
		return nil
	}

	r := make([]uint64, len(s))

	for i, id := range s {
		r[i] = uint64(id)
	}

	return r
}

func (s ItemIDs) ToInt64s() []int64 {
	if s == nil {
		return nil
	}

	r := make([]int64, len(s))

	for i, id := range s {
		r[i] = int64(id)
	}

	return r
}

func (s *ItemIDs) SetInt64s(arr []int64) {
	if arr == nil {
		*s = nil
		return
	}

	*s = make(ItemIDs, len(arr))

	for i, n := range arr {
		(*s)[i] = nspb.ItemID(n)
	}
}

// Scan implements the sql.Scanner interface.
func (s *ItemIDs) Scan(src interface{}) error {
	return IntArrayScan(src, s)
}

// Value implements the driver.Valuer interface.
func (s ItemIDs) Value() (driver.Value, error) {
	return IntArrayValue(s)
}
