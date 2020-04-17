package dataaccess

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
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

type ImportPlayerData struct {
	AccountID       nspb.AccountID
	SteamID         nspb.SteamID
	Label           string
	Name            string
	PersonaName     string
	AvatarURL       string
	AvatarMediumURL string
	AvatarFullURL   string
	ProfileURL      string
	CountryCode     string
	TeamID          nspb.TeamID
	IsLocked        bool
	LockedUntil     sql.NullTime
	FantasyRole     nspb.FantasyRole
}

type ImportPlayerResult struct {
	FollowedPlayer *nsm.FollowedPlayer
	Player         *nsm.Player
	ProPlayer      *nsm.ProPlayer
	Created        bool
}

func (s *Saver) ImportPlayer(ctx context.Context, data *ImportPlayerData) (*ImportPlayerResult, error) {
	tx, txerr := s.mq.Begin(ctx, nil)

	if txerr != nil {
		return nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	followed := &nsm.FollowedPlayer{
		AccountID: data.AccountID,
		Label:     data.Name,
	}

	created, err := s.FollowPlayer(ctx, followed, &FollowPlayerOptions{
		Update: false,
	})

	if err != nil && !xerrors.Is(err, ErrPlayerAlreadyFollowed) {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error saving followed player: %w", err)
	}

	player := &nsm.Player{
		AccountID:       data.AccountID,
		SteamID:         data.SteamID,
		Name:            data.Name,
		PersonaName:     data.PersonaName,
		AvatarURL:       data.AvatarURL,
		AvatarMediumURL: data.AvatarMediumURL,
		AvatarFullURL:   data.AvatarFullURL,
		ProfileURL:      data.ProfileURL,
		CountryCode:     data.CountryCode,
	}

	q := tx.
		Q().
		Select().
		Eq(nsm.PlayerTable.Col("account_id"), player.AccountID)

	if _, err = tx.M().Upsert(ctx, player, q); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error saving player: %w", err)
	}

	pro := &nsm.ProPlayer{
		AccountID:   data.AccountID,
		TeamID:      nsm.ID(data.TeamID),
		IsLocked:    data.IsLocked,
		FantasyRole: data.FantasyRole,
		LockedUntil: data.LockedUntil,
	}

	q = tx.
		Q().
		Select().
		Eq(nsm.ProPlayerTable.Col("account_id"), pro.AccountID)

	if _, err = tx.M().Upsert(ctx, pro, q); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error saving pro player: %w", err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	r := &ImportPlayerResult{
		FollowedPlayer: followed,
		Player:         player,
		ProPlayer:      pro,
		Created:        created,
	}

	return r, nil
}
