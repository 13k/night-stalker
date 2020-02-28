package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
)

func TestLiveMatchStatsPlayers_MatchIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStatsPlayers
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStatsPlayers{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject: nscol.LiveMatchStatsPlayers{
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

func TestLiveMatchStatsPlayers_AccountIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStatsPlayers
		Expected nscol.AccountIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStatsPlayers{},
			Expected: nscol.AccountIDs{},
		},
		{
			Subject: nscol.LiveMatchStatsPlayers{
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

func TestLiveMatchStatsPlayers_GroupByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStatsPlayers
		Expected map[nspb.MatchID]nscol.LiveMatchStatsPlayers
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStatsPlayers{},
			Expected: map[nspb.MatchID]nscol.LiveMatchStatsPlayers{},
		},
		{
			Subject: nscol.LiveMatchStatsPlayers{
				{MatchID: 1, AccountID: 1},
				{MatchID: 1, AccountID: 2},
				{MatchID: 2, AccountID: 3},
				{MatchID: 3, AccountID: 3},
				{MatchID: 3, AccountID: 4},
			},
			Expected: map[nspb.MatchID]nscol.LiveMatchStatsPlayers{
				1: {
					{MatchID: 1, AccountID: 1},
					{MatchID: 1, AccountID: 2},
				},
				2: {
					{MatchID: 2, AccountID: 3},
				},
				3: {
					{MatchID: 3, AccountID: 3},
					{MatchID: 3, AccountID: 4},
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
						"case %d: matchID %d: group %d: expected MatchID %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.MatchID,
						actualElem.MatchID,
					)
				}

				if actualElem.AccountID != expectedElem.AccountID {
					t.Fatalf(
						"case %d: matchID %d: group %d: expected AccountID %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.AccountID,
						actualElem.AccountID,
					)
				}
			}
		}
	}
}

func TestLiveMatchStatsPlayers_GroupByAccountID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatchStatsPlayers
		Expected map[nspb.AccountID]nscol.LiveMatchStatsPlayers
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatchStatsPlayers{},
			Expected: map[nspb.AccountID]nscol.LiveMatchStatsPlayers{},
		},
		{
			Subject: nscol.LiveMatchStatsPlayers{
				{MatchID: 1, AccountID: 1},
				{MatchID: 1, AccountID: 2},
				{MatchID: 2, AccountID: 3},
				{MatchID: 3, AccountID: 3},
				{MatchID: 3, AccountID: 4},
			},
			Expected: map[nspb.AccountID]nscol.LiveMatchStatsPlayers{
				1: {
					{MatchID: 1, AccountID: 1},
				},
				2: {
					{MatchID: 1, AccountID: 2},
				},
				3: {
					{MatchID: 2, AccountID: 3},
					{MatchID: 3, AccountID: 3},
				},
				4: {
					{MatchID: 3, AccountID: 4},
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.GroupByAccountID()

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

		for accountID, expectedGroup := range testCase.Expected {
			actualGroup := actual[accountID]

			if expectedGroup == nil {
				if actualGroup != nil {
					t.Fatalf("case %d: accountID %d: expected nil", testCaseIdx, accountID)
				}
			} else {
				if actualGroup == nil {
					t.Fatalf("case %d: accountID %d: expected non-nil", testCaseIdx, accountID)
				}
			}

			expectedGroupLen := len(expectedGroup)
			actualGroupLen := len(actualGroup)

			if actualGroupLen != expectedGroupLen {
				t.Fatalf("case %d: accountID %d: expected len %d, got %d", testCaseIdx, accountID, expectedGroupLen, actualGroupLen)
			}

			for groupIdx, expectedElem := range expectedGroup {
				actualElem := actualGroup[groupIdx]

				if actualElem.MatchID != expectedElem.MatchID {
					t.Fatalf(
						"case %d: accountID %d: group %d: expected MatchID %d, got %d",
						testCaseIdx,
						accountID,
						groupIdx,
						expectedElem.MatchID,
						actualElem.MatchID,
					)
				}

				if actualElem.AccountID != expectedElem.AccountID {
					t.Fatalf(
						"case %d: accountID %d: group %d: expected AccountID %d, got %d",
						testCaseIdx,
						accountID,
						groupIdx,
						expectedElem.AccountID,
						actualElem.AccountID,
					)
				}
			}
		}
	}
}
