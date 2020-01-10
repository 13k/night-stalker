package livematches

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

func TestCleanResponseGames(t *testing.T) {
	testCases := []struct {
		Subject  []*protocol.CSourceTVGameSmall
		Expected []*protocol.CSourceTVGameSmall
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  []*protocol.CSourceTVGameSmall{},
			Expected: []*protocol.CSourceTVGameSmall{},
		},
		{
			Subject: []*protocol.CSourceTVGameSmall{
				nil,
			},
			Expected: []*protocol.CSourceTVGameSmall{},
		},
		{
			Subject: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{},
			},
			Expected: []*protocol.CSourceTVGameSmall{},
		},
		{
			Subject: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(1),
				},
			},
		},
		{
			Subject: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(2),
				},
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(2),
				},
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(1),
				},
			},
		},
		{
			Subject: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.0),
					SortScore:      proto.Uint32(4),
				},
				&protocol.CSourceTVGameSmall{},
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.0),
					SortScore:      proto.Uint32(3),
				},
				nil,
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.1),
					SortScore:      proto.Uint32(2),
				},
				&protocol.CSourceTVGameSmall{},
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.1),
					SortScore:      proto.Uint32(1),
				},
			},
			Expected: []*protocol.CSourceTVGameSmall{
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(2),
					LastUpdateTime: proto.Float32(2.1),
					SortScore:      proto.Uint32(2),
				},
				&protocol.CSourceTVGameSmall{
					MatchId:        proto.Uint64(1),
					LastUpdateTime: proto.Float32(1.1),
					SortScore:      proto.Uint32(1),
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual := cleanResponseGames(testCase.Subject)

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
