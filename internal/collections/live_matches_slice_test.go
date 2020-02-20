package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func TestLiveMatchesSlice_Insert(t *testing.T) {
	type insertCase struct {
		Index int
		Match *models.LiveMatch
	}

	testCases := []struct {
		Subject  nscol.LiveMatchesSlice
		Insert   []*insertCase
		Expected nscol.LiveMatchesSlice
	}{
		// out-of-bounds
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
			},
			Insert: []*insertCase{
				{
					Index: -1,
					Match: &models.LiveMatch{
						MatchID: nspb.MatchID(2),
					},
				},
				{
					Index: 2,
					Match: &models.LiveMatch{
						MatchID: nspb.MatchID(2),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
			},
		},
		// start
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
			Insert: []*insertCase{
				{
					Index: 0,
					Match: &models.LiveMatch{
						MatchID: nspb.MatchID(3),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(3),
				},
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
		},
		// end
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
			Insert: []*insertCase{
				{
					Index: 2,
					Match: &models.LiveMatch{
						MatchID: nspb.MatchID(3),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
				{
					MatchID: nspb.MatchID(3),
				},
			},
		},
		// middle
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
			Insert: []*insertCase{
				{
					Index: 1,
					Match: &models.LiveMatch{
						MatchID: nspb.MatchID(3),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(3),
				},
				{
					MatchID: nspb.MatchID(2),
				},
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

func TestLiveMatchesSlice_Remove(t *testing.T) {
	type removeCase struct {
		Index    int
		Expected *models.LiveMatch
	}

	testCases := []struct {
		Subject  nscol.LiveMatchesSlice
		Remove   []*removeCase
		Expected nscol.LiveMatchesSlice
	}{
		// out-of-bounds
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
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
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
			},
		},
		// single element
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
			},
			Remove: []*removeCase{
				{
					Index: 0,
					Expected: &models.LiveMatch{
						MatchID: nspb.MatchID(1),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{},
		},
		// start
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
			Remove: []*removeCase{
				{
					Index: 0,
					Expected: &models.LiveMatch{
						MatchID: nspb.MatchID(1),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(2),
				},
			},
		},
		// end
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
			},
			Remove: []*removeCase{
				{
					Index: 1,
					Expected: &models.LiveMatch{
						MatchID: nspb.MatchID(2),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
			},
		},
		// middle
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
				{
					MatchID: nspb.MatchID(3),
				},
			},
			Remove: []*removeCase{
				{
					Index: 1,
					Expected: &models.LiveMatch{
						MatchID: nspb.MatchID(2),
					},
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(3),
				},
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

func TestLiveMatchesSlice_RemoveDeactivated(t *testing.T) {
	testCases := []struct {
		Subject nscol.LiveMatchesSlice
		Removed nscol.LiveMatchesSlice
		Result  nscol.LiveMatchesSlice
	}{
		{
			Subject: nil,
			Removed: nil,
			Result:  nil,
		},
		{
			Subject: nscol.LiveMatchesSlice{},
			Removed: nscol.LiveMatchesSlice{},
			Result:  nscol.LiveMatchesSlice{},
		},
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID:        nspb.MatchID(1),
					DeactivateTime: nil,
				},
				{
					MatchID:        nspb.MatchID(2),
					DeactivateTime: models.NullUnixTimestamp(2),
				},
				{
					MatchID:        nspb.MatchID(3),
					DeactivateTime: nil,
				},
				{
					MatchID:        nspb.MatchID(4),
					DeactivateTime: models.NullUnixTimestamp(4),
				},
				{
					MatchID:        nspb.MatchID(5),
					DeactivateTime: models.NullUnixTimestamp(5),
				},
			},
			Removed: nscol.LiveMatchesSlice{
				{
					MatchID:        nspb.MatchID(2),
					DeactivateTime: models.NullUnixTimestamp(2),
				},
				{
					MatchID:        nspb.MatchID(4),
					DeactivateTime: models.NullUnixTimestamp(4),
				},
				{
					MatchID:        nspb.MatchID(5),
					DeactivateTime: models.NullUnixTimestamp(5),
				},
			},
			Result: nscol.LiveMatchesSlice{
				{
					MatchID:        nspb.MatchID(1),
					DeactivateTime: nil,
				},
				{
					MatchID:        nspb.MatchID(3),
					DeactivateTime: nil,
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

			if actual.DeactivateTime == nil {
				t.Fatalf("case %d: index %d: expected removed non-nil DeactivateTime", testCaseIdx, i)
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

			if actual.DeactivateTime != nil {
				t.Fatalf("case %d: index %d: expected result nil DeactivateTime", testCaseIdx, i)
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

func TestLiveMatchesSlice_Batches(t *testing.T) {
	testCases := []struct {
		Subject   nscol.LiveMatchesSlice
		BatchSize int
		Expected  []nscol.LiveMatchesSlice
	}{
		{
			Subject:   nil,
			BatchSize: 1,
			Expected:  nil,
		},
		{
			Subject:   nscol.LiveMatchesSlice{},
			BatchSize: 1,
			Expected:  []nscol.LiveMatchesSlice{},
		},
		{
			Subject: nscol.LiveMatchesSlice{
				{
					MatchID: nspb.MatchID(1),
				},
				{
					MatchID: nspb.MatchID(2),
				},
				{
					MatchID: nspb.MatchID(3),
				},
				{
					MatchID: nspb.MatchID(4),
				},
				{
					MatchID: nspb.MatchID(5),
				},
			},
			BatchSize: 2,
			Expected: []nscol.LiveMatchesSlice{
				{
					{
						MatchID: nspb.MatchID(1),
					},
					{
						MatchID: nspb.MatchID(2),
					},
				},
				{
					{
						MatchID: nspb.MatchID(3),
					},
					{
						MatchID: nspb.MatchID(4),
					},
				},
				{
					{
						MatchID: nspb.MatchID(5),
					},
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
