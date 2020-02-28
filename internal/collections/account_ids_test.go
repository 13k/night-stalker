package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
)

func TestAccountIDs_AddUnique(t *testing.T) {
	testCases := []struct {
		Subject  nscol.AccountIDs
		Add      nscol.AccountIDs
		Expected nscol.AccountIDs
	}{
		{
			Subject:  nil,
			Add:      nil,
			Expected: nil,
		},
		{
			Subject:  nil,
			Add:      nscol.AccountIDs{},
			Expected: nil,
		},
		{
			Subject:  nscol.AccountIDs{},
			Add:      nil,
			Expected: nscol.AccountIDs{},
		},
		{
			Subject:  nscol.AccountIDs{},
			Add:      nscol.AccountIDs{},
			Expected: nscol.AccountIDs{},
		},
		{
			Subject:  nscol.AccountIDs{},
			Add:      nscol.AccountIDs{1},
			Expected: nscol.AccountIDs{1},
		},
		{
			Subject:  nscol.AccountIDs{1},
			Add:      nscol.AccountIDs{1},
			Expected: nscol.AccountIDs{1},
		},
		{
			Subject:  nscol.AccountIDs{1, 1, 1},
			Add:      nscol.AccountIDs{1, 1, 1},
			Expected: nscol.AccountIDs{1, 1, 1},
		},
		{
			Subject:  nscol.AccountIDs{1, 1, 2, 2, 3, 3},
			Add:      nscol.AccountIDs{1, 2, 3},
			Expected: nscol.AccountIDs{1, 1, 2, 2, 3, 3},
		},
		{
			Subject:  nscol.AccountIDs{1, 1, 3, 3, 5, 5},
			Add:      nscol.AccountIDs{1, 2, 3, 4, 5, 6},
			Expected: nscol.AccountIDs{1, 1, 3, 3, 5, 5, 2, 4, 6},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.AddUnique(testCase.Add...)

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
