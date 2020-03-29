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

	for _, stats := range s {
		m[stats.MatchID] = append(m[stats.MatchID], stats)
	}

	return m
}

func (s LiveMatchStats) KeyByMatchID() map[nspb.MatchID]*models.LiveMatchStats {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]*models.LiveMatchStats)

	for _, stats := range s {
		m[stats.MatchID] = stats
	}

	return m
}
