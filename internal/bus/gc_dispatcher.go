package bus

import (
	"github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"
)

type GCDispatcherSendMessage struct {
	MsgType protocol.EDOTAGCMsg
	Message proto.Message
}

type GCDispatcherReceivedMessage struct {
	MsgType protocol.EDOTAGCMsg
	Message proto.Message
}
