package gc

import (
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-dota2"
	d2pb "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"
)

const (
	msgTypeFindTopSourceTVGamesResponse = d2pb.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse
	msgTypeMatchesMinimalResponse       = d2pb.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse
)

func IsKnownPacket(packet *gc.GCPacket) bool {
	if packet.AppId != dota2.AppID {
		return false
	}

	switch t := d2pb.EDOTAGCMsg(packet.MsgType); t {
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
