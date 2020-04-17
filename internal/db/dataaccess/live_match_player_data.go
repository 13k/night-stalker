package dataaccess

import (
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchPlayerData struct {
	FollowedPlayer       *nsm.FollowedPlayer
	Player               *nsm.Player
	LiveMatchPlayer      *nsm.LiveMatchPlayer
	LiveMatchStatsPlayer *nsm.LiveMatchStatsPlayer
}
