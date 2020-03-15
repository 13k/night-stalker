package bus

import (
	"time"
)

type ErrPublishTimeout struct {
	Message Message
	Timeout time.Duration
}

func (*ErrPublishTimeout) Error() string {
	return "bus publish timeout"
}

type ErrSubscriptionExpired struct {
	Subscription *Subscription
}

func (*ErrSubscriptionExpired) Error() string {
	return "bus subscription expired"
}
