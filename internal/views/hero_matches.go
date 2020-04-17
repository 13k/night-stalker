package views

import (
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewHeroMatches(data *nsdbda.HeroMatchesData) (*nspb.HeroMatches, error) {
	if data == nil {
		return nil, nil
	}

	pb := &nspb.HeroMatches{
		Hero: NewHero(data.Hero),
	}

	var err error

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
