package bus

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

type GCDispatcherSendMessage struct {
	MsgType protocol.EDOTAGCMsg
	Message proto.Message
}

type GCDispatcherReceivedMessage struct {
	MsgType protocol.EDOTAGCMsg
	Message proto.Message
}
