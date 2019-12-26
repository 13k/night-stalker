package json

import (
	"bytes"
	"strconv"
)

/*
StringUint is an integer type used for JSON encoding.

Used to deserialize big unsigned integer values encoded as strings (base 10)
into uint64 and serialize uint64 into string (base 10).
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

// UnmarshalJSON deserializes JSON numeric string (base 10) into uint64 value.
func (si *StringUint) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*si = 0
		return nil
	}

	data = bytes.Trim(data, `"`)
	i, err := strconv.ParseUint(string(data), 10, 64)

	if err != nil {
		return err
	}

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
