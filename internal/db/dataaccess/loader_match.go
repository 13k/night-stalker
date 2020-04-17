package dataaccess

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nstime "github.com/13k/night-stalker/internal/time"
	nsm "github.com/13k/night-stalker/models"
)

const (
	defaultMatchHistoryMinDuration = 5 * time.Minute
	defaultMatchHistoryAgeDays     = 7
)

var (
	defaultMatchHistoryFilters = MatchFilters{
		MinDuration: defaultMatchHistoryMinDuration,
		Since:       nstime.TravelFrom(time.Now()).DaysAgo(defaultMatchHistoryAgeDays).BeginningOfDay().T(),
	}
)

func (l *Loader) FindMatchIDs(ctx context.Context, filters MatchFilters) (nscol.MatchIDs, error) {
	if err := filters.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid filters: %w", err)
	}

	tMatch := nsm.MatchTable
	tMatchPlayer := nsm.MatchPlayerTable
	tLiveMatchPlayer := nsm.LiveMatchPlayerTable
	tFollowedPlayer := nsm.FollowedPlayerTable

	var matchIDs nscol.MatchIDs
	var queries []*nsdb.SelectQuery

	baseQ := l.mq.
		Q().
		Select().
		Filter(filters).
		Trace()

	if !filters.Players.Empty() {
		queries = append(
			queries,
			baseQ.
				InnerJoinEq(
					tMatch.PK(),
					tMatchPlayer.Col("match_id"),
				).
				InnerJoinEq(
					tMatchPlayer.Col("account_id"),
					tFollowedPlayer.Col("account_id"),
				).
				Filter(MatchPlayerFilters{PlayerStatsFilters: filters.Players}),
			baseQ.
				InnerJoinEq(
					tMatch.PK(),
					tLiveMatchPlayer.Col("match_id"),
				).
				InnerJoinEq(
					tLiveMatchPlayer.Col("account_id"),
					tFollowedPlayer.Col("account_id"),
				).
				Filter(LiveMatchPlayerFilters{PlayerStatsFilters: filters.Players}),
		)
	}

	if len(queries) == 0 {
		queries = append(queries, baseQ)
	}

	for _, q := range queries {
		ids, err := l.mq.M().PluckID(ctx, nsm.MatchModel, q)

		if err != nil {
			return nil, xerrors.Errorf("error finding match IDs: %w", err)
		}

		matchIDs = matchIDs.AddUnique(nscol.NewMatchIDsModelIDs(ids...)...)
	}

	return matchIDs, nil
}

func (l *Loader) MatchesData(ctx context.Context, matchIDs ...nspb.MatchID) (MatchesData, error) {
	if len(matchIDs) == 0 {
		return nil, xerrors.Errorf("invalid matchIDs: %w", ErrEmptyMatchIDs)
	}

	tMatch := nsm.MatchTable
	tLiveMatch := nsm.LiveMatchTable
	tLiveMatchStats := nsm.LiveMatchStatsTable

	var matches nscol.Matches

	q := l.mq.
		Q().
		Select().
		Prepared(true).
		Trace()

	if len(matchIDs) == 1 {
		q = q.Eq(tMatch.PK(), matchIDs[0])
	} else {
		q = q.In(tMatch.PK(), matchIDs)
	}

	if err := l.mq.M().FindAll(ctx, nsm.MatchModel, q, &matches); err != nil {
		return nil, xerrors.Errorf("error loading matches: %w", err)
	}

	if err := l.mq.M().Eagerload(ctx, "Players", matches.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading match players: %w", err)
	}

	var liveMatches nscol.LiveMatches

	q = l.mq.
		Q().
		Select().
		Prepared(true).
		Trace()

	if len(matchIDs) == 1 {
		q = q.Eq(tLiveMatch.Col("match_id"), matchIDs[0])
	} else {
		q = q.In(tLiveMatch.Col("match_id"), matchIDs)
	}

	if err := l.mq.M().FindAll(ctx, nsm.LiveMatchModel, q, &liveMatches); err != nil {
		return nil, xerrors.Errorf("error loading live matches: %w", err)
	}

	if err := l.mq.M().Eagerload(ctx, "Players", liveMatches.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading live match players: %w", err)
	}

	var stats nscol.LiveMatchStats

	q = l.mq.
		Q().
		Select().
		Prepared(true).
		Trace()

	if len(matchIDs) == 1 {
		q = q.Eq(tLiveMatchStats.Col("match_id"), matchIDs[0])
	} else {
		q = q.In(tLiveMatchStats.Col("match_id"), matchIDs)
	}

	if err := l.mq.M().FindAll(ctx, nsm.LiveMatchStatsModel, q, &stats); err != nil {
		return nil, xerrors.Errorf("error loading live match stats: %w", err)
	}

	if err := l.mq.M().Eagerload(ctx, "Players", stats.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading live match stats players: %w", err)
	}

	var allMatchIDs nscol.MatchIDs

	allMatchIDs = allMatchIDs.AddUnique(matches.MatchIDs()...)
	allMatchIDs = allMatchIDs.AddUnique(liveMatches.MatchIDs()...)
	allMatchIDs = allMatchIDs.AddUnique(stats.MatchIDs()...)

	matchesByMatchID := matches.KeyByMatchID()
	liveMatchesByMatchID := liveMatches.KeyByMatchID()
	statsByMatchID := stats.GroupByMatchID()

	data := make([]*MatchData, len(allMatchIDs))

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

		playersData := make([]*MatchPlayerData, len(accountIDs))

		for j, accountID := range accountIDs {
			var matchPlayer *nsm.MatchPlayer
			var livePlayer *nsm.LiveMatchPlayer

			matchPlayersGroup := matchPlayersByAccountID[accountID]
			livePlayersGroup := livePlayersByAccountID[accountID]
			statsPlayersGroup := statsPlayersByAccountID[accountID]

			if len(matchPlayersGroup) > 0 {
				matchPlayer = matchPlayersGroup[0]
			}

			if len(livePlayersGroup) > 0 {
				livePlayer = livePlayersGroup[0]
			}

			playersData[j] = NewMatchPlayerData(
				matchPlayer,
				livePlayer,
				statsPlayersGroup,
			)
		}

		matchPlayersData, err := NewMatchPlayersData(playersData...)

		if err != nil {
			return nil, xerrors.Errorf("error creating match players data: %w", err)
		}

		data[i] = NewMatchData(
			match,
			liveMatch,
			stats,
			matchPlayersData,
		)
	}

	matchesData, err := NewMatchesData(data...)

	if err != nil {
		return nil, xerrors.Errorf("error creating matches data: %w", err)
	}

	return matchesData, nil
}
