package comm

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

const (
	processorName = "comm.chat"
)

type ChatOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Chat)(nil)

type Chat struct {
	options           ChatOptions
	ctx               context.Context
	log               *nslog.Logger
	bus               *nsbus.Bus
	busSubSteamEvents *nsbus.Subscription
}

func NewChat(options ChatOptions) *Chat {
	p := &Chat{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
	}

	p.busSubscribe()

	return p
}

func (p *Chat) ChildSpec() oversight.ChildProcessSpecification {
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

func (p *Chat) Start(ctx context.Context) error {
	p.ctx = ctx

	p.busSubscribe()

	return p.loop()
}

func (p *Chat) busSubscribe() {
	if p.busSubSteamEvents == nil {
		p.busSubSteamEvents = p.bus.Sub(nsbus.TopicSteamEvents)
	}
}

func (p *Chat) busUnsubscribe() {
	if p.busSubSteamEvents != nil {
		p.bus.Unsub(p.busSubSteamEvents)
		p.busSubSteamEvents = nil
	}
}

func (p *Chat) loop() error {
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
		case busmsg, ok := <-p.busSubSteamEvents.C:
			if !ok {
				return nil
			}

			if steammsg, ok := busmsg.Payload.(*nsbus.SteamEventMessage); ok {
				if chatmsg, ok := steammsg.Event.(*steam.ChatMsgEvent); ok {
					p.handleChatMessage(chatmsg)
				}
			}
		}
	}
}

func (p *Chat) stop() {
	p.busUnsubscribe()
	p.log.Warn("stop")
}

func (p *Chat) handleChatMessage(chatmsg *steam.ChatMsgEvent) {
	if chatmsg.IsMessage() {
		p.log.WithOFields(
			"room", chatmsg.ChatRoomId,
			"from", chatmsg.ChatterId,
			"message", chatmsg.Message,
		).Info("chat message")
	}
}
