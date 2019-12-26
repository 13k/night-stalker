package livematches

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/cskr/pubsub"
	"github.com/faceit/go-steam"
	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"
	"github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nsctx "github.com/13k/night-stalker/internal/context"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nsrds "github.com/13k/night-stalker/internal/redis"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "live_matches"

	msgTypeFindTopSourceTVGames  = protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames
	msgTypeMatchesMinimalRequest = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest
)

var _ nsproc.Processor = (*Watcher)(nil)

type Watcher struct {
	seq                  *querySeq
	interval             time.Duration
	ctx                  context.Context
	log                  *nslog.Logger
	db                   *gorm.DB
	redis                *redis.Client
	steam                *steam.Client
	dota                 *dota2.Dota2
	bus                  *pubsub.PubSub
	busSubSession        chan interface{}
	busSubMatchesMinimal chan interface{}
	busSubSourceTVGames  chan interface{}
}

func NewWatcher(interval time.Duration) *Watcher {
	return &Watcher{
		interval: interval,
		seq:      newQuerySeq(),
	}
}

func (p *Watcher) ChildSpec(stimeout time.Duration) oversight.ChildProcessSpecification {
	return oversight.ChildProcessSpecification{
		Name:     processorName,
		Restart:  oversight.Transient(),
		Start:    p.Start,
		Shutdown: oversight.Timeout(stimeout),
	}
}

func (p *Watcher) Start(ctx context.Context) error {
	if p.log = nsctx.GetLogger(ctx); p.log == nil {
		return nsproc.ErrProcessorContextLogger
	}

	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	if p.redis = nsctx.GetRedis(ctx); p.db == nil {
		return nsproc.ErrProcessorContextRedis
	}

	if p.steam = nsctx.GetSteam(ctx); p.steam == nil {
		return nsproc.ErrProcessorContextSteamClient
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return nsproc.ErrProcessorContextDotaClient
	}

	if p.bus = nsctx.GetPubSub(ctx); p.bus == nil {
		return nsproc.ErrProcessorContextPubSub
	}

	p.ctx = ctx
	p.log = p.log.WithPackage(processorName)
	p.busSubSession = p.bus.Sub(nsbus.TopicSession)
	p.busSubMatchesMinimal = p.bus.Sub(nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse)
	p.busSubSourceTVGames = p.bus.Sub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse)

	return p.loop()
}

func (p *Watcher) loop() error {
	ready := false
	ticker := time.NewTicker(p.interval)

	defer func() {
		ticker.Stop()
		p.log.Warn("stop")
	}()

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg, ok := <-p.busSubSession:
			if !ok {
				return nil
			}

			if change, ok := busmsg.(*nsbus.SessionChangeMessage); ok {
				ready = change.IsReady

				if ready {
					p.log.Debug("started querying live matches")
					p.tick()
				} else {
					p.flush()
					p.log.Warn("stopped querying live matches")
				}
			}
		case busmsg := <-p.busSubMatchesMinimal:
			if msg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				p.handleGCMessage(msg)
			}
		case busmsg := <-p.busSubSourceTVGames:
			if msg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				p.handleGCMessage(msg)
			}
		case <-ticker.C:
			if ready {
				p.tick()
			}
		}
	}
}

// TODO: investigate CMsgGCPlayerInfoRequest (k_EMsgGCPlayerInfoRequest) [async]
func (p *Watcher) tick() {
	p.flush()

	if err := p.query(); err != nil {
		p.log.WithError(err).Error("error querying live matches")
	}
}

func (p *Watcher) query() error {
	if p.ctx.Err() != nil {
		return p.ctx.Err()
	}

	page := &queryPage{}

	if p.seq.index == 0 {
		return p.queryPage(page)
	}

	page.start = p.seq.total - p.seq.psize

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
	}).Info("requesting live matches")

	busmsg := &nsbus.GCDispatcherSendMessage{
		MsgType: msgTypeFindTopSourceTVGames,
		Message: req,
	}

	p.bus.Pub(busmsg, nsbus.TopicGCDispatcherSend)

	return nil
}

func (p *Watcher) requestMatchesMinimal(matchIDs []uint64) { //nolint: unused
	req := &protocol.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}

	p.log.WithField("count", len(matchIDs)).Debug("requesting matches minimal")

	busmsg := &nsbus.GCDispatcherSendMessage{
		MsgType: msgTypeMatchesMinimalRequest,
		Message: req,
	}

	p.bus.Pub(busmsg, nsbus.TopicGCDispatcherSend)
}

func (p *Watcher) handleGCMessage(received *nsbus.GCDispatcherReceivedMessage) {
	switch msg := received.Message.(type) {
	case *protocol.CMsgGCToClientFindTopSourceTVGamesResponse:
		p.handleFindTopSourceTVGamesResponse(msg)
	case *protocol.CMsgClientToGCMatchesMinimalResponse:
		p.handleMatchesMinimalResponse(msg)
	}
}

func (p *Watcher) handleMatchesMinimalResponse(msg *protocol.CMsgClientToGCMatchesMinimalResponse) {
	p.log.WithField("count", len(msg.GetMatches())).Debug("received matches minimal")

	for _, match := range msg.GetMatches() {
		p.log.WithFields(logrus.Fields{
			"match_id": match.GetMatchId(),
		}).Debug("received match info")
	}
}

func (p *Watcher) handleFindTopSourceTVGamesResponse(
	msg *protocol.CMsgGCToClientFindTopSourceTVGamesResponse,
) {
	query := &queryPage{
		index: msg.GetGameListIndex(),
		start: msg.GetStartGame(),
		res:   msg,
	}

	p.handleResponse(query)
}

func (p *Watcher) handleResponse(page *queryPage) {
	if p.seq.IsCached(page) {
		p.log.WithFields(logrus.Fields{
			"seq_index": p.seq.index,
			"res_index": page.index,
			"res_start": page.start,
		}).Debug("cached live matches response")

		return
	}

	p.log.WithFields(logrus.Fields{
		"index": page.index,
		"start": page.start,
	}).Info("received live matches")

	if err := p.saveResponse(page.res); err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{
			"index": page.index,
			"start": page.start,
		}).Error("error saving live matches response")

		return
	}

	games := page.res.GetGameList()
	matchIDs := make([]uint64, len(games))
	rZMatchIDs := make([]*redis.Z, len(games))
	rKey := nsrds.KeyLiveMatches(int(page.index))

	for i, game := range games {
		matchIDs[i] = game.GetMatchId()

		rZMatchIDs[i] = &redis.Z{
			Score:  float64(game.GetSortScore()),
			Member: game.GetMatchId(),
		}
	}

	if err := p.redis.ZAdd(rKey, rZMatchIDs...).Err(); err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{
			"index": page.index,
			"start": page.start,
		}).Error("error caching live match IDs")

		return
	}

	// p.requestMatchesMinimal(matchIDs)

	if p.seq.index == 0 {
		p.seq.Init(page)

		if err := p.query(); err != nil {
			p.log.WithError(err).WithFields(logrus.Fields{
				"index": page.index,
				"start": page.start,
			}).Error("error requesting live matches")

			return
		}
	}

	p.seq.Cache(page)

	if p.seq.IsFull() {
		p.flush()
	}
}

// TODO: remove old redis zsets
func (p *Watcher) flush() {
	defer p.seq.Reset()

	if p.seq.IsEmpty() {
		return
	}

	index := p.seq.MaxIndex()

	var matches []*protocol.CSourceTVGameSmall

	for _, page := range p.seq.Pages() {
		matches = append(matches, page.res.GetGameList()...)
	}

	p.log.WithFields(logrus.Fields{
		"index": index,
		"count": len(matches),
	}).Debug("flushing matches")

	busmsg := &nsbus.LiveMatchesDotaMessage{
		Index:   index,
		Matches: matches,
	}

	p.bus.Pub(busmsg, nsbus.TopicLiveMatches)

	if err := p.redis.Set(nsrds.KeyLiveMatchesIndex, index, 0).Err(); err != nil {
		p.log.
			WithError(err).
			WithField("index", index).
			Error("error caching live matches index")

		return
	}

	rKey := nsrds.KeyLiveMatches(int(index))
	p.redis.Publish(nsrds.TopicLiveMatchesUpdate, rKey)
}

func (p *Watcher) saveResponse(resp *protocol.CMsgGCToClientFindTopSourceTVGamesResponse) error {
	for _, match := range resp.GetGameList() {
		if p.ctx.Err() != nil {
			return p.ctx.Err()
		}

		l := p.log.WithFields(logrus.Fields{
			"match_id": match.GetMatchId(),
			"lobby_id": match.GetLobbyId(),
		})

		tx := p.db.Begin()
		liveMatch := models.LiveMatchDotaProto(match)
		result := tx.
			Where(models.LiveMatch{MatchID: liveMatch.MatchID}).
			Assign(liveMatch).
			FirstOrCreate(liveMatch)

		if err := result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Error("error upserting match")
			return err
		}

		for _, player := range match.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()
				return p.ctx.Err()
			}

			livePlayer := models.LiveMatchPlayerDotaProto(player)
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
				return err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}
	}

	// p.log.WithField("count", len(resp.GetGameList())).Debug("updated live matches")
	return nil
}
