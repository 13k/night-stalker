package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsm "github.com/13k/night-stalker/models"
)

func (l *Loader) LiveMatchStats(ctx context.Context, filters LiveMatchStatsFilters) (nscol.LiveMatchStats, error) {
	var stats nscol.LiveMatchStats

	if err := l.mq.M().Filter(ctx, nsm.LiveMatchStatsModel, filters, &stats); err != nil {
		return nil, xerrors.Errorf("error loading live match stats: %w", err)
	}

	return stats, nil
}
