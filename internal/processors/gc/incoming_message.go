package gc

import (
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

type IncomingMessage struct {
	Type     protocol.EDOTAGCMsg
	Message  proto.Message
	BusTopic string
}

func NewIncomingMessage(msgType protocol.EDOTAGCMsg) *IncomingMessage {
	var message proto.Message
	var topic string

	switch msgType {
	case msgTypeMatchesMinimalResponse:
		message = &protocol.CMsgClientToGCMatchesMinimalResponse{}
		topic = nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse
	case msgTypeFindTopSourceTVGamesResponse:
		message = &protocol.CMsgGCToClientFindTopSourceTVGamesResponse{}
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
