package session

import (
	"context"
)

type session struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func newSession(parentCtx context.Context) *session {
	ctx, cancel := context.WithCancel(parentCtx)

	return &session{
		ctx:    ctx,
		cancel: cancel,
	}
}
