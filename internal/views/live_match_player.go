package views

import (
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewLiveMatchPlayer(data *nsdbda.LiveMatchPlayerData) *nspb.LiveMatch_Player {
	pb := &nspb.LiveMatch_Player{
		AccountId: uint32(data.FollowedPlayer.AccountID),
		Name:      data.FollowedPlayer.Label,
		Label:     data.FollowedPlayer.Label,
		Slug:      data.FollowedPlayer.Slug,
		HeroId:    uint64(data.LiveMatchPlayer.HeroID),
	}

	if player := data.Player; player != nil {
		if pb.AccountId == 0 {
			pb.AccountId = uint32(player.AccountID)
		}

		pb.Name = player.Name
		pb.PersonaName = player.PersonaName
		pb.AvatarUrl = player.AvatarURL
		pb.AvatarMediumUrl = player.AvatarMediumURL
		pb.AvatarFullUrl = player.AvatarFullURL
	}

	if statsPlayer := data.LiveMatchStatsPlayer; statsPlayer != nil {
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

	pb.IsPro = data.ProPlayer != nil

	return pb
}
