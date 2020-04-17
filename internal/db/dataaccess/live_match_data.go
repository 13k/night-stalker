package dataaccess

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchData struct {
	LiveMatch       *nsm.LiveMatch
	LiveMatchStats  *nsm.LiveMatchStats
	FollowedPlayers map[nspb.AccountID]*nsm.FollowedPlayer
	Players         map[nspb.AccountID]*nsm.Player
	ProPlayers      map[nspb.AccountID]*nsm.ProPlayer
}
