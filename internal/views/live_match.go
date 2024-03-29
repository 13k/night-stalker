package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatch(
	liveMatch *models.LiveMatch,
	stats *models.LiveMatchStats,
	followed map[nspb.AccountID]*models.FollowedPlayer,
	players map[nspb.AccountID]*models.Player,
	proPlayers map[nspb.AccountID]*models.ProPlayer,
) (*nspb.LiveMatch, error) {
	pb, err := LiveMatchFromModel(liveMatch)

	if err != nil {
		err = xerrors.Errorf("error creating LiveMatch view: %w", err)
		return nil, err
	}

	statsPlayers := make(map[nspb.AccountID]*models.LiveMatchStatsPlayer)

	if stats != nil {
		pb.GameState = stats.GameState
		pb.GameTimestamp = stats.GameTimestamp
		pb.GameTime = stats.GameTime

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
			pb.RadiantScore = radiantTeam.Score
			pb.RadiantNetWorth = radiantTeam.NetWorth

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
			pb.DireScore = direTeam.Score
			pb.DireNetWorth = direTeam.NetWorth

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

		if radiantTeam != nil && direTeam != nil {
			pb.RadiantLead = int32(radiantTeam.NetWorth) - int32(direTeam.NetWorth)
		}

		for _, player := range stats.Players {
			statsPlayers[player.AccountID] = player
		}
	}

	for _, livePlayer := range liveMatch.Players {
		followedPlayer, ok := followed[livePlayer.AccountID]

		if !ok {
			continue
		}

		pb.Players = append(pb.Players, NewLiveMatchPlayer(
			followedPlayer,
			players[livePlayer.AccountID],
			proPlayers[livePlayer.AccountID],
			livePlayer,
			statsPlayers[livePlayer.AccountID],
		))
	}

	return pb, nil
}

func LiveMatchFromModel(m *models.LiveMatch) (*nspb.LiveMatch, error) {
	pb := &nspb.LiveMatch{
		MatchId:                    uint64(m.MatchID),
		ServerId:                   uint64(m.ServerID),
		LobbyId:                    uint64(m.LobbyID),
		LobbyType:                  m.LobbyType,
		LeagueId:                   uint64(m.LeagueID),
		SeriesId:                   uint64(m.SeriesID),
		GameMode:                   m.GameMode,
		GameTime:                   m.GameTime,
		AverageMmr:                 m.AverageMMR,
		Delay:                      m.Delay,
		Spectators:                 m.Spectators,
		SortScore:                  m.SortScore,
		RadiantLead:                m.RadiantLead,
		RadiantTeamId:              uint64(m.RadiantTeamID),
		RadiantTeamName:            m.RadiantTeamName,
		RadiantTeamLogo:            uint64(m.RadiantTeamLogo),
		RadiantScore:               m.RadiantScore,
		DireTeamId:                 uint64(m.DireTeamID),
		DireTeamName:               m.DireTeamName,
		DireTeamLogo:               uint64(m.DireTeamLogo),
		DireScore:                  m.DireScore,
		BuildingState:              uint32(m.BuildingState),
		WeekendTourneyTournamentId: m.WeekendTourneyTournamentID,
		WeekendTourneyDivision:     m.WeekendTourneyDivision,
		WeekendTourneySkillLevel:   m.WeekendTourneySkillLevel,
		WeekendTourneyBracketRound: m.WeekendTourneyBracketRound,
	}

	var err error

	if pb.ActivateTime, err = nssql.NullTimeProto(m.ActivateTime); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.DeactivateTime, err = nssql.NullTimeProto(m.DeactivateTime); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.LastUpdateTime, err = nssql.NullTimeProto(m.LastUpdateTime); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	return pb, nil
}
