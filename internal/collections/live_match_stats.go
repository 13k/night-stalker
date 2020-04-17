package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchStats []*nsm.LiveMatchStats

func (s LiveMatchStats) Records() []nsm.Record {
	if s == nil {
		return nil
	}

	records := make([]nsm.Record, len(s))

	for i, m := range s {
		records[i] = m
	}

	return records
}

func (s LiveMatchStats) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, st := range s {
		matchIDs[i] = nspb.MatchID(st.MatchID)
	}

	return matchIDs
}

func (s LiveMatchStats) GroupByMatchID() map[nspb.MatchID]LiveMatchStats {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]LiveMatchStats)

	for _, st := range s {
		matchID := nspb.MatchID(st.MatchID)
		m[matchID] = append(m[matchID], st)
	}

	return m
}

func (s LiveMatchStats) KeyByMatchID() map[nspb.MatchID]*nsm.LiveMatchStats {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]*nsm.LiveMatchStats)

	for _, st := range s {
		m[nspb.MatchID(st.MatchID)] = st
	}

	return m
}
