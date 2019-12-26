package json

import (
	"strconv"
	"time"
)

/*
UnixTime is a time.Time wrapper type used for JSON encoding.

Used to deserialize unix timestamps encoded as integer into time.Time and
serialize time.Time to integer.

See UnmarshalJSON for details on deserialization.
*/
type UnixTime struct {
	*time.Time
}

// UnmarshalJSON deserializes JSON integer into Time value.
func (ut *UnixTime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		ut.Time = nil
		return nil
	}

	secs, err := strconv.ParseInt(string(data), 10, 64)

	if err != nil {
		return err
	}

	if secs == 0 {
		ut.Time = nil
	} else {
		t := time.Unix(secs, 0)
		ut.Time = &t
	}

	return nil
}

// MarshalJSON serializes the Time value to JSON integer.
func (ut UnixTime) MarshalJSON() ([]byte, error) {
	if ut.Time == nil {
		return nil, nil
	}

	data := make([]byte, 0)
	data = strconv.AppendInt(data, ut.Unix(), 10)

	return data, nil
}
