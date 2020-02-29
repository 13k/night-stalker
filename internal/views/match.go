package views

import (
	"sort"

	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewMatch(data *MatchData) (*nspb.Match, error) {
	if err := data.Validate(); err != nil {
		err = xerrors.Errorf("invalid MatchData: %w", err)
		return nil, err
	}

	var err error

	pb := &nspb.Match{
		MatchId: data.MatchID,
	}

	if data.Match != nil {
		pb.GameMode = data.Match.GameMode
		pb.LeagueId = data.Match.LeagueID
		pb.SeriesType = data.Match.SeriesType
		pb.SeriesGame = data.Match.SeriesGame
		pb.Duration = data.Match.Duration
		pb.Outcome = data.Match.Outcome
		pb.RadiantScore = data.Match.RadiantScore
		pb.DireScore = data.Match.DireScore

		if pb.StartTime, err = models.NullTimestampProto(data.Match.StartTime); err != nil {
			err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
			return nil, err
		}
	}

	if data.LiveMatch != nil {
		if pb.GameMode == 0 {
			pb.GameMode = data.LiveMatch.GameMode
		}

		pb.LobbyId = data.LiveMatch.LobbyID
		pb.LobbyType = data.LiveMatch.LobbyType
		pb.SeriesId = data.LiveMatch.SeriesID
		pb.AverageMmr = data.LiveMatch.AverageMMR
		pb.RadiantTeamId = uint64(data.LiveMatch.RadiantTeamID)
		pb.RadiantTeamName = data.LiveMatch.RadiantTeamName
		pb.RadiantTeamLogo = uint64(data.LiveMatch.RadiantTeamLogo)
		pb.DireTeamId = uint64(data.LiveMatch.DireTeamID)
		pb.DireTeamName = data.LiveMatch.DireTeamName
		pb.DireTeamLogo = uint64(data.LiveMatch.DireTeamLogo)

		if pb.ActivateTime, err = models.NullTimestampProto(data.LiveMatch.ActivateTime); err != nil {
			err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
			return nil, err
		}

		if pb.DeactivateTime, err = models.NullTimestampProto(data.LiveMatch.DeactivateTime); err != nil {
			err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
			return nil, err
		}

		if pb.LastUpdateTime, err = models.NullTimestampProto(data.LiveMatch.LastUpdateTime); err != nil {
			err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
			return nil, err
		}
	}

	for _, stats := range data.LiveMatchStats {
		if pb.LeagueId == 0 {
			pb.LeagueId = stats.LeagueID
		}

		if pb.GameMode == 0 {
			pb.GameMode = stats.GameMode
		}

		var radiantTeam *models.LiveMatchStatsTeam
		var direTeam *models.LiveMatchStatsTeam

		for _, team := range stats.Teams {
			switch team.GameTeam {
			case nspb.GameTeam_GAME_TEAM_GOODGUYS:
				radiantTeam = team
			case nspb.GameTeam_GAME_TEAM_BADGUYS:
				direTeam = team
			}
		}

		if radiantTeam != nil {
			if pb.RadiantTeamId == 0 {
				pb.RadiantTeamId = uint64(radiantTeam.TeamID)
			}

			if pb.RadiantTeamName == "" {
				pb.RadiantTeamName = radiantTeam.Name
			}

			if pb.RadiantTeamTag == "" {
				pb.RadiantTeamTag = radiantTeam.Tag
			}

			if pb.RadiantTeamLogo == 0 {
				pb.RadiantTeamLogo = uint64(radiantTeam.LogoID)
			}

			if pb.RadiantTeamLogoUrl == "" {
				pb.RadiantTeamLogoUrl = radiantTeam.LogoURL
			}
		}

		if direTeam != nil {
			if pb.DireTeamId == 0 {
				pb.DireTeamId = uint64(direTeam.TeamID)
			}

			if pb.DireTeamName == "" {
				pb.DireTeamName = direTeam.Name
			}

			if pb.DireTeamTag == "" {
				pb.DireTeamTag = direTeam.Tag
			}

			if pb.DireTeamLogo == 0 {
				pb.DireTeamLogo = uint64(direTeam.LogoID)
			}

			if pb.DireTeamLogoUrl == "" {
				pb.DireTeamLogoUrl = direTeam.LogoURL
			}
		}
	}

	for _, playerData := range data.PlayersData {
		pbMatchPlayer, err := NewMatchPlayer(playerData)

		if err != nil {
			err = xerrors.Errorf("error creating Match_Player view: %w", err)
			return nil, err
		}

		pb.Players = append(pb.Players, pbMatchPlayer)
	}

	return pb, nil
}

func NewMatches(data MatchesData) ([]*nspb.Match, error) {
	if len(data) == 0 {
		return nil, nil
	}

	matches := make([]*nspb.Match, 0, len(data))

	for _, matchData := range data {
		match, err := NewMatch(matchData)

		if err != nil {
			err = xerrors.Errorf("error creating Match view: %w", err)
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func NewSortedMatches(data MatchesData) ([]*nspb.Match, error) {
	matches, err := NewMatches(data)

	if err != nil {
		err = xerrors.Errorf("error creating Match views: %w", err)
		return nil, err
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].MatchId < matches[j].MatchId
	})

	return matches, nil
}

func NewMatchPlayer(data *MatchPlayerData) (*nspb.Match_Player, error) {
	if err := data.Validate(); err != nil {
		err = xerrors.Errorf("invalid MatchPlayerData: %w", err)
		return nil, err
	}

	pb := &nspb.Match_Player{
		AccountId: data.AccountID,
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
		pb.Items = data.MatchPlayer.Items
	}

	if data.LiveMatchPlayer != nil {
		pb.HeroId = uint64(data.LiveMatchPlayer.HeroID)
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