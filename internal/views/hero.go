package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewHero(h *models.Hero) *nspb.Hero {
	return &nspb.Hero{
		Id:                 uint64(h.ID),
		Name:               h.Name,
		Slug:               h.Slug,
		LocalizedName:      h.LocalizedName,
		Aliases:            h.Aliases,
		Roles:              h.Roles,
		RoleLevels:         h.RoleLevels,
		Complexity:         int64(h.Complexity),
		Legs:               int64(h.Legs),
		AttributePrimary:   h.AttributePrimary,
		AttackCapabilities: h.AttackCapabilities,
	}
}

func NewHeroes(heroes []*models.Hero) []*nspb.Hero {
	if len(heroes) == 0 {
		return nil
	}

	view := make([]*nspb.Hero, len(heroes))

	for i, hero := range heroes {
		view[i] = NewHero(hero)
	}

	return view
}

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
