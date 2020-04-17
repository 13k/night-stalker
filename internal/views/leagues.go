package views

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

func NewLeagues(leagues nscol.Leagues) (*nspb.Leagues, error) {
	if len(leagues) == 0 {
		return nil, nil
	}

	view := &nspb.Leagues{
		Leagues: make([]*nspb.League, len(leagues)),
	}

	var err error

	for i, league := range leagues {
		view.Leagues[i], err = NewLeague(league)

		if err != nil {
			return nil, xerrors.Errorf("error creating League view: %w", err)
		}
	}

	return view, nil
}
