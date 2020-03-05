package bus

import (
	"fmt"
	"time"

	"github.com/olebedev/emitter"
	"github.com/sirupsen/logrus"

	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	defaultCap        = uint(32)
	defaultPubTimeout = 10 * time.Second
)

type Message struct {
	Topic   string
	Pattern string
	Payload interface{}
}

type Options struct {
	Cap        uint
	Log        *nslog.Logger
	PubTimeout time.Duration
}

type Bus struct {
	*emitter.Emitter

	options Options
	log     *nslog.Logger
	subs    map[<-chan Message]<-chan emitter.Event
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
		subs:    make(map[<-chan Message]<-chan emitter.Event),
	}
}

func (b *Bus) Shutdown() {
	b.log.Debug("shutdown")
	b.Emitter.Off("*")
}

func (b *Bus) Pub(message Message) {
	select {
	case <-b.Emitter.Emit(message.Topic, message.Payload):
		return
	case <-time.After(b.options.PubTimeout):
		b.log.WithFields(logrus.Fields{
			"payload": fmt.Sprintf("%T", message.Payload),
			"topic":   message.Topic,
			"timeout": b.options.PubTimeout,
		}).Warn("publish timeout")
	}
}

func (b *Bus) PubSync(message Message) {
	<-b.Emitter.Emit(message.Topic, message.Payload)
}

func (b *Bus) Sub(topic string) <-chan Message {
	b.log.WithField("topic", topic).Debug("sub")
	eventsCh := b.Emitter.On(topic)
	messagesCh := make(chan Message, b.Emitter.Cap)
	b.subs[messagesCh] = eventsCh

	go func() {
		for ev := range eventsCh {
			var payload interface{}

			if len(ev.Args) > 0 {
				payload = ev.Args[0]
			}

			messagesCh <- Message{
				Topic:   ev.Topic,
				Pattern: ev.OriginalTopic,
				Payload: payload,
			}
		}
	}()

	return messagesCh
}

func (b *Bus) Unsub(topic string, channels ...<-chan Message) {
	b.log.WithField("topic", topic).Debug("unsub")
	eventChannels := make([]<-chan emitter.Event, len(channels))

	for i, ch := range channels {
		eventChannels[i] = b.subs[ch]
	}

	b.Emitter.Off(topic, eventChannels...)
}
