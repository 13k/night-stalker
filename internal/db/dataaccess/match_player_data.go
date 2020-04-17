package dataaccess

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type MatchPlayerData struct {
	MatchID               nspb.MatchID
	AccountID             nspb.AccountID
	MatchPlayer           *nsm.MatchPlayer
	LiveMatchPlayer       *nsm.LiveMatchPlayer
	LiveMatchStatsPlayers nscol.LiveMatchStatsPlayers
}

func NewMatchPlayerData(
	matchPlayer *nsm.MatchPlayer,
	livePlayer *nsm.LiveMatchPlayer,
	statsPlayers nscol.LiveMatchStatsPlayers,
) *MatchPlayerData {
	data := &MatchPlayerData{
		MatchPlayer:           matchPlayer,
		LiveMatchPlayer:       livePlayer,
		LiveMatchStatsPlayers: statsPlayers,
	}

	if matchPlayer != nil {
		data.MatchID = nspb.MatchID(matchPlayer.MatchID)
		data.AccountID = matchPlayer.AccountID
	}

	if livePlayer != nil {
		if data.MatchID == 0 {
			data.MatchID = nspb.MatchID(livePlayer.MatchID)
		}

		if data.AccountID == 0 {
			data.AccountID = livePlayer.AccountID
		}
	}

	if data.MatchID == 0 || data.AccountID == 0 {
		for _, p := range statsPlayers {
			if data.MatchID == 0 {
				data.MatchID = nspb.MatchID(p.MatchID)
			}

			if data.AccountID == 0 {
				data.AccountID = p.AccountID
			}

			if data.MatchID != 0 && data.AccountID != 0 {
				break
			}
		}
	}

	return data
}

func (d *MatchPlayerData) Validate() error {
	if d.MatchID == 0 {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrMissingMatchID)
	}

	if d.MatchPlayer != nil && nspb.MatchID(d.MatchPlayer.MatchID) != d.MatchID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentMatchIDs)
	}

	if d.LiveMatchPlayer != nil && nspb.MatchID(d.LiveMatchPlayer.MatchID) != d.MatchID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentMatchIDs)
	}

	for _, s := range d.LiveMatchStatsPlayers {
		if nspb.MatchID(s.MatchID) != d.MatchID {
			return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentMatchIDs)
		}
	}

	if d.AccountID == 0 {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrMissingAccountID)
	}

	if d.MatchPlayer != nil && d.MatchPlayer.AccountID != d.AccountID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentAccountIDs)
	}

	if d.LiveMatchPlayer != nil && d.LiveMatchPlayer.AccountID != d.AccountID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentAccountIDs)
	}

	for _, s := range d.LiveMatchStatsPlayers {
		if s.AccountID != d.AccountID {
			return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentAccountIDs)
		}
	}

	return nil
}

type MatchPlayersData map[nspb.AccountID]*MatchPlayerData

func NewMatchPlayersData(data ...*MatchPlayerData) (MatchPlayersData, error) {
	matchPlayersData := make(MatchPlayersData)

	for _, d := range data {
		if err := d.Validate(); err != nil {
			return nil, xerrors.Errorf("invalid MatchPlayersData: %w", err)
		}

		matchPlayersData[d.AccountID] = d
	}

	return matchPlayersData, nil
}

func (d MatchPlayersData) Validate() error {
	for _, pd := range d {
		if err := pd.Validate(); err != nil {
			return xerrors.Errorf("invalid MatchPlayersData: %w", err)
		}
	}

	return nil
}
