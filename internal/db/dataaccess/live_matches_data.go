package dataaccess

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchesData struct {
	LiveMatches                nscol.LiveMatches
	LiveMatchStatsByMatchID    map[nspb.MatchID]*nsm.LiveMatchStats
	FollowedPlayersByAccountID map[nspb.AccountID]*nsm.FollowedPlayer
	PlayersByAccountID         map[nspb.AccountID]*nsm.Player
	ProPlayersByAccountID      map[nspb.AccountID]*nsm.ProPlayer
}
