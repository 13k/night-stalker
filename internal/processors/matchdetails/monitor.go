package matchdetails

import (
	"context"
	"runtime/debug"
	"sync"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "match_details"
	batchSize     = 10

	msgTypeMatchesMinimalRequest = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest
)

type MonitorOptions struct {
	Logger          *nslog.Logger
	PoolSize        int
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Monitor)(nil)

type Monitor struct {
	options              *MonitorOptions
	ctx                  context.Context
	log                  *nslog.Logger
	db                   *gorm.DB
	workerPool           *ants.Pool
	bus                  *nsbus.Bus
	busSubLiveMatches    chan interface{}
	busSubMatchesMinimal chan interface{}
	matchesMtx           sync.RWMutex
	matches              nscol.LiveMatchesSlice
}

func NewMonitor(options *MonitorOptions) *Monitor {
	return &Monitor{
		options: options,
		log:     options.Logger.WithPackage(processorName),
	}
}

func (p *Monitor) ChildSpec() oversight.ChildProcessSpecification {
	var shutdown oversight.Shutdown

	if p.options.ShutdownTimeout > 0 {
		shutdown = oversight.Timeout(p.options.ShutdownTimeout)
	} else {
		shutdown = oversight.Infinity()
	}

	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: shutdown,
	}
}

func (p *Monitor) Start(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	p.ctx = ctx
	p.busSubLiveMatches = p.bus.Sub(nsbus.TopicLiveMatches)
	p.busSubMatchesMinimal = p.bus.Sub(nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse)

	var err error

	if p.workerPool, err = ants.NewPool(p.options.PoolSize); err != nil {
		p.log.WithError(err).Error("error starting worker pool")
		return err
	}

	return p.loop()
}

func (p *Monitor) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	tick := time.NewTicker(p.options.Interval)

	defer func() {
		tick.Stop()
		p.workerPool.Release()
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-tick.C:
			p.tick()
		case busmsg, ok := <-p.busSubLiveMatches:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.(*nsbus.LiveMatchesMessage); ok {
				p.handleLiveMatches(msg)
			}
		case busmsg := <-p.busSubMatchesMinimal:
			if dispatcherMsg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dispatcherMsg.Message.(*protocol.CMsgClientToGCMatchesMinimalResponse); ok {
					p.handleMatchesMinimalResponse(res)
				}
			}
		}
	}
}

func (p *Monitor) handleLiveMatches(msg *nsbus.LiveMatchesMessage) {
	p.log.WithFields(logrus.Fields{
		"count": len(msg.Matches),
	}).Debug("received live matches")

	p.matchesMtx.Lock()
	defer p.matchesMtx.Unlock()
	p.matches = msg.Matches
}

func (p *Monitor) tick() {
	p.matchesMtx.RLock()
	defer p.matchesMtx.RUnlock()

	if len(p.matches) == 0 {
		return
	}

	p.log.WithFields(logrus.Fields{
		"count": len(p.matches),
	}).Debug("requesting match details")

	batches := p.matches.Batches(batchSize)

	for _, batch := range batches {
		p.submitLiveMatchesBatch(batch)
	}
}

func (p *Monitor) submitLiveMatchesBatch(matches []*models.LiveMatch) {
	if err := p.ctx.Err(); err != nil {
		p.log.WithError(err).Error()
		return
	}

	err := p.workerPool.Submit(func() {
		p.work(matches)
	})

	if err != nil {
		p.log.WithError(err).Error("error creating worker")
	}
}

func (p *Monitor) work(matches []*models.LiveMatch) {
	matchIDs := make([]nspb.MatchID, len(matches))

	for i, match := range matches {
		matchIDs[i] = match.MatchID
	}

	p.requestMatchesMinimal(matchIDs...)
}

func (p *Monitor) requestMatchesMinimal(matchIDs ...nspb.MatchID) {
	req := &protocol.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}

	busmsg := &nsbus.GCDispatcherSendMessage{
		MsgType: msgTypeMatchesMinimalRequest,
		Message: req,
	}

	p.bus.Pub(busmsg, nsbus.TopicGCDispatcherSend)
}

func (p *Monitor) handleMatchesMinimalResponse(msg *protocol.CMsgClientToGCMatchesMinimalResponse) {
	if len(msg.GetMatches()) == 0 {
		return
	}

	p.log.
		WithField("count", len(msg.GetMatches())).
		Debug("received matches minimal")

	matches, err := p.saveMatches(msg.GetMatches())

	if err != nil {
		p.log.WithError(err).Error("error saving matches")
		return
	}

	matchIDs := make([]nspb.MatchID, len(matches))

	for i, match := range matches {
		matchIDs[i] = match.ID
	}

	p.busPublishMatchesFinished(matchIDs...)
}

func (p *Monitor) saveMatches(minMatches []*protocol.CMsgDOTAMatchMinimal) ([]*models.Match, error) {
	matches := make([]*models.Match, len(minMatches))

	for i, pbMatch := range minMatches {
		if p.ctx.Err() != nil {
			return nil, p.ctx.Err()
		}

		l := p.log.WithField("match_id", pbMatch.GetMatchId())
		match := models.MatchDotaProto(pbMatch)

		tx := p.db.Begin()

		result := tx.
			Where(models.Match{ID: match.ID}).
			Assign(match).
			FirstOrCreate(match)

		if err := result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Error("error upserting match")
			return nil, err
		}

		for _, pbPlayer := range pbMatch.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()
				return nil, p.ctx.Err()
			}

			matchPlayer := models.MatchPlayerDotaProto(pbPlayer)
			matchPlayer.MatchID = match.ID

			criteria := &models.MatchPlayer{
				MatchID:   matchPlayer.MatchID,
				AccountID: matchPlayer.AccountID,
			}

			result = tx.
				Where(criteria).
				Assign(matchPlayer).
				FirstOrCreate(matchPlayer)

			if err := result.Error; err != nil {
				tx.Rollback()
				l.WithError(err).Error("error upserting match player")
				return nil, err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}

		matches[i] = match
	}

	return matches, nil
}

func (p *Monitor) busPublishMatchesFinished(matchIDs ...nspb.MatchID) {
	busmsg := &nsbus.LiveMatchesFinishedMessage{MatchIDs: matchIDs}
	p.bus.Pub(busmsg, nsbus.TopicLiveMatchesFinished)
}
