package bus

import (
	"github.com/olebedev/emitter"
)

type Subscription struct {
	Topic string
	C     <-chan Message

	ev <-chan emitter.Event
}
