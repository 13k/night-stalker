package matchdetails

import (
	"context"
	"fmt"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsbussub "github.com/13k/night-stalker/internal/bus/subscribers"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nserr "github.com/13k/night-stalker/internal/errors"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "match_details"
	batchSize     = 25

	msgTypeMatchesMinimalRequest = d2pb.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest
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
	options                  MonitorOptions
	ctx                      context.Context
	log                      *nslog.Logger
	db                       *gorm.DB
	bus                      *nsbus.Bus
	dota                     *nsdota2.Client
	liveMatches              *nsbussub.LiveMatches
	busSubMatchesMinimalResp *nsbus.Subscription
}

func NewMonitor(options MonitorOptions) *Monitor {
	return &Monitor{
		options:     options,
		log:         options.Log.WithPackage(processorName),
		bus:         options.Bus,
		liveMatches: nsbussub.NewLiveMatchesSubscriber(options.Bus),
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
		Start:    p.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (p *Monitor) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Monitor) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	p.busSubscribe()
	p.liveMatches.Start(p.ctx)

	return p.loop()
}

func (p *Monitor) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Monitor) busSubscribe() {
	if p.busSubMatchesMinimalResp == nil {
		p.busSubMatchesMinimalResp = p.bus.Sub(nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse)
	}
}

func (p *Monitor) busUnsubscribe() {
	if p.busSubMatchesMinimalResp != nil {
		p.bus.Unsub(p.busSubMatchesMinimalResp)
		p.busSubMatchesMinimalResp = nil
	}
}

func (p *Monitor) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaClient)
	}

	p.ctx = ctx

	return nil
}

func (p *Monitor) loop() error {
	t := time.NewTicker(p.options.Interval)

	defer p.stop(t)

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-t.C:
			p.tick()
		case busmsg, ok := <-p.busSubMatchesMinimalResp.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubMatchesMinimalResp,
				})
			}

			if dspmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dspmsg.Message.(*d2pb.CMsgClientToGCMatchesMinimalResponse); ok {
					p.handleMatchesMinimalResponse(res)
				}
			}
		}
	}
}

func (p *Monitor) tick() {
	if !p.dota.Session.IsReady() {
		return
	}

	batches := p.liveMatches.Batches(batchSize)

	if len(batches) == 0 {
		return
	}

	p.log.WithOFields(
		"batch_size", batchSize,
		"batch_count", len(batches),
		"total", p.liveMatches.Len(),
	).Debug("requesting matches details")

	for _, batch := range batches {
		if err := p.busPubRequestMatchesMinimal(batch.MatchIDs()); err != nil {
			p.handleError(xerrors.Errorf("error publishing request match details: %w", err))
		}
	}
}

func (p *Monitor) handleMatchesMinimalResponse(msg *d2pb.CMsgClientToGCMatchesMinimalResponse) {
	if len(msg.GetMatches()) == 0 {
		return
	}

	p.log.
		WithField("count", len(msg.GetMatches())).
		Debug("received matches details")

	matches, err := p.saveMatches(msg.GetMatches())

	if err != nil {
		p.handleError(xerrors.Errorf("error saving matches: %w", err))
		return
	}

	if err := p.busPubLiveMatchesRemove(matches.MatchIDs()); err != nil {
		p.handleError(xerrors.Errorf("error publishing live matches change: %w", err))
		return
	}
}

func (p *Monitor) saveMatches(minMatches []*d2pb.CMsgDOTAMatchMinimal) (nscol.Matches, error) {
	matches := make(nscol.Matches, len(minMatches))

	for i, pbMatch := range minMatches {
		if p.ctx.Err() != nil {
			return nil, &errMatchSave{
				MatchID: pbMatch.GetMatchId(),
				Err:     nserr.Wrap("error saving match", p.ctx.Err()),
			}
		}

		match := models.MatchDotaProto(pbMatch)

		tx := p.db.Begin()

		result := tx.
			Where(models.Match{ID: match.ID}).
			Assign(match).
			FirstOrCreate(match)

		if err := result.Error; err != nil {
			tx.Rollback()

			return nil, &errMatchSave{
				MatchID: pbMatch.GetMatchId(),
				Err:     nserr.Wrap("error saving match", err),
			}
		}

		for _, pbPlayer := range pbMatch.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()

				return nil, &errMatchSave{
					MatchID: pbMatch.GetMatchId(),
					Err:     nserr.Wrap("error saving match", p.ctx.Err()),
				}
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

				return nil, &errMatchSave{
					MatchID: pbMatch.GetMatchId(),
					Err:     nserr.Wrap("error saving match", err),
				}
			}
		}

		if err := tx.Commit().Error; err != nil {
			return nil, &errMatchSave{
				MatchID: pbMatch.GetMatchId(),
				Err:     nserr.Wrap("error saving match", err),
			}
		}

		matches[i] = match
	}

	return matches, nil
}

func (p *Monitor) busPubRequestMatchesMinimal(matchIDs nscol.MatchIDs) error {
	if !p.dota.Session.IsReady() {
		return nil
	}

	req := &d2pb.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs.ToUint64s(),
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

func (p *Monitor) handleError(err error) {
	msg := fmt.Sprintf("%s error", processorName)
	l := p.log

	if e := (&errMatchSave{}); xerrors.As(err, &e) {
		msg = "error saving match"
		l = l.WithField("match_id", e.MatchID)
	}

	l.WithError(err).Error(msg)
}
