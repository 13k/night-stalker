package livematches

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
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
	ctx                     context.Context
	options                 CollectorOptions
	matches                 *nscol.LiveMatches
	log                     *nslog.Logger
	db                      *gorm.DB
	rds                     *redis.Client
	bus                     *nsbus.Bus
	busSubLiveMatchesAll    <-chan nsbus.Message
	busSubLiveMatchStatsAll <-chan nsbus.Message
}

func NewCollector(options CollectorOptions) *Collector {
	proc := &Collector{
		options: options,
		log:     options.Log.WithPackage(processorName),
		bus:     options.Bus,
	}

	proc.busSubscribe()

	return proc
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
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: shutdown,
	}
}

func (p *Collector) Start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return err
	}

	if err := p.seedLiveMatches(); err != nil {
		return err
	}

	return p.loop()
}

func (p *Collector) busSubscribe() {
	if p.busSubLiveMatchesAll == nil {
		p.busSubLiveMatchesAll = p.bus.Sub(nsbus.TopicPatternLiveMatchesAll)
	}

	if p.busSubLiveMatchStatsAll == nil {
		p.busSubLiveMatchStatsAll = p.bus.Sub(nsbus.TopicPatternLiveMatchStatsAll)
	}
}

func (p *Collector) busUnsubscribe() {
	if p.busSubLiveMatchesAll != nil {
		p.bus.Unsub(nsbus.TopicPatternLiveMatchesAll, p.busSubLiveMatchesAll)
	}

	if p.busSubLiveMatchStatsAll != nil {
		p.bus.Unsub(nsbus.TopicLiveMatchStatsAdd, p.busSubLiveMatchStatsAll)
	}
}

func (p *Collector) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.rds = nsctx.GetRedis(ctx); p.rds == nil {
		return nsproc.ErrProcessorContextRedis
	}

	p.ctx = ctx

	return nil
}

func (p *Collector) seedLiveMatches() error {
	if p.matches != nil {
		return nil
	}

	matchIDs, err := p.rdsLiveMatchIDs()

	if err != nil {
		p.log.WithError(err).Error("error loading live matches ids")
		return err
	}

	if len(matchIDs) == 0 {
		p.matches = nscol.NewLiveMatches()
		return nil
	}

	liveMatches, err := p.loadLiveMatches(matchIDs)

	if err != nil {
		p.log.WithError(err).Error("error loading live matches")
		return err
	}

	p.matches = nscol.NewLiveMatches(liveMatches...)

	p.log.WithField("count", len(liveMatches)).Debug("seeded matches")

	return nil
}

func (p *Collector) rdsLiveMatchIDs() (nscol.MatchIDs, error) {
	result := p.rds.ZRevRange(nsrds.KeyLiveMatches, 0, -1)

	if err := result.Err(); err != nil {
		p.log.
			WithField("key", nsrds.KeyLiveMatches).
			WithError(err).
			Error("error fetching cached live matches index")

		return nil, err
	}

	matchIDs := make(nscol.MatchIDs, len(result.Val()))

	if err := result.ScanSlice(&matchIDs); err != nil {
		p.log.WithError(err).Error("error parsing live match IDs")
		return nil, err
	}

	return matchIDs, nil
}

func (p *Collector) loadLiveMatches(matchIDs nscol.MatchIDs) (nscol.LiveMatchesSlice, error) {
	var matches nscol.LiveMatchesSlice

	err := p.db.
		Where("match_id IN (?)", matchIDs).
		Order("sort_score DESC").
		Find(&matches).
		Error

	if err != nil {
		p.log.WithError(err).Error("database live matches")
		return nil, err
	}

	return matches, nil
}

func (p *Collector) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	defer func() {
		p.busUnsubscribe()
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSubLiveMatchesAll:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				p.handleLiveMatchesChange(msg)
			}
		case busmsg, ok := <-p.busSubLiveMatchStatsAll:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchStatsChangeMessage); ok {
				p.handleLiveMatchStatsChange(msg)
			}
		}
	}
}

func (p *Collector) handleLiveMatchesChange(msg *nsbus.LiveMatchesChangeMessage) {
	switch msg.Op {
	case nspb.CollectionOp_ADD, nspb.CollectionOp_UPDATE:
		p.add(msg.Matches)
	case nspb.CollectionOp_REMOVE:
		p.remove(msg.MatchIDs)
	default:
		return
	}
}

func (p *Collector) add(matches nscol.LiveMatchesSlice) {
	if len(matches) == 0 {
		return
	}

	if err := p.rdsAddLiveMatches(matches); err != nil {
		p.log.WithError(err).Error("failed to append matches to redis")
		return
	}

	prevLen := p.matches.Len()
	change := p.matches.Add(matches...)

	if len(change) > 0 {
		p.log.WithFields(logrus.Fields{
			"count":         len(change),
			"total_before":  prevLen,
			"total_current": p.matches.Len(),
		}).Debug("matches added")

		p.notifyLiveMatchesAdd(change)
	}
}

func (p *Collector) remove(matchIDs nscol.MatchIDs) {
	if len(matchIDs) == 0 {
		return
	}

	if err := p.rdsRemoveLiveMatches(matchIDs); err != nil {
		p.log.WithError(err).Error("failed to remove matches from redis")
		return
	}

	prevLen := p.matches.Len()
	change := p.matches.Remove(matchIDs...)

	if len(change) > 0 {
		p.log.WithFields(logrus.Fields{
			"count":         len(change),
			"total_before":  prevLen,
			"total_current": p.matches.Len(),
		}).Debug("matches removed")

		p.notifyLiveMatchesRemove(change)
	}
}

func (p *Collector) handleLiveMatchStatsChange(msg *nsbus.LiveMatchStatsChangeMessage) {
	switch msg.Op {
	case nspb.CollectionOp_ADD:
		p.addStats(msg.Stats)
	default:
		return
	}
}

func (p *Collector) addStats(stats nscol.LiveMatchStatsSlice) {
	l := p.log.WithField("count", len(stats))

	if err := p.rdsPubLiveMatchStatsAdd(stats); err != nil {
		l.WithError(err).Error("error publishing match stats update")
		return
	}

	l.Debug("received match stats")
}

func (p *Collector) notifyLiveMatchesAdd(liveMatches nscol.LiveMatchesSlice) {
	p.busPublishAllMatches()

	if err := p.rdsPubLiveMatchesAdd(liveMatches); err != nil {
		p.log.WithError(err).Error("error publishing live matches change")
		return
	}
}

func (p *Collector) notifyLiveMatchesRemove(matchIDs nscol.MatchIDs) {
	p.busPublishAllMatches()

	if err := p.rdsPubLiveMatchesRemove(matchIDs); err != nil {
		p.log.WithError(err).Error("error publishing live matches change")
		return
	}
}

func (p *Collector) busPublishAllMatches() {
	p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesReplace,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_REPLACE,
			Matches: p.matches.All(),
		},
	})
}

func (p *Collector) rdsAddLiveMatches(liveMatches nscol.LiveMatchesSlice) error {
	if err := p.rds.ZAdd(nsrds.KeyLiveMatches, liveMatches.ToRedisZValues()...).Err(); err != nil {
		return err
	}

	zValuesByTime := make([]*redis.Z, len(liveMatches))

	for i, liveMatch := range liveMatches {
		var activateTimeUnix int64

		if liveMatch.ActivateTime != nil {
			activateTimeUnix = liveMatch.ActivateTime.UTC().Unix()
		}

		zValuesByTime[i] = &redis.Z{
			Member: liveMatch.MatchID,
			Score:  float64(activateTimeUnix),
		}
	}

	return p.rds.ZAdd("live_matches_by_time", zValuesByTime...).Err()
}

func (p *Collector) rdsRemoveLiveMatches(matchIDs nscol.MatchIDs) error {
	ifaceMatchIDs := matchIDs.ToInterfaces()

	if err := p.rds.ZRem(nsrds.KeyLiveMatches, ifaceMatchIDs...).Err(); err != nil {
		return err
	}

	return p.rds.ZRem("live_matches_by_time", ifaceMatchIDs...).Err()
}

func (p *Collector) rdsPubLiveMatchesAdd(liveMatches nscol.LiveMatchesSlice) error {
	return p.rds.Publish(nsrds.TopicLiveMatchesAdd, liveMatches.MatchIDs().Join(",")).Err()
}

func (p *Collector) rdsPubLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	return p.rds.Publish(nsrds.TopicLiveMatchesRemove, matchIDs.Join(",")).Err()
}

func (p *Collector) rdsPubLiveMatchStatsAdd(stats nscol.LiveMatchStatsSlice) error {
	return p.rds.Publish(nsrds.TopicLiveMatchStatsAdd, stats.MatchIDs().Join(",")).Err()
}