package collections

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type FollowedPlayers []*models.FollowedPlayer

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

func (s FollowedPlayers) KeyByAccountID() map[nspb.AccountID]*models.FollowedPlayer {
	if s == nil {
		return nil
	}

	m := make(map[nspb.AccountID]*models.FollowedPlayer)

	for _, p := range s {
		m[p.AccountID] = p
	}

	return m
}
