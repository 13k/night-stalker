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
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

const (
	processorName = "gc.dispatcher"
	queueSize     = 32
	queueTimeout  = 10 * time.Second

	msgTypeFindTopSourceTVGamesResponse = protocol.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse
	msgTypeMatchesMinimalResponse       = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse
)

func UnmarshalPacket(packet *gc.GCPacket, message proto.Message) error {
	return proto.Unmarshal(packet.Body, message)
}

type DispatcherOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Dispatcher)(nil)

type Dispatcher struct {
	options    DispatcherOptions
	ctx        context.Context
	log        *nslog.Logger
	steam      *steam.Client
	bus        *nsbus.Bus
	busSubSend <-chan nsbus.Message
	handling   bool
	rx         chan *gc.GCPacket
	tx         chan *nsbus.GCDispatcherSendMessage
}

func NewDispatcher(options DispatcherOptions) *Dispatcher {
	proc := &Dispatcher{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
		rx:      make(chan *gc.GCPacket, queueSize),
		tx:      make(chan *nsbus.GCDispatcherSendMessage, queueSize),
	}

	proc.busSubscribe()

	return proc
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
	if err := p.setupContext(ctx); err != nil {
		return err
	}

	p.setupGCPacketHandling()

	go p.recvLoop()
	go p.sendLoop()

	return p.loop()
}

func (p *Dispatcher) busSubscribe() {
	if p.busSubSend == nil {
		p.busSubSend = p.bus.Sub(nsbus.TopicGCDispatcherSend)
	}
}

func (p *Dispatcher) busUnsubscribe() {
	if p.busSubSend != nil {
		p.bus.Unsub(nsbus.TopicGCDispatcherSend, p.busSubSend)
	}
}

func (p *Dispatcher) setupContext(ctx context.Context) error {
	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	p.ctx = ctx

	return nil
}

func (p *Dispatcher) setupGCPacketHandling() {
	if !p.handling {
		p.steam.GC.RegisterPacketHandler(p)
		p.handling = true
	}
}

func (p *Dispatcher) recvLoop() {
	defer func() {
		p.log.Debug("rx stop")
	}()

	p.log.Debug("rx start")

	for {
		select {
		case <-p.ctx.Done():
			return
		case packet, ok := <-p.rx:
			if !ok {
				return
			}

			msgType := protocol.EDOTAGCMsg(packet.MsgType)
			l := p.log.WithField("msg_type", msgType.String())

			if err := p.recv(msgType, packet); err != nil {
				l.WithError(err).Error("error receiving GC packet")
			}
		}
	}
}

func (p *Dispatcher) sendLoop() {
	defer func() {
		p.log.Debug("tx stop")
	}()

	p.log.Debug("tx start")

	for {
		select {
		case <-p.ctx.Done():
			return
		case sendmsg, ok := <-p.tx:
			if !ok {
				return
			}

			p.send(sendmsg.MsgType, sendmsg.Message)
		}
	}
}

func (p *Dispatcher) stop() {
	p.busUnsubscribe()
	close(p.rx)
	close(p.tx)
	p.rx = make(chan *gc.GCPacket, queueSize)
	p.tx = make(chan *nsbus.GCDispatcherSendMessage, queueSize)
}

func (p *Dispatcher) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	defer func() {
		p.stop()
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSubSend:
			if !ok {
				return nil
			}

			if sendmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherSendMessage); ok {
				p.enqueueTx(sendmsg)
			}
		}
	}
}

// runs in the steam client goroutine
func (p *Dispatcher) HandleGCPacket(packet *gc.GCPacket) {
	if packet.AppId != dota2.AppID {
		return
	}

	p.enqueueRx(packet)
}

func (p *Dispatcher) enqueueRx(packet *gc.GCPacket) {
	t := time.NewTicker(queueTimeout)

	defer t.Stop()

	select {
	case p.rx <- packet:
		return
	case <-t.C:
		p.log.WithFields(logrus.Fields{
			"msg_type": protocol.EDOTAGCMsg(packet.MsgType).String(),
			"wait":     queueTimeout.String(),
		}).Warn("ignored incoming packet (queue is full)")
	}
}

func (p *Dispatcher) enqueueTx(sendmsg *nsbus.GCDispatcherSendMessage) {
	t := time.NewTicker(queueTimeout)

	defer t.Stop()

	select {
	case p.tx <- sendmsg:
		return
	case <-t.C:
		p.log.WithFields(logrus.Fields{
			"msg_type": sendmsg.MsgType.String(),
			"wait":     queueTimeout.String(),
		}).Warn("ignored outgoing message (queue is full)")
	}
}

func (p *Dispatcher) recv(msgType protocol.EDOTAGCMsg, packet *gc.GCPacket) error {
	var topic string
	var message proto.Message

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

	p.log.WithField("msg_type", msgType.String()).Debug("receiving")

	if err := UnmarshalPacket(packet, message); err != nil {
		return err
	}

	p.bus.Pub(nsbus.Message{
		Topic: topic,
		Payload: &nsbus.GCDispatcherReceivedMessage{
			MsgType: msgType,
			Message: message,
		},
	})

	return nil
}

func (p *Dispatcher) send(msgType protocol.EDOTAGCMsg, message proto.Message) {
	p.log.WithField("msg_type", msgType.String()).Debug("sending")
	p.steam.GC.Write(gc.NewGCMsgProtobuf(dota2.AppID, uint32(msgType), message))
}
