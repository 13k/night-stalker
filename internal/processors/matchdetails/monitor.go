package matchdetails

import (
	"context"
	"runtime/debug"
	"sync"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "match_details"
	batchSize     = 25

	msgTypeMatchesMinimalRequest = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest
)

type MonitorOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	PoolSize        int
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Monitor)(nil)

type Monitor struct {
	options               MonitorOptions
	ctx                   context.Context
	log                   *nslog.Logger
	db                    *gorm.DB
	bus                   *nsbus.Bus
	busLiveMatchesReplace *nsbus.Subscription
	busMatchesMinimalResp *nsbus.Subscription
	liveMatchesMtx        sync.RWMutex
	liveMatches           nscol.LiveMatches
}

func NewMonitor(options MonitorOptions) *Monitor {
	p := &Monitor{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
	}

	p.busSubscribe()

	return p
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
	if err := p.setupContext(ctx); err != nil {
		return err
	}

	p.busSubscribe()

	return p.loop()
}

func (p *Monitor) busSubscribe() {
	if p.busLiveMatchesReplace == nil {
		p.busLiveMatchesReplace = p.bus.Sub(nsbus.TopicLiveMatchesReplace)
	}

	if p.busMatchesMinimalResp == nil {
		p.busMatchesMinimalResp = p.bus.Sub(nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse)
	}
}

func (p *Monitor) busUnsubscribe() {
	if p.busLiveMatchesReplace != nil {
		p.bus.Unsub(p.busLiveMatchesReplace)
		p.busLiveMatchesReplace = nil
	}

	if p.busMatchesMinimalResp != nil {
		p.bus.Unsub(p.busMatchesMinimalResp)
		p.busMatchesMinimalResp = nil
	}
}

func (p *Monitor) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	p.ctx = ctx

	return nil
}

func (p *Monitor) setLiveMatches(liveMatches nscol.LiveMatches) {
	p.liveMatchesMtx.Lock()
	defer p.liveMatchesMtx.Unlock()
	p.liveMatches = liveMatches
}

func (p *Monitor) getLiveMatchesCount() int {
	p.liveMatchesMtx.RLock()
	defer p.liveMatchesMtx.RUnlock()
	return len(p.liveMatches)
}

func (p *Monitor) getLiveMatchesBatches() []nscol.LiveMatches {
	p.liveMatchesMtx.RLock()
	defer p.liveMatchesMtx.RUnlock()

	if len(p.liveMatches) == 0 {
		return nil
	}

	return p.liveMatches.Batches(batchSize)
}

func (p *Monitor) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	t := time.NewTicker(p.options.Interval)

	defer p.stop(t)

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-t.C:
			p.tick()
		case busmsg, ok := <-p.busLiveMatchesReplace.C:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				p.handleLiveMatchesChange(msg)
			}
		case busmsg, ok := <-p.busMatchesMinimalResp.C:
			if !ok {
				return nil
			}

			if dspmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dspmsg.Message.(*protocol.CMsgClientToGCMatchesMinimalResponse); ok {
					p.handleMatchesMinimalResponse(res)
				}
			}
		}
	}
}

func (p *Monitor) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.log.Warn("stop")
}

func (p *Monitor) tick() {
	batches := p.getLiveMatchesBatches()

	if len(batches) == 0 {
		return
	}

	p.log.WithFields(logrus.Fields{
		"batch_size": batchSize,
		"batches":    len(batches),
		"total":      p.getLiveMatchesCount(),
	}).Debug("requesting matches details")

	for _, batch := range batches {
		if err := p.busPubRequestMatchesMinimal(batch.MatchIDs()); err != nil {
			p.log.WithError(err).Error()
		}
	}
}

func (p *Monitor) handleLiveMatchesChange(msg *nsbus.LiveMatchesChangeMessage) {
	if msg.Op != nspb.CollectionOp_COLLECTION_OP_REPLACE {
		p.log.WithField("op", msg.Op.String()).Warn("ignored live matches change message")
		return
	}

	p.log.WithFields(logrus.Fields{
		"count": len(msg.Matches),
	}).Debug("received live matches")

	p.setLiveMatches(msg.Matches)
}

func (p *Monitor) handleMatchesMinimalResponse(msg *protocol.CMsgClientToGCMatchesMinimalResponse) {
	if len(msg.GetMatches()) == 0 {
		return
	}

	p.log.
		WithField("count", len(msg.GetMatches())).
		Debug("received matches details")

	matches, err := p.saveMatches(msg.GetMatches())

	if err != nil {
		p.log.WithError(err).Error("error saving matches")
		return
	}

	if err := p.busPubLiveMatchesRemove(matches.MatchIDs()); err != nil {
		p.log.WithError(err).Error()
	}
}

func (p *Monitor) saveMatches(minMatches []*protocol.CMsgDOTAMatchMinimal) (nscol.Matches, error) {
	matches := make(nscol.Matches, len(minMatches))

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

func (p *Monitor) busPubRequestMatchesMinimal(matchIDs nscol.MatchIDs) error {
	req := &protocol.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}

	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicGCDispatcherSend,
		Payload: &nsbus.GCDispatcherSendMessage{
			MsgType: msgTypeMatchesMinimalRequest,
			Message: req,
		},
	})
}

func (p *Monitor) busPubLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesRemove,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:       nspb.CollectionOp_COLLECTION_OP_REMOVE,
			MatchIDs: matchIDs,
		},
	})
}
