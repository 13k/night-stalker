package collections

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type Matches []*models.Match

func (s Matches) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, match := range s {
		matchIDs[i] = match.ID
	}

	return matchIDs
}

func (s Matches) KeyByMatchID() map[nspb.MatchID]*models.Match {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]*models.Match)

	for _, match := range s {
		m[match.ID] = match
	}

	return m
}
