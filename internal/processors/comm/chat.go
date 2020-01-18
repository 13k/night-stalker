package comm

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
)

const (
	processorName = "comm.chat"
)

type ChatOptions struct {
	Logger          *nslog.Logger
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Chat)(nil)

type Chat struct {
	options           *ChatOptions
	ctx               context.Context
	log               *nslog.Logger
	bus               *nsbus.Bus
	busSubSteamEvents chan interface{}
}

func NewChat(options *ChatOptions) *Chat {
	return &Chat{
		options: options,
		log:     options.Logger.WithPackage(processorName),
	}
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
	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	p.ctx = ctx
	p.busSubSteamEvents = p.bus.Sub(nsbus.TopicSteamEvents)

	return p.loop()
}

func (p *Chat) loop() error {
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
		case busmsg, ok := <-p.busSubSteamEvents:
			if !ok {
				return nil
			}

			if steammsg, ok := busmsg.(*nsbus.SteamEventMessage); ok {
				if chatmsg, ok := steammsg.Event.(*steam.ChatMsgEvent); ok {
					p.handleChatMessage(chatmsg)
				}
			}
		}
	}
}

func (p *Chat) handleChatMessage(chatmsg *steam.ChatMsgEvent) {
	if chatmsg.IsMessage() {
		p.log.WithFields(logrus.Fields{
			"from":    chatmsg.ChatterId,
			"room":    chatmsg.ChatRoomId,
			"message": chatmsg.Message,
		}).Info("chat message")
	}
}
