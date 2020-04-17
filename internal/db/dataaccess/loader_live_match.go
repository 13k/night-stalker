package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

func (l *Loader) LiveMatches(ctx context.Context, filters LiveMatchFilters) (nscol.LiveMatches, error) {
	if err := filters.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid filters: %w", err)
	}

	var liveMatches nscol.LiveMatches

	if err := l.mq.M().Filter(ctx, nsm.LiveMatchModel, filters, &liveMatches); err != nil {
		return nil, xerrors.Errorf("error loading live matches: %w", err)
	}

	return liveMatches, nil
}

type LiveMatchesParams struct {
	MatchIDs         nscol.MatchIDs
	WithFollowedOnly bool
}

func (l *Loader) LiveMatchesData(ctx context.Context, params *LiveMatchesParams) (*LiveMatchesData, error) {
	if len(params.MatchIDs) == 0 {
		return nil, nil
	}

	liveMatches, err := l.LiveMatches(ctx, LiveMatchFilters{
		MatchIDs:         params.MatchIDs,
		WithFollowedOnly: params.WithFollowedOnly,
		OrderBy:          nsdb.OrderFields.Desc("sort_score"),
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading live matches: %w", err)
	}

	if len(liveMatches) == 0 {
		return nil, nil
	}

	if err = l.mq.M().Eagerload(ctx, "Players", liveMatches.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading live match players: %w", err)
	}

	var accountIDs nscol.AccountIDs

	for _, liveMatch := range liveMatches {
		livePlayers := nscol.LiveMatchPlayers(liveMatch.Players)
		accountIDs = accountIDs.AddUnique(livePlayers.AccountIDs()...)
	}

	if len(accountIDs) == 0 {
		return nil, nil
	}

	stats, err := l.LiveMatchStats(ctx, LiveMatchStatsFilters{
		MatchIDs: liveMatches.MatchIDs(),
		Latest:   1,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading live match stats: %w", err)
	}

	if err = l.mq.M().Eagerload(ctx, "Players", stats.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading live match stats players: %w", err)
	}

	if err = l.mq.M().Eagerload(ctx, "Teams", stats.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading live match stats teams: %w", err)
	}

	statsByMatchID := stats.KeyByMatchID()

	followedPlayers, err := l.FollowedPlayers(ctx, PlayerMetaFilters{
		AccountIDs: accountIDs,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading followed players: %w", err)
	}

	accountIDs = followedPlayers.AccountIDs()

	if len(accountIDs) == 0 {
		return nil, nil
	}

	followedByAccountID := followedPlayers.KeyByAccountID()

	players, err := l.Players(ctx, PlayerMetaFilters{
		AccountIDs: accountIDs,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading players: %w", err)
	}

	playersByAccountID := players.KeyByAccountID()

	proPlayers, err := l.ProPlayers(ctx, PlayerMetaFilters{
		AccountIDs: accountIDs,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading pro players: %w", err)
	}

	prosByAccountID := proPlayers.KeyByAccountID()

	data := &LiveMatchesData{
		LiveMatches:                liveMatches,
		LiveMatchStatsByMatchID:    statsByMatchID,
		FollowedPlayersByAccountID: followedByAccountID,
		PlayersByAccountID:         playersByAccountID,
		ProPlayersByAccountID:      prosByAccountID,
	}

	return data, nil
}
