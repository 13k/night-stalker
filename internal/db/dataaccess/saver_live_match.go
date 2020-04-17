package dataaccess

import (
	"context"

	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

func (s *Saver) UpsertLiveMatchProto(ctx context.Context, pb *d2pb.CSourceTVGameSmall) (*nsm.LiveMatch, error) {
	tx, txerr := s.mq.Begin(ctx, nil)

	if txerr != nil {
		return nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	liveMatch := nsm.NewLiveMatchProto(pb)

	q := tx.
		Q().
		Select().
		Eq(nsm.LiveMatchTable.Col("match_id"), liveMatch.MatchID).
		Trace()

	if _, err := tx.M().Upsert(ctx, liveMatch, q); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error saving live match: %w", err)
	}

	liveMatch.Players = make([]*nsm.LiveMatchPlayer, len(pb.GetPlayers()))

	for i, gamePlayer := range pb.GetPlayers() {
		livePlayer := nsm.NewLiveMatchPlayerAssocProto(liveMatch, gamePlayer)

		q := tx.
			Q().
			Select().
			Eq(nsm.LiveMatchPlayerTable.Col("live_match_id"), livePlayer.LiveMatchID).
			Eq(nsm.LiveMatchPlayerTable.Col("account_id"), livePlayer.AccountID).
			Trace()

		if _, err := tx.M().Upsert(ctx, livePlayer, q); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error saving live match player: %w", err)
		}

		liveMatch.Players[i] = livePlayer
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return liveMatch, nil
}
