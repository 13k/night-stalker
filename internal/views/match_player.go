package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewMatchPlayer(data *MatchPlayerData) (*nspb.Match_Player, error) {
	if err := data.Validate(); err != nil {
		err = xerrors.Errorf("invalid MatchPlayerData: %w", err)
		return nil, err
	}

	pb := &nspb.Match_Player{
		AccountId: uint32(data.AccountID),
	}

	if data.MatchPlayer != nil {
		if pb.HeroId == 0 {
			pb.HeroId = uint64(data.MatchPlayer.HeroID)
		}

		pb.PlayerSlot = uint32(data.MatchPlayer.PlayerSlot)
		pb.ProName = data.MatchPlayer.ProName
		pb.Kills = data.MatchPlayer.Kills
		pb.Deaths = data.MatchPlayer.Deaths
		pb.Assists = data.MatchPlayer.Assists
		pb.Items = data.MatchPlayer.Items.ToUint64s()
	}

	if data.LiveMatchPlayer != nil {
		if pb.HeroId == 0 {
			pb.HeroId = uint64(data.LiveMatchPlayer.HeroID)
		}
	}

	for _, statsPlayer := range data.LiveMatchStatsPlayers {
		if pb.HeroId == 0 {
			pb.HeroId = uint64(statsPlayer.HeroID)
		}

		if pb.PlayerSlot == 0 {
			pb.PlayerSlot = uint32(statsPlayer.PlayerSlot)
		}
	}

	return pb, nil
}
