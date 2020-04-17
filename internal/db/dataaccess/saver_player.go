package dataaccess

import (
	"context"
	"errors"

	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

var (
	ErrPlayerAlreadyFollowed = errors.New("player already followed")
)

type FollowPlayerOptions struct {
	Update bool
}

func (s *Saver) FollowPlayer(
	ctx context.Context,
	player *nsm.FollowedPlayer,
	options *FollowPlayerOptions,
) (bool, error) {
	if player.AccountID == 0 {
		return false, xerrors.New("cannot follow player without account ID")
	}

	record := &nsm.FollowedPlayer{}

	q := s.mq.
		Q().
		Select().
		Eq(nsm.FollowedPlayerTable.Col("account_id"), player.AccountID).
		Trace()

	exists, err := s.mq.M().Find(ctx, record, q)

	if err != nil {
		return false, xerrors.Errorf("error finding followed player: %w", err)
	}

	if exists {
		if !options.Update {
			return false, xerrors.Errorf("error following player: %w", ErrPlayerAlreadyFollowed)
		}

		if record.AssignPartial(player) {
			if err := s.mq.M().Update(ctx, player); err != nil {
				return false, xerrors.Errorf("error updating followed player: %w", err)
			}

			player.Assign(record)
		}

		return false, nil
	}

	if err := s.mq.M().Create(ctx, player); err != nil {
		return false, xerrors.Errorf("error creating followed player: %w", err)
	}

	return true, nil
}
