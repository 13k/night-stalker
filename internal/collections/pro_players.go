package collections

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type ProPlayers []*models.ProPlayer

func (s ProPlayers) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, len(s))

	for i, p := range s {
		ids[i] = p.AccountID
	}

	return ids
}

func (s ProPlayers) GroupByAccountID() map[nspb.AccountID]ProPlayers {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]ProPlayers)

	for _, p := range s {
		m[p.AccountID] = append(m[p.AccountID], p)
	}

	return m
}

func (s ProPlayers) KeyByAccountID() map[nspb.AccountID]*models.ProPlayer {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*models.ProPlayer)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
