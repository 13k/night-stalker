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
	nspb "github.com/13k/night-stalker/internal/protocol"
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
	busLiveMatchesReplace <-chan nsbus.Message
	activeReqsMtx         sync.Mutex
	activeReqs            map[nspb.MatchID]bool
	matchesMtx            sync.RWMutex
	matches               nscol.LiveMatches
	results               nscol.LiveMatchStats
	resultsCh             chan *models.LiveMatchStats
}

func NewMonitor(options MonitorOptions) *Monitor {
	proc := &Monitor{
		options:    options,
		log:        options.Log.WithPackage(processorName),
		bus:        options.Bus,
		activeReqs: make(map[nspb.MatchID]bool),
		resultsCh:  make(chan *models.LiveMatchStats, resultsQueueSize),
	}

	proc.busSubscribe()

	return proc
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

	return p.loop()
}

func (p *Monitor) busSubscribe() {
	if p.busLiveMatchesReplace == nil {
		p.busLiveMatchesReplace = p.bus.Sub(nsbus.TopicLiveMatchesReplace)
	}
}

func (p *Monitor) busUnsubscribe() {
	if p.busLiveMatchesReplace != nil {
		p.bus.Unsub(nsbus.TopicLiveMatchesReplace, p.busLiveMatchesReplace)
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
		p.busUnsubscribe()
		p.workerPool.Release()
		p.log.Warn("stop")
	}()

	go p.resultsLoop()

	p.log.Info("start")

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-tick.C:
			p.tick()
		case busmsg, ok := <-p.busLiveMatchesReplace:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.Payload.(*nsbus.LiveMatchesChangeMessage); ok {
				p.handleLiveMatchesChange(msg)
			}
		}
	}
}

func (p *Monitor) resultsLoop() {
	tick := time.NewTicker(resultsBufferFlushInterval)

	defer func() {
		tick.Stop()
		p.log.Debug("stop results")
	}()

	p.log.Debug("start results")

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-tick.C:
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
	if msg.Op != nspb.CollectionOp_REPLACE {
		p.log.WithField("op", msg.Op.String()).Warn("ignored live matches change message")
		return
	}

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
	}).Debug("requesting stats")

	for _, match := range p.matches {
		p.submitLiveMatch(match)
	}
}

func (p *Monitor) submitLiveMatch(match *models.LiveMatch) {
	l := p.log.WithField("match_id", match.MatchID)

	if err := p.ctx.Err(); err != nil {
		l.WithError(err).Error()
		return
	}

	err := p.workerPool.Submit(func() {
		p.work(match)
	})

	if err != nil {
		l.WithError(err).Error("error creating worker")
	}
}

func (p *Monitor) work(match *models.LiveMatch) {
	l := p.log.WithFields(logrus.Fields{
		"match_id":        match.MatchID,
		"server_steam_id": match.ServerSteamID,
	})

	var skip bool
	p.activeReqsMtx.Lock()
	skip = p.activeReqs[match.MatchID]
	p.activeReqsMtx.Unlock()

	if skip {
		l.Warn("request in progress")
		return
	}

	p.activeReqsMtx.Lock()
	p.activeReqs[match.MatchID] = true
	p.activeReqsMtx.Unlock()

	defer func() {
		p.activeReqsMtx.Lock()
		delete(p.activeReqs, match.MatchID)
		p.activeReqsMtx.Unlock()
	}()

	result, err := p.requestMatchStats(match)

	if err != nil {
		l.WithError(err).Error("error requesting API")
		return
	}

	if result.GetMatch() == nil {
		return
	}

	stats, err := p.createLiveMatchStats(result)

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

	p.busPublishLiveMatchStatsAdd(p.results...)
	p.results = nil
}

func (p *Monitor) requestMatchStats(match *models.LiveMatch) (*protocol.CMsgDOTARealtimeGameStatsTerse, error) {
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
	params.Set("server_steam_id", strconv.FormatUint(match.ServerSteamID.ToUint64(), 10))

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
	result *protocol.CMsgDOTARealtimeGameStatsTerse,
) (*models.LiveMatchStats, error) {
	stats := models.LiveMatchStatsDotaProto(result)

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

func (p *Monitor) busPublishLiveMatchStatsAdd(stats ...*models.LiveMatchStats) {
	p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchStatsAdd,
		Payload: &nsbus.LiveMatchStatsChangeMessage{
			Op:    nspb.CollectionOp_ADD,
			Stats: stats,
		},
	})
}
