package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatchStatsPlayers []*models.LiveMatchStatsPlayer

func (s LiveMatchStatsPlayers) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, p := range s {
		matchIDs[i] = p.MatchID
	}

	return matchIDs
}

func (s LiveMatchStatsPlayers) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, len(s))

	for i, p := range s {
		ids[i] = p.AccountID
	}

	return ids
}

func (s LiveMatchStatsPlayers) GroupByMatchID() map[nspb.MatchID]LiveMatchStatsPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]LiveMatchStatsPlayers)

	for _, p := range s {
		m[p.MatchID] = append(m[p.MatchID], p)
	}

	return m
}

func (s LiveMatchStatsPlayers) GroupByAccountID() map[nspb.AccountID]LiveMatchStatsPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]LiveMatchStatsPlayers)

	for _, p := range s {
		m[p.AccountID] = append(m[p.AccountID], p)
	}

	return m
}
