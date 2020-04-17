package dataaccess

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type SearchData struct {
	Heroes                nscol.Heroes
	FollowedPlayers       nscol.FollowedPlayers
	PlayersByAccountID    map[nspb.AccountID]*nsm.Player
	ProPlayersByAccountID map[nspb.AccountID]*nsm.ProPlayer
}
