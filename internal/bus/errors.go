package bus

import (
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

type PublishTimeoutError struct {
	Message Message
	Timeout time.Duration

	message string
}

func NewPublishTimeoutErrorX(msg Message, timeout time.Duration) error {
	err := &PublishTimeoutError{
		Message: msg,
		Timeout: timeout,
	}

	return xerrors.Errorf("bus error: %w", err)
}

func (err *PublishTimeoutError) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf(
			"timeout when publishing message to topic '%s' (payload type: %T)",
			err.Message.Topic,
			err.Message.Payload,
		)
	}

	return err.message
}

type SubscriptionExpiredError struct {
	Subscription *Subscription

	message string
}

func NewSubscriptionExpiredErrorX(sub *Subscription) error {
	err := &SubscriptionExpiredError{Subscription: sub}
	return xerrors.Errorf("bus error: %w", err)
}

func (err *SubscriptionExpiredError) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf("bus subscription expired (topic '%s')", err.Subscription.Topic)
	}

	return err.message
}
