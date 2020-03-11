package json

import (
	"bytes"
	"strconv"

	nssconv "github.com/13k/night-stalker/internal/strconv"
)

/*
StringUint is an integer type used for JSON encoding.

It's used to decode JSON numeric values or numeric strings (base 10 or float) into uint64 and
serialize uint64 into string (base 10).
*/
type StringUint uint64

// Uint64 returns the value converted to uint64.
func (si StringUint) Uint64() uint64 {
	return uint64(si)
}

// Uint32 returns the value converted to uint32.
func (si StringUint) Uint32() uint32 {
	return uint32(si)
}

// UnmarshalJSON deserializes JSON numeric value or numeric string into uint64 value.
func (si *StringUint) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*si = 0
		return nil
	}

	data = bytes.Trim(data, `"`)
	i := nssconv.SafeParseUint(string(data))
	*si = StringUint(i)

	return nil
}

// MarshalJSON serializes the uint64 value to JSON numeric string (base 10).
func (si StringUint) MarshalJSON() ([]byte, error) {
	data := []byte{'"'}
	data = strconv.AppendUint(data, uint64(si), 10)
	data = append(data, '"')
	return data, nil
}
