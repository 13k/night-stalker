package processors

import (
	"cirello.io/oversight"
)

type Processor interface {
	ChildSpec() oversight.ChildProcessSpecification
}
