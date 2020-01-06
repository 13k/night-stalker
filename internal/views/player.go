package views

import (
	"sort"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewPlayer(
	followed *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
	livePlayers []*models.LiveMatchPlayer,
	statsPlayersByMatchID map[nspb.MatchID][]*models.LiveMatchStatsPlayer,
) (*nspb.Player, error) {
	pb := &nspb.Player{}

	pb.AccountId = followed.AccountID
	pb.Name = followed.Label

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

	if pb.Name == "" {
	statsPlayersLoop:
		for _, matchStatsPlayers := range statsPlayersByMatchID {
			for _, statsPlayer := range matchStatsPlayers {
				if statsPlayer.Name != "" {
					pb.Name = statsPlayer.Name
					break statsPlayersLoop
				}
			}
		}
	}

	if proPlayer != nil {
		pb.IsPro = true
		pb.Team = NewPlayerTeam(proPlayer.Team)
	}

	for _, livePlayer := range livePlayers {
		if livePlayer.LiveMatch == nil {
			continue
		}

		pbMatch, err := NewPlayerMatch(
			livePlayer,
			statsPlayersByMatchID[livePlayer.LiveMatch.MatchID],
		)

		if err != nil {
			return nil, err
		}

		pb.Matches = append(pb.Matches, pbMatch)
	}

	sort.Slice(pb.Matches, func(i, j int) bool {
		return pb.Matches[i].MatchId < pb.Matches[j].MatchId
	})

	return pb, nil
}

func NewPlayerTeam(team *models.Team) *nspb.Player_Team {
	if team == nil {
		return nil
	}

	return &nspb.Player_Team{
		Id:      uint64(team.ID),
		Name:    team.Name,
		Tag:     team.Tag,
		LogoUrl: team.LogoURL,
	}
}

func NewPlayerMatch(
	livePlayer *models.LiveMatchPlayer,
	statsPlayers []*models.LiveMatchStatsPlayer,
) (*nspb.Player_Match, error) {
	liveMatch := livePlayer.LiveMatch

	pb := &nspb.Player_Match{
		MatchId:         liveMatch.MatchID,
		HeroId:          uint64(livePlayer.HeroID),
		LobbyId:         liveMatch.LobbyID,
		LobbyType:       liveMatch.LobbyType,
		LeagueId:        liveMatch.LeagueID,
		SeriesId:        liveMatch.SeriesID,
		GameMode:        liveMatch.GameMode,
		AverageMmr:      liveMatch.AverageMMR,
		RadiantTeamId:   uint64(liveMatch.RadiantTeamID),
		RadiantTeamName: liveMatch.RadiantTeamName,
		RadiantTeamLogo: uint64(liveMatch.RadiantTeamLogo),
		DireTeamId:      uint64(liveMatch.DireTeamID),
		DireTeamName:    liveMatch.DireTeamName,
		DireTeamLogo:    uint64(liveMatch.DireTeamLogo),
	}

	var err error

	if pb.ActivateTime, err = models.NullTimestampProto(liveMatch.ActivateTime); err != nil {
		return nil, err
	}

	if pb.DeactivateTime, err = models.NullTimestampProto(liveMatch.DeactivateTime); err != nil {
		return nil, err
	}

	if pb.LastUpdateTime, err = models.NullTimestampProto(liveMatch.LastUpdateTime); err != nil {
		return nil, err
	}

	for _, statsPlayer := range statsPlayers {
		if pb.HeroId == 0 {
			if statsPlayer.HeroID != 0 {
				pb.HeroId = uint64(statsPlayer.HeroID)
			}
		}

		if stats := statsPlayer.LiveMatchStats; stats != nil {
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
	}

	return pb, nil
}
