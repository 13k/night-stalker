package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

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
