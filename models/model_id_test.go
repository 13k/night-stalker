package models_test

import (
	"testing"

	nsm "github.com/13k/night-stalker/models"
)

func TestNewUniqueIDs(t *testing.T) {
	testCases := []struct {
		Subject  nsm.IDs
		Expected nsm.IDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nsm.IDs{},
			Expected: nil,
		},
		{
			Subject:  nsm.IDs{1},
			Expected: nsm.IDs{1},
		},
		{
			Subject:  nsm.IDs{1, 1, 1},
			Expected: nsm.IDs{1},
		},
		{
			Subject:  nsm.IDs{1, 1, 2, 2, 3, 3, 1, 2, 3},
			Expected: nsm.IDs{1, 2, 3},
		},
		{
			Subject:  nsm.IDs{1, 1, 3, 3, 5, 5, 1, 2, 3, 4, 5, 6},
			Expected: nsm.IDs{1, 3, 5, 2, 4, 6},
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual := nsm.NewUniqueIDs(testCase.Subject...)

		if testCase.Expected == nil {
			if actual != nil {
				t.Fatalf("case %d: expected nil", testCaseIdx)
			}
		} else {
			if actual == nil {
				t.Fatalf("case %d: expected non-nil", testCaseIdx)
			}
		}

		expectedLen := len(testCase.Expected)
		actualLen := len(actual)

		if actualLen != expectedLen {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for i, expectedID := range testCase.Expected {
			actualID := actual[i]

			if actualID != expectedID {
				t.Fatalf("case %d: index %d: expected %d, got %d", testCaseIdx, i, expectedID, actualID)
			}
		}
	}
}
