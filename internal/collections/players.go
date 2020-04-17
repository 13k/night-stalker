package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type Players []*nsm.Player

func (s Players) Records() []nsm.Record {
	if s == nil {
		return nil
	}

	records := make([]nsm.Record, len(s))

	for i, m := range s {
		records[i] = m
	}

	return records
}

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

func (s Players) KeyByAccountID() map[nspb.AccountID]*nsm.Player {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*nsm.Player)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
