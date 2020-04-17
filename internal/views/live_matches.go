package views

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewLiveMatches(data *nsdbda.LiveMatchesData) (*nspb.LiveMatches, error) {
	if data == nil {
		return nil, nil
	}

	if len(data.LiveMatches) == 0 {
		return nil, nil
	}

	pbLiveMatches := make([]*nspb.LiveMatch, len(data.LiveMatches))

	for i, liveMatch := range data.LiveMatches {
		pbLiveMatch, err := NewLiveMatch(&nsdbda.LiveMatchData{
			LiveMatch:       liveMatch,
			LiveMatchStats:  data.LiveMatchStatsByMatchID[nspb.MatchID(liveMatch.MatchID)],
			FollowedPlayers: data.FollowedPlayersByAccountID,
			Players:         data.PlayersByAccountID,
		})

		if err != nil {
			return nil, xerrors.Errorf("error creating LiveMatch view: %w", err)
		}

		pbLiveMatches[i] = pbLiveMatch
	}

	view := &nspb.LiveMatches{
		Matches: pbLiveMatches,
	}

	return view, nil
}

func NewShallowLiveMatches(matchIDs nscol.MatchIDs) *nspb.LiveMatches {
	if len(matchIDs) == 0 {
		return nil
	}

	pbLiveMatches := make([]*nspb.LiveMatch, len(matchIDs))

	for i, matchID := range matchIDs {
		pbLiveMatches[i] = &nspb.LiveMatch{
			MatchId: uint64(matchID),
		}
	}

	return &nspb.LiveMatches{
		Matches: pbLiveMatches,
	}
}
