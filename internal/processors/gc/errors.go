package gc

import (
	"time"

	gc "github.com/faceit/go-steam/protocol/gamecoordinator"

	nsbus "github.com/13k/night-stalker/internal/bus"
)

type recvTimeoutError struct {
	Packet  *gc.GCPacket
	Timeout time.Duration
}

func (*recvTimeoutError) Error() string {
	return "receive timeout"
}

type sendTimeoutError struct {
	BusMessage *nsbus.GCDispatcherSendMessage
	Timeout    time.Duration
}

func (*sendTimeoutError) Error() string {
	return "send timeout"
}
