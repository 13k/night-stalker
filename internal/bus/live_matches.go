package bus

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

type LiveMatchesDotaMessage struct {
	Index   uint32
	Matches []*protocol.CSourceTVGameSmall
}

type LiveMatchesChangeMessage struct {
	Change *nspb.LiveMatchesChange
}
