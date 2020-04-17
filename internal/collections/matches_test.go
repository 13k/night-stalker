package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

func TestMatches_MatchIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Matches
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Matches{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject: nscol.Matches{
				{ID: 1},
				{ID: 1},
				{ID: 2},
				{ID: 3},
				{ID: 3},
			},
			Expected: nscol.MatchIDs{1, 1, 2, 3, 3},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.MatchIDs()

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

func TestMatches_KeyByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Matches
		Expected map[nspb.MatchID]*nsm.Match
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Matches{},
			Expected: map[nspb.MatchID]*nsm.Match{},
		},
		{
			Subject: nscol.Matches{
				{ID: 1},
				{ID: 2},
				{ID: 3},
				{ID: 3},
				{ID: 4},
				{ID: 4},
				{ID: 4},
			},
			Expected: map[nspb.MatchID]*nsm.Match{
				1: {ID: 1},
				2: {ID: 2},
				3: {ID: 3},
				4: {ID: 4},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.KeyByMatchID()

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

		for matchID, expectedElem := range testCase.Expected {
			actualElem := actual[matchID]

			if expectedElem == nil {
				if actualElem != nil {
					t.Fatalf("case %d: matchID %d: expected nil", testCaseIdx, matchID)
				}
			} else {
				if actualElem == nil {
					t.Fatalf("case %d: matchID %d: expected non-nil", testCaseIdx, matchID)
				}
			}

			if actualElem.ID != expectedElem.ID {
				t.Fatalf(
					"case %d: matchID %d: expected ID %d, got %d",
					testCaseIdx,
					matchID,
					expectedElem.ID,
					actualElem.ID,
				)
			}
		}
	}
}
