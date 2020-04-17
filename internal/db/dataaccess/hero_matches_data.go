package dataaccess

import (
	nsm "github.com/13k/night-stalker/models"
)

type HeroMatchesData struct {
	Hero         *nsm.Hero
	KnownPlayers PlayersData
	MatchesData  MatchesData
}
