package collections

import (
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
