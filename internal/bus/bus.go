package bus

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/olebedev/emitter"
	"golang.org/x/xerrors"

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
		return xerrors.Errorf("bus error: %w", &ErrPublishTimeout{
			Message: message,
			Timeout: b.options.PubTimeout,
		})
	}
}

func (b *Bus) PubSync(message Message) {
	<-b.Emitter.Emit(message.Topic, message.Payload)
}

func (b *Bus) Sub(topic string) *Subscription {
	b.log.WithOFields(
		"topic", topic,
		"caller", getCaller(),
	).Trace("sub")

	events := b.Emitter.On(topic)
	messages := make(chan Message, b.Emitter.Cap)

	go func() {
		for ev := range events {
			var payload interface{}

			if len(ev.Args) > 0 {
				payload = ev.Args[0]
			}

			messages <- Message{
				Topic:   ev.OriginalTopic,
				Pattern: ev.Topic,
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
	caller := getCaller()

	for _, sub := range subs {
		b.log.WithOFields(
			"topic", sub.Topic,
			"caller", caller,
		).Trace("unsub")

		b.Emitter.Off(sub.Topic, sub.ev)
	}
}

func getCaller() string {
	var framePtrs [3]uintptr

	runtime.Callers(2, framePtrs[:])

	frames := runtime.CallersFrames(framePtrs[:])

	if _, ok := frames.Next(); !ok {
		return ""
	}

	fr, ok := frames.Next()

	if !ok {
		return ""
	}

	file := strings.TrimPrefix(fr.File, "/home/k/Projects/dota2/night-stalker/")

	return fmt.Sprintf("%s:%d", file, fr.Line)
}
