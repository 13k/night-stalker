package dataaccess

import (
	"context"

	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

func (s *Saver) UpsertMatchProto(ctx context.Context, pb *d2pb.CMsgDOTAMatchMinimal) (*nsm.Match, error) {
	tx, txerr := s.mq.Begin(ctx, nil)

	if txerr != nil {
		return nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	match := nsm.NewMatchProto(pb)

	q := tx.
		Q().
		Select().
		Eq(nsm.MatchTable.PK(), match.ID)

	if _, err := tx.M().Upsert(ctx, match, q); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error saving match: %w", err)
	}

	match.Players = make([]*nsm.MatchPlayer, len(pb.GetPlayers()))

	for i, pbPlayer := range pb.GetPlayers() {
		matchPlayer := nsm.NewMatchPlayerAssocProto(match, pbPlayer)

		q := tx.
			Q().
			Select().
			Eq(nsm.MatchPlayerTable.Col("match_id"), matchPlayer.MatchID).
			Eq(nsm.MatchPlayerTable.Col("account_id"), matchPlayer.AccountID)

		if _, err := tx.M().Upsert(ctx, matchPlayer, q); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error saving match player: %w", err)
		}

		match.Players[i] = matchPlayer
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return match, nil
}
