package strconv

import (
	"math"
	"math/big"
)

// SafeParseUint parses a string numeric value into a uint64.
//
// For integer values, it accepts the formats defined in `math/big.Int.SetString()`.
//
// For float values, it accepts the formats defined in `math/big.Float.Parse()`. The value is parsed
// with a 64 bits precision and is rounded down (`math/big.ToZero`).
//
// Returns zero if it cannot parse the value (invalid/malformed, negative or overflows uint64).
func SafeParseUint(s string) uint64 {
	if z, ok := new(big.Int).SetString(s, 0); ok {
		if z.IsUint64() {
			return z.Uint64()
		}

		// valid integer but either negative or overflow
		return 0
	}

	if z, _, err := big.ParseFloat(s, 0, 64, big.ToZero); err == nil {
		i, acc := z.Uint64()

		// overflow
		if i == math.MaxUint64 && acc == big.Below {
			return 0
		}

		return i
	}

	// invalid
	return 0
}
