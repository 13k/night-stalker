package views

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewSearchPlayer(
	followed *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
) *nspb.Search_Player {
	pb := &nspb.Search_Player{
		AccountId: uint32(followed.AccountID),
		Name:      followed.Label,
		Slug:      followed.Slug,
		IsPro:     proPlayer != nil,
	}

	if player != nil {
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
