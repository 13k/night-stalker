package views

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type MatchData struct {
	MatchID        nspb.MatchID
	Match          *models.Match
	LiveMatch      *models.LiveMatch
	LiveMatchStats []*models.LiveMatchStats
	PlayersData    MatchPlayersData
}

func NewMatchData(
	match *models.Match,
	liveMatch *models.LiveMatch,
	stats []*models.LiveMatchStats,
	playersData MatchPlayersData,
) *MatchData {
	data := &MatchData{
		Match:          match,
		LiveMatch:      liveMatch,
		LiveMatchStats: stats,
		PlayersData:    playersData,
	}

	if match != nil {
		data.MatchID = match.ID
	}

	if data.MatchID == 0 && liveMatch != nil {
		data.MatchID = liveMatch.MatchID
	}

	if data.MatchID == 0 {
		for _, s := range stats {
			data.MatchID = s.MatchID

			if data.MatchID != 0 {
				break
			}
		}
	}

	return data
}

func (d *MatchData) Validate() error {
	if d.MatchID == 0 {
		return xerrors.Errorf("invalid MatchData: %w", ErrMissingMatchID)
	}

	if d.Match != nil && d.Match.ID != d.MatchID {
		return xerrors.Errorf("invalid MatchData: %w", ErrInconsistentMatchIDs)
	}

	if d.LiveMatch != nil && d.LiveMatch.MatchID != d.MatchID {
		return xerrors.Errorf("invalid MatchData: %w", ErrInconsistentMatchIDs)
	}

	for _, s := range d.LiveMatchStats {
		if s.MatchID != d.MatchID {
			return xerrors.Errorf("invalid MatchData: %w", ErrInconsistentMatchIDs)
		}
	}

	return d.PlayersData.Validate()
}

func (d *MatchData) AccountIDs() nscol.AccountIDs {
	if d.PlayersData == nil {
		return nil
	}

	accountIDs := make(nscol.AccountIDs, 0, len(d.PlayersData))

	for accountID := range d.PlayersData {
		accountIDs = append(accountIDs, accountID)
	}

	return accountIDs
}

type MatchesData map[nspb.MatchID]*MatchData

func NewMatchesData(data ...*MatchData) (MatchesData, error) {
	matchesData := make(MatchesData)

	for _, datum := range data {
		if err := datum.Validate(); err != nil {
			return nil, err
		}

		matchesData[datum.MatchID] = datum
	}

	return matchesData, nil
}

func (d MatchesData) MatchIDs() nscol.MatchIDs {
	matchIDs := make(nscol.MatchIDs, 0, len(d))

	for matchID := range d {
		matchIDs = append(matchIDs, matchID)
	}

	return matchIDs
}

func (d MatchesData) AccountIDs() nscol.AccountIDs {
	var accountIDs nscol.AccountIDs

	for _, matchData := range d {
		accountIDs = accountIDs.AddUnique(matchData.AccountIDs()...)
	}

	return accountIDs
}

type MatchPlayerData struct {
	MatchID               nspb.MatchID
	AccountID             nspb.AccountID
	MatchPlayer           *models.MatchPlayer
	LiveMatchPlayer       *models.LiveMatchPlayer
	LiveMatchStatsPlayers nscol.LiveMatchStatsPlayers
}

func NewMatchPlayerData(
	matchPlayer *models.MatchPlayer,
	livePlayer *models.LiveMatchPlayer,
	statsPlayers nscol.LiveMatchStatsPlayers,
) *MatchPlayerData {
	data := &MatchPlayerData{
		MatchPlayer:           matchPlayer,
		LiveMatchPlayer:       livePlayer,
		LiveMatchStatsPlayers: statsPlayers,
	}

	if matchPlayer != nil {
		data.MatchID = matchPlayer.MatchID
		data.AccountID = matchPlayer.AccountID
	}

	if livePlayer != nil {
		if data.MatchID == 0 {
			data.MatchID = livePlayer.MatchID
		}

		if data.AccountID == 0 {
			data.AccountID = livePlayer.AccountID
		}
	}

	if data.MatchID == 0 || data.AccountID == 0 {
		for _, p := range statsPlayers {
			if data.MatchID == 0 {
				data.MatchID = p.MatchID
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

	if d.MatchPlayer != nil && d.MatchPlayer.MatchID != d.MatchID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentMatchIDs)
	}

	if d.LiveMatchPlayer != nil && d.LiveMatchPlayer.MatchID != d.MatchID {
		return xerrors.Errorf("invalid MatchPlayerData: %w", ErrInconsistentMatchIDs)
	}

	for _, s := range d.LiveMatchStatsPlayers {
		if s.MatchID != d.MatchID {
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
			err = xerrors.Errorf("invalid MatchPlayersData: %w", err)
			return nil, err
		}

		matchPlayersData[d.AccountID] = d
	}

	return matchPlayersData, nil
}

func (d MatchPlayersData) Validate() error {
	for _, pd := range d {
		if err := pd.Validate(); err != nil {
			err = xerrors.Errorf("invalid MatchPlayersData: %w", err)
			return err
		}
	}

	return nil
}
