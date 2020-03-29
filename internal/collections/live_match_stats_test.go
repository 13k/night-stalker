package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func TestLiveMatchStats_MatchIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStats
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStats{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject: nscol.LiveMatchStats{
				{MatchID: 1},
				{MatchID: 1},
				{MatchID: 2},
				{MatchID: 3},
				{MatchID: 3},
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

func TestLiveMatchStats_GroupByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStats
		Expected map[nspb.MatchID]nscol.LiveMatchStats
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStats{},
			Expected: map[nspb.MatchID]nscol.LiveMatchStats{},
		},
		{
			Subject: nscol.LiveMatchStats{
				{MatchID: 1, ID: 1},
				{MatchID: 1, ID: 2},
				{MatchID: 2, ID: 3},
				{MatchID: 3, ID: 3},
				{MatchID: 3, ID: 4},
			},
			Expected: map[nspb.MatchID]nscol.LiveMatchStats{
				1: {
					{MatchID: 1, ID: 1},
					{MatchID: 1, ID: 2},
				},
				2: {
					{MatchID: 2, ID: 3},
				},
				3: {
					{MatchID: 3, ID: 3},
					{MatchID: 3, ID: 4},
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.GroupByMatchID()

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

		for matchID, expectedGroup := range testCase.Expected {
			actualGroup := actual[matchID]

			if expectedGroup == nil {
				if actualGroup != nil {
					t.Fatalf("case %d: matchID %d: expected nil", testCaseIdx, matchID)
				}
			} else {
				if actualGroup == nil {
					t.Fatalf("case %d: matchID %d: expected non-nil", testCaseIdx, matchID)
				}
			}

			expectedGroupLen := len(expectedGroup)
			actualGroupLen := len(actualGroup)

			if actualGroupLen != expectedGroupLen {
				t.Fatalf("case %d: matchID %d: expected len %d, got %d", testCaseIdx, matchID, expectedGroupLen, actualGroupLen)
			}

			for groupIdx, expectedElem := range expectedGroup {
				actualElem := actualGroup[groupIdx]

				if actualElem.MatchID != expectedElem.MatchID {
					t.Fatalf(
						"case %d: matchID %d: index %d: expected MatchID %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.MatchID,
						actualElem.MatchID,
					)
				}

				if actualElem.ID != expectedElem.ID {
					t.Fatalf(
						"case %d: matchID %d: index %d: expected ID %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.ID,
						actualElem.ID,
					)
				}
			}
		}
	}
}

func TestLiveMatchStats_KeyByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStats
		Expected map[nspb.MatchID]*models.LiveMatchStats
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStats{},
			Expected: map[nspb.MatchID]*models.LiveMatchStats{},
		},
		{
			Subject: nscol.LiveMatchStats{
				{ID: 1, MatchID: 1},
				{ID: 2, MatchID: 2},
				{ID: 3, MatchID: 3},
				{ID: 4, MatchID: 3},
				{ID: 5, MatchID: 4},
			},
			Expected: map[nspb.MatchID]*models.LiveMatchStats{
				1: {ID: 1, MatchID: 1},
				2: {ID: 2, MatchID: 2},
				3: {ID: 4, MatchID: 3},
				4: {ID: 5, MatchID: 4},
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

			if actualElem.MatchID != expectedElem.MatchID {
				t.Fatalf(
					"case %d: matchID %d: expected MatchID %d, got %d",
					testCaseIdx,
					matchID,
					expectedElem.MatchID,
					actualElem.MatchID,
				)
			}
		}
	}
}
