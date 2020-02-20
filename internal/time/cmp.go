package time

import (
	"time"
)

func EqualPtr(left, right *time.Time) bool {
	// nil, nil
	// <addr>, <addr>
	if left == right {
		return true
	}

	// nil, <addr>
	// <addr>, nil
	if left == nil || right == nil {
		return false
	}

	// <addr1>, <addr2>
	return left.Equal(*right)
}
