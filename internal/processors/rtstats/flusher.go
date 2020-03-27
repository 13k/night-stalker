package rtstats

import (
	"context"
	"sync"
	"time"

	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nslog "github.com/13k/night-stalker/internal/logger"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type flusherOptions struct {
	Log      *nslog.Logger
	Bus      *nsbus.Bus
	Interval time.Duration
	Cap      int
}

type flusher struct {
	options *flusherOptions
	log     *nslog.Logger
	bus     *nsbus.Bus
	buf     nscol.LiveMatchStats
	mtx     *sync.Mutex
	ctx     context.Context
	values  chan *models.LiveMatchStats
	errors  chan error
}

func newFlusher(options *flusherOptions) *flusher {
	return &flusher{
		options: options,
		log:     options.Log.WithPackage("flusher"),
		bus:     options.Bus,
		mtx:     &sync.Mutex{},
	}
}

func (f *flusher) Errors() <-chan error {
	return f.errors
}

func (f *flusher) Start(ctx context.Context) {
	f.ctx = ctx
	f.values = make(chan *models.LiveMatchStats)
	f.errors = make(chan error)

	go f.loop()
}

func (f *flusher) stop(t *time.Ticker) {
	t.Stop()

	close(f.values)
	f.values = nil

	close(f.errors)
	f.errors = nil

	f.ctx = nil
	f.log.Trace("stop")
}

func (f *flusher) loop() {
	t := time.NewTicker(f.options.Interval)

	defer f.stop(t)

	f.log.Trace("start")

	for {
		select {
		case <-f.ctx.Done():
			return
		case <-t.C:
			f.Flush()
		case v, ok := <-f.values:
			if !ok {
				return
			}

			f.add(v)
		}
	}
}

func (f *flusher) Add(stats *models.LiveMatchStats) {
	f.values <- stats
}

func (f *flusher) add(stats *models.LiveMatchStats) {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	f.buf = append(f.buf, stats)

	if len(f.buf) >= f.options.Cap {
		f.safeFlush()
	}
}

func (f *flusher) Flush() {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	if len(f.buf) == 0 {
		return
	}

	f.safeFlush()
}

// TODO: it should reset flush timer
func (f *flusher) safeFlush() {
	f.log.WithField("count", len(f.buf)).Trace("flush")

	busMsg := nsbus.Message{
		Topic: nsbus.TopicLiveMatchStatsAdd,
		Payload: &nsbus.LiveMatchStatsChangeMessage{
			Op:    nspb.CollectionOp_COLLECTION_OP_ADD,
			Stats: f.buf,
		},
	}

	if err := f.bus.Pub(busMsg); err != nil {
		f.errors <- xerrors.Errorf("error publishing live match stats change: %w", err)
		return
	}

	f.buf = nil
}
