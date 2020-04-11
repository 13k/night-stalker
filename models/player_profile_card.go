package models

import (
	"database/sql"

	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var PlayerProfileCardTable = NewTable("player_profile_cards")

type PlayerProfileCard struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID        nspb.AccountID `db:"account_id"`
	BadgePoints      uint32         `db:"badge_points"`
	LeaderboardRank  uint32         `db:"leaderboard_rank"`
	RankTier         uint32         `db:"rank_tier"`
	RankTierScore    uint32         `db:"rank_tier_score"`
	RankTierMMRType  uint32         `db:"rank_tier_mmr_type"`
	PreviousRankTier uint32         `db:"previous_rank_tier"`
	IsPlusSubscriber bool           `db:"is_plus_subscriber"`
	PlusStartAt      sql.NullTime   `db:"plus_start_at"`
	EventID          uint32         `db:"event_id"`
	EventPoints      uint32         `db:"event_points"`

	Timestamps
	SoftDelete
}

func NewPlayerProfileCardProto(pb *d2pb.CMsgDOTAProfileCard) *PlayerProfileCard {
	return &PlayerProfileCard{
		AccountID:        nspb.AccountID(pb.GetAccountId()),
		BadgePoints:      pb.GetBadgePoints(),
		EventPoints:      pb.GetEventPoints(),
		EventID:          pb.GetEventId(),
		RankTier:         pb.GetRankTier(),
		LeaderboardRank:  pb.GetLeaderboardRank(),
		RankTierScore:    pb.GetRankTierScore(),
		PreviousRankTier: pb.GetPreviousRankTier(),
		RankTierMMRType:  pb.GetRankTierMmrType(),
		IsPlusSubscriber: pb.GetIsPlusSubscriber(),
		PlusStartAt:      nssql.NullTimeUnix(int64(pb.GetPlusOriginalStartDate())),
	}
}
