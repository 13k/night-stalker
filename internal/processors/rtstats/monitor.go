package rtstats

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"cirello.io/oversight"
	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "rtstats"
)

type MonitorOptions struct {
	Logger          *nslog.Logger
	PoolSize        int
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Monitor)(nil)

type Monitor struct {
	options           *MonitorOptions
	ctx               context.Context
	log               *nslog.Logger
	db                *gorm.DB
	redis             *redis.Client
	workerPool        *ants.Pool
	api               *geyserd2.Client
	apiMatchStats     *geyserd2.DOTA2MatchStats
	bus               *nsbus.Bus
	busSubLiveMatches chan interface{}
	activeReqsMtx     sync.Mutex
	activeReqs        map[nspb.MatchID]bool
}

func NewMonitor(options *MonitorOptions) *Monitor {
	return &Monitor{
		options:    options,
		log:        options.Logger.WithPackage(processorName),
		activeReqs: make(map[nspb.MatchID]bool),
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

	if p.redis = nsctx.GetRedis(ctx); p.redis == nil {
		return nsproc.ErrProcessorContextRedis
	}

	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	if p.api = nsctx.GetDotaAPI(ctx); p.api == nil {
		return nsproc.ErrProcessorContextDotaAPI
	}

	p.ctx = ctx
	p.busSubLiveMatches = p.bus.Sub(nsbus.TopicLiveMatches)

	var err error

	if p.apiMatchStats, err = p.api.DOTA2MatchStats(); err != nil {
		p.log.WithError(err).Error("error creating API interface")
		return err
	}

	if p.workerPool, err = ants.NewPool(p.options.PoolSize); err != nil {
		p.log.WithError(err).Error("error starting worker pool")
		return err
	}

	return p.loop()
}

func (p *Monitor) loop() error {
	defer func() {
		p.workerPool.Release()
		p.log.Warn("stop")
	}()

	p.log.Info("start")

	for {
		select {
		case busmsg, ok := <-p.busSubLiveMatches:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.(*nsbus.LiveMatchesMessage); ok {
				p.handleLiveMatches(msg)
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *Monitor) handleLiveMatches(msg *nsbus.LiveMatchesMessage) {
	p.log.WithFields(logrus.Fields{
		"count": len(msg.Matches),
	}).Debug("processing live matches")

	for _, match := range msg.Matches {
		p.handleLiveMatch(match)
	}
}

func (p *Monitor) handleLiveMatch(match *models.LiveMatch) {
	l := p.log.WithField("match_id", match.MatchID)

	if match.MatchID == 0 {
		l.WithFields(logrus.Fields{
			"nil_match":       match == nil,
			"server_steam_id": match.ServerSteamID,
			"lobby_id":        match.LobbyID,
		}).Warn("ignoring match with zero match_id")

		return
	}

	if err := p.ctx.Err(); err != nil {
		l.WithError(err).Error("error processing live match")
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

	apiStats, err := p.requestMatchStats(match)

	if err != nil {
		l.WithError(err).Error("error requesting API")
		return
	}

	if apiStats.GetMatch() == nil {
		return
	}

	stats, err := p.createMatchStats(apiStats)

	if err != nil {
		l.WithError(err).Error("error saving stats to database")
		return
	}

	if err := p.publishChange(stats); err != nil {
		l.WithError(err).Error("error publishing stats change")
		return
	}

	l.Debug("saved and published stats")
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

	result := &protocol.CMsgDOTARealtimeGameStatsTerse{}

	req.SetOptions(reqOptions).SetResult(result)

	resp, err := req.Execute()

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("invalid response status: %s", resp.Status())
	}

	return result, nil
}

func (p *Monitor) createMatchStats(apiStats *protocol.CMsgDOTARealtimeGameStatsTerse) (*models.LiveMatchStats, error) {
	stats := models.LiveMatchStatsDotaProto(apiStats)

	for _, team := range apiStats.GetTeams() {
		stats.Teams = append(stats.Teams, models.LiveMatchStatsTeamDotaProto(team))

		for _, player := range team.GetPlayers() {
			stats.Players = append(stats.Players, models.LiveMatchStatsPlayerDotaProto(player))
		}
	}

	for _, pickban := range apiStats.GetMatch().GetPicks() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(false, pickban))
	}

	for _, pickban := range apiStats.GetMatch().GetBans() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(true, pickban))
	}

	for _, building := range apiStats.GetBuildings() {
		stats.Buildings = append(stats.Buildings, models.LiveMatchStatsBuildingDotaProto(building))
	}

	if err := p.db.Save(stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

func (p *Monitor) publishChange(stats *models.LiveMatchStats) error {
	return p.redis.Publish(nsrds.TopicMatchStatsUpdate, stats.MatchID).Err()
}
