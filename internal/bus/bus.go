package bus

import (
	"github.com/cskr/pubsub"
)

// Bus is a thin wrapper around pubsub.PubSub that publishes messages in short-lived goroutines to
// avoid deadlocks between interdependent publishers and subscribers.
type Bus struct {
	*pubsub.PubSub
}

func New(cap int) *Bus {
	return &Bus{
		PubSub: pubsub.New(cap),
	}
}

func (bus *Bus) Pub(msg interface{}, topics ...string) {
	go bus.PubSub.Pub(msg, topics...)
}

func (bus *Bus) TryPub(msg interface{}, topics ...string) {
	go bus.PubSub.TryPub(msg, topics...)
}
