package gc

import (
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	d2pb "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

type IncomingMessage struct {
	Type     d2pb.EDOTAGCMsg
	Message  proto.Message
	BusTopic string
}

func NewIncomingMessage(msgType d2pb.EDOTAGCMsg) *IncomingMessage {
	var message proto.Message
	var topic string

	switch msgType {
	case msgTypeMatchesMinimalResponse:
		message = &d2pb.CMsgClientToGCMatchesMinimalResponse{}
		topic = nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse
	case msgTypeFindTopSourceTVGamesResponse:
		message = &d2pb.CMsgGCToClientFindTopSourceTVGamesResponse{}
		topic = nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse
	default:
		return nil
	}

	return &IncomingMessage{
		Type:     msgType,
		Message:  message,
		BusTopic: topic,
	}
}

func (m *IncomingMessage) UnmarshalPacket(packet *gc.GCPacket) error {
	return UnmarshalPacket(packet, m.Message)
}
