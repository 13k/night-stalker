package dataaccess

import (
	nsm "github.com/13k/night-stalker/models"
)

type SearchPlayerData struct {
	FollowedPlayer *nsm.FollowedPlayer
	Player         *nsm.Player
	ProPlayer      *nsm.ProPlayer
}
