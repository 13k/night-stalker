package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

func (l *Loader) FollowedPlayers(ctx context.Context, filters PlayerMetaFilters) (nscol.FollowedPlayers, error) {
	f := &FollowedPlayerFilters{PlayerMetaFilters: filters}

	if err := f.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid filters: %w", err)
	}

	var players nscol.FollowedPlayers

	if err := l.mq.M().Filter(ctx, nsm.FollowedPlayerModel, f, &players); err != nil {
		return nil, xerrors.Errorf("error loading followed players: %w", err)
	}

	return players, nil
}

func (l *Loader) Players(ctx context.Context, filters PlayerMetaFilters) (nscol.Players, error) {
	f := &PlayerFilters{PlayerMetaFilters: filters}

	if err := f.Validate(); err != nil {
		return nil, xerrors.Errorf("invalid filters: %w", err)
	}

	var players nscol.Players

	if err := l.mq.M().Filter(ctx, nsm.PlayerModel, f, &players); err != nil {
		return nil, xerrors.Errorf("error loading players: %w", err)
	}

	return players, nil
}

func (l *Loader) FilterFollowedPlayersAccountIDs(
	ctx context.Context,
	accountIDs ...nspb.AccountID,
) (nscol.AccountIDs, error) {
	if len(accountIDs) == 0 {
		return nil, xerrors.Errorf("invalid accountIDs: %w", ErrEmptyAccountIDs)
	}

	var resultAccountIDs nscol.AccountIDs

	q := l.mq.
		Q().
		Select().
		In(nsm.FollowedPlayerTable.Col("account_id"), accountIDs).
		Prepared(true).
		Trace()

	err := l.mq.M().PluckCol(ctx, nsm.FollowedPlayerModel, "account_id", q, &resultAccountIDs)

	if err != nil {
		return nil, xerrors.Errorf("error filtering followed players account IDs: %w", err)
	}

	return resultAccountIDs, nil
}

func (l *Loader) PlayersData(ctx context.Context, accountIDs ...nspb.AccountID) (PlayersData, error) {
	if len(accountIDs) == 0 {
		return nil, xerrors.Errorf("invalid accountIDs: %w", ErrEmptyAccountIDs)
	}

	tFollowedPlayer := nsm.FollowedPlayerTable
	tPlayer := nsm.PlayerTable

	var followedPlayers nscol.FollowedPlayers

	q := l.mq.
		Q().
		Select().
		In(tFollowedPlayer.Col("account_id"), accountIDs).
		Prepared(true).
		Trace()

	err := l.mq.M().FindAll(ctx, nsm.FollowedPlayerModel, q, &followedPlayers)

	if err != nil {
		return nil, xerrors.Errorf("error loading followed players: %w", err)
	}

	if len(followedPlayers) == 0 {
		return nil, nil
	}

	// only fetch data of followed players
	accountIDs = followedPlayers.AccountIDs()

	var players nscol.Players

	q = l.mq.
		Q().
		Select().
		In(tPlayer.Col("account_id"), accountIDs).
		Prepared(true).
		Trace()

	err = l.mq.M().FindAll(ctx, nsm.PlayerModel, q, &players)

	if err != nil {
		return nil, xerrors.Errorf("error loading players: %w", err)
	}

	if err = l.mq.M().Eagerload(ctx, "Team", players.Records()...); err != nil {
		return nil, xerrors.Errorf("error loading pro players' teams: %w", err)
	}

	playersByAccountID := players.KeyByAccountID()
	data := make([]*PlayerData, len(followedPlayers))

	for i, followedPlayer := range followedPlayers {
		accountID := followedPlayer.AccountID

		data[i] = NewPlayerData(
			followedPlayer,
			playersByAccountID[accountID],
		)
	}

	playersData, err := NewPlayersData(data...)

	if err != nil {
		return nil, xerrors.Errorf("error creating players data view: %w", err)
	}

	return playersData, nil
}

func (l *Loader) MatchesDataForPlayer(ctx context.Context, accountID nspb.AccountID) (MatchesData, error) {
	filters := defaultMatchHistoryFilters

	filters.Players = PlayerStatsFilters{
		AccountIDs: nscol.AccountIDs{accountID},
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

type PlayerMatchesParams struct {
	AccountID nspb.AccountID
}

func (l *Loader) PlayerMatchesData(ctx context.Context, params *PlayerMatchesParams) (*PlayerMatchesData, error) {
	playersData, err := l.PlayersData(ctx, params.AccountID)

	if err != nil {
		return nil, xerrors.Errorf("error loading players data: %w", err)
	}

	if playersData == nil {
		return nil, nil
	}

	playerData := playersData[params.AccountID]

	if playerData == nil {
		return nil, nil
	}

	matchesData, err := l.MatchesDataForPlayer(ctx, params.AccountID)

	if err != nil {
		return nil, xerrors.Errorf("error loading player matches data: %w", err)
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

	data := &PlayerMatchesData{
		PlayerData:   playerData,
		KnownPlayers: knownPlayersData,
		MatchesData:  matchesData,
	}

	return data, nil
}
