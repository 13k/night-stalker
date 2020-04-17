package views

import (
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewSearchPlayer(data *nsdbda.SearchPlayerData) *nspb.Search_Player {
	pb := &nspb.Search_Player{
		AccountId: uint32(data.FollowedPlayer.AccountID),
		Name:      data.FollowedPlayer.Label,
		Slug:      data.FollowedPlayer.Slug,
		IsPro:     data.Player.TeamID != 0,
	}

	if player := data.Player; player != nil {
		if pb.AccountId == 0 {
			pb.AccountId = uint32(player.AccountID)
		}

		// pb.Name = player.Name
		pb.PersonaName = player.PersonaName
		pb.AvatarUrl = player.AvatarURL
		pb.AvatarMediumUrl = player.AvatarMediumURL
		pb.AvatarFullUrl = player.AvatarFullURL
	}

	return pb
}
