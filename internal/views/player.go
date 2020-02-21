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
	matchPlayers []*models.MatchPlayer,
	statsPlayers []*models.LiveMatchStatsPlayer,
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
		for _, statsPlayer := range statsPlayers {
			if statsPlayer.Name != "" {
				pb.Name = statsPlayer.Name
				break
			}
		}
	}

	if proPlayer != nil {
		pb.IsPro = true
		pb.Team = NewPlayerTeam(proPlayer.Team)
	}

	matchPlayersByMatchID := make(map[nspb.MatchID]*models.MatchPlayer)

	for _, matchPlayer := range matchPlayers {
		matchPlayersByMatchID[matchPlayer.Match.ID] = matchPlayer
	}

	statsPlayersByMatchID := make(map[nspb.MatchID][]*models.LiveMatchStatsPlayer)

	for _, statsPlayer := range statsPlayers {
		matchID := statsPlayer.LiveMatchStats.MatchID
		statsPlayersByMatchID[matchID] = append(statsPlayersByMatchID[matchID], statsPlayer)
	}

	for _, livePlayer := range livePlayers {
		matchID := livePlayer.LiveMatch.MatchID

		pbMatch, err := NewPlayerMatch(
			livePlayer,
			matchPlayersByMatchID[matchID],
			statsPlayersByMatchID[matchID],
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
	matchPlayer *models.MatchPlayer,
	statsPlayers []*models.LiveMatchStatsPlayer,
) (*nspb.Player_Match, error) {
	pb := &nspb.Player_Match{
		PlayerDetails: NewPlayerMatchPlayerDetails(livePlayer, matchPlayer, statsPlayers),
	}

	var err error

	if livePlayer != nil {
		liveMatch := livePlayer.LiveMatch

		if liveMatch != nil {
			pb.MatchId = liveMatch.MatchID
			pb.LobbyId = liveMatch.LobbyID
			pb.LobbyType = liveMatch.LobbyType
			pb.SeriesId = liveMatch.SeriesID
			pb.GameMode = liveMatch.GameMode
			pb.AverageMmr = liveMatch.AverageMMR
			pb.RadiantTeamId = uint64(liveMatch.RadiantTeamID)
			pb.RadiantTeamName = liveMatch.RadiantTeamName
			pb.RadiantTeamLogo = uint64(liveMatch.RadiantTeamLogo)
			pb.DireTeamId = uint64(liveMatch.DireTeamID)
			pb.DireTeamName = liveMatch.DireTeamName
			pb.DireTeamLogo = uint64(liveMatch.DireTeamLogo)

			if pb.ActivateTime, err = models.NullTimestampProto(liveMatch.ActivateTime); err != nil {
				return nil, err
			}

			if pb.DeactivateTime, err = models.NullTimestampProto(liveMatch.DeactivateTime); err != nil {
				return nil, err
			}

			if pb.LastUpdateTime, err = models.NullTimestampProto(liveMatch.LastUpdateTime); err != nil {
				return nil, err
			}

			match := liveMatch.Match

			if match != nil {
				if pb.MatchId == 0 {
					pb.MatchId = match.ID
				}

				if pb.GameMode == 0 {
					pb.GameMode = match.GameMode
				}

				pb.LeagueId = match.LeagueID
				pb.SeriesType = match.SeriesType
				pb.SeriesGame = match.SeriesGame
				pb.Duration = match.Duration
				pb.Outcome = match.Outcome
				pb.RadiantScore = match.RadiantScore
				pb.DireScore = match.DireScore

				if pb.StartTime, err = models.NullTimestampProto(match.StartTime); err != nil {
					return nil, err
				}
			}
		}
	}

	for _, statsPlayer := range statsPlayers {
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

func NewPlayerMatchPlayerDetails(
	livePlayer *models.LiveMatchPlayer,
	matchPlayer *models.MatchPlayer,
	statsPlayers []*models.LiveMatchStatsPlayer,
) *nspb.Player_Match_PlayerDetails {
	pb := &nspb.Player_Match_PlayerDetails{}

	if livePlayer != nil {
		pb.HeroId = uint64(livePlayer.HeroID)
	}

	if matchPlayer != nil {
		if pb.HeroId == 0 {
			pb.HeroId = uint64(matchPlayer.HeroID)
		}

		pb.PlayerSlot = uint32(matchPlayer.PlayerSlot)
		pb.ProName = matchPlayer.ProName
		pb.Kills = matchPlayer.Kills
		pb.Deaths = matchPlayer.Deaths
		pb.Assists = matchPlayer.Assists
		pb.Items = matchPlayer.Items
	}

	for _, statsPlayer := range statsPlayers {
		if pb.HeroId == 0 {
			pb.HeroId = uint64(statsPlayer.HeroID)
		}

		if pb.PlayerSlot == 0 {
			pb.PlayerSlot = uint32(statsPlayer.PlayerSlot)
		}
	}

	return pb
}
