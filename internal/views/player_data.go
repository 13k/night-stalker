package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type PlayerData struct {
	AccountID      nspb.AccountID
	FollowedPlayer *models.FollowedPlayer
	Player         *models.Player
	ProPlayer      *models.ProPlayer
}

func NewPlayerData(
	followedPlayer *models.FollowedPlayer,
	player *models.Player,
	proPlayer *models.ProPlayer,
) *PlayerData {
	data := &PlayerData{
		FollowedPlayer: followedPlayer,
		Player:         player,
		ProPlayer:      proPlayer,
	}

	if data.AccountID == 0 {
		data.AccountID = followedPlayer.AccountID
	}

	if data.AccountID == 0 {
		data.AccountID = player.AccountID
	}

	if data.AccountID == 0 {
		data.AccountID = proPlayer.AccountID
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

	if d.ProPlayer != nil && d.ProPlayer.AccountID != d.AccountID {
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
