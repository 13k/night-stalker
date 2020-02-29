package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatchPlayers []*models.LiveMatchPlayer

func (s LiveMatchPlayers) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, p := range s {
		matchIDs[i] = p.MatchID
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
		m[p.MatchID] = append(m[p.MatchID], p)
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
