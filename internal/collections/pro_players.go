package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type ProPlayers []*nsm.ProPlayer

func (s ProPlayers) Records() []nsm.Record {
	if s == nil {
		return nil
	}

	records := make([]nsm.Record, len(s))

	for i, m := range s {
		records[i] = m
	}

	return records
}

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

func (s ProPlayers) KeyByAccountID() map[nspb.AccountID]*nsm.ProPlayer {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*nsm.ProPlayer)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
