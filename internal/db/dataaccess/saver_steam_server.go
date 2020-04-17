package dataaccess

import (
	"context"

	"github.com/faceit/go-steam/netutil"
	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

func (s *Saver) UpsertSteamServersAddresses(
	ctx context.Context,
	addresses []*netutil.PortAddr,
) ([]*nsm.SteamServer, error) {
	var err error

	tx, txerr := s.mq.Begin(ctx, nil)

	if txerr != nil {
		return nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	servers := make([]*nsm.SteamServer, len(addresses))

	for i, addr := range addresses {
		server := &nsm.SteamServer{
			Address: addr.String(),
		}

		q := tx.
			Q().
			Select().
			Eq(nsm.SteamServerTable.Col("address"), server.Address).
			Trace()

		if _, err = tx.M().Upsert(ctx, server, q); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error saving steam server: %w", err)
		}

		servers[i] = server
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return servers, nil
}
