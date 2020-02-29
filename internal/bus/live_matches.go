package bus

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type LiveMatchesChangeMessage struct {
	Op       nspb.CollectionOp
	Matches  nscol.LiveMatches
	MatchIDs nscol.MatchIDs
}

type LiveMatchStatsChangeMessage struct {
	Op    nspb.CollectionOp
	Stats nscol.LiveMatchStats
}
