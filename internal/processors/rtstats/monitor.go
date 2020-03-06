package rtstats

import (
	"context"
	"fmt"
	"net/url"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"cirello.io/oversight"
	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
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
	processorName              = "rtstats"
	resultsQueueSize           = 10
	resultsBufferSize          = 10
	resultsBufferFlushInterval = 5 * time.Second
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
	workerPool            *ants.Pool
	api                   *geyserd2.Client
	apiMatchStats         *geyserd2.DOTA2MatchStats
	bus                   *nsbus.Bus
	busLiveMatchesReplace *nsbus.Subscription
	activeReqsMtx         sync.Mutex
	activeReqs            map[nspb.MatchID]bool
	liveMatchesMtx        sync.RWMutex
	liveMatches           nscol.LiveMatches
	results               nscol.LiveMatchStats
	resultsCh             chan *models.LiveMatchStats
}

func NewMonitor(options MonitorOptions) *Monitor {
	p := &Monitor{
		options:    options,
		log:        options.Log.WithPackage(processorName),
		bus:        options.Bus,
		activeReqs: make(map[nspb.MatchID]bool),
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

	if err := p.setupAPI(); err != nil {
		return err
	}

	if err := p.setupWorkerPool(); err != nil {
		return err
	}

	p.setupResults()
	p.busSubscribe()

	go p.resultsLoop()

	return p.loop()
}

func (p *Monitor) busSubscribe() {
	if p.busLiveMatchesReplace == nil {
		p.busLiveMatchesReplace = p.bus.Sub(nsbus.TopicLiveMatchesReplace)
	}
}

func (p *Monitor) busUnsubscribe() {
	if p.busLiveMatchesReplace != nil {
		p.bus.Unsub(p.busLiveMatchesReplace)
		p.busLiveMatchesReplace = nil
	}
}

func (p *Monitor) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.api = nsctx.GetDotaAPI(ctx); p.api == nil {
		return nsproc.ErrProcessorContextDotaAPI
	}

	p.ctx = ctx

	return nil
}

func (p *Monitor) setupAPI() error {
	if p.apiMatchStats != nil {
		return nil
	}

	var err error

	if p.apiMatchStats, err = p.api.DOTA2MatchStats(); err != nil {
		p.log.WithError(err).Error("error creating API interface")
		return err
	}

	return nil
}

func (p *Monitor) setupWorkerPool() error {
	if p.workerPool != nil {
		return nil
	}

	var err error

	if p.workerPool, err = ants.NewPool(p.options.PoolSize); err != nil {
		p.log.WithError(err).Error("error starting worker pool")
		return err
	}

	return nil
}

func (p *Monitor) teardownWorkerPool() {
	if p.workerPool != nil {
		p.workerPool.Release()
		p.workerPool = nil
	}
}

func (p *Monitor) setupResults() {
	if p.resultsCh == nil {
		p.resultsCh = make(chan *models.LiveMatchStats, resultsQueueSize)
	}
}

func (p *Monitor) teardownResults() {
	if p.resultsCh != nil {
		close(p.resultsCh)
		p.resultsCh = nil
	}
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
		}
	}
}

func (p *Monitor) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.teardownWorkerPool()
	p.teardownResults()
	p.log.Warn("stop")
}

func (p *Monitor) resultsLoop() {
	t := time.NewTicker(resultsBufferFlushInterval)

	defer func() {
		t.Stop()
		p.log.Debug("stop results")
	}()

	p.log.Debug("start results")

	for {
		select {
		case <-t.C:
			p.flushResults()
		case stats, ok := <-p.resultsCh:
			if !ok {
				return
			}

			p.results = append(p.results, stats)

			if len(p.results) >= resultsBufferSize {
				p.flushResults()
			}
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

	p.liveMatchesMtx.Lock()
	defer p.liveMatchesMtx.Unlock()
	p.liveMatches = msg.Matches
}

func (p *Monitor) tick() {
	p.liveMatchesMtx.RLock()
	defer p.liveMatchesMtx.RUnlock()

	if len(p.liveMatches) == 0 {
		return
	}

	p.log.WithFields(logrus.Fields{
		"count": len(p.liveMatches),
	}).Debug("requesting stats")

	for _, liveMatch := range p.liveMatches {
		p.submitLiveMatch(liveMatch)
	}
}

func (p *Monitor) submitLiveMatch(liveMatch *models.LiveMatch) {
	l := p.log.WithField("match_id", liveMatch.MatchID)

	if err := p.ctx.Err(); err != nil {
		l.WithError(err).Error()
		return
	}

	err := p.workerPool.Submit(func() {
		p.work(liveMatch)
	})

	if err != nil {
		l.WithError(err).Error("error creating worker")
	}
}

func (p *Monitor) work(liveMatch *models.LiveMatch) {
	l := p.log.WithFields(logrus.Fields{
		"match_id":        liveMatch.MatchID,
		"server_steam_id": liveMatch.ServerSteamID,
	})

	var skip bool
	p.activeReqsMtx.Lock()
	skip = p.activeReqs[liveMatch.MatchID]
	p.activeReqsMtx.Unlock()

	if skip {
		l.Warn("request in progress")
		return
	}

	p.activeReqsMtx.Lock()
	p.activeReqs[liveMatch.MatchID] = true
	p.activeReqsMtx.Unlock()

	defer func() {
		p.activeReqsMtx.Lock()
		delete(p.activeReqs, liveMatch.MatchID)
		p.activeReqsMtx.Unlock()
	}()

	result, err := p.requestMatchStats(liveMatch)

	if err != nil {
		l.WithError(err).Error("error requesting API")
		return
	}

	if result.GetMatch() == nil {
		return
	}

	if result.GetMatch().GetMatchid() != liveMatch.MatchID {
		return
	}

	stats, err := p.createLiveMatchStats(liveMatch, result)

	if err != nil {
		l.WithError(err).Error("error saving stats to database")
		return
	}

	p.resultsCh <- stats
}

func (p *Monitor) flushResults() {
	if len(p.results) == 0 {
		return
	}

	if err := p.busPublishLiveMatchStatsAdd(p.results...); err != nil {
		p.log.WithError(err).Error()
	}

	p.results = nil
}

func (p *Monitor) requestMatchStats(liveMatch *models.LiveMatch) (*protocol.CMsgDOTARealtimeGameStatsTerse, error) {
	req, err := p.apiMatchStats.GetRealtimeStats()

	if err != nil {
		p.log.WithError(err).Error("error creating API request")
		return nil, err
	}

	headers := map[string]string{
		"Accept":     "application/json",
		"Connection": "keep-alive",
	}

	params := url.Values{}
	params.Set("server_steam_id", strconv.FormatUint(liveMatch.ServerSteamID.ToUint64(), 10))

	reqOptions := geyser.RequestOptions{
		Context: p.ctx,
		Params:  params,
		Headers: headers,
	}

	result := &apiResult{}
	req.SetOptions(reqOptions).SetResult(result)

	resp, err := req.Execute()

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("invalid response status: %s", resp.Status())
	}

	return result.ToProto(), nil
}

func (p *Monitor) createLiveMatchStats(
	liveMatch *models.LiveMatch,
	result *protocol.CMsgDOTARealtimeGameStatsTerse,
) (*models.LiveMatchStats, error) {
	stats := models.NewLiveMatchStats(liveMatch, result)

	for _, team := range result.GetTeams() {
		stats.Teams = append(stats.Teams, models.LiveMatchStatsTeamDotaProto(team))

		for _, player := range team.GetPlayers() {
			stats.Players = append(stats.Players, models.NewLiveMatchStatsPlayer(stats, player))
		}
	}

	for _, pickban := range result.GetMatch().GetPicks() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(false, pickban))
	}

	for _, pickban := range result.GetMatch().GetBans() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(true, pickban))
	}

	for _, building := range result.GetBuildings() {
		stats.Buildings = append(stats.Buildings, models.LiveMatchStatsBuildingDotaProto(building))
	}

	if err := p.db.Save(stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

func (p *Monitor) busPublishLiveMatchStatsAdd(stats ...*models.LiveMatchStats) error {
	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchStatsAdd,
		Payload: &nsbus.LiveMatchStatsChangeMessage{
			Op:    nspb.CollectionOp_COLLECTION_OP_ADD,
			Stats: stats,
		},
	})
}
