package time_test

import (
	"testing"
	"time"

	nstime "github.com/13k/night-stalker/internal/time"
)

func TestEqualPtr(t *testing.T) {
	now := time.Now()
	future := now.Add(10 * time.Second)

	testCases := []struct {
		Left     *time.Time
		Right    *time.Time
		Expected bool
	}{
		{
			Left:     nil,
			Right:    nil,
			Expected: true,
		},
		{
			Left:     &now,
			Right:    nil,
			Expected: false,
		},
		{
			Left:     nil,
			Right:    &now,
			Expected: false,
		},
		{
			Left:     &now,
			Right:    &now,
			Expected: true,
		},
		{
			Left:     &now,
			Right:    &future,
			Expected: false,
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual := nstime.EqualPtr(testCase.Left, testCase.Right)

		if actual != testCase.Expected {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}
