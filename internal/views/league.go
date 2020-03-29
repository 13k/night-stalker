package views

import (
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
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

	if pb.LastActivityAt, err = nssql.NullTimeProto(league.LastActivityAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.StartAt, err = nssql.NullTimeProto(league.StartAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	if pb.FinishAt, err = nssql.NullTimeProto(league.FinishAt); err != nil {
		err = xerrors.Errorf("error converting Time to protobuf Timestamp: %w", err)
		return nil, err
	}

	return pb, nil
}
