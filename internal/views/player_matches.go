package views

import (
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewPlayerMatches(data *nsdbda.PlayerMatchesData) (*nspb.PlayerMatches, error) {
	if data == nil {
		return nil, nil
	}

	pbPlayer, err := NewPlayer(data.PlayerData)

	if err != nil {
		return nil, xerrors.Errorf("error creating Player view: %w", err)
	}

	pb := &nspb.PlayerMatches{
		Player: pbPlayer,
	}

	pb.KnownPlayers, err = NewSortedPlayers(data.KnownPlayers)

	if err != nil {
		return nil, xerrors.Errorf("error creating Player views: %w", err)
	}

	pb.Matches, err = NewSortedMatches(data.MatchesData)

	if err != nil {
		return nil, xerrors.Errorf("error creating Match views: %w", err)
	}

	return pb, nil
}
