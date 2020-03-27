package gc

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	gc "github.com/faceit/go-steam/protocol/gamecoordinator"
	protov1 "github.com/golang/protobuf/proto" //nolint: staticcheck
	"github.com/paralin/go-dota2"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	nssteam "github.com/13k/night-stalker/internal/steam"
)

const (
	processorName = "gc.dispatch"
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
	steam      *nssteam.Client
	dota       *nsdota2.Client
	handling   map[*steam.GameCoordinator]bool
	bus        *nsbus.Bus
	rx         chan *gc.GCPacket
	tx         chan *nsbus.GCDispatcherSendMessage
	busSubSend *nsbus.Subscription
}

func NewDispatcher(options DispatcherOptions) *Dispatcher {
	return &Dispatcher{
		options:  options,
		log:      options.Log.WithPackage(processorName),
		bus:      options.Bus,
		handling: map[*steam.GameCoordinator]bool{},
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
		Start:    p.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (p *Dispatcher) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Dispatcher) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	p.setupGCPacketHandling()
	p.setupRxTx()
	p.busSubscribe()

	go p.recvLoop()
	go p.sendLoop()

	return p.loop()
}

func (p *Dispatcher) stop() {
	p.busUnsubscribe()
	p.teardownRxTx()
	p.ctx = nil
	p.log.Warn("stop")
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
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextSteamClient)
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaClient)
	}

	p.ctx = ctx

	return nil
}

func (p *Dispatcher) setupGCPacketHandling() {
	if !p.handling[p.steam.GC] {
		p.steam.GC.RegisterPacketHandler(p)
		p.handling[p.steam.GC] = true
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

		msgType := d2pb.EDOTAGCMsg(packet.MsgType)

		if err := p.recv(msgType, packet); err != nil {
			p.handleError(xerrors.Errorf("error receiving packet: %w", err))
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
			p.handleError(xerrors.Errorf("error sending message: %w", err))
		}
	}
}

func (p *Dispatcher) loop() error {
	defer p.stop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSubSend.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubSend,
				})
			}

			if sendmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherSendMessage); ok {
				if err := p.enqueueTx(sendmsg); err != nil {
					p.handleError(xerrors.Errorf("error enqueuing tx: %w", err))
				}
			}
		}
	}
}

// HandleGCPacket implements the steam.GCPacketHandler interface.
//
// It runs in the steam client goroutine.
//
// It's not possible to de-register this handler once registered, so the owner must be long lived.
func (p *Dispatcher) HandleGCPacket(packet *gc.GCPacket) {
	if p.ctx == nil {
		p.log.Warn("received packet while stopped")
		return
	}

	if !IsKnownPacket(packet) {
		return
	}

	if err := p.enqueueRx(packet); err != nil {
		p.handleError(xerrors.Errorf("error enqueuing rx: %w", err))
	}
}

func (p *Dispatcher) enqueueRx(packet *gc.GCPacket) error {
	if p.ctx.Err() != nil {
		return xerrors.Errorf("error enqueuing rx: %w", p.ctx.Err())
	}

	select {
	case p.rx <- packet:
		return nil
	case <-time.After(queueTimeout):
		return xerrors.Errorf("queue timeout: %w", &recvQueueTimeoutError{
			Packet:  packet,
			Timeout: queueTimeout,
		})
	}
}

func (p *Dispatcher) enqueueTx(sendmsg *nsbus.GCDispatcherSendMessage) error {
	if p.ctx.Err() != nil {
		return xerrors.Errorf("error enqueuing tx: %w", p.ctx.Err())
	}

	if !p.dota.Session.IsReady() {
		p.log.WithField("msg_type", sendmsg.MsgType).Warn("tried to enqueue tx while disconnected")
		return nil
	}

	select {
	case p.tx <- sendmsg:
		return nil
	case <-time.After(queueTimeout):
		return xerrors.Errorf("queue timeout: %w", &sendQueueTimeoutError{
			BusMessage: sendmsg,
			Timeout:    queueTimeout,
		})
	}
}

func (p *Dispatcher) recv(msgType d2pb.EDOTAGCMsg, packet *gc.GCPacket) error {
	if p.ctx.Err() != nil {
		return xerrors.Errorf("error receiving message: %w", &recvError{
			MsgType: msgType,
			Packet:  packet,
			Err:     p.ctx.Err(),
		})
	}

	incoming := NewIncomingMessage(msgType)

	if incoming == nil {
		return nil
	}

	if err := incoming.UnmarshalPacket(packet); err != nil {
		return xerrors.Errorf("error unmarshaling packet: %w", &recvError{
			MsgType: msgType,
			Packet:  packet,
			Err:     err,
		})
	}

	if err := p.busPubReceivedMessage(incoming); err != nil {
		return xerrors.Errorf("error publishing bus message: %w", &recvError{
			MsgType: msgType,
			Packet:  packet,
			Err:     err,
		})
	}

	p.log.WithField("msg_type", incoming.Type).Trace("received message")

	return nil
}

func (p *Dispatcher) send(msgType d2pb.EDOTAGCMsg, message proto.Message) error {
	if p.ctx.Err() != nil {
		return xerrors.Errorf("error sending message: %w", &sendError{
			MsgType: msgType,
			Message: message,
			Err:     p.ctx.Err(),
		})
	}

	if !p.dota.Session.IsReady() {
		p.log.WithField("msg_type", msgType).Warn("tried to send while disconnected")
		return nil
	}

	messageV1 := protov1.MessageV1(message)
	p.steam.GC.Write(gc.NewGCMsgProtobuf(dota2.AppID, uint32(msgType), messageV1))

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

func (p *Dispatcher) handleError(err error) {
	if e := (&recvQueueTimeoutError{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"msg_type", d2pb.EDOTAGCMsg(e.Packet.MsgType),
			"timeout", e.Timeout,
		).Warn("ignored incoming packet (queue is full)")
	} else if e := (&sendQueueTimeoutError{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"msg_type", e.BusMessage.MsgType,
			"timeout", e.Timeout,
		).Warn("ignored outgoing message (queue is full)")
	} else if e := (&recvError{}); xerrors.As(err, &e) {
		p.log.
			WithField("msg_type", e.MsgType).
			WithError(e.Err).
			Error("error receiving message")
	} else if e := (&sendError{}); xerrors.As(err, &e) {
		p.log.
			WithField("msg_type", e.MsgType).
			WithError(e.Err).
			Error("error sending message")
	} else {
		p.log.WithError(err).Error("dispatcher error")
	}

	p.log.Errorx(err)
}
