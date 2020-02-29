package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatchStats []*models.LiveMatchStats

func (s LiveMatchStats) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, stats := range s {
		matchIDs[i] = stats.MatchID
	}

	return matchIDs
}

func (s LiveMatchStats) GroupByMatchID() map[nspb.MatchID]LiveMatchStats {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]LiveMatchStats)

	for _, ss := range s {
		m[ss.MatchID] = append(m[ss.MatchID], ss)
	}

	return m
}
