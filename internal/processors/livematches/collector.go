package livematches

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	nsrt "github.com/13k/night-stalker/internal/runtime"
)

const (
	processorName = "live_matches"
)

type CollectorOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Collector)(nil)

type Collector struct {
	ctx                  context.Context
	options              CollectorOptions
	matches              *nscol.LiveMatchesContainer
	log                  *nslog.Logger
	db                   *gorm.DB
	rds                  *nsrds.Redis
	bus                  *nsbus.Bus
	busLiveMatchesAll    *nsbus.Subscription
	busLiveMatchStatsAll *nsbus.Subscription
}

func NewCollector(options CollectorOptions) *Collector {
	p := &Collector{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
	}

	p.busSubscribe()

	return p
}

func (p *Collector) ChildSpec() oversight.ChildProcessSpecification {
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

func (p *Collector) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Collector) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	if err := p.seedLiveMatches(); err != nil {
		return xerrors.Errorf("error seeding live matches: %w", err)
	}

	p.busSubscribe()

	return p.loop()
}

func (p *Collector) stop() {
	p.busUnsubscribe()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Collector) busSubscribe() {
	if p.busLiveMatchesAll == nil {
		p.busLiveMatchesAll = p.bus.Sub(nsbus.TopicPatternLiveMatchesAll)
	}

	if p.busLiveMatchStatsAll == nil {
		p.busLiveMatchStatsAll = p.bus.Sub(nsbus.TopicPatternLiveMatchStatsAll)
	}
}

func (p *Collector) busUnsubscribe() {
	if p.busLiveMatchesAll != nil {
		p.bus.Unsub(p.busLiveMatchesAll)
		p.busLiveMatchesAll = nil
	}

	if p.busLiveMatchStatsAll != nil {
		p.bus.Unsub(p.busLiveMatchStatsAll)
		p.busLiveMatchStatsAll = nil
	}
}

func (p *Collector) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	if p.rds = nsctx.GetRedis(ctx); p.rds == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextRedis)
	}

	p.ctx = ctx

	return nil
}

func (p *Collector) seedLiveMatches() error {
	if p.matches != nil {
		return nil
	}

	matchIDs, err := p.rds.LiveMatchIDs()

	if err != nil {
		return xerrors.Errorf("error loading live matches ids: %w", err)
	}

	if len(matchIDs) == 0 {
		p.matches = nscol.NewLiveMatchesContainer()
		return nil
	}

	liveMatches, err := p.loadLiveMatches(matchIDs)

	if err != nil {
		return xerrors.Errorf("error loading live matches: %w", err)
	}

	p.matches = nscol.NewLiveMatchesContainer(liveMatches...)

	p.log.WithField("count", len(liveMatches)).Trace("seeded matches")

	return nil
}

func (p *Collector) loadLiveMatches(matchIDs nscol.MatchIDs) (nscol.LiveMatches, error) {
	var matches nscol.LiveMatches

	err := p.db.
		Where("match_id IN (?)", matchIDs).
		Order("sort_score DESC").
		Find(&matches).
		Error

	if err != nil {
		return nil, xerrors.Errorf("error loading live matches: %w", err)
	}

	return matches, nil
}

func (p *Collector) loop() error {
	defer p.stop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busLiveMatchesAll.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busLiveMatchesAll,
				})
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				p.handleLiveMatchesChange(msg)
			}
		case busmsg, ok := <-p.busLiveMatchStatsAll.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busLiveMatchStatsAll,
				})
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchStatsChangeMessage); ok {
				p.handleLiveMatchStatsChange(msg)
			}
		}
	}
}

func (p *Collector) handleLiveMatchesChange(msg *nsbus.LiveMatchesChangeMessage) {
	var err error

	switch msg.Op {
	case nspb.CollectionOp_COLLECTION_OP_ADD, nspb.CollectionOp_COLLECTION_OP_UPDATE:
		err = p.add(msg.Matches)
	case nspb.CollectionOp_COLLECTION_OP_REMOVE:
		err = p.remove(msg.MatchIDs)
	default:
		return
	}

	if err != nil {
		p.handleError(xerrors.Errorf("error handling live matches change: %w", err))
	}
}

func (p *Collector) add(matches nscol.LiveMatches) error {
	if len(matches) == 0 {
		return nil
	}

	if err := p.rds.AddLiveMatches(matches); err != nil {
		return xerrors.Errorf("error adding live matches to redis: %w", err)
	}

	beforeLen := p.matches.Len()
	change := p.matches.Add(matches...)
	afterLen := p.matches.Len()

	if len(change) > 0 {
		p.log.WithOFields(
			"before", beforeLen,
			"after", afterLen,
			"change", afterLen-beforeLen,
		).Debug("matches added")

		if err := p.notifyLiveMatchesAdd(change); err != nil {
			return xerrors.Errorf("error notifying live matches change: %w", err)
		}
	}

	return nil
}

func (p *Collector) remove(matchIDs nscol.MatchIDs) error {
	if len(matchIDs) == 0 {
		return nil
	}

	if err := p.rds.RemoveLiveMatches(matchIDs); err != nil {
		return xerrors.Errorf("error removing live matches from redis: %w", err)
	}

	beforeLen := p.matches.Len()
	change := p.matches.Remove(matchIDs...)
	afterLen := p.matches.Len()

	if len(change) > 0 {
		p.log.WithOFields(
			"before", beforeLen,
			"after", afterLen,
			"change", afterLen-beforeLen,
		).Debug("matches removed")

		if err := p.notifyLiveMatchesRemove(change); err != nil {
			return xerrors.Errorf("error notifying live matches change: %w", err)
		}
	}

	return nil
}

func (p *Collector) handleLiveMatchStatsChange(msg *nsbus.LiveMatchStatsChangeMessage) {
	var err error

	switch msg.Op {
	case nspb.CollectionOp_COLLECTION_OP_ADD:
		err = p.addStats(msg.Stats)
	default:
		return
	}

	if err != nil {
		p.handleError(xerrors.Errorf("error handling live match stats change: %w", err))
	}
}

func (p *Collector) addStats(stats nscol.LiveMatchStats) error {
	if err := p.rds.PubLiveMatchStatsAdd(stats); err != nil {
		return xerrors.Errorf("error publishing match stats update: %w", err)
	}

	p.log.WithField("count", len(stats)).Debug("stats added")

	return nil
}

func (p *Collector) notifyLiveMatchesAdd(liveMatches nscol.LiveMatches) error {
	if err := p.busPubAllMatches(); err != nil {
		return xerrors.Errorf("error publishing to bus: %w", err)
	}

	if err := p.rds.PubLiveMatchesAdd(liveMatches); err != nil {
		return xerrors.Errorf("error publishing live matches change: %w", err)
	}

	return nil
}

func (p *Collector) notifyLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	if err := p.busPubAllMatches(); err != nil {
		return xerrors.Errorf("error publishing to bus: %w", err)
	}

	if err := p.rds.PubLiveMatchesRemove(matchIDs); err != nil {
		return xerrors.Errorf("error publishing live matches change: %w", err)
	}

	return nil
}

func (p *Collector) busPubAllMatches() error {
	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesReplace,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_COLLECTION_OP_REPLACE,
			Matches: p.matches.All(),
		},
	})
}

func (p *Collector) handleError(err error) {
	if e := (&errRedisOp{}); xerrors.As(err, &e) {
		p.log.WithField("key", e.Key).WithError(e.Err).Error("redis error")
	} else if e := (&errRedisPubsub{}); xerrors.As(err, &e) {
		p.log.WithField("topic", e.Topic).WithError(e.Err).Error("redis error")
	} else {
		p.log.WithError(err).Error("live_matches error")
	}

	p.log.Errorx(err)
}
