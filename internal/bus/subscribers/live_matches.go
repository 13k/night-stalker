package subscribers

import (
	"context"
	"sync"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type LiveMatches struct {
	ctx                      context.Context
	bus                      *nsbus.Bus
	busSubLiveMatchesReplace *nsbus.Subscription
	mtx                      sync.RWMutex
	liveMatches              nscol.LiveMatches
}

func NewLiveMatchesSubscriber(bus *nsbus.Bus) *LiveMatches {
	return &LiveMatches{bus: bus}
}

func (s *LiveMatches) Start(ctx context.Context) {
	s.start(ctx)
}

func (s *LiveMatches) Get() nscol.LiveMatches {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.liveMatches
}

func (s *LiveMatches) Set(liveMatches nscol.LiveMatches) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.liveMatches = liveMatches
}

func (s *LiveMatches) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return len(s.liveMatches)
}

func (s *LiveMatches) Batches(batchSize int) []nscol.LiveMatches {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if len(s.liveMatches) == 0 {
		return nil
	}

	return s.liveMatches.Batches(batchSize)
}

func (s *LiveMatches) start(ctx context.Context) {
	s.ctx = ctx

	if s.busSubLiveMatchesReplace == nil {
		s.busSubLiveMatchesReplace = s.bus.Sub(nsbus.TopicLiveMatchesReplace)
	}

	go s.loop()
}

func (s *LiveMatches) stop() {
	if s.busSubLiveMatchesReplace != nil {
		s.bus.Unsub(s.busSubLiveMatchesReplace)
		s.busSubLiveMatchesReplace = nil
	}

	s.ctx = nil
}

func (s *LiveMatches) loop() {
	defer s.stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case busmsg, ok := <-s.busSubLiveMatchesReplace.C:
			if !ok {
				return
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				s.handleLiveMatchesChange(msg)
			}
		}
	}
}

func (s *LiveMatches) handleLiveMatchesChange(msg *nsbus.LiveMatchesChangeMessage) {
	if msg.Op != nspb.CollectionOp_COLLECTION_OP_REPLACE {
		return
	}

	s.Set(msg.Matches)
}
