package processors

import (
	"time"

	"cirello.io/oversight"
)

type Processor interface {
	ChildSpec(stimeout time.Duration) oversight.ChildProcessSpecification
}
