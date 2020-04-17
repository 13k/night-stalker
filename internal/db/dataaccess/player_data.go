package dataaccess

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type PlayerData struct {
	AccountID      nspb.AccountID
	FollowedPlayer *nsm.FollowedPlayer
	Player         *nsm.Player
}

func NewPlayerData(
	followedPlayer *nsm.FollowedPlayer,
	player *nsm.Player,
) *PlayerData {
	data := &PlayerData{
		FollowedPlayer: followedPlayer,
		Player:         player,
	}

	if data.AccountID == 0 {
		data.AccountID = followedPlayer.AccountID
	}

	if data.AccountID == 0 && player != nil {
		data.AccountID = player.AccountID
	}

	return data
}

func (d *PlayerData) Validate() error {
	if d.AccountID == 0 {
		return xerrors.Errorf("invalid PlayerData: %w", ErrMissingAccountID)
	}

	if d.FollowedPlayer != nil && d.FollowedPlayer.AccountID != d.AccountID {
		return xerrors.Errorf("invalid PlayerData: %w", ErrInconsistentAccountIDs)
	}

	if d.Player != nil && d.Player.AccountID != d.AccountID {
		return xerrors.Errorf("invalid PlayerData: %w", ErrInconsistentAccountIDs)
	}

	return nil
}

type PlayersData map[nspb.AccountID]*PlayerData

func NewPlayersData(playersData ...*PlayerData) (PlayersData, error) {
	data := make(PlayersData)

	for _, d := range playersData {
		if err := d.Validate(); err != nil {
			return nil, err
		}

		data[d.AccountID] = d
	}

	return data, nil
}
