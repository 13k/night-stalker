package bus

import (
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
)

type LiveMatchesMessage struct {
	Matches nscol.LiveMatchesSlice
}

type LiveMatchesFinishedMessage struct {
	MatchIDs []nspb.MatchID
}

type LiveMatchesChangeMessage struct {
	Change *nspb.LiveMatchesChange
}
