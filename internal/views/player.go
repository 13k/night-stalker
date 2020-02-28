package views

import (
	"sort"

	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protocol"
)

func NewPlayer(data *PlayerData) (*nspb.Player, error) {
	if err := data.Validate(); err != nil {
		err = xerrors.Errorf("invalid PlayerData: %w", err)
		return nil, err
	}

	pb := &nspb.Player{
		AccountId: data.AccountID,
	}

	if pb.Name == "" && data.FollowedPlayer != nil {
		pb.Name = data.FollowedPlayer.Label
	}

	if data.Player != nil {
		if pb.Name == "" {
			pb.Name = data.Player.Name
		}

		pb.PersonaName = data.Player.PersonaName
		pb.AvatarUrl = data.Player.AvatarURL
		pb.AvatarMediumUrl = data.Player.AvatarMediumURL
		pb.AvatarFullUrl = data.Player.AvatarFullURL
	}

	if data.ProPlayer != nil {
		pb.IsPro = true

		if data.ProPlayer.Team != nil {
			pb.Team = NewTeam(data.ProPlayer.Team)
		}
	}

	return pb, nil
}

func NewPlayers(data PlayersData) ([]*nspb.Player, error) {
	if len(data) == 0 {
		return nil, nil
	}

	players := make([]*nspb.Player, 0, len(data))

	for _, playerData := range data {
		player, err := NewPlayer(playerData)

		if err != nil {
			err = xerrors.Errorf("error creating Player view: %w", err)
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

func NewSortedPlayers(data PlayersData) ([]*nspb.Player, error) {
	players, err := NewPlayers(data)

	if err != nil {
		err = xerrors.Errorf("error creating Player views: %w", err)
		return nil, err
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].AccountId < players[j].AccountId
	})

	return players, nil
}

func NewPlayerMatches(
	data *PlayerData,
	knownPlayers PlayersData,
	matchesData MatchesData,
) (*nspb.PlayerMatches, error) {
	pbPlayer, err := NewPlayer(data)

	if err != nil {
		err = xerrors.Errorf("error creating Player view: %w", err)
		return nil, err
	}

	pb := &nspb.PlayerMatches{
		Player: pbPlayer,
	}

	pb.KnownPlayers, err = NewSortedPlayers(knownPlayers)

	if err != nil {
		err = xerrors.Errorf("error creating Player views: %w", err)
		return nil, err
	}

	pb.Matches, err = NewSortedMatches(matchesData)

	if err != nil {
		err = xerrors.Errorf("error creating Match views: %w", err)
		return nil, err
	}

	return pb, nil
}
