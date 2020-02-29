package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatches(
	matches []*models.LiveMatch,
	stats map[nspb.MatchID]*models.LiveMatchStats,
	followed map[nspb.AccountID]*models.FollowedPlayer,
	players map[nspb.AccountID]*models.Player,
	proPlayers map[nspb.AccountID]*models.ProPlayer,
) (*nspb.LiveMatches, error) {
	pbMatches := make([]*nspb.LiveMatch, len(matches))

	for i, match := range matches {
		pbMatch, err := NewLiveMatch(
			match,
			stats[match.MatchID],
			followed,
			players,
			proPlayers,
		)

		if err != nil {
			err = xerrors.Errorf("error creating LiveMatch view: %w", err)
			return nil, err
		}

		pbMatches[i] = pbMatch
	}

	view := &nspb.LiveMatches{
		Matches: pbMatches,
	}

	return view, nil
}
