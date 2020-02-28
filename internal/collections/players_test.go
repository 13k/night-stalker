package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func TestPlayers_AccountIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Players
		Expected nscol.AccountIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Players{},
			Expected: nscol.AccountIDs{},
		},
		{
			Subject: nscol.Players{
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

func TestPlayers_GroupByAccountID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Players
		Expected map[nspb.AccountID]nscol.Players
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Players{},
			Expected: map[nspb.AccountID]nscol.Players{},
		},
		{
			Subject: nscol.Players{
				{AccountID: 1},
				{AccountID: 2},
				{AccountID: 3},
				{AccountID: 3},
				{AccountID: 4},
			},
			Expected: map[nspb.AccountID]nscol.Players{
				1: {
					{AccountID: 1},
				},
				2: {
					{AccountID: 2},
				},
				3: {
					{AccountID: 3},
					{AccountID: 3},
				},
				4: {
					{AccountID: 4},
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

func TestPlayers_KeyByAccountID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Players
		Expected map[nspb.AccountID]*models.Player
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Players{},
			Expected: map[nspb.AccountID]*models.Player{},
		},
		{
			Subject: nscol.Players{
				{ID: 1, AccountID: 1},
				{ID: 2, AccountID: 2},
				{ID: 3, AccountID: 3},
				{ID: 4, AccountID: 3},
				{ID: 5, AccountID: 4},
			},
			Expected: map[nspb.AccountID]*models.Player{
				1: {ID: 1, AccountID: 1},
				2: {ID: 2, AccountID: 2},
				3: {ID: 4, AccountID: 3},
				4: {ID: 5, AccountID: 4},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.KeyByAccountID()

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

		for accountID, expectedElem := range testCase.Expected {
			actualElem := actual[accountID]

			if expectedElem == nil {
				if actualElem != nil {
					t.Fatalf("case %d: accountID %d: expected nil", testCaseIdx, accountID)
				}
			} else {
				if actualElem == nil {
					t.Fatalf("case %d: accountID %d: expected non-nil", testCaseIdx, accountID)
				}
			}

			if actualElem.ID != expectedElem.ID {
				t.Fatalf(
					"case %d: accountID %d: expected ID %d, got %d",
					testCaseIdx,
					accountID,
					expectedElem.ID,
					actualElem.ID,
				)
			}

			if actualElem.AccountID != expectedElem.AccountID {
				t.Fatalf(
					"case %d: accountID %d: expected AccountID %d, got %d",
					testCaseIdx,
					accountID,
					expectedElem.AccountID,
					actualElem.AccountID,
				)
			}
		}
	}
}
