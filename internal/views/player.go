package views

import (
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewPlayer(data *nsdbda.PlayerData) (*nspb.Player, error) {
	if err := data.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid PlayerData: %w", err)
	}

	pb := &nspb.Player{
		AccountId: uint32(data.AccountID),
	}

	if data.FollowedPlayer != nil {
		pb.Name = data.FollowedPlayer.Label
		pb.Slug = data.FollowedPlayer.Slug
	}

	if data.Player != nil {
		if pb.Name == "" {
			pb.Name = data.Player.Name
		}

		pb.PersonaName = data.Player.PersonaName
		pb.AvatarUrl = data.Player.AvatarURL
		pb.AvatarMediumUrl = data.Player.AvatarMediumURL
		pb.AvatarFullUrl = data.Player.AvatarFullURL
	}

	if data.ProPlayer != nil {
		pb.IsPro = true

		if data.ProPlayer.Team != nil {
			pb.Team = NewTeam(data.ProPlayer.Team)
		}
	}

	return pb, nil
}
