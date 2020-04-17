package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsm "github.com/13k/night-stalker/models"
)

func TestTeams_KeyByID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.Teams
		Expected map[nsm.ID]*nsm.Team
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.Teams{},
			Expected: make(map[nsm.ID]*nsm.Team),
		},
		{
			Subject: nscol.Teams{
				{ID: 1},
				{ID: 2},
				{ID: 3},
				{ID: 3},
				{ID: 4},
				{ID: 4},
				{ID: 4},
			},
			Expected: map[nsm.ID]*nsm.Team{
				1: {ID: 1},
				2: {ID: 2},
				3: {ID: 3},
				4: {ID: 4},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.KeyByID()

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

		for id, expectedElem := range testCase.Expected {
			actualElem := actual[id]

			if expectedElem == nil {
				if actualElem != nil {
					t.Fatalf("case %d: ID %d: expected nil", testCaseIdx, id)
				}
			} else {
				if actualElem == nil {
					t.Fatalf("case %d: ID %d: expected non-nil", testCaseIdx, id)
				}
			}

			if actualElem.ID != expectedElem.ID {
				t.Fatalf(
					"case %d: ID %d: expected ID %d, got %d",
					testCaseIdx,
					id,
					expectedElem.ID,
					actualElem.ID,
				)
			}
		}
	}
}
