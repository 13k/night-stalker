package livematches

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/faceit/go-steam"
	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "live_matches"

	msgTypeFindTopSourceTVGames = protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames
)

type WatcherOptions struct {
	Logger          *nslog.Logger
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Watcher)(nil)

type Watcher struct {
	options                   *WatcherOptions
	matches                   *nscol.LiveMatches
	discoveryPage             *discoveryPage
	ctx                       context.Context
	log                       *nslog.Logger
	db                        *gorm.DB
	redis                     *redis.Client
	steam                     *steam.Client
	dota                      *dota2.Dota2
	bus                       *nsbus.Bus
	busSubSourceTVGames       chan interface{}
	busSubLiveMatchesFinished chan interface{}
}

func NewWatcher(options *WatcherOptions) *Watcher {
	return &Watcher{
		options:       options,
		log:           options.Logger.WithPackage(processorName),
		matches:       nscol.NewLiveMatches(),
		discoveryPage: &discoveryPage{},
	}
}

func (p *Watcher) ChildSpec() oversight.ChildProcessSpecification {
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

func (p *Watcher) Start(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.redis = nsctx.GetRedis(ctx); p.redis == nil {
		return nsproc.ErrProcessorContextRedis
	}

	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return nsproc.ErrProcessorContextDotaClient
	}

	if p.bus = nsctx.GetBus(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextBus
	}

	p.ctx = ctx
	p.busSubSourceTVGames = p.bus.Sub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse)
	p.busSubLiveMatchesFinished = p.bus.Sub(nsbus.TopicLiveMatchesFinished)

	return p.loop()
}

func (p *Watcher) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	ticker := time.NewTicker(p.options.Interval)

	defer func() {
		ticker.Stop()
		p.log.Warn("stop")
	}()

	p.log.Info("start")
	p.tick()

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-ticker.C:
			p.tick()
		case busmsg := <-p.busSubLiveMatchesFinished:
			if msg, ok := busmsg.(*nsbus.LiveMatchesFinishedMessage); ok {
				p.handleMatchesFinished(msg)
			}
		case busmsg := <-p.busSubSourceTVGames:
			if dispatcherMsg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dispatcherMsg.Message.(*protocol.CMsgGCToClientFindTopSourceTVGamesResponse); ok {
					p.handleFindTopSourceTVGamesResponse(res)
				}
			}
		}
	}
}

// TODO: investigate CMsgGCPlayerInfoRequest (k_EMsgGCPlayerInfoRequest) [async]
func (p *Watcher) tick() {
	if err := p.query(); err != nil {
		p.log.WithError(err).Error("error querying live matches")
	}
}

func (p *Watcher) query() error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	page := &queryPage{
		start: p.discoveryPage.LastPageStart(),
	}

	return p.queryPage(page)
}

func (p *Watcher) queryPage(page *queryPage) error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	req := &protocol.CMsgClientToGCFindTopSourceTVGames{
		GameListIndex: proto.Uint32(page.index),
		StartGame:     proto.Uint32(page.start),
	}

	p.log.WithFields(logrus.Fields{
		"index": req.GetGameListIndex(),
		"start": req.GetStartGame(),
	}).Debug("requesting live matches")

	busmsg := &nsbus.GCDispatcherSendMessage{
		MsgType: msgTypeFindTopSourceTVGames,
		Message: req,
	}

	p.bus.Pub(busmsg, nsbus.TopicGCDispatcherSend)

	return nil
}

func (p *Watcher) handleMatchesFinished(msg *nsbus.LiveMatchesFinishedMessage) {
	p.log.
		WithField("match_ids", msg.MatchIDs).
		Debug("matches finished")

	if err := p.redisRemoveLiveMatches(msg.MatchIDs...); err != nil {
		p.log.WithError(err).Error("error removing matches from redis")
		return
	}

	change := p.matches.Remove(msg.MatchIDs...)

	if change > 0 {
		p.busPublishMatches()

		if err := p.redisPublishMatchesUpdate(); err != nil {
			p.log.WithError(err).Error("error publishing live matches update")
			return
		}
	}
}

func (p *Watcher) handleFindTopSourceTVGamesResponse(msg *protocol.CMsgGCToClientFindTopSourceTVGamesResponse) {
	go p.handleResponse(newQueryPage(msg))
}

func (p *Watcher) handleResponse(page *queryPage) {
	l := p.log.WithFields(logrus.Fields{
		"index": page.index,
		"start": page.start,
		"total": page.total,
		"count": page.psize,
	})

	if p.discoveryPage.Empty() {
		l.Debug("received discovery response")

		p.discoveryPage.SetPage(page)

		if err := p.query(); err != nil {
			l.WithError(err).Error("error requesting live matches")
			return
		}

		return
	}

	games := cleanResponseGames(page.res.GetGameList())
	realCount := len(games)
	l = l.WithField("real_count", realCount)

	if page.psize != realCount {
		l.WithField("cleaned", page.psize-realCount).
			Warn("cleaned duplicate or invalid matches")
	}

	l.Debug("received live matches")

	matches, err := p.saveGames(games)

	if err != nil {
		l.WithError(err).Error("error saving live matches response")
		return
	}

	if err := p.redisAppendLiveMatches(matches...); err != nil {
		l.WithError(err).Error("error adding matches to redis")
		return
	}

	change := p.matches.Add(matches...)

	if change > 0 {
		p.busPublishMatches()

		if err := p.redisPublishMatchesUpdate(); err != nil {
			p.log.WithError(err).Error("error publishing live matches update")
			return
		}
	}
}

func (p *Watcher) saveGames(games []*protocol.CSourceTVGameSmall) ([]*models.LiveMatch, error) {
	liveMatches := make([]*models.LiveMatch, 0, len(games))

	for _, game := range games {
		if p.ctx.Err() != nil {
			return nil, p.ctx.Err()
		}

		l := p.log.WithFields(logrus.Fields{
			"match_id":        game.GetMatchId(),
			"server_steam_id": game.GetServerSteamId(),
		})

		liveMatch := models.LiveMatchDotaProto(game)

		tx := p.db.Begin()

		result := tx.
			Where(models.LiveMatch{MatchID: liveMatch.MatchID}).
			Assign(liveMatch).
			FirstOrCreate(liveMatch)

		if err := result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Error("error upserting match")
			return nil, err
		}

		for _, gamePlayer := range game.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()
				return nil, p.ctx.Err()
			}

			livePlayer := models.LiveMatchPlayerDotaProto(gamePlayer)
			livePlayer.LiveMatchID = liveMatch.ID

			criteria := &models.LiveMatchPlayer{
				LiveMatchID: livePlayer.LiveMatchID,
				AccountID:   livePlayer.AccountID,
			}

			result = tx.
				Where(criteria).
				Assign(livePlayer).
				FirstOrCreate(livePlayer)

			if err := result.Error; err != nil {
				tx.Rollback()
				l.WithError(err).Error("error upserting live match player")
				return nil, err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}

		liveMatches = append(liveMatches, liveMatch)
	}

	return liveMatches, nil
}

func (p *Watcher) busPublishMatches() {
	busmsg := &nsbus.LiveMatchesMessage{Matches: p.matches.All()}
	p.bus.Pub(busmsg, nsbus.TopicLiveMatches)
}

func (p *Watcher) redisAppendLiveMatches(matches ...*models.LiveMatch) error {
	rZMatchIDs := make([]*redis.Z, len(matches))

	for i, match := range matches {
		rZMatchIDs[i] = &redis.Z{
			Score:  match.SortScore,
			Member: match.MatchID,
		}
	}

	return p.redis.ZAdd(nsrds.KeyLiveMatches, rZMatchIDs...).Err()
}

func (p *Watcher) redisRemoveLiveMatches(matchIDs ...nspb.MatchID) error {
	ifaceMatchIDs := make([]interface{}, len(matchIDs))

	for i, matchID := range matchIDs {
		ifaceMatchIDs[i] = matchID
	}

	return p.redis.ZRem(nsrds.KeyLiveMatches, ifaceMatchIDs...).Err()
}

func (p *Watcher) redisPublishMatchesUpdate() error {
	result := p.redis.Publish(nsrds.TopicLiveMatchesUpdate, nsrds.KeyLiveMatches)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}
