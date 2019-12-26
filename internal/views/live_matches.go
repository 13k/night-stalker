package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatches(
	matches []*models.LiveMatch,
	stats map[nspb.MatchID]*models.LiveMatchStats,
	followed map[nspb.AccountID]*models.FollowedPlayer,
	players map[nspb.AccountID]*models.Player,
	proPlayers map[nspb.AccountID]*models.ProPlayer,
) ([]*nspb.LiveMatch, error) {
	view := make([]*nspb.LiveMatch, len(matches))

	for i, match := range matches {
		pbMatch, err := NewLiveMatch(
			match,
			stats[match.MatchID],
			followed,
			players,
			proPlayers,
		)

		if err != nil {
			return nil, err
		}

		view[i] = pbMatch
	}

	return view, nil
}
