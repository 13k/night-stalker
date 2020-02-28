package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
)

func TestFollowedPlayers_AccountIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.FollowedPlayers
		Expected nscol.AccountIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.FollowedPlayers{},
			Expected: nscol.AccountIDs{},
		},
		{
			Subject: nscol.FollowedPlayers{
				{AccountID: 1},
				{AccountID: 1},
				{AccountID: 2},
				{AccountID: 3},
				{AccountID: 3},
			},
			Expected: nscol.AccountIDs{1, 1, 2, 3, 3},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.AccountIDs()

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
