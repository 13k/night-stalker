package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewLeague(league *models.League) (*nspb.League, error) {
	pb := &nspb.League{
		Id:             uint64(league.ID),
		Name:           league.Name,
		Tier:           league.Tier,
		Region:         league.Region,
		Status:         league.Status,
		TotalPrizePool: league.TotalPrizePool,
	}

	var err error

	if pb.LastActivityAt, err = models.NullTimestampProto(league.LastActivityAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.StartAt, err = models.NullTimestampProto(league.StartAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.FinishAt, err = models.NullTimestampProto(league.FinishAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	return pb, nil
}

func NewLeagues(leagues []*models.League) ([]*nspb.League, error) {
	if len(leagues) == 0 {
		return nil, nil
	}

	views := make([]*nspb.League, len(leagues))

	var err error

	for i, league := range leagues {
		views[i], err = NewLeague(league)

		if err != nil {
			err = xerrors.Errorf("error creating League view: %w", err)
			return nil, err
		}
	}

	return views, nil
}
