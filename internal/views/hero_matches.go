package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewHeroMatches(
	hero *models.Hero,
	knownPlayers PlayersData,
	matchesData MatchesData,
) (*nspb.HeroMatches, error) {
	pb := &nspb.HeroMatches{
		Hero: NewHero(hero),
	}

	var err error

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
