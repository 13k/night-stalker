package web

import (
	"errors"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

var (
	errEmptyAccountIDs = errors.New("empty account IDs")
)

func (app *App) filterFollowedPlayersAccountIDs(accountIDs ...nspb.AccountID) (nscol.AccountIDs, error) {
	if len(accountIDs) == 0 {
		err := xerrors.Errorf("invalid accountIDs: %w", errEmptyAccountIDs)
		return nil, err
	}

	resultAccountIDs := nscol.AccountIDs{}

	err := app.db.
		Debug().
		Model(models.FollowedPlayerModel).
		Where("account_id IN (?)", accountIDs).
		Pluck("account_id", &resultAccountIDs).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading followed players account IDs: %w", err)
		return nil, err
	}

	return resultAccountIDs, nil
}

func (app *App) loadPlayersData(accountIDs ...nspb.AccountID) (nsviews.PlayersData, error) {
	if len(accountIDs) == 0 {
		err := xerrors.Errorf("invalid accountIDs: %w", errEmptyAccountIDs)
		return nil, err
	}

	var followedPlayers nscol.FollowedPlayers

	err := app.db.
		Debug().
		Where("account_id IN (?)", accountIDs).
		Find(&followedPlayers).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading followed players: %w", err)
		return nil, err
	}

	if len(followedPlayers) == 0 {
		return nil, nil
	}

	// only fetch data of followed players
	accountIDs = followedPlayers.AccountIDs()

	var players nscol.Players

	err = app.db.
		Debug().
		Where("account_id IN (?)", accountIDs).
		Find(&players).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading players: %w", err)
		return nil, err
	}

	var proPlayers nscol.ProPlayers

	err = app.db.
		Debug().
		Where("account_id IN (?)", accountIDs).
		Preload("Team").
		Find(&proPlayers).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading pro players: %w", err)
		return nil, err
	}

	playersByAccountID := players.KeyByAccountID()
	proPlayersByAccountID := proPlayers.KeyByAccountID()
	data := make([]*nsviews.PlayerData, len(followedPlayers))

	for i, followedPlayer := range followedPlayers {
		accountID := followedPlayer.AccountID

		data[i] = nsviews.NewPlayerData(
			followedPlayer,
			playersByAccountID[accountID],
			proPlayersByAccountID[accountID],
		)
	}

	playersData, err := nsviews.NewPlayersData(data...)

	if err != nil {
		err = xerrors.Errorf("error creating players data view: %w", err)
		return nil, err
	}

	return playersData, nil
}

func (app *App) loadPlayerMatchesData(accountID nspb.AccountID) (nsviews.MatchesData, error) {
	matchIDs, err := app.findMatchIDs(&findMatchIDsFilters{
		PlayerFilters: &nsdb.PlayerFilters{
			AccountIDs: nscol.AccountIDs{accountID},
		},
	})

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	if len(matchIDs) == 0 {
		return nil, nil
	}

	matchesData, err := app.loadMatchesData(matchIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading matches data: %w", err)
		return nil, err
	}

	return matchesData, nil
}
