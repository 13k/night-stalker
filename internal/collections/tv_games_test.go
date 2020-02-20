package collections_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"

	nscol "github.com/13k/night-stalker/internal/collections"
)

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
