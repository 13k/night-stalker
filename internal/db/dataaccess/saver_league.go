package dataaccess

import (
	"context"

	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

func (s *Saver) UpsertLeagueProto(ctx context.Context, pb *d2pb.CMsgDOTALeagueInfo) (*nsm.League, bool, error) {
	league := nsm.NewLeagueProto(pb)

	q := s.mq.
		Q().
		Select().
		Eq(nsm.LeagueTable.PK(), league.ID)

	created, err := s.mq.M().Upsert(ctx, league, q)

	if err != nil {
		return nil, false, xerrors.Errorf("error saving league: %w", err)
	}

	return league, created, nil
}
