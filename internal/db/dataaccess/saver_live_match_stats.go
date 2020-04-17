package dataaccess

import (
	"context"

	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsm "github.com/13k/night-stalker/models"
)

func (s *Saver) CreateLiveMatchStatsAssocProto(
	ctx context.Context,
	liveMatch *nsm.LiveMatch,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse,
) (*nsm.LiveMatchStats, error) {
	tx, txerr := s.mq.Begin(ctx, nil)

	if txerr != nil {
		return nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	stats := nsm.NewLiveMatchStatsAssocProto(liveMatch, pb)

	if err := tx.M().Create(ctx, stats); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return nil, xerrors.Errorf("error creating live match stats: %w", err)
	}

	lteams := len(pb.GetTeams())
	stats.Teams = make([]*nsm.LiveMatchStatsTeam, lteams)
	stats.Players = make([]*nsm.LiveMatchStatsPlayer, 0, lteams*5) // estimate each team has 5 players

	for it, t := range pb.GetTeams() {
		team := nsm.NewLiveMatchStatsTeamAssocProto(stats, t)

		if err := tx.M().Create(ctx, team); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error creating live match stats team: %w", err)
		}

		stats.Teams[it] = team

		for _, p := range t.GetPlayers() {
			player := nsm.NewLiveMatchStatsPlayerAssocProto(stats, p)

			if err := tx.M().Create(ctx, player); err != nil {
				if txerr := tx.Rollback(); txerr != nil {
					return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
				}

				return nil, xerrors.Errorf("error creating live match stats team player: %w", err)
			}

			stats.Players = append(stats.Players, player)
		}
	}

	stats.Draft = make([]*nsm.LiveMatchStatsPickBan, 0, len(pb.GetMatch().GetPicks())+len(pb.GetMatch().GetBans()))

	for _, pbd := range pb.GetMatch().GetPicks() {
		pickBan := nsm.LiveMatchStatsPickBanAssocProto(stats, false, pbd)

		if err := tx.M().Create(ctx, pickBan); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error creating live match stats pickban: %w", err)
		}

		stats.Draft = append(stats.Draft, pickBan)
	}

	for _, pb := range pb.GetMatch().GetBans() {
		pickBan := nsm.LiveMatchStatsPickBanAssocProto(stats, true, pb)

		if err := tx.M().Create(ctx, pickBan); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error creating live match stats pickban: %w", err)
		}

		stats.Draft = append(stats.Draft, pickBan)
	}

	stats.Buildings = make([]*nsm.LiveMatchStatsBuilding, len(pb.GetBuildings()))

	for i, b := range pb.GetBuildings() {
		building := nsm.NewLiveMatchStatsBuildingAssocProto(stats, b)

		if err := tx.M().Create(ctx, building); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return nil, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return nil, xerrors.Errorf("error creating live match stats building: %w", err)
		}

		stats.Buildings[i] = building
	}

	if txerr := tx.Commit(); txerr != nil {
		return nil, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return stats, nil
}
