package collections

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type Players []*models.Player

func (s Players) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, len(s))

	for i, p := range s {
		ids[i] = p.AccountID
	}

	return ids
}

func (s Players) GroupByAccountID() map[nspb.AccountID]Players {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]Players)

	for _, p := range s {
		m[p.AccountID] = append(m[p.AccountID], p)
	}

	return m
}

func (s Players) KeyByAccountID() map[nspb.AccountID]*models.Player {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*models.Player)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
