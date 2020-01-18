package gc

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

const (
	processorName = "gc.dispatcher"

	msgTypeFindTopSourceTVGamesResponse = protocol.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse
	msgTypeMatchesMinimalResponse       = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse
)

func UnmarshalPacket(packet *gc.GCPacket, message proto.Message) error {
	return proto.Unmarshal(packet.Body, message)
}

type DispatcherOptions struct {
	Logger          *nslog.Logger
	BufferSize      int
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Dispatcher)(nil)

type Dispatcher struct {
	options    *DispatcherOptions
	ctx        context.Context
	log        *nslog.Logger
	steam      *steam.Client
	queue      chan *gc.GCPacket
	bus        *nsbus.Bus
	busSubSend chan interface{}
}

func NewDispatcher(options *DispatcherOptions) *Dispatcher {
	return &Dispatcher{
		options: options,
		log:     options.Logger.WithPackage(processorName),
		queue:   make(chan *gc.GCPacket, options.BufferSize),
	}
}

func (p *Dispatcher) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if p.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(p.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: shutdown,
	}
}

func (p *Dispatcher) Start(ctx context.Context) error {
	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	p.ctx = ctx
	p.busSubSend = p.bus.Sub(nsbus.TopicGCDispatcherSend)

	p.steam.GC.RegisterPacketHandler(p)

	return p.loop()
}

// runs in the steam client goroutine
func (p *Dispatcher) HandleGCPacket(packet *gc.GCPacket) {
	if packet.AppId != dota2.AppID {
		return
	}

	msgType := protocol.EDOTAGCMsg(packet.MsgType)

	select {
	case p.queue <- packet:
	default:
		p.log.WithField("msg_type", msgType.String()).Warn("ignored packet (not handling packets or too busy)")
	}
}

func (p *Dispatcher) write(msgType protocol.EDOTAGCMsg, message proto.Message) {
	p.steam.GC.Write(gc.NewGCMsgProtobuf(dota2.AppID, uint32(msgType), message))
}

func (p *Dispatcher) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	defer func() {
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case packet := <-p.queue:
			p.handlePacket(packet)
		case busmsg := <-p.busSubSend:
			if sendmsg, ok := busmsg.(*nsbus.GCDispatcherSendMessage); ok {
				p.write(sendmsg.MsgType, sendmsg.Message)
			}
		}
	}
}

func (p *Dispatcher) handlePacket(packet *gc.GCPacket) {
	msgType := protocol.EDOTAGCMsg(packet.MsgType)
	busmsg := &nsbus.GCDispatcherReceivedMessage{MsgType: msgType}

	var topic string

	switch msgType {
	case msgTypeMatchesMinimalResponse:
		busmsg.Message = &protocol.CMsgClientToGCMatchesMinimalResponse{}
		topic = nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse
	case msgTypeFindTopSourceTVGamesResponse:
		busmsg.Message = &protocol.CMsgGCToClientFindTopSourceTVGamesResponse{}
		topic = nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse
	default:
		return
	}

	p.log.WithField("msg_type", msgType.String()).Debug("handling")

	if err := UnmarshalPacket(packet, busmsg.Message); err != nil {
		p.log.WithError(err).Error("error unmarshaling GC packet")
		return
	}

	p.bus.TryPub(busmsg, topic)
}
