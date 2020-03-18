package gc

import (
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"
)

const (
	msgTypeFindTopSourceTVGamesResponse = protocol.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse
	msgTypeMatchesMinimalResponse       = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse
)

func IsKnownPacket(packet *gc.GCPacket) bool {
	if packet.AppId != dota2.AppID {
		return false
	}

	switch t := protocol.EDOTAGCMsg(packet.MsgType); t {
	case
		msgTypeFindTopSourceTVGamesResponse,
		msgTypeMatchesMinimalResponse:
		return true
	}

	return false
}

func UnmarshalPacket(packet *gc.GCPacket, message proto.Message) error {
	return proto.Unmarshal(packet.Body, message)
}
