package collections_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"

	nscol "github.com/13k/night-stalker/internal/collections"
)

func TestTVGames_Remove(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		Index    int
		Expected nscol.TVGames
		Removed  *protocol.CSourceTVGameSmall
	}{
		// nil
		{
			Subject:  nil,
			Index:    -1,
			Expected: nil,
			Removed:  nil,
		},
		{
			Subject:  nil,
			Index:    0,
			Expected: nil,
			Removed:  nil,
		},
		{
			Subject:  nil,
			Index:    1,
			Expected: nil,
			Removed:  nil,
		},
		// empty
		{
			Subject:  nscol.TVGames{},
			Index:    -1,
			Expected: nscol.TVGames{},
			Removed:  nil,
		},
		{
			Subject:  nscol.TVGames{},
			Index:    0,
			Expected: nscol.TVGames{},
			Removed:  nil,
		},
		{
			Subject:  nscol.TVGames{},
			Index:    1,
			Expected: nscol.TVGames{},
			Removed:  nil,
		},
		// out-of-bounds
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Index: -1,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Removed: nil,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Index: 1,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Removed: nil,
		},
		// single element
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Index:    0,
			Expected: nscol.TVGames{},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(1),
			},
		},
		// start
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
			},
			Index: 0,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(2)},
			},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(1),
			},
		},
		// end
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
			},
			Index: 1,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(2),
			},
		},
		// middle
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
			},
			Index: 1,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(3)},
			},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(2),
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subjectLenBefore := len(testCase.Subject)
		actualSlice, actualRemoved := testCase.Subject.Remove(testCase.Index)
		subjectLenAfter := len(testCase.Subject)

		if subjectLenBefore != subjectLenAfter {
			t.Fatalf(
				"case %d: expected subject length to not change, but changed from: %d, to: %d",
				testCaseIdx,
				subjectLenBefore,
				subjectLenAfter,
			)
		}

		if testCase.Removed == nil {
			if actualRemoved != nil {
				t.Fatalf("case %d: expected nil entry, got %#v", testCaseIdx, actualRemoved)
			}
		} else {
			if actualRemoved == nil {
				t.Fatalf("case %d: expected non-nil entry", testCaseIdx)
			}

			if actualRemoved.GetMatchId() != testCase.Removed.GetMatchId() {
				t.Fatalf("case %d: expected entry %#v, got %#v", testCaseIdx, testCase.Removed, actualRemoved)
			}
		}

		if testCase.Expected == nil {
			if actualSlice != nil {
				t.Fatalf("case %d: expected nil slice, got %#v", testCaseIdx, actualSlice)
			}
		} else {
			if actualSlice == nil {
				t.Fatalf("case %d: expected non-nil slice", testCaseIdx)
			}
		}

		if actualLen := len(actualSlice); actualLen != len(testCase.Expected) {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, len(testCase.Expected), actualLen)
		}

		for gameIdx, expectedGame := range testCase.Expected {
			actualGame := actualSlice[gameIdx]

			if actualGame == nil {
				t.Fatalf("case %d: index %d: expected non-nil", testCaseIdx, gameIdx)
			}

			if actualGame.GetMatchId() != expectedGame.GetMatchId() {
				t.Fatalf(
					"case %d: index %d: expected MatchId to be %d, got %d",
					testCaseIdx,
					gameIdx,
					expectedGame.GetMatchId(),
					actualGame.GetMatchId(),
				)
			}
		}
	}
}

func TestTVGames_Clean(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		Expected nscol.TVGames
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.TVGames{},
			Expected: nscol.TVGames{},
		},
		{
			Subject: nscol.TVGames{
				nil,
			},
			Expected: nscol.TVGames{},
		},
		{
			Subject: nscol.TVGames{
				&protocol.CSourceTVGameSmall{},
			},
			Expected: nscol.TVGames{},
		},
		{
			Subject: nscol.TVGames{
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: nscol.TVGames{
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(1),
				},
			},
		},
		{
			Subject: nscol.TVGames{
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(2),
				},
				{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: nscol.TVGames{
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(2),
				},
				{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(1),
				},
			},
		},
		{
			Subject: nscol.TVGames{
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(4),
				},
				{},
				{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(3),
				},
				nil,
				{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.1),
					SortScore:      proto.Uint32(2),
				},
				{},
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.1),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: nscol.TVGames{
				{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.1),
					SortScore:      proto.Uint32(2),
				},
				{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.1),
					SortScore:      proto.Uint32(1),
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual := testCase.Subject.Clean()

		if testCase.Expected == nil && actual != nil {
			t.Fatalf("case %d: expected nil, got %T", testCaseIdx, actual)
		}

		expectedLen := len(testCase.Expected)

		if actualLen := len(actual); actualLen != expectedLen {
			t.Fatalf("case %d: expected length to be %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for gameIdx, expectedGame := range testCase.Expected {
			actualGame := actual[gameIdx]

			expectedMatchID := expectedGame.GetMatchId()

			if actualMatchID := actualGame.GetMatchId(); actualMatchID != expectedMatchID {
				t.Fatalf(
					"case %d: game %d: expected MatchId to be %d, got %d",
					testCaseIdx,
					gameIdx,
					expectedMatchID,
					actualMatchID,
				)
			}

			expectedUpdateTime := expectedGame.GetLastUpdateTime()

			if actualUpdateTime := actualGame.GetLastUpdateTime(); actualUpdateTime != expectedUpdateTime {
				t.Fatalf(
					"case %d: game %d: expected LastUpdateTime to be %f, got %f",
					testCaseIdx,
					gameIdx,
					expectedUpdateTime,
					actualUpdateTime,
				)
			}

			expectedSortScore := expectedGame.GetSortScore()

			if actualSortScore := actualGame.GetSortScore(); actualSortScore != expectedSortScore {
				t.Fatalf(
					"case %d: game %d: expected SortScore to be %d, got %d",
					testCaseIdx,
					gameIdx,
					expectedSortScore,
					actualSortScore,
				)
			}
		}
	}
}
