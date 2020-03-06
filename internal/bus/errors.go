package bus

import (
	"fmt"
	"time"
)

type PublishTimeoutError struct {
	Message Message
	Timeout time.Duration

	message string
}

func (err *PublishTimeoutError) Error() string {
	if err.message == "" {
		err.message = fmt.Sprintf(
			"timeout when publishing message to topic %q (payload type: %T)",
			err.Message.Topic,
			err.Message.Payload,
		)
	}

	return err.message
}
