package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLeagues(leagues []*models.League) (*nspb.Leagues, error) {
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
			err = xerrors.Errorf("error creating League view: %w", err)
			return nil, err
		}
	}

	return view, nil
}
