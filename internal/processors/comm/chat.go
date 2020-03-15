package comm

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrt "github.com/13k/night-stalker/internal/runtime"
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
	options        ChatOptions
	ctx            context.Context
	log            *nslog.Logger
	steam          *steam.Client
	bus            *nsbus.Bus
	busSteamEvents *nsbus.Subscription
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
		Start:    p.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (p *Chat) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Chat) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	p.busSubscribe()

	return p.loop()
}

func (p *Chat) stop() {
	p.busUnsubscribe()
	p.log.Warn("stop")
}

func (p *Chat) busSubscribe() {
	if p.busSteamEvents == nil {
		p.busSteamEvents = p.bus.Sub(nsbus.TopicSteamEvents)
	}
}

func (p *Chat) busUnsubscribe() {
	if p.busSteamEvents != nil {
		p.bus.Unsub(p.busSteamEvents)
		p.busSteamEvents = nil
	}
}

func (p *Chat) setupContext(ctx context.Context) error {
	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextSteamClient)
	}

	p.ctx = ctx

	return nil
}

func (p *Chat) loop() error {
	defer p.stop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSteamEvents.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSteamEvents,
				})
			}

			if steammsg, ok := busmsg.Payload.(*nsbus.SteamEventMessage); ok {
				if chatmsg, ok := steammsg.Event.(*steam.ChatMsgEvent); ok {
					p.handleChatMessage(chatmsg)
				}
			}
		}
	}
}

func (p *Chat) handleChatMessage(chatmsg *steam.ChatMsgEvent) {
	if !chatmsg.IsMessage() {
		return
	}

	p.handleMessage(chatmsg)
}

func (p *Chat) handleMessage(chatmsg *steam.ChatMsgEvent) {
	p.log.WithOFields(
		"room", chatmsg.ChatRoomId,
		"from", chatmsg.ChatterId,
		"message", chatmsg.Message,
	).Info("chat message")

	if chatmsg.ChatterId != 76561197982474165 {
		return
	}

	switch chatmsg.Message {
	case "!disconnect":
		p.steam.Disconnect()
	}
}

func (p *Chat) handleError(err error) {
	p.log.WithError(err).Error("chat error")
}
