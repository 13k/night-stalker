package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewHeroes(heroes []*models.Hero) []*nspb.Hero {
	view := make([]*nspb.Hero, len(heroes))

	for i, hero := range heroes {
		view[i] = HeroProto(hero)
	}

	return view
}
