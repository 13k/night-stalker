package models

import (
	"time"

	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var PlayerProfileCardModel Model = (*PlayerProfileCard)(nil)

type PlayerProfileCardID uint64

type PlayerProfileCard struct {
	ID               PlayerProfileCardID `gorm:"column:id;primary_key"`
	AccountID        nspb.AccountID      `gorm:"column:account_id;unique_index;not null"`
	BadgePoints      uint32              `gorm:"column:badge_points"`
	LeaderboardRank  uint32              `gorm:"column:leaderboard_rank"`
	RankTier         uint32              `gorm:"column:rank_tier"`
	RankTierScore    uint32              `gorm:"column:rank_tier_score"`
	RankTierMMRType  uint32              `gorm:"column:rank_tier_mmr_type"`
	PreviousRankTier uint32              `gorm:"column:previous_rank_tier"`
	IsPlusSubscriber bool                `gorm:"column:is_plus_subscriber"`
	PlusStartAt      *time.Time          `gorm:"column:plus_start_at"`
	EventID          uint32              `gorm:"column:event_id"`
	EventPoints      uint32              `gorm:"column:event_points"`
	Timestamps
}

func (*PlayerProfileCard) TableName() string {
	return "player_profile_cards"
}

func PlayerProfileCardProto(pb *protocol.CMsgDOTAProfileCard) *PlayerProfileCard {
	return &PlayerProfileCard{
		AccountID:        pb.GetAccountId(),
		BadgePoints:      pb.GetBadgePoints(),
		EventPoints:      pb.GetEventPoints(),
		EventID:          pb.GetEventId(),
		RankTier:         pb.GetRankTier(),
		LeaderboardRank:  pb.GetLeaderboardRank(),
		RankTierScore:    pb.GetRankTierScore(),
		PreviousRankTier: pb.GetPreviousRankTier(),
		RankTierMMRType:  pb.GetRankTierMmrType(),
		IsPlusSubscriber: pb.GetIsPlusSubscriber(),
		PlusStartAt:      NullUnixTimestamp(int64(pb.GetPlusOriginalStartDate())),
	}
}
