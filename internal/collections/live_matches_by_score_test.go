package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func TestLiveMatchesByScore_SearchIndex(t *testing.T) {
	testCases := []struct {
		Subject  nscol.LiveMatches
		Match    *models.LiveMatch
		Expected int
	}{
		{
			Subject: nil,
			Match: &models.LiveMatch{
				MatchID:   1,
				SortScore: float64(1.0),
			},
			Expected: 0,
		},
		// greater than all
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   4,
				SortScore: float64(4.0),
			},
			Expected: 0,
		},
		// less than all
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   4,
				SortScore: float64(0.0),
			},
			Expected: 3,
		},
		// equal 2
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   2,
				SortScore: float64(2.0),
			},
			Expected: 1,
		},
		// greater than 2
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   4,
				SortScore: float64(2.1),
			},
			Expected: 1,
		},
		// less than 2
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Match: &models.LiveMatch{
				MatchID:   4,
				SortScore: float64(1.9),
			},
			Expected: 2,
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Subject...)
		actual := subject.SearchIndex(testCase.Match)

		if actual != testCase.Expected {
			t.Fatalf("case %d: expected %d, got %d", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestLiveMatchesByScore_FindIndex(t *testing.T) {
	type findCase struct {
		MatchID  nspb.MatchID
		Expected int
	}

	testCases := []struct {
		Subject nscol.LiveMatches
		Find    []*findCase
	}{
		{
			Subject: nil,
			Find: []*findCase{
				{
					MatchID:  1,
					Expected: -1,
				},
			},
		},
		{
			Subject: nscol.LiveMatches{},
			Find: []*findCase{
				{
					MatchID:  1,
					Expected: -1,
				},
			},
		},
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   5,
					SortScore: float64(5.0),
				},
				{
					MatchID:   4,
					SortScore: float64(1.0),
				},
				{
					MatchID:   3,
					SortScore: float64(1.0),
				},
				{
					MatchID:   2,
					SortScore: float64(1.0),
				},
				{
					MatchID:   1,
					SortScore: float64(0.0),
				},
			},
			Find: []*findCase{
				{
					MatchID:  6,
					Expected: -1,
				},
				{
					MatchID:  0,
					Expected: -1,
				},
				{
					MatchID:  5,
					Expected: 0,
				},
				{
					MatchID:  4,
					Expected: 1,
				},
				{
					MatchID:  3,
					Expected: 2,
				},
				{
					MatchID:  2,
					Expected: 3,
				},
				{
					MatchID:  1,
					Expected: 4,
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Subject...)

		for findCaseIdx, findCase := range testCase.Find {
			actual := subject.FindIndex(findCase.MatchID)

			if actual != findCase.Expected {
				t.Fatalf("case %d: findCase %d: expected %d, got %d", testCaseIdx, findCaseIdx, findCase.Expected, actual)
			}
		}
	}
}

func TestLiveMatchesByScore_Add(t *testing.T) {
	type addCase struct {
		Match    *models.LiveMatch
		Expected int
	}

	testCases := []struct {
		Subject  nscol.LiveMatches
		Add      []*addCase
		Expected nscol.LiveMatches
	}{
		// insert
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   2,
						SortScore: float64(2.0),
					},
					Expected: 0,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
		},
		// noop
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   1,
						SortScore: float64(1.0),
					},
					Expected: -1,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
		},
		// update/reorder
		{
			Subject: nscol.LiveMatches{
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   1,
						SortScore: float64(5.0),
					},
					Expected: 0,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   1,
					SortScore: float64(5.0),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
			},
		},
		// complex
		{
			Subject: nil,
			Add: []*addCase{
				{
					Match: &models.LiveMatch{
						MatchID:   1,
						SortScore: float64(1.0),
					},
					Expected: 0,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   2,
						SortScore: float64(2.0),
					},
					Expected: 0,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   1,
						SortScore: float64(1.0),
					},
					Expected: -1,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   3,
						SortScore: float64(3.0),
					},
					Expected: 0,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   1,
						SortScore: float64(5.0),
					},
					Expected: 0,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   2,
						SortScore: float64(2.0),
					},
					Expected: -1,
				},
				{
					Match: &models.LiveMatch{
						MatchID:   3,
						SortScore: float64(3.1),
					},
					Expected: 1,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   1,
					SortScore: float64(5.0),
				},
				{
					MatchID:   3,
					SortScore: float64(3.1),
				},
				{
					MatchID:   2,
					SortScore: float64(2.0),
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nscol.NewLiveMatchesByScore(testCase.Subject...)

		for addCaseIdx, addCase := range testCase.Add {
			actual := subject.Add(addCase.Match)

			if actual != addCase.Expected {
				t.Fatalf(
					"case %d: addCase %d: expected index %d, got %d",
					testCaseIdx,
					addCaseIdx,
					addCase.Expected,
					actual,
				)
			}

			if addCase.Expected >= 0 {
				expectedMatch := addCase.Match
				actualMatch := subject.At(actual)

				if actualMatch.MatchID != expectedMatch.MatchID {
					t.Fatalf(
						"case %d: addCase %d: expected MatchID to be %d, got %d",
						testCaseIdx,
						addCaseIdx,
						expectedMatch.MatchID,
						actualMatch.MatchID,
					)
				}

				if actualMatch.SortScore != expectedMatch.SortScore {
					t.Fatalf(
						"case %d: addCase %d: expected SortScore to be %f, got %f",
						testCaseIdx,
						addCaseIdx,
						expectedMatch.SortScore,
						actualMatch.SortScore,
					)
				}
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
		MatchID  nspb.MatchID
		Expected nspb.MatchID
	}

	testCases := []struct {
		Matches  nscol.LiveMatches
		Remove   []*removeCase
		Expected nscol.LiveMatches
	}{
		// remove
		{
			Matches: nscol.LiveMatches{
				{
					MatchID:   4,
					SortScore: float64(4.0),
				},
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   2,
					SortScore: float64(3.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Remove: []*removeCase{
				{
					MatchID:  2,
					Expected: 2,
				},
				{
					MatchID:  5,
					Expected: 0,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   4,
					SortScore: float64(4.0),
				},
				{
					MatchID:   3,
					SortScore: float64(3.0),
				},
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
		},
		// noop
		{
			Matches: nscol.LiveMatches{
				{
					MatchID:   1,
					SortScore: float64(1.0),
				},
			},
			Remove: []*removeCase{
				{
					MatchID:  2,
					Expected: 0,
				},
			},
			Expected: nscol.LiveMatches{
				{
					MatchID:   1,
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
