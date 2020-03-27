package bus

import (
	d2pb "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"
)

type GCDispatcherSendMessage struct {
	MsgType d2pb.EDOTAGCMsg
	Message proto.Message
}

type GCDispatcherReceivedMessage struct {
	MsgType d2pb.EDOTAGCMsg
	Message proto.Message
}
