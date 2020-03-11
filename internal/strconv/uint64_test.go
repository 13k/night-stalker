package strconv_test

import (
	"math"
	"testing"

	nssconv "github.com/13k/night-stalker/internal/strconv"
)

func TestSafeParseUint(t *testing.T) {
	testCases := []struct {
		Subject  string
		Expected uint64
	}{
		{"1", 1},                                 // int
		{"+1", 1},                                // int explicit sign
		{"-1", 0},                                // int negative
		{"0", 0},                                 // int zero
		{"+0", 0},                                // int zero explicit sign
		{"-0", 0},                                // int zero negative
		{"18446744073709551615", math.MaxUint64}, // int max
		{"18446744073709551616", 0},              // int max+1
		{"18446744073709551614", uint64(math.MaxUint64) - 1}, // int max-1

		{"1.0", 1},  // float exact
		{"+1.0", 1}, // float exact explicit sign
		{"-1.0", 0}, // float exact negative
		{"1.", 1},   // float exact
		{"+1.", 1},  // float exact explicit sign
		{"-1.", 0},  // float exact negative
		{"0.0", 0},  // float exact zero
		{"+0.0", 0}, // float exact zero explicit sign
		{"-0.0", 0}, // float exact zero negative
		{".0", 0},   // float exact zero
		{"+.0", 0},  // float exact zero explicit sign
		{"-.0", 0},  // float exact zero negative
		{"0.", 0},   // float exact zero
		{"+0.", 0},  // float exact zero explicit sign
		{"-0.", 0},  // float exact zero negative
		{"18446744073709551615.0", math.MaxUint64},             // float exact max
		{"18446744073709551616.0", 0},                          // float exact max+1
		{"18446744073709551614.0", uint64(math.MaxUint64) - 1}, // float exact max-1

		{"1.234", 1},  // float decimal
		{"+1.234", 1}, // float decimal explicit sign
		{"-1.234", 0}, // float decimal negative
		{"0.123", 0},  // float decimal zero
		{"+0.123", 0}, // float decimal zero explicit sign
		{"-0.123", 0}, // float decimal zero negative
		{".123", 0},   // float decimal zero
		{"+.123", 0},  // float decimal zero explicit sign
		{"-.123", 0},  // float decimal zero negative
		{"18446744073709551615.1", math.MaxUint64},             // float decimal max
		{"18446744073709551616.1", 0},                          // float decimal max+1
		{"18446744073709551614.9", uint64(math.MaxUint64) - 1}, // float decimal max-1

		{"1.234e+3", 1234}, // float exponent exact
		{"1.234e+1", 12},   // float exponent decimal
		{"1.234e-19", 0},   // float exponent decimal
		{"1.234e-1", 0},    // float exponent decimal
		{"1.9e+19", 0},     // float exponent overflow
		{"-1.9e+19", 0},    // float exponent negative overflow
		{"0.0e+19", 0},     // float exponent zero exact

		{"leet", 0}, // invalid
		{"1eet", 0}, // invalid
		{"l33t", 0}, // invalid
		{"lee7", 0}, // invalid
		{"+", 0},    // invalid
		{"-", 0},    // invalid
	}

	for testCaseIdx, testCase := range testCases {
		actual := nssconv.SafeParseUint(testCase.Subject)

		if actual != testCase.Expected {
			t.Fatalf("case %d: expected %d, got %d (subject %q)", testCaseIdx, testCase.Expected, actual, testCase.Subject)
		}
	}
}
