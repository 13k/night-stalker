package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchPlayers []*nsm.LiveMatchPlayer

func (s LiveMatchPlayers) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, p := range s {
		matchIDs[i] = nspb.MatchID(p.MatchID)
	}

	return matchIDs
}

func (s LiveMatchPlayers) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, len(s))

	for i, p := range s {
		ids[i] = p.AccountID
	}

	return ids
}

func (s LiveMatchPlayers) GroupByMatchID() map[nspb.MatchID]LiveMatchPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]LiveMatchPlayers)

	for _, p := range s {
		matchID := nspb.MatchID(p.MatchID)
		m[matchID] = append(m[matchID], p)
	}

	return m
}

func (s LiveMatchPlayers) GroupByAccountID() map[nspb.AccountID]LiveMatchPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]LiveMatchPlayers)

	for _, p := range s {
		m[p.AccountID] = append(m[p.AccountID], p)
	}

	return m
}
