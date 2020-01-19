package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func TestLiveMatchesByScore_SearchIndex(t *testing.T) {
	testCases := []struct {
		Matches  nscol.LiveMatchesSlice
		Match    *models.LiveMatch
		Expected int
	}{
		{
			Matches: nil,
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(1),
				SortScore: float64(1.0),
			},
			Expected: 0,
		},
		// greater than all
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(4),
				SortScore: float64(4.0),
			},
			Expected: 0,
		},
		// less than all
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(4),
				SortScore: float64(0.0),
			},
			Expected: 3,
		},
		// equal 2
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(2),
				SortScore: float64(2.0),
			},
			Expected: 1,
		},
		// greater than 2
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(4),
				SortScore: float64(2.1),
			},
			Expected: 1,
		},
		// less than 2
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   nspb.MatchID(4),
				SortScore: float64(1.9),
			},
			Expected: 2,
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Matches...)
		actual := subject.SearchIndex(testCase.Match)

		if actual != testCase.Expected {
			t.Fatalf("case %d: expected %d, got %d", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestLiveMatchesByScore_Add(t *testing.T) {
	type addCase struct {
		Match    *models.LiveMatch
		Expected bool
	}

	testCases := []struct {
		Matches  nscol.LiveMatchesSlice
		Add      []*addCase
		Expected nscol.LiveMatchesSlice
	}{
		// insert
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(2),
						SortScore: float64(2.0),
					},
					Expected: true,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
		},
		// noop
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(1),
						SortScore: float64(1.0),
					},
					Expected: false,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
		},
		// update/reorder
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(1),
						SortScore: float64(5.0),
					},
					Expected: true,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(5.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
			},
		},
		// complex
		{
			Matches: nil,
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(1),
						SortScore: float64(1.0),
					},
					Expected: true,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(2),
						SortScore: float64(2.0),
					},
					Expected: true,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(1),
						SortScore: float64(1.0),
					},
					Expected: false,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(3),
						SortScore: float64(3.0),
					},
					Expected: true,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(1),
						SortScore: float64(5.0),
					},
					Expected: true,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(2),
						SortScore: float64(2.0),
					},
					Expected: false,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   nspb.MatchID(3),
						SortScore: float64(3.1),
					},
					Expected: true,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(5.0),
				},
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.1),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Matches...)

		for addCaseIdx, addCase := range testCase.Add {
			if actual := subject.Add(addCase.Match); actual != addCase.Expected {
				t.Fatalf("case %d: addCase %d: expected %v, got %v", testCaseIdx, addCaseIdx, addCase.Expected, actual)
			}
		}

		if actualLen := subject.Len(); actualLen != len(testCase.Expected) {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, len(testCase.Expected), actualLen)
		}

		for matchIdx, expectedMatch := range testCase.Expected {
			actualMatch := subject.At(matchIdx)

			if actualMatch == nil {
				t.Fatalf("case %d: expected *LiveMatch, got nil", testCaseIdx)
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

			if actualMatch.SortScore != expectedMatch.SortScore {
				t.Fatalf(
					"case %d: index %d: expected SortScore to be %f, got %f",
					testCaseIdx,
					matchIdx,
					expectedMatch.SortScore,
					actualMatch.SortScore,
				)
			}
		}
	}
}

func TestLiveMatchesByScore_Remove(t *testing.T) {
	type removeCase struct {
		MatchID    nspb.MatchID
		Expected bool
	}

	testCases := []struct {
		Matches  nscol.LiveMatchesSlice
		Remove      []*removeCase
		Expected nscol.LiveMatchesSlice
	}{
		// remove
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(2),
					SortScore: float64(2.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Remove: []*removeCase{
				{
					MatchID: nspb.MatchID(2),
					Expected: true,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(3),
					SortScore: float64(3.0),
				},
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
		},
		// noop
		{
			Matches: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
			Remove: []*removeCase{
				{
					MatchID: nspb.MatchID(2),
					Expected: false,
				},
			},
			Expected: nscol.LiveMatchesSlice{
				{
					MatchID:   nspb.MatchID(1),
					SortScore: float64(1.0),
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Matches...)

		for removeCaseIdx, removeCase := range testCase.Remove {
			if actual := subject.Remove(removeCase.MatchID); actual != removeCase.Expected {
				t.Fatalf("case %d: removeCase %d: expected %v, got %v", testCaseIdx, removeCaseIdx, removeCase.Expected, actual)
			}
		}

		if actualLen := subject.Len(); actualLen != len(testCase.Expected) {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, len(testCase.Expected), actualLen)
		}

		for matchIdx, expectedMatch := range testCase.Expected {
			actualMatch := subject.At(matchIdx)

			if actualMatch == nil {
				t.Fatalf("case %d: expected *LiveMatch, got nil", testCaseIdx)
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

			if actualMatch.SortScore != expectedMatch.SortScore {
				t.Fatalf(
					"case %d: index %d: expected SortScore to be %f, got %f",
					testCaseIdx,
					matchIdx,
					expectedMatch.SortScore,
					actualMatch.SortScore,
				)
			}
		}
	}
}
