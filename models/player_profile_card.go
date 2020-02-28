package models

import (
	"time"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

var PlayerProfileCardModel Model = (*PlayerProfileCard)(nil)

type PlayerProfileCardID uint64

// PlayerProfileCard ...
type PlayerProfileCard struct {
	ID                     PlayerProfileCardID `gorm:"column:id;primary_key"`
	AccountID              nspb.AccountID      `gorm:"column:account_id;unique_index;not null"`
	BackgroundDefIndex     uint32              `gorm:"column:background_def_index"`
	BadgePoints            uint32              `gorm:"column:badge_points"`
	EventPoints            uint32              `gorm:"column:event_points"`
	EventID                uint32              `gorm:"column:event_id"`
	RankTier               uint32              `gorm:"column:rank_tier"`
	LeaderboardRank        uint32              `gorm:"column:leaderboard_rank"`
	RankTierScore          uint32              `gorm:"column:rank_tier_score"`
	PreviousRankTier       uint32              `gorm:"column:previous_rank_tier"`
	RankTierMmrType        uint32              `gorm:"column:rank_tier_mmr_type"`
	RankTierCore           uint32              `gorm:"column:rank_tier_core"`
	RankTierCoreScore      uint32              `gorm:"column:rank_tier_core_score"`
	LeaderboardRankCore    uint32              `gorm:"column:leaderboard_rank_core"`
	RankTierSupport        uint32              `gorm:"column:rank_tier_support"`
	RankTierSupportScore   uint32              `gorm:"column:rank_tier_support_score"`
	LeaderboardRankSupport uint32              `gorm:"column:leaderboard_rank_support"`
	IsPlusSubscriber       bool                `gorm:"column:is_plus_subscriber"`
	PlusOriginalStartDate  *time.Time          `gorm:"column:plus_original_start_date"`
	Timestamps
}

func (*PlayerProfileCard) TableName() string {
	return "player_profile_cards"
}

func PlayerProfileCardProto(pb *protocol.CMsgDOTAProfileCard) *PlayerProfileCard {
	return &PlayerProfileCard{
		AccountID:              pb.GetAccountId(),
		BackgroundDefIndex:     pb.GetBackgroundDefIndex(),
		BadgePoints:            pb.GetBadgePoints(),
		EventPoints:            pb.GetEventPoints(),
		EventID:                pb.GetEventId(),
		RankTier:               pb.GetRankTier(),
		LeaderboardRank:        pb.GetLeaderboardRank(),
		RankTierScore:          pb.GetRankTierScore(),
		PreviousRankTier:       pb.GetPreviousRankTier(),
		RankTierMmrType:        pb.GetRankTierMmrType(),
		RankTierCore:           pb.GetRankTierCore(),
		RankTierCoreScore:      pb.GetRankTierCoreScore(),
		LeaderboardRankCore:    pb.GetLeaderboardRankCore(),
		RankTierSupport:        pb.GetRankTierSupport(),
		RankTierSupportScore:   pb.GetRankTierSupportScore(),
		LeaderboardRankSupport: pb.GetLeaderboardRankSupport(),
		IsPlusSubscriber:       pb.GetIsPlusSubscriber(),
		PlusOriginalStartDate:  NullUnixTimestamp(int64(pb.GetPlusOriginalStartDate())),
	}
}
