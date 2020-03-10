package bus

import (
	"time"

	"github.com/olebedev/emitter"

	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	defaultCap        = uint(32)
	defaultPubTimeout = 10 * time.Second
)

type Options struct {
	Cap        uint
	Log        *nslog.Logger
	PubTimeout time.Duration
}

type Bus struct {
	*emitter.Emitter

	options Options
	log     *nslog.Logger
}

func New(options Options) *Bus {
	if options.Cap == 0 {
		options.Cap = defaultCap
	}

	if options.PubTimeout == 0 {
		options.PubTimeout = defaultPubTimeout
	}

	return &Bus{
		Emitter: emitter.New(options.Cap),
		options: options,
		log:     options.Log.WithPackage("bus"),
	}
}

func (b *Bus) Shutdown() {
	b.log.Debug("shutdown")
	b.Emitter.Off("*")
}

func (b *Bus) Pub(message Message) error {
	select {
	case <-b.Emitter.Emit(message.Topic, message.Payload):
		return nil
	case <-time.After(b.options.PubTimeout):
		return NewPublishTimeoutErrorX(message, b.options.PubTimeout)
	}
}

func (b *Bus) PubSync(message Message) {
	<-b.Emitter.Emit(message.Topic, message.Payload)
}

func (b *Bus) Sub(topic string) *Subscription {
	b.log.WithField("topic", topic).Trace("sub")

	events := b.Emitter.On(topic)
	messages := make(chan Message, b.Emitter.Cap)

	go func() {
		for ev := range events {
			var payload interface{}

			if len(ev.Args) > 0 {
				payload = ev.Args[0]
			}

			messages <- Message{
				Topic:   ev.Topic,
				Pattern: ev.OriginalTopic,
				Payload: payload,
			}
		}
	}()

	return &Subscription{
		Topic: topic,
		C:     messages,
		ev:    events,
	}
}

func (b *Bus) Unsub(subs ...*Subscription) {
	for _, sub := range subs {
		b.log.WithField("topic", sub.Topic).Trace("unsub")
		b.Emitter.Off(sub.Topic, sub.ev)
	}
}
