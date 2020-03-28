package gc

import (
	"time"

	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	d2pb "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nserr "github.com/13k/night-stalker/internal/errors"
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
	*nserr.Err
	MsgType d2pb.EDOTAGCMsg
	Packet  *gc.GCPacket
}

type sendError struct {
	*nserr.Err
	MsgType d2pb.EDOTAGCMsg
	Message proto.Message
}
