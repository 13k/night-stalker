package gc

import (
	"time"

	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

type recvQueueTimeoutError struct {
	Packet  *gc.GCPacket
	Timeout time.Duration
}

func (*recvQueueTimeoutError) Error() string {
	return "receive timeout"
}

type sendQueueTimeoutError struct {
	BusMessage *nsbus.GCDispatcherSendMessage
	Timeout    time.Duration
}

func (*sendQueueTimeoutError) Error() string {
	return "send timeout"
}

type recvError struct {
	MsgType protocol.EDOTAGCMsg
	Packet  *gc.GCPacket
	Err     error
}

func (*recvError) Error() string {
	return "receive error"
}

type sendError struct {
	MsgType protocol.EDOTAGCMsg
	Message proto.Message
	Err     error
}

func (*sendError) Error() string {
	return "send error"
}
