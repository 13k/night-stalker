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
	livePlayers map[nspb.MatchID]*models.LiveMatchPlayer,
	statsPlayers map[nspb.MatchID]*models.LiveMatchStatsPlayer,
) *nspb.Player {
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

	for matchID := range livePlayers {
		pb.Matches = append(pb.Matches, NewPlayerMatch(
			matchID,
			livePlayers[matchID],
			statsPlayers[matchID],
		))
	}

	sort.Slice(pb.Matches, func(i, j int) bool {
		return pb.Matches[i].MatchId < pb.Matches[j].MatchId
	})

	return pb
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
	matchID nspb.MatchID,
	livePlayer *models.LiveMatchPlayer,
	statsPlayer *models.LiveMatchStatsPlayer,
) *nspb.Player_Match {
	match := &nspb.Player_Match{
		MatchId: matchID,
		HeroId:  uint64(livePlayer.HeroID),
	}

	if match.HeroId == 0 && statsPlayer != nil {
		match.HeroId = uint64(statsPlayer.HeroID)
	}

	return match
}
