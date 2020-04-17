package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type MatchPlayers []*nsm.MatchPlayer

func (s MatchPlayers) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, p := range s {
		matchIDs[i] = nspb.MatchID(p.MatchID)
	}

	return matchIDs
}

func (s MatchPlayers) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, 0, len(s))

	for _, p := range s {
		// MatchPlayer can have zero AccountID (player with private profile)
		if p.AccountID != 0 {
			ids = append(ids, p.AccountID)
		}
	}

	return ids
}

func (s MatchPlayers) GroupByMatchID() map[nspb.MatchID]MatchPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]MatchPlayers)

	for _, p := range s {
		matchID := nspb.MatchID(p.MatchID)
		m[matchID] = append(m[matchID], p)
	}

	return m
}

func (s MatchPlayers) GroupByAccountID() map[nspb.AccountID]MatchPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]MatchPlayers)

	for _, p := range s {
		m[p.AccountID] = append(m[p.AccountID], p)
	}

	return m
}
