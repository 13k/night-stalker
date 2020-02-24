package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewSearch(
	heroes []*models.Hero,
	followed []*models.FollowedPlayer,
	players []*models.Player,
	proPlayers []*models.ProPlayer,
) (*nspb.Search, error) {
	pb := &nspb.Search{
		HeroIds: make([]uint64, len(heroes)),
		Players: make([]*nspb.Search_Player, len(followed)),
	}

	for i, hero := range heroes {
		pb.HeroIds[i] = uint64(hero.ID)
	}

	playersByAccountID := make(map[nspb.AccountID]*models.Player)

	for _, player := range players {
		playersByAccountID[player.AccountID] = player
	}

	proPlayersByAccountID := make(map[nspb.AccountID]*models.ProPlayer)

	for _, proPlayer := range proPlayers {
		proPlayersByAccountID[proPlayer.AccountID] = proPlayer
	}

	for i, fp := range followed {
		pb.Players[i] = NewSearchPlayer(
			fp,
			playersByAccountID[fp.AccountID],
			proPlayersByAccountID[fp.AccountID],
		)
	}

	return pb, nil
}

func NewSearchPlayer(
	followed *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
) *nspb.Search_Player {
	pb := &nspb.Search_Player{
		AccountId: followed.AccountID,
		Name:      followed.Label,
		IsPro:     proPlayer != nil,
	}

	if player != nil {
		if pb.AccountId == 0 {
			pb.AccountId = player.AccountID
		}

		// pb.Name = player.Name
		pb.PersonaName = player.PersonaName
		pb.AvatarUrl = player.AvatarURL
		pb.AvatarMediumUrl = player.AvatarMediumURL
		pb.AvatarFullUrl = player.AvatarFullURL
	}

	return pb
}
