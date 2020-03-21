package models

import (
	"github.com/lib/pq"
	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LiveMatchStatsModel Model = (*LiveMatchStats)(nil)

type LiveMatchStatsID uint64

// LiveMatchStats ...
type LiveMatchStats struct {
	ID                         LiveMatchStatsID  `gorm:"column:id;primary_key"`
	LiveMatchID                LiveMatchID       `gorm:"column:live_match_id;index;not null"`
	MatchID                    nspb.MatchID      `gorm:"column:match_id;index;not null"`
	ServerID                   nspb.SteamID      `gorm:"column:server_id;index;not null"`
	LeagueID                   nspb.LeagueID     `gorm:"column:league_id"`
	LeagueNodeID               nspb.LeagueNodeID `gorm:"column:league_node_id"`
	GameTimestamp              uint32            `gorm:"column:game_timestamp"`
	GameTime                   int32             `gorm:"column:game_time"`
	GameMode                   nspb.GameMode     `gorm:"column:game_mode"`
	GameState                  nspb.GameState    `gorm:"column:game_state"`
	DeltaFrame                 bool              `gorm:"column:delta_frame"`
	GraphGold                  pq.Int64Array     `gorm:"column:graph_gold"`
	GraphXP                    pq.Int64Array     `gorm:"column:graph_xp"`
	GraphKill                  pq.Int64Array     `gorm:"column:graph_kill"`
	GraphTower                 pq.Int64Array     `gorm:"column:graph_tower"`
	GraphRax                   pq.Int64Array     `gorm:"column:graph_rax"`
	SteamBroadcasterAccountIDs nssql.AccountIDs  `gorm:"column:steam_broadcaster_account_ids"`
	Timestamps

	LiveMatch *LiveMatch
	Match     *Match
	Teams     []*LiveMatchStatsTeam
	Players   []*LiveMatchStatsPlayer
	Draft     []*LiveMatchStatsPickBan
	Buildings []*LiveMatchStatsBuilding
}

func (*LiveMatchStats) TableName() string {
	return "live_match_stats"
}

func NewLiveMatchStats(liveMatch *LiveMatch, pb *protocol.CMsgDOTARealtimeGameStatsTerse) *LiveMatchStats {
	m := LiveMatchStatsDotaProto(pb)
	m.LiveMatchID = liveMatch.ID
	return m
}

func LiveMatchStatsDotaProto(pb *protocol.CMsgDOTARealtimeGameStatsTerse) *LiveMatchStats {
	return &LiveMatchStats{
		MatchID:                    nspb.MatchID(pb.GetMatch().GetMatchid()),
		ServerID:                   nspb.SteamID(pb.GetMatch().GetServerSteamId()),
		LeagueID:                   nspb.LeagueID(pb.GetMatch().GetLeagueId()),
		LeagueNodeID:               nspb.LeagueNodeID(pb.GetMatch().GetLeagueNodeId()),
		GameTimestamp:              pb.GetMatch().GetTimestamp(),
		GameTime:                   pb.GetMatch().GetGameTime(),
		GameMode:                   nspb.GameMode(pb.GetMatch().GetGameMode()),
		GameState:                  nspb.GameState(pb.GetMatch().GetGameState()),
		DeltaFrame:                 pb.GetDeltaFrame(),
		GraphGold:                  Int32Array(pb.GetGraphData().GetGraphGold()),
		SteamBroadcasterAccountIDs: nssql.NewAccountIDs(pb.GetMatch().GetSteamBroadcasterAccountIds()),
		// GraphXP:                    Int32Array(pb.GetGraphData().GetGraphXp()),
		// GraphKill:                  Int32Array(pb.GetGraphData().GetGraphKill()),
		// GraphTower:                 Int32Array(pb.GetGraphData().GetGraphTower()),
		// GraphRax:                   Int32Array(pb.GetGraphData().GetGraphRax()),
	}
}
