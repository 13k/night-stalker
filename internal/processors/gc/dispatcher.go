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
	processorName = "gc_dispatch"
	queueSize     = 4
	queueTimeout  = 10 * time.Second
)

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
	handling   bool
	bus        *nsbus.Bus
	busSubSend *nsbus.Subscription
	rx         chan *gc.GCPacket
	tx         chan *nsbus.GCDispatcherSendMessage
}

func NewDispatcher(options DispatcherOptions) *Dispatcher {
	p := &Dispatcher{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
	}

	p.busSubscribe()

	return p
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
	p.setupRxTx()
	p.busSubscribe()

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
		p.bus.Unsub(p.busSubSend)
		p.busSubSend = nil
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

func (p *Dispatcher) setupRxTx() {
	if p.rx == nil {
		p.rx = make(chan *gc.GCPacket, queueSize)
	}

	if p.tx == nil {
		p.tx = make(chan *nsbus.GCDispatcherSendMessage, queueSize)
	}
}

func (p *Dispatcher) teardownRxTx() {
	if p.rx != nil {
		close(p.rx)
		p.rx = nil
	}

	if p.tx != nil {
		close(p.tx)
		p.tx = nil
	}
}

func (p *Dispatcher) recvLoop() {
	defer func() {
		p.log.Debug("rx stop")
	}()

	p.log.Debug("rx start")

	for {
		packet, ok := <-p.rx

		if !ok {
			return
		}

		msgType := protocol.EDOTAGCMsg(packet.MsgType)

		if err := p.recv(msgType, packet); err != nil {
			p.log.
				WithField("msg_type", msgType).
				WithError(err).
				Error("error receiving message")
		}
	}
}

func (p *Dispatcher) sendLoop() {
	defer func() {
		p.log.Debug("tx stop")
	}()

	p.log.Debug("tx start")

	for {
		sendmsg, ok := <-p.tx

		if !ok {
			return
		}

		if err := p.send(sendmsg.MsgType, sendmsg.Message); err != nil {
			p.log.
				WithField("msg_type", sendmsg.MsgType).
				WithError(err).
				Error("error sending message")
		}
	}
}

func (p *Dispatcher) stop() {
	p.busUnsubscribe()
	p.teardownRxTx()
	p.log.Warn("stop")
}

func (p *Dispatcher) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	defer p.stop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSubSend.C:
			if !ok {
				return nil
			}

			if sendmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherSendMessage); ok {
				if err := p.enqueueTx(sendmsg); err != nil {
					p.handleQueueError(err)
				}
			}
		}
	}
}

// runs in the steam client goroutine
func (p *Dispatcher) HandleGCPacket(packet *gc.GCPacket) {
	if p.ctx == nil {
		return
	}

	if !IsKnownPacket(packet) {
		return
	}

	if err := p.enqueueRx(packet); err != nil {
		p.handleQueueError(err)
	}
}

func (p *Dispatcher) enqueueRx(packet *gc.GCPacket) error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	select {
	case p.rx <- packet:
		return nil
	case <-time.After(queueTimeout):
		return &recvTimeoutError{
			Packet:  packet,
			Timeout: queueTimeout,
		}
	}
}

func (p *Dispatcher) enqueueTx(sendmsg *nsbus.GCDispatcherSendMessage) error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	select {
	case p.tx <- sendmsg:
		return nil
	case <-time.After(queueTimeout):
		return &sendTimeoutError{
			BusMessage: sendmsg,
			Timeout:    queueTimeout,
		}
	}
}

func (p *Dispatcher) handleQueueError(err error) {
	switch e := err.(type) {
	case *recvTimeoutError:
		p.log.WithOFields(
			"msg_type", protocol.EDOTAGCMsg(e.Packet.MsgType),
			"timeout", e.Timeout,
		).Warn("ignored incoming packet (queue is full)")
	case *sendTimeoutError:
		p.log.WithOFields(
			"msg_type", e.BusMessage.MsgType,
			"timeout", e.Timeout,
		).Warn("ignored outgoing message (queue is full)")
	default:
		p.log.WithError(err).Error("queue error")
	}
}

func (p *Dispatcher) recv(msgType protocol.EDOTAGCMsg, packet *gc.GCPacket) error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	incoming := NewIncomingMessage(msgType)

	if incoming == nil {
		return nil
	}

	if err := incoming.UnmarshalPacket(packet); err != nil {
		return err
	}

	if err := p.busPubReceivedMessage(incoming); err != nil {
		return err
	}

	p.log.WithField("msg_type", incoming.Type).Trace("received message")

	return nil
}

func (p *Dispatcher) send(msgType protocol.EDOTAGCMsg, message proto.Message) error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	p.steam.GC.Write(gc.NewGCMsgProtobuf(dota2.AppID, uint32(msgType), message))

	p.log.WithField("msg_type", msgType).Trace("sent message")

	return nil
}

func (p *Dispatcher) busPubReceivedMessage(incoming *IncomingMessage) error {
	return p.bus.Pub(nsbus.Message{
		Topic: incoming.BusTopic,
		Payload: &nsbus.GCDispatcherReceivedMessage{
			MsgType: incoming.Type,
			Message: incoming.Message,
		},
	})
}
