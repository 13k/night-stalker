package dataaccess

import (
	"context"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
)

type SearchParams struct {
	Query string
}

func (l *Loader) SearchData(ctx context.Context, params *SearchParams) (*SearchData, error) {
	heroes, err := l.Heroes(ctx, HeroFilters{
		Query: params.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading heroes: %w", err)
	}

	followed, err := l.FollowedPlayers(ctx, PlayerMetaFilters{
		Query: params.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading followed players: %w", err)
	}

	var players nscol.Players
	var proPlayers nscol.ProPlayers

	if len(followed) > 0 {
		accountIDs := followed.AccountIDs()

		players, err = l.Players(ctx, PlayerMetaFilters{
			AccountIDs: accountIDs,
		})

		if err != nil {
			return nil, xerrors.Errorf("error loading players: %w", err)
		}

		proPlayers, err = l.ProPlayers(ctx, PlayerMetaFilters{
			AccountIDs: accountIDs,
		})

		if err != nil {
			return nil, xerrors.Errorf("error loading pro players: %w", err)
		}
	}

	data := &SearchData{
		Heroes:                heroes,
		FollowedPlayers:       followed,
		PlayersByAccountID:    players.KeyByAccountID(),
		ProPlayersByAccountID: proPlayers.KeyByAccountID(),
	}

	return data, nil
}
