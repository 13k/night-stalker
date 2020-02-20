package collections

import (
	"github.com/13k/night-stalker/models"
)

type LiveMatchStatsSlice []*models.LiveMatchStats

func (s LiveMatchStatsSlice) MatchIDs() MatchIDs {
	matchIDs := make(MatchIDs, len(s))

	for i, stats := range s {
		matchIDs[i] = stats.MatchID
	}

	return matchIDs
}
