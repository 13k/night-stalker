package runtime

import (
	"bytes"
	"fmt"

	"github.com/go-stack/stack"

	nslog "github.com/13k/night-stalker/internal/logger"
)

func FormatMultilineCallStack(trace stack.CallStack) []byte {
	buf := &bytes.Buffer{}

	for _, call := range trace {
		fnName := fmt.Sprintf("%n", call)
		fmt.Fprintf(buf, "%-40s %+v\n", fnName, call)
	}

	return buf.Bytes()
}

func LogPanic(log *nslog.Logger, v interface{}, skip int) {
	callStack := stack.Trace().TrimRuntime()

	if skip > 0 && skip < len(callStack) {
		callStack = callStack.TrimBelow(callStack[skip])
	}

	log.WithField("panic", v).Error("recovered panic")
	log.Errorf("stack trace:\n%s", FormatMultilineCallStack(callStack))
}

func RecoverError(log *nslog.Logger, err *error) {
	if v := recover(); v != nil {
		LogPanic(log, v, 3)

		if e, ok := v.(error); ok {
			*err = e
		} else {
			panic(v)
		}
	}
}
