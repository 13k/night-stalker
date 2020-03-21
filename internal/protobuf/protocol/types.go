package protocol

import (
	"github.com/faceit/go-steam/steamid"
)

type (
	AccountID    uint32
	MatchID      uint64
	LobbyID      uint64
	LeagueID     uint64
	LeagueNodeID uint32
	SeriesID     uint64
	HeroID       uint64
	TeamID       uint64
	ItemID       uint64
)

type SteamID = steamid.SteamId
