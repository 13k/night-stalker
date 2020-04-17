package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type FollowedPlayers []*nsm.FollowedPlayer

func (s FollowedPlayers) AccountIDs() AccountIDs {
	if s == nil {
		return nil
	}

	ids := make(AccountIDs, len(s))

	for i, p := range s {
		ids[i] = p.AccountID
	}

	return ids
}

func (s FollowedPlayers) KeyByAccountID() map[nspb.AccountID]*nsm.FollowedPlayer {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*nsm.FollowedPlayer)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
