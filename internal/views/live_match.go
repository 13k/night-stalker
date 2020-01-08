package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLiveMatch(
	match *models.LiveMatch,
	stats *models.LiveMatchStats,
	followed map[nspb.AccountID]*models.FollowedPlayer,
	players map[nspb.AccountID]*models.Player,
	proPlayers map[nspb.AccountID]*models.ProPlayer,
) (*nspb.LiveMatch, error) {
	pb, err := LiveMatchProto(match)

	if err != nil {
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

	for _, livePlayer := range match.Players {
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

func LiveMatchProto(m *models.LiveMatch) (*nspb.LiveMatch, error) {
	pb := &nspb.LiveMatch{
		MatchId:                    m.MatchID,
		ServerSteamId:              uint64(m.ServerSteamID),
		LobbyId:                    m.LobbyID,
		LobbyType:                  m.LobbyType,
		LeagueId:                   m.LeagueID,
		SeriesId:                   m.SeriesID,
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

	if pb.ActivateTime, err = models.NullTimestampProto(m.ActivateTime); err != nil {
		return nil, err
	}

	if pb.DeactivateTime, err = models.NullTimestampProto(m.DeactivateTime); err != nil {
		return nil, err
	}

	if pb.LastUpdateTime, err = models.NullTimestampProto(m.LastUpdateTime); err != nil {
		return nil, err
	}

	return pb, nil
}

func NewLiveMatchPlayer(
	followed *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
	livePlayer *models.LiveMatchPlayer,
	statsPlayer *models.LiveMatchStatsPlayer,
) *nspb.LiveMatch_Player {
	pb := &nspb.LiveMatch_Player{
		AccountId: followed.AccountID,
		Name:      followed.Label,
		Label:     followed.Label,
		HeroId:    uint64(livePlayer.HeroID),
	}

	if player != nil {
		if pb.AccountId == 0 {
			pb.AccountId = player.AccountID
		}

		pb.Name = player.Name
		pb.PersonaName = player.PersonaName
		pb.AvatarUrl = player.AvatarURL
		pb.AvatarMediumUrl = player.AvatarMediumURL
		pb.AvatarFullUrl = player.AvatarFullURL
	}

	if statsPlayer != nil {
		if pb.AccountId == 0 {
			pb.AccountId = statsPlayer.AccountID
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
