package protocol

import (
	"github.com/faceit/go-steam/steamid"
)

type (
	AbilityID    uint32
	AccountID    uint32
	HeroID       uint64
	ItemID       uint64
	LeagueID     uint64
	LeagueNodeID uint32
	LobbyID      uint64
	MatchID      uint64
	SeriesID     uint64
	TeamID       uint64

	SteamID = steamid.SteamId
)
