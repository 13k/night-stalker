package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

func (l *Loader) Heroes(ctx context.Context, filters HeroFilters) (nscol.Heroes, error) {
	var heroes nscol.Heroes

	// allow empty filters (loads all heroes)
	if err := l.mq.M().Filter(ctx, nsm.HeroModel, filters, &heroes); err != nil {
		return nil, xerrors.Errorf("error loading heroes: %w", err)
	}

	return heroes, nil
}

func (l *Loader) MatchesDataForHero(ctx context.Context, id nspb.HeroID) (MatchesData, error) {
	filters := defaultMatchHistoryFilters

	filters.Players = PlayerStatsFilters{
		HeroIDs: nscol.HeroIDs{id},
	}

	matchIDs, err := l.FindMatchIDs(ctx, filters)

	if err != nil {
		return nil, xerrors.Errorf("error finding match IDs: %w", err)
	}

	if len(matchIDs) == 0 {
		return nil, nil
	}

	matchesData, err := l.MatchesData(ctx, matchIDs...)

	if err != nil {
		return nil, xerrors.Errorf("error loading matches data: %w", err)
	}

	return matchesData, nil
}

type HeroMatchesParams struct {
	HeroID nspb.HeroID
}

func (l *Loader) HeroMatchesData(ctx context.Context, params *HeroMatchesParams) (*HeroMatchesData, error) {
	heroes, err := l.Heroes(ctx, HeroFilters{
		HeroIDs: nscol.HeroIDs{params.HeroID},
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading heroes data: %w", err)
	}

	if len(heroes) == 0 {
		return nil, nil
	}

	hero := heroes[0]
	matchesData, err := l.MatchesDataForHero(ctx, params.HeroID)

	if err != nil {
		return nil, xerrors.Errorf("error loading hero matches data: %w", err)
	}

	playersAccountIDs := matchesData.AccountIDs()

	var followedAccountIDs nscol.AccountIDs

	if len(playersAccountIDs) > 0 {
		followedAccountIDs, err = l.FilterFollowedPlayersAccountIDs(ctx, playersAccountIDs...)

		if err != nil {
			return nil, xerrors.Errorf("error filtering followed players account IDs: %w", err)
		}
	}

	var knownPlayersData PlayersData

	if len(followedAccountIDs) > 0 {
		knownPlayersData, err = l.PlayersData(ctx, followedAccountIDs...)

		if err != nil {
			return nil, xerrors.Errorf("error loading players data: %w", err)
		}
	}

	data := &HeroMatchesData{
		Hero:         hero,
		KnownPlayers: knownPlayersData,
		MatchesData:  matchesData,
	}

	return data, nil
}
