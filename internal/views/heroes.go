package views

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewHeroes(heroes []*models.Hero) *nspb.Heroes {
	if len(heroes) == 0 {
		return nil
	}

	view := &nspb.Heroes{
		Heroes: make([]*nspb.Hero, len(heroes)),
	}

	for i, hero := range heroes {
		view.Heroes[i] = NewHero(hero)
	}

	return view
}
