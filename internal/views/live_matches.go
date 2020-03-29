package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatches(
	liveMatches []*models.LiveMatch,
	stats map[nspb.MatchID]*models.LiveMatchStats,
	followed map[nspb.AccountID]*models.FollowedPlayer,
	players map[nspb.AccountID]*models.Player,
	proPlayers map[nspb.AccountID]*models.ProPlayer,
) (*nspb.LiveMatches, error) {
	pbLiveMatches := make([]*nspb.LiveMatch, len(liveMatches))

	for i, liveMatch := range liveMatches {
		pbLiveMatch, err := NewLiveMatch(
			liveMatch,
			stats[liveMatch.MatchID],
			followed,
			players,
			proPlayers,
		)

		if err != nil {
			err = xerrors.Errorf("error creating LiveMatch view: %w", err)
			return nil, err
		}

		pbLiveMatches[i] = pbLiveMatch
	}

	view := &nspb.LiveMatches{
		Matches: pbLiveMatches,
	}

	return view, nil
}
