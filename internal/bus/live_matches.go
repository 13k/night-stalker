package bus

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
)

type LiveMatchesChangeMessage struct {
	Op       nspb.CollectionOp
	Matches  nscol.LiveMatchesSlice
	MatchIDs nscol.MatchIDs
}

type LiveMatchStatsChangeMessage struct {
	Op    nspb.CollectionOp
	Stats nscol.LiveMatchStatsSlice
}
