package collections_test

import (
	"database/sql"
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/13k/night-stalker/models"
)

func TestLiveMatches_Swap(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatches
		Args     [2]int
		Expected nscol.LiveMatches
	}{
		{
			Subject: nscol.LiveMatches{
				{ID: 1},
				{ID: 2},
			},
			Args: [2]int{0, 0},
			Expected: nscol.LiveMatches{
				{ID: 1},
				{ID: 2},
			},
		},
		{
			Subject: nscol.LiveMatches{
				{ID: 1},
				{ID: 2},
			},
			Args: [2]int{0, 1},
			Expected: nscol.LiveMatches{
				{ID: 2},
				{ID: 1},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		args := testCase.Args
		subject := testCase.Subject

		subject.Swap(args[0], args[1])

		expectedLen := len(testCase.Expected)
		actualLen := len(subject)

		if actualLen != expectedLen {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for elemIdx, expectedElem := range testCase.Expected {
			actualElem := subject[elemIdx]

			if actualElem.ID != expectedElem.ID {
				t.Fatalf("case %d: index %d: expected %d, got %d", testCaseIdx, elemIdx, expectedElem.ID, actualElem.ID)
			}
		}
	}
}

func TestLiveMatches_MatchIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatches
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatches{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject: nscol.LiveMatches{
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

func TestLiveMatches_KeyByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatches
		Expected map[nspb.MatchID]*models.LiveMatch
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.LiveMatches{},
			Expected: map[nspb.MatchID]*models.LiveMatch{},
		},
		{
			Subject: nscol.LiveMatches{
				{ID: 1, MatchID: 1},
				{ID: 2, MatchID: 2},
				{ID: 3, MatchID: 3},
				{ID: 4, MatchID: 3},
				{ID: 5, MatchID: 4},
			},
			Expected: map[nspb.MatchID]*models.LiveMatch{
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

func TestLiveMatches_Insert(t *testing.T) {
	type insertCase struct {
		Index int
		Match *models.LiveMatch
	}

	testCases := []struct {
		Subject  nscol.LiveMatches
		Insert   []*insertCase
		Expected nscol.LiveMatches
	}{
		// out-of-bounds
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
			},
			Insert: []*insertCase{
				{
					Index: -1,
					Match: &models.LiveMatch{MatchID: 2},
				},
				{
					Index: 2,
					Match: &models.LiveMatch{MatchID: 2},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
			},
		},
		// start
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
			},
			Insert: []*insertCase{
				{
					Index: 0,
					Match: &models.LiveMatch{MatchID: 3},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 3},
				{MatchID: 1},
				{MatchID: 2},
			},
		},
		// end
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
			},
			Insert: []*insertCase{
				{
					Index: 2,
					Match: &models.LiveMatch{MatchID: 3},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
				{MatchID: 3},
			},
		},
		// middle
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
			},
			Insert: []*insertCase{
				{
					Index: 1,
					Match: &models.LiveMatch{MatchID: 3},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 3},
				{MatchID: 2},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject

		for _, insertCase := range testCase.Insert {
			subject.Insert(insertCase.Index, insertCase.Match)
		}

		if actualLen := subject.Len(); actualLen != len(testCase.Expected) {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, len(testCase.Expected), actualLen)
		}

		for matchIdx, expectedMatch := range testCase.Expected {
			actualMatch := subject[matchIdx]

			if actualMatch == nil {
				t.Fatalf("case %d: index %d: expected non-nil", testCaseIdx, matchIdx)
			}

			if actualMatch.MatchID != expectedMatch.MatchID {
				t.Fatalf(
					"case %d: index %d: expected MatchID to be %d, got %d",
					testCaseIdx,
					matchIdx,
					expectedMatch.MatchID,
					actualMatch.MatchID,
				)
			}
		}
	}
}

func TestLiveMatches_Remove(t *testing.T) {
	type removeCase struct {
		Index    int
		Expected *models.LiveMatch
	}

	testCases := []struct {
		Subject  nscol.LiveMatches
		Remove   []*removeCase
		Expected nscol.LiveMatches
	}{
		// out-of-bounds
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
			},
			Remove: []*removeCase{
				{
					Index:    -1,
					Expected: nil,
				},
				{
					Index:    1,
					Expected: nil,
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
			},
		},
		// single element
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
			},
			Remove: []*removeCase{
				{
					Index:    0,
					Expected: &models.LiveMatch{MatchID: 1},
				},
			},
			Expected: nscol.LiveMatches{},
		},
		// start
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
			},
			Remove: []*removeCase{
				{
					Index:    0,
					Expected: &models.LiveMatch{MatchID: 1},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 2},
			},
		},
		// end
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
			},
			Remove: []*removeCase{
				{
					Index:    1,
					Expected: &models.LiveMatch{MatchID: 2},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
			},
		},
		// middle
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
				{MatchID: 3},
			},
			Remove: []*removeCase{
				{
					Index:    1,
					Expected: &models.LiveMatch{MatchID: 2},
				},
			},
			Expected: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 3},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject

		for removeCaseIdx, removeCase := range testCase.Remove {
			actual := subject.Remove(removeCase.Index)

			if removeCase.Expected == nil {
				if actual != nil {
					t.Fatalf("case %d: removeCase %d: expected nil, got %#v", testCaseIdx, removeCaseIdx, actual)
				}
			} else {
				if actual == nil {
					t.Fatalf("case %d: removeCase %d: expected non-nil", testCaseIdx, removeCaseIdx)
				}

				if actual.MatchID != removeCase.Expected.MatchID {
					t.Fatalf("case %d: removeCase %d: expected %#v, got %#v", testCaseIdx, removeCaseIdx, removeCase.Expected, actual)
				}
			}
		}

		if actualLen := subject.Len(); actualLen != len(testCase.Expected) {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, len(testCase.Expected), actualLen)
		}

		for matchIdx, expectedMatch := range testCase.Expected {
			actualMatch := subject[matchIdx]

			if actualMatch == nil {
				t.Fatalf("case %d: index %d: expected non-nil", testCaseIdx, matchIdx)
			}

			if actualMatch.MatchID != expectedMatch.MatchID {
				t.Fatalf(
					"case %d: index %d: expected MatchID to be %d, got %d",
					testCaseIdx,
					matchIdx,
					expectedMatch.MatchID,
					actualMatch.MatchID,
				)
			}
		}
	}
}

func TestLiveMatches_RemoveDeactivated(t *testing.T) {
	testCases := []struct {
		Subject nscol.LiveMatches
		Removed nscol.LiveMatches
		Result  nscol.LiveMatches
	}{
		{
			Subject: nil,
			Removed: nil,
			Result:  nil,
		},
		{
			Subject: nscol.LiveMatches{},
			Removed: nscol.LiveMatches{},
			Result:  nscol.LiveMatches{},
		},
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:        1,
					DeactivateTime: sql.NullTime{},
				},
				{
					MatchID:        2,
					DeactivateTime: nssql.NullTimeUnix(2),
				},
				{
					MatchID:        3,
					DeactivateTime: sql.NullTime{},
				},
				{
					MatchID:        4,
					DeactivateTime: nssql.NullTimeUnix(4),
				},
				{
					MatchID:        5,
					DeactivateTime: nssql.NullTimeUnix(5),
				},
			},
			Removed: nscol.LiveMatches{
				{
					MatchID:        2,
					DeactivateTime: nssql.NullTimeUnix(2),
				},
				{
					MatchID:        4,
					DeactivateTime: nssql.NullTimeUnix(4),
				},
				{
					MatchID:        5,
					DeactivateTime: nssql.NullTimeUnix(5),
				},
			},
			Result: nscol.LiveMatches{
				{
					MatchID:        1,
					DeactivateTime: sql.NullTime{},
				},
				{
					MatchID:        3,
					DeactivateTime: sql.NullTime{},
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		removed := testCase.Subject.RemoveDeactivated()

		if testCase.Removed == nil {
			if removed != nil {
				t.Fatalf("case %d: expected nil, got %#v", testCaseIdx, removed)
			}
		} else {
			if removed == nil {
				t.Fatalf("case %d: expected non-nil", testCaseIdx)
			}
		}

		expectedRemovedLen := testCase.Removed.Len()

		if actualRemovedLen := removed.Len(); actualRemovedLen != expectedRemovedLen {
			t.Fatalf("case %d: expected removed length to be %d, got %d", testCaseIdx, expectedRemovedLen, actualRemovedLen)
		}

		for i, expected := range testCase.Removed {
			actual := removed[i]

			if actual == nil {
				t.Fatalf("case %d: index %d: expected removed non-nil", testCaseIdx, i)
			}

			if !actual.DeactivateTime.Valid {
				t.Fatalf("case %d: index %d: expected removed valid DeactivateTime", testCaseIdx, i)
			}

			if actual.MatchID != expected.MatchID {
				t.Fatalf(
					"case %d: index %d: expected removed MatchID to be %d, got %d",
					testCaseIdx,
					i,
					expected.MatchID,
					actual.MatchID,
				)
			}
		}

		expectedResultLen := testCase.Result.Len()

		if actualResultLen := testCase.Subject.Len(); actualResultLen != expectedResultLen {
			t.Fatalf("case %d: expected result length to be %d, got %d", testCaseIdx, expectedResultLen, actualResultLen)
		}

		for i, expected := range testCase.Result {
			actual := testCase.Subject[i]

			if actual == nil {
				t.Fatalf("case %d: index %d: expected result non-nil", testCaseIdx, i)
			}

			if actual.DeactivateTime.Valid {
				t.Fatalf("case %d: index %d: expected result invalid DeactivateTime", testCaseIdx, i)
			}

			if actual.MatchID != expected.MatchID {
				t.Fatalf(
					"case %d: index %d: expected result MatchID to be %d, got %d",
					testCaseIdx,
					i,
					expected.MatchID,
					actual.MatchID,
				)
			}
		}
	}
}

func TestLiveMatches_Batches(t *testing.T) {
	testCases := []struct {
		Subject   nscol.LiveMatches
		BatchSize int
		Expected  []nscol.LiveMatches
	}{
		{
			Subject:   nil,
			BatchSize: 1,
			Expected:  nil,
		},
		{
			Subject:   nscol.LiveMatches{},
			BatchSize: 1,
			Expected:  []nscol.LiveMatches{},
		},
		{
			Subject: nscol.LiveMatches{
				{MatchID: 1},
				{MatchID: 2},
				{MatchID: 3},
				{MatchID: 4},
				{MatchID: 5},
			},
			BatchSize: 2,
			Expected: []nscol.LiveMatches{
				{
					{MatchID: 1},
					{MatchID: 2},
				},
				{
					{MatchID: 3},
					{MatchID: 4},
				},
				{
					{MatchID: 5},
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual := testCase.Subject.Batches(testCase.BatchSize)
		expectedLen := len(testCase.Expected)

		if testCase.Expected == nil && actual != nil {
			t.Fatalf("case %d: expected nil, got slice [len: %d, cap: %d]", testCaseIdx, len(actual), cap(actual))
		}

		if testCase.Expected != nil && actual == nil {
			t.Fatalf("case %d: expected non-nil", testCaseIdx)
		}

		if actualLen := len(actual); actualLen != expectedLen {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for batchIdx, expectedBatch := range testCase.Expected {
			expectedBatchLen := len(expectedBatch)
			actualBatch := actual[batchIdx]

			if actualBatchLen := len(actualBatch); actualBatchLen != expectedBatchLen {
				t.Fatalf("case %d: batch %d: expected len %d, got %d", testCaseIdx, batchIdx, expectedBatchLen, actualBatchLen)
			}

			for matchIdx, expectedMatch := range expectedBatch {
				actualMatch := actualBatch[matchIdx]

				if actualMatch == nil {
					t.Fatalf("case %d: batch %d: index %d: expected *LiveMatch, got nil", testCaseIdx, batchIdx, matchIdx)
				}

				if actualMatch.MatchID != expectedMatch.MatchID {
					t.Fatalf(
						"case %d: batch %d: index %d: expected MatchID to be %d, got %d",
						testCaseIdx,
						batchIdx,
						matchIdx,
						expectedMatch.MatchID,
						actualMatch.MatchID,
					)
				}
			}
		}
	}
}
