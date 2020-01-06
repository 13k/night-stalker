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
	"github.com/cskr/pubsub"
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

var _ nsproc.Processor = (*Monitor)(nil)

type Monitor struct {
	ctx               context.Context
	log               *nslog.Logger
	db                *gorm.DB
	redis             *redis.Client
	workerPool        *ants.Pool
	poolSize          int
	api               *geyserd2.Client
	apiMatchStats     *geyserd2.DOTA2MatchStats
	bus               *pubsub.PubSub
	busSubLiveMatches chan interface{}
	activeReqsMtx     sync.Mutex
	activeReqs        map[nspb.MatchID]bool
}

func NewMonitor(poolSize int) *Monitor {
	return &Monitor{
		poolSize:   poolSize,
		activeReqs: make(map[nspb.MatchID]bool),
	}
}

func (p *Monitor) ChildSpec(stimeout time.Duration) oversight.ChildProcessSpecification {
	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: oversight.Timeout(stimeout),
	}
}

func (p *Monitor) Start(ctx context.Context) error {
	if p.log = nsctx.GetLogger(ctx); p.log == nil {
		return nsproc.ErrProcessorContextLogger
	}

	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.redis = nsctx.GetRedis(ctx); p.db == nil {
		return nsproc.ErrProcessorContextRedis
	}

	if p.bus = nsctx.GetPubSub(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextPubSub
	}

	if p.api = nsctx.GetDotaAPI(ctx); p.api == nil {
		return nsproc.ErrProcessorContextDotaAPI
	}

	p.ctx = ctx
	p.log = p.log.WithPackage(processorName)
	p.busSubLiveMatches = p.bus.Sub(nsbus.TopicLiveMatches)

	var err error

	if p.apiMatchStats, err = p.api.DOTA2MatchStats(); err != nil {
		p.log.WithError(err).Error("error creating API interface")
		return err
	}

	if p.workerPool, err = ants.NewPool(p.poolSize); err != nil {
		p.log.WithError(err).Error("error starting worker pool")
		return err
	}

	return p.loop()
}

func (p *Monitor) loop() error {
	defer func() {
		p.workerPool.Release()
	}()

	for {
		select {
		case busmsg, ok := <-p.busSubLiveMatches:
			if !ok {
				return nil
			}

			if msg, ok := busmsg.(*nsbus.LiveMatchesDotaMessage); ok {
				p.handleLiveMatches(msg)
			}
		case <-p.ctx.Done():
			return nil
		}
	}
}

func (p *Monitor) handleLiveMatches(msg *nsbus.LiveMatchesDotaMessage) {
	p.log.WithFields(logrus.Fields{
		"index": msg.Index,
		"count": len(msg.Matches),
	}).Debug("processing live matches")

	for _, match := range msg.Matches {
		p.handleLiveMatch(match)
	}
}

func (p *Monitor) handleLiveMatch(match *protocol.CSourceTVGameSmall) {
	l := p.log.WithField("match_id", match.GetMatchId())

	if err := p.ctx.Err(); err != nil {
		l.WithError(err).Error("error processing live match")
		return
	}

	if !p.needsUpdate(match) {
		return
	}

	err := p.workerPool.Submit(func() {
		p.work(match)
	})

	if err != nil {
		l.WithError(err).Error("error creating worker")
	}
}

func (p *Monitor) work(match *protocol.CSourceTVGameSmall) {
	l := p.log.WithField("match_id", match.GetMatchId())

	var skip bool
	p.activeReqsMtx.Lock()
	skip = p.activeReqs[match.GetMatchId()]
	p.activeReqsMtx.Unlock()

	if skip {
		l.Warn("request in progress")
		return
	}

	p.activeReqsMtx.Lock()
	p.activeReqs[match.GetMatchId()] = true
	p.activeReqsMtx.Unlock()

	defer func() {
		p.activeReqsMtx.Lock()
		delete(p.activeReqs, match.GetMatchId())
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

	l.Debug("saved stats")
}

func (p *Monitor) needsUpdate(match *protocol.CSourceTVGameSmall) bool {
	l := p.log.WithField("match_id", match.GetMatchId())

	stats, err := p.loadMatchStats(match)

	if err != nil {
		l.WithError(err).Error("error fetching stats from database")
		return false
	}

	if p.db.NewRecord(stats) {
		return true
	}

	var count int

	lmsp := &models.LiveMatchStatsPlayer{LiveMatchStatsID: stats.ID}
	dbResult := p.db.
		Model(models.LiveMatchStatsPlayerModel).
		Where(lmsp).
		Where("hero_id = ?", 0).
		Count(&count)

	if err = dbResult.Error; err != nil {
		l.WithError(err).Error("error counting stats players in database")
		return false
	}

	return count > 0
}

func (p *Monitor) loadMatchStats(match *protocol.CSourceTVGameSmall) (*models.LiveMatchStats, error) {
	stats := &models.LiveMatchStats{}

	dbResult := p.db.
		Select([]string{"id"}).
		Where(&models.LiveMatchStats{MatchID: match.GetMatchId()}).
		Order("updated_at DESC").
		FirstOrInit(stats)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return stats, nil
}

func (p *Monitor) requestMatchStats(
	match *protocol.CSourceTVGameSmall,
) (*protocol.CMsgDOTARealtimeGameStatsTerse, error) {
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
	params.Set("server_steam_id", strconv.FormatUint(match.GetServerSteamId(), 10))

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
	p.log.WithField("match_id", stats.MatchID).Debug("publishing match stats update")
	return p.redis.Publish(nsrds.TopicMatchStatsUpdate, stats.MatchID).Err()
}
