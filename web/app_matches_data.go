package web

import (
	"errors"
	"time"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nstime "github.com/13k/night-stalker/internal/time"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

var (
	errEmptyFilters  = errors.New("empty filters")
	errEmptyMatchIDs = errors.New("empty match IDs")
)

type findMatchIDsFilters struct {
	*nsdb.PlayerFilters
	Since time.Time
}

func (app *App) findMatchIDs(filters *findMatchIDsFilters) (nscol.MatchIDs, error) {
	if filters == nil || filters.Empty() {
		err := xerrors.Errorf("invalid filters: %w", errEmptyFilters)
		return nil, err
	}

	if filters.Since.IsZero() {
		filters.Since = nstime.TravelFrom(time.Now()).DaysAgo(15).BeginningOfDay().T()
	}

	var matchIDs nscol.MatchIDs
	resultMatchIDs := nscol.MatchIDs{}

	baseScope := app.db.
		Debug().
		Model(models.MatchModel)

	baseScope = nsdb.GtEq(baseScope, models.MatchModel, "created_at", filters.Since)
	baseScope = nsdb.Group(baseScope, models.MatchModel, "id")

	scope := baseScope.
		Joins("INNER JOIN match_players ON (matches.id = match_players.match_id)").
		Joins(`
INNER JOIN followed_players ON
	(match_players.account_id = followed_players.account_id OR match_players.account_id = 0)
		`)

	scope = filters.Filter(scope, models.MatchPlayerModel)
	err := scope.Pluck("matches.id", &resultMatchIDs).Error

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	matchIDs = matchIDs.AddUnique(resultMatchIDs...)
	resultMatchIDs = nscol.MatchIDs{}

	scope = baseScope.
		Joins("INNER JOIN live_match_players ON (matches.id = live_match_players.match_id)").
		Joins("INNER JOIN followed_players ON (live_match_players.account_id = followed_players.account_id)")

	scope = filters.Filter(scope, models.LiveMatchPlayerModel)
	err = scope.Pluck("matches.id", &resultMatchIDs).Error

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	matchIDs = matchIDs.AddUnique(resultMatchIDs...)
	resultMatchIDs = nscol.MatchIDs{}

	scope = baseScope.
		Joins("INNER JOIN live_match_stats_players ON (matches.id = live_match_stats_players.match_id)").
		Joins("INNER JOIN followed_players ON (live_match_stats_players.account_id = followed_players.account_id)")

	scope = filters.Filter(scope, models.LiveMatchStatsPlayerModel)
	err = scope.Pluck("matches.id", &resultMatchIDs).Error

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	matchIDs = matchIDs.AddUnique(resultMatchIDs...)

	return matchIDs, nil
}

func (app *App) loadMatchesData(matchIDs ...nspb.MatchID) (nsviews.MatchesData, error) {
	if len(matchIDs) == 0 {
		err := xerrors.Errorf("invalid matchIDs: %w", errEmptyMatchIDs)
		return nil, err
	}

	var matches nscol.Matches

	err := app.db.
		Debug().
		Where("id IN (?)", matchIDs).
		Preload("Players").
		Find(&matches).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading matches: %w", err)
		return nil, err
	}

	var liveMatches nscol.LiveMatches

	err = app.db.
		Debug().
		Where("match_id IN (?)", matchIDs).
		Preload("Players").
		Find(&liveMatches).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return nil, err
	}

	var stats nscol.LiveMatchStats

	err = app.db.
		Debug().
		Where("match_id IN (?)", matchIDs).
		Preload("Players").
		Find(&stats).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading live match stats: %w", err)
		return nil, err
	}

	var allMatchIDs nscol.MatchIDs

	allMatchIDs = allMatchIDs.AddUnique(matches.MatchIDs()...)
	allMatchIDs = allMatchIDs.AddUnique(liveMatches.MatchIDs()...)
	allMatchIDs = allMatchIDs.AddUnique(stats.MatchIDs()...)

	matchesByMatchID := matches.KeyByMatchID()
	liveMatchesByMatchID := liveMatches.KeyByMatchID()
	statsByMatchID := stats.GroupByMatchID()

	data := make([]*nsviews.MatchData, len(allMatchIDs))

	for i, matchID := range allMatchIDs {
		match := matchesByMatchID[matchID]
		liveMatch := liveMatchesByMatchID[matchID]
		stats := statsByMatchID[matchID]

		var matchPlayers nscol.MatchPlayers
		var livePlayers nscol.LiveMatchPlayers
		var statsPlayers nscol.LiveMatchStatsPlayers

		if match != nil {
			matchPlayers = match.Players
		}

		if liveMatch != nil {
			livePlayers = liveMatch.Players
		}

		for _, s := range stats {
			for _, p := range s.Players {
				statsPlayers = append(statsPlayers, p)
			}
		}

		var accountIDs nscol.AccountIDs

		accountIDs = accountIDs.AddUnique(matchPlayers.AccountIDs()...)
		accountIDs = accountIDs.AddUnique(livePlayers.AccountIDs()...)
		accountIDs = accountIDs.AddUnique(statsPlayers.AccountIDs()...)

		matchPlayersByAccountID := matchPlayers.GroupByAccountID()
		livePlayersByAccountID := livePlayers.GroupByAccountID()
		statsPlayersByAccountID := statsPlayers.GroupByAccountID()

		playersData := make([]*nsviews.MatchPlayerData, len(accountIDs))

		for j, accountID := range accountIDs {
			var matchPlayer *models.MatchPlayer
			var livePlayer *models.LiveMatchPlayer

			matchPlayersGroup := matchPlayersByAccountID[accountID]
			livePlayersGroup := livePlayersByAccountID[accountID]
			statsPlayersGroup := statsPlayersByAccountID[accountID]

			if len(matchPlayersGroup) > 0 {
				matchPlayer = matchPlayersGroup[0]
			}

			if len(livePlayersGroup) > 0 {
				livePlayer = livePlayersGroup[0]
			}

			playersData[j] = nsviews.NewMatchPlayerData(
				matchPlayer,
				livePlayer,
				statsPlayersGroup,
			)
		}

		var matchPlayersData nsviews.MatchPlayersData

		matchPlayersData, err = nsviews.NewMatchPlayersData(playersData...)

		if err != nil {
			err = xerrors.Errorf("error creating match players data: %w", err)
			return nil, err
		}

		data[i] = nsviews.NewMatchData(
			match,
			liveMatch,
			stats,
			matchPlayersData,
		)
	}

	matchesData, err := nsviews.NewMatchesData(data...)

	if err != nil {
		err = xerrors.Errorf("error creating matches data: %w", err)
		return nil, err
	}

	return matchesData, nil
}
