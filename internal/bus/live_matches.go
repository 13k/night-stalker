package bus

import (
	nsproto "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

type LiveMatchesDotaMessage struct {
	Index   uint32
	Matches []*protocol.CSourceTVGameSmall
}

type LiveMatchesProtoMessage struct {
	Matches []*nsproto.LiveMatch
}
