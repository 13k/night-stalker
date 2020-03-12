package session

import (
	"context"
	"time"

	"github.com/paralin/go-dota2"

	nslog "github.com/13k/night-stalker/internal/logger"
)

const (
	helloRetryCount    = 360
	helloRetryInterval = 10 * time.Second
)

type dotaGreeter struct {
	log    *nslog.Logger
	dota   *dota2.Dota2
	ctx    context.Context
	cancel context.CancelFunc
}

func newDotaGreeter(log *nslog.Logger, dota *dota2.Dota2) *dotaGreeter {
	return &dotaGreeter{
		log:  log,
		dota: dota,
	}
}

func (c *dotaGreeter) hello() error {
	c.ctx, c.cancel = context.WithCancel(context.Background())

	retryCount := 0
	t := time.NewTicker(helloRetryInterval)

	defer func() {
		c.log.Trace("hello() stop")
		t.Stop()
	}()

	c.log.Trace("hello() start")

	for {
		c.dota.SayHello()
		retryCount++

		select {
		case <-c.ctx.Done():
			return nil
		case <-t.C:
			if retryCount >= helloRetryCount {
				return NewErrDotaGCWelcomeTimeoutX(retryCount, helloRetryInterval)
			}
		}
	}
}

func (c *dotaGreeter) welcome() {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
}
