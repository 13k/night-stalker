package collections

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
)

type AccountIDs []nspb.AccountID

func (s AccountIDs) AddUnique(ids ...nspb.AccountID) AccountIDs {
	if len(ids) == 0 {
		return s
	}

	unique := make(map[nspb.AccountID]bool)

	for _, sid := range s {
		unique[sid] = true
	}

	for _, id := range ids {
		if !unique[id] {
			s = append(s, id)
			unique[id] = true
		}
	}

	return s
}
