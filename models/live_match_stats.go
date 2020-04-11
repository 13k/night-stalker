package models

import (
	"github.com/lib/pq"
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LiveMatchStatsTable = NewTable("live_match_stats")

type LiveMatchStats struct {
	ID `db:"id" goqu:"defaultifempty"`

	ServerID                   nspb.SteamID      `db:"server_id"`
	LeagueNodeID               nspb.LeagueNodeID `db:"league_node_id"`
	GameTimestamp              uint32            `db:"game_timestamp"`
	GameTime                   int32             `db:"game_time"`
	GameMode                   nspb.GameMode     `db:"game_mode"`
	GameState                  nspb.GameState    `db:"game_state"`
	DeltaFrame                 bool              `db:"delta_frame"`
	GraphGold                  pq.Int64Array     `db:"graph_gold"`
	GraphXP                    pq.Int64Array     `db:"graph_xp"`
	GraphKill                  pq.Int64Array     `db:"graph_kill"`
	GraphTower                 pq.Int64Array     `db:"graph_tower"`
	GraphRax                   pq.Int64Array     `db:"graph_rax"`
	SteamBroadcasterAccountIDs nssql.AccountIDs  `db:"steam_broadcaster_account_ids"`

	LiveMatchID ID `db:"live_match_id"`
	MatchID     ID `db:"match_id"`
	LeagueID    ID `db:"league_id"`

	Timestamps
	SoftDelete

	LiveMatch *LiveMatch                `db:"-" model:"belongs_to"`
	Match     *Match                    `db:"-" model:"belongs_to"`
	League    *League                   `db:"-" model:"belongs_to"`
	Teams     []*LiveMatchStatsTeam     `db:"-" model:"has_many"`
	Players   []*LiveMatchStatsPlayer   `db:"-" model:"has_many"`
	Draft     []*LiveMatchStatsPickBan  `db:"-" model:"has_many"`
	Buildings []*LiveMatchStatsBuilding `db:"-" model:"has_many"`
}

func NewLiveMatchStatsAssocProto(
	liveMatch *LiveMatch,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse,
) *LiveMatchStats {
	m := NewLiveMatchStatsProto(pb)
	m.LiveMatch = liveMatch
	m.LiveMatchID = liveMatch.ID
	m.MatchID = liveMatch.MatchID
	return m
}

func NewLiveMatchStatsProto(pb *d2pb.CMsgDOTARealtimeGameStatsTerse) *LiveMatchStats {
	return &LiveMatchStats{
		MatchID:                    ID(pb.GetMatch().GetMatchid()),
		ServerID:                   nspb.SteamID(pb.GetMatch().GetServerSteamId()),
		LeagueID:                   ID(pb.GetMatch().GetLeagueId()),
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
