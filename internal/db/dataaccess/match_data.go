package dataaccess

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type MatchData struct {
	MatchID        nspb.MatchID
	Match          *nsm.Match
	LiveMatch      *nsm.LiveMatch
	LiveMatchStats nscol.LiveMatchStats
	PlayersData    MatchPlayersData
}

func NewMatchData(
	match *nsm.Match,
	liveMatch *nsm.LiveMatch,
	stats nscol.LiveMatchStats,
	playersData MatchPlayersData,
) *MatchData {
	data := &MatchData{
		Match:          match,
		LiveMatch:      liveMatch,
		LiveMatchStats: stats,
		PlayersData:    playersData,
	}

	if match != nil {
		data.MatchID = nspb.MatchID(match.ID)
	}

	if data.MatchID == 0 && liveMatch != nil {
		data.MatchID = nspb.MatchID(liveMatch.MatchID)
	}

	if data.MatchID == 0 {
		for _, s := range stats {
			data.MatchID = nspb.MatchID(s.MatchID)

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

	if d.Match != nil && nspb.MatchID(d.Match.ID) != d.MatchID {
		return xerrors.Errorf("invalid MatchData: %w", ErrInconsistentMatchIDs)
	}

	if d.LiveMatch != nil && nspb.MatchID(d.LiveMatch.MatchID) != d.MatchID {
		return xerrors.Errorf("invalid MatchData: %w", ErrInconsistentMatchIDs)
	}

	for _, s := range d.LiveMatchStats {
		if nspb.MatchID(s.MatchID) != d.MatchID {
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
