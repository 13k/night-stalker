package views

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatchPlayer(
	followed *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
	livePlayer *models.LiveMatchPlayer,
	statsPlayer *models.LiveMatchStatsPlayer,
) *nspb.LiveMatch_Player {
	pb := &nspb.LiveMatch_Player{
		AccountId: uint32(followed.AccountID),
		Name:      followed.Label,
		Label:     followed.Label,
		Slug:      followed.Slug,
		HeroId:    uint64(livePlayer.HeroID),
	}

	if player != nil {
		if pb.AccountId == 0 {
			pb.AccountId = uint32(player.AccountID)
		}

		pb.Name = player.Name
		pb.PersonaName = player.PersonaName
		pb.AvatarUrl = player.AvatarURL
		pb.AvatarMediumUrl = player.AvatarMediumURL
		pb.AvatarFullUrl = player.AvatarFullURL
	}

	if statsPlayer != nil {
		if pb.AccountId == 0 {
			pb.AccountId = uint32(statsPlayer.AccountID)
		}

		if pb.HeroId == 0 {
			pb.HeroId = uint64(statsPlayer.HeroID)
		}

		if pb.Name == "" {
			pb.Name = statsPlayer.Name
		}

		pb.PlayerSlot = uint32(statsPlayer.PlayerSlot)
		pb.Team = statsPlayer.GameTeam
		pb.Level = statsPlayer.Level
		pb.Kills = statsPlayer.Kills
		pb.Deaths = statsPlayer.Deaths
		pb.Assists = statsPlayer.Assists
		pb.Denies = statsPlayer.Denies
		pb.LastHits = statsPlayer.LastHits
		pb.Gold = statsPlayer.Gold
		pb.NetWorth = statsPlayer.NetWorth
	}

	pb.IsPro = proPlayer != nil

	return pb
}
