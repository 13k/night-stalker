package collections_test

import (
	"testing"

	"github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func TestTVGames_MatchIDs(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.TVGames{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
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

func TestTVGames_FindIndexByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		MatchID  nspb.MatchID
		Expected int
	}{
		{
			Subject:  nil,
			MatchID:  0,
			Expected: -1,
		},
		{
			Subject:  nscol.TVGames{},
			MatchID:  0,
			Expected: -1,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(4)},
			},
			MatchID:  0,
			Expected: -1,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(4)},
			},
			MatchID:  1,
			Expected: 0,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(4)},
			},
			MatchID:  2,
			Expected: 2,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(4)},
			},
			MatchID:  3,
			Expected: 3,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3)},
				{MatchId: proto.Uint64(4)},
			},
			MatchID:  4,
			Expected: 5,
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.FindIndexByMatchID(testCase.MatchID)

		if actual != testCase.Expected {
			t.Fatalf("case %d: expected %d, got %d", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestTVGames_GroupByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		Expected map[nspb.MatchID]nscol.TVGames
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.TVGames{},
			Expected: map[nspb.MatchID]nscol.TVGames{},
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1), ServerSteamId: proto.Uint64(1)},
				{MatchId: proto.Uint64(1), ServerSteamId: proto.Uint64(2)},
				{MatchId: proto.Uint64(2), ServerSteamId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3), ServerSteamId: proto.Uint64(3)},
				{MatchId: proto.Uint64(3), ServerSteamId: proto.Uint64(4)},
			},
			Expected: map[nspb.MatchID]nscol.TVGames{
				1: {
					{MatchId: proto.Uint64(1), ServerSteamId: proto.Uint64(1)},
					{MatchId: proto.Uint64(1), ServerSteamId: proto.Uint64(2)},
				},
				2: {
					{MatchId: proto.Uint64(2), ServerSteamId: proto.Uint64(3)},
				},
				3: {
					{MatchId: proto.Uint64(3), ServerSteamId: proto.Uint64(3)},
					{MatchId: proto.Uint64(3), ServerSteamId: proto.Uint64(4)},
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

				if actualElem.GetMatchId() != expectedElem.GetMatchId() {
					t.Fatalf(
						"case %d: matchID %d: group %d: expected MatchId %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.GetMatchId(),
						actualElem.GetMatchId(),
					)
				}

				if actualElem.GetServerSteamId() != expectedElem.GetServerSteamId() {
					t.Fatalf(
						"case %d: matchID %d: index %d: expected ServerSteamId %d, got %d",
						testCaseIdx,
						matchID,
						groupIdx,
						expectedElem.GetServerSteamId(),
						actualElem.GetServerSteamId(),
					)
				}
			}
		}
	}
}

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

func TestTVGames_RemoveByMatchID(t *testing.T) {
	testCases := []struct {
		Subject  nscol.TVGames
		MatchID  nspb.MatchID
		Expected nscol.TVGames
		Removed  *protocol.CSourceTVGameSmall
	}{
		// nil
		{
			Subject:  nil,
			MatchID:  0,
			Expected: nil,
			Removed:  nil,
		},
		{
			Subject:  nil,
			MatchID:  1,
			Expected: nil,
			Removed:  nil,
		},
		// empty
		{
			Subject:  nscol.TVGames{},
			MatchID:  0,
			Expected: nscol.TVGames{},
			Removed:  nil,
		},
		{
			Subject:  nscol.TVGames{},
			MatchID:  1,
			Expected: nscol.TVGames{},
			Removed:  nil,
		},
		// not found
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			MatchID: 0,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			Removed: nil,
		},
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
			},
			MatchID: 2,
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
			MatchID:  1,
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
			MatchID: 1,
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
			MatchID: 2,
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
			MatchID: 2,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(3)},
			},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(2),
			},
		},
		// with duplicates
		{
			Subject: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
			},
			MatchID: 2,
			Expected: nscol.TVGames{
				{MatchId: proto.Uint64(1)},
				{MatchId: proto.Uint64(2)},
				{MatchId: proto.Uint64(3)},
			},
			Removed: &protocol.CSourceTVGameSmall{
				MatchId: proto.Uint64(2),
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subjectLenBefore := len(testCase.Subject)
		actualSlice, actualRemoved := testCase.Subject.RemoveByMatchID(testCase.MatchID)
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
