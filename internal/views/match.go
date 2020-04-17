package views

import (
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsm "github.com/13k/night-stalker/models"
)

func NewMatch(data *nsdbda.MatchData) (*nspb.Match, error) {
	if err := data.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid MatchData: %w", err)
	}

	var err error

	pb := &nspb.Match{
		MatchId: uint64(data.MatchID),
	}

	if data.Match != nil {
		pb.GameMode = data.Match.GameMode
		pb.LeagueId = uint64(data.Match.LeagueID)
		pb.SeriesType = uint32(data.Match.SeriesType)
		pb.SeriesGame = data.Match.SeriesGame
		pb.Duration = data.Match.Duration
		pb.Outcome = data.Match.Outcome
		pb.RadiantScore = data.Match.RadiantScore
		pb.DireScore = data.Match.DireScore

		if pb.StartTime, err = nssql.NullTimeProto(data.Match.StartTime); err != nil {
			return nil, xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		}
	}

	if data.LiveMatch != nil {
		if pb.GameMode == 0 {
			pb.GameMode = data.LiveMatch.GameMode
		}

		pb.LobbyId = uint64(data.LiveMatch.LobbyID)
		pb.LobbyType = data.LiveMatch.LobbyType
		pb.SeriesId = uint64(data.LiveMatch.SeriesID)
		pb.AverageMmr = data.LiveMatch.AverageMMR
		pb.RadiantTeamId = uint64(data.LiveMatch.RadiantTeamID)
		pb.RadiantTeamName = data.LiveMatch.RadiantTeamName
		pb.RadiantTeamLogo = uint64(data.LiveMatch.RadiantTeamLogo)
		pb.DireTeamId = uint64(data.LiveMatch.DireTeamID)
		pb.DireTeamName = data.LiveMatch.DireTeamName
		pb.DireTeamLogo = uint64(data.LiveMatch.DireTeamLogo)

		if pb.ActivateTime, err = nssql.NullTimeProto(data.LiveMatch.ActivateTime); err != nil {
			return nil, xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		}

		if pb.DeactivateTime, err = nssql.NullTimeProto(data.LiveMatch.DeactivateTime); err != nil {
			return nil, xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		}

		if pb.LastUpdateTime, err = nssql.NullTimeProto(data.LiveMatch.LastUpdateTime); err != nil {
			return nil, xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		}
	}

	for _, stats := range data.LiveMatchStats {
		if pb.LeagueId == 0 {
			pb.LeagueId = uint64(stats.LeagueID)
		}

		if pb.GameMode == 0 {
			pb.GameMode = stats.GameMode
		}

		var radiantTeam *nsm.LiveMatchStatsTeam
		var direTeam *nsm.LiveMatchStatsTeam

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
			return nil, xerrors.Errorf("error creating Match_Player view: %w", err)
		}

		pb.Players = append(pb.Players, pbMatchPlayer)
	}

	return pb, nil
}
