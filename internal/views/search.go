package views

import (
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewSearch(data *nsdbda.SearchData) *nspb.Search {
	pb := &nspb.Search{
		HeroIds: make([]uint64, len(data.Heroes)),
		Players: make([]*nspb.Search_Player, len(data.FollowedPlayers)),
	}

	for i, hero := range data.Heroes {
		pb.HeroIds[i] = uint64(hero.ID)
	}

	for i, followed := range data.FollowedPlayers {
		pb.Players[i] = NewSearchPlayer(&nsdbda.SearchPlayerData{
			FollowedPlayer: followed,
			Player:         data.PlayersByAccountID[followed.AccountID],
			ProPlayer:      data.ProPlayersByAccountID[followed.AccountID],
		})
	}

	return pb
}
