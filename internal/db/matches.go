package db

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nstime "github.com/13k/night-stalker/internal/time"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

const (
	defaultMatchIDsMinDuration = 5 * time.Minute
	defaultMatchIDsAgeDays     = 7
)

var (
	ErrEmptyMatchIDs = errors.New("empty match IDs")
)

type FindMatchIDsFilters struct {
	*PlayerFilters
	MinDuration time.Duration
	MaxDuration time.Duration
	Since       time.Time
}

func FindMatchIDs(db *gorm.DB, filters FindMatchIDsFilters) (nscol.MatchIDs, error) {
	if err := filters.Validate(); err != nil {
		err = xerrors.Errorf("invalid filters: %w", err)
		return nil, err
	}

	if filters.MinDuration == 0 {
		filters.MinDuration = defaultMatchIDsMinDuration
	}

	if filters.Since.IsZero() {
		filters.Since = nstime.TravelFrom(time.Now()).DaysAgo(defaultMatchIDsAgeDays).BeginningOfDay().T()
	}

	var matchIDs nscol.MatchIDs
	resultMatchIDs := nscol.MatchIDs{}

	baseScope := db.
		Debug().
		Model(models.MatchModel)

	baseScope = GtEq(baseScope, models.MatchModel, "created_at", filters.Since)
	baseScope = GtEq(baseScope, models.MatchModel, "duration", int64(filters.MinDuration/time.Second))
	baseScope = Group(baseScope, models.MatchModel, "id")

	if filters.MaxDuration != 0 {
		baseScope = LtEq(baseScope, models.MatchModel, "duration", int64(filters.MaxDuration/time.Second))
	}

	scope := baseScope.
		Joins(`INNER JOIN match_players ON (matches.id = match_players.match_id)`).
		Joins(`INNER JOIN followed_players ON (match_players.account_id = followed_players.account_id)`)

	scope = filters.Filter(scope, models.MatchPlayerModel)
	err := scope.Pluck("matches.id", &resultMatchIDs).Error

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	matchIDs = matchIDs.AddUnique(resultMatchIDs...)
	resultMatchIDs = nscol.MatchIDs{}

	scope = baseScope.
		Joins(`INNER JOIN live_match_players ON (matches.id = live_match_players.match_id)`).
		Joins(`INNER JOIN followed_players ON (live_match_players.account_id = followed_players.account_id)`)

	scope = filters.Filter(scope, models.LiveMatchPlayerModel)
	err = scope.Pluck(`matches.id`, &resultMatchIDs).Error

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	matchIDs = matchIDs.AddUnique(resultMatchIDs...)

	return matchIDs, nil
}

func LoadMatchesData(db *gorm.DB, matchIDs ...nspb.MatchID) (nsviews.MatchesData, error) {
	if len(matchIDs) == 0 {
		err := xerrors.Errorf("invalid matchIDs: %w", ErrEmptyMatchIDs)
		return nil, err
	}

	var matches nscol.Matches

	err := db.
		Debug().
		Where(`id IN (?)`, matchIDs).
		Preload("Players").
		Find(&matches).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading matches: %w", err)
		return nil, err
	}

	var liveMatches nscol.LiveMatches

	err = db.
		Debug().
		Where(`match_id IN (?)`, matchIDs).
		Preload("Players").
		Find(&liveMatches).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return nil, err
	}

	var stats nscol.LiveMatchStats

	err = db.
		Debug().
		Where(`match_id IN (?)`, matchIDs).
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
