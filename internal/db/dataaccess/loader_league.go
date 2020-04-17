package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsm "github.com/13k/night-stalker/models"
)

func (l *Loader) Leagues(ctx context.Context, filters LeagueFilters) (nscol.Leagues, error) {
	if err := filters.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid filters: %w", err)
	}

	var leagues nscol.Leagues

	if err := l.mq.M().Filter(ctx, nsm.LeagueModel, filters, &leagues); err != nil {
		return nil, xerrors.Errorf("error loading leagues: %w", err)
	}

	return leagues, nil
}
