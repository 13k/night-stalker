package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type Matches []*nsm.Match

func (s Matches) Records() []nsm.Record {
	if s == nil {
		return nil
	}

	records := make([]nsm.Record, len(s))

	for i, m := range s {
		records[i] = m
	}

	return records
}

func (s Matches) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, match := range s {
		matchIDs[i] = nspb.MatchID(match.ID)
	}

	return matchIDs
}

func (s Matches) KeyByMatchID() map[nspb.MatchID]*nsm.Match {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]*nsm.Match)

	for _, match := range s {
		m[nspb.MatchID(match.ID)] = match
	}

	return m
}
