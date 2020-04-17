package views

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewHeroes(heroes nscol.Heroes) *nspb.Heroes {
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
