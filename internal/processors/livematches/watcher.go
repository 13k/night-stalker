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
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "live_matches"

	msgTypeFindTopSourceTVGames  = protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames
	msgTypeMatchesMinimalRequest = protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest
)

type WatcherOptions struct {
	Logger          *nslog.Logger
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Watcher)(nil)

type Watcher struct {
	options              *WatcherOptions
	seq                  *querySeq
	ctx                  context.Context
	log                  *nslog.Logger
	db                   *gorm.DB
	redis                *redis.Client
	steam                *steam.Client
	dota                 *dota2.Dota2
	bus                  *pubsub.PubSub
	busSubMatchesMinimal chan interface{}
	busSubSourceTVGames  chan interface{}
}

func NewWatcher(options *WatcherOptions) *Watcher {
	return &Watcher{
		options: options,
		log:     options.Logger.WithPackage(processorName),
		seq:     newQuerySeq(),
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
	p.busSubMatchesMinimal = p.bus.Sub(nsbus.TopicGCDispatcherReceivedMatchesMinimalResponse)
	p.busSubSourceTVGames = p.bus.Sub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse)

	return p.loop()
}

func (p *Watcher) loop() error {
	ticker := time.NewTicker(p.options.Interval)

	defer func() {
		p.flush()
		ticker.Stop()
		p.log.Warn("stop")
	}()

	p.log.Info("start")
	p.tick()

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case busmsg := <-p.busSubMatchesMinimal:
			if msg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				p.handleGCMessage(msg)
			}
		case busmsg := <-p.busSubSourceTVGames:
			if msg, ok := busmsg.(*nsbus.GCDispatcherReceivedMessage); ok {
				p.handleGCMessage(msg)
			}
		case <-ticker.C:
			p.tick()
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
	}).Debug("requesting live matches")

	busmsg := &nsbus.GCDispatcherSendMessage{
		MsgType: msgTypeFindTopSourceTVGames,
		Message: req,
	}

	p.bus.Pub(busmsg, nsbus.TopicGCDispatcherSend)

	return nil
}

func (p *Watcher) requestMatchesMinimal(matchIDs []nspb.MatchID) { //nolint: unused
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
	page := &queryPage{
		res:   msg,
		index: msg.GetGameListIndex(),
		start: msg.GetStartGame(),
		total: msg.GetNumGames(),
		games: msg.GetGameList(),
	}

	p.handleResponse(page)
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

	l := p.log.WithFields(logrus.Fields{
		"index": page.index,
		"start": page.start,
		"total": page.total,
	})

	l.Debug("received live matches")

	page.games = cleanResponseGames(page.res.GetGameList())

	matches, err := p.savePage(page)

	if err != nil {
		l.WithError(err).Error("error saving live matches response")
		return
	}

	page.matches = matches

	if err := p.appendRedisLiveMatches(page.index, matches); err != nil {
		l.WithError(err).Error("error caching live match IDs")
		return
	}

	p.seq.Cache(page)

	if p.seq.index == 0 {
		p.seq.Init(page)

		if err := p.query(); err != nil {
			l.WithError(err).Error("error requesting live matches")
			return
		}
	}

	if p.seq.IsFull() {
		p.flush()
	}
}

func (p *Watcher) savePage(page *queryPage) ([]*models.LiveMatch, error) {
	var matches []*models.LiveMatch

	for _, game := range page.games {
		if p.ctx.Err() != nil {
			return nil, p.ctx.Err()
		}

		l := p.log.WithFields(logrus.Fields{
			"match_id": game.GetMatchId(),
			"lobby_id": game.GetLobbyId(),
		})

		tx := p.db.Begin()
		liveMatch := models.LiveMatchDotaProto(game)
		result := tx.
			Where(models.LiveMatch{MatchID: liveMatch.MatchID}).
			Assign(liveMatch).
			FirstOrCreate(liveMatch)

		if err := result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Error("error upserting match")
			return nil, err
		}

		for _, player := range game.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()
				return nil, p.ctx.Err()
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
				return nil, err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}

		matches = append(matches, liveMatch)
	}

	return matches, nil
}

func (p *Watcher) appendRedisLiveMatches(index uint32, matches []*models.LiveMatch) error {
	rZMatchIDs := make([]*redis.Z, len(matches))
	rKey := nsrds.KeyLiveMatches(int(index))

	for i, match := range matches {
		rZMatchIDs[i] = &redis.Z{
			Score:  match.SortScore,
			Member: match.MatchID,
		}
	}

	return p.redis.ZAdd(rKey, rZMatchIDs...).Err()
}

func (p *Watcher) flush() {
	defer p.seq.Reset()

	if p.seq.IsEmpty() {
		return
	}

	index := p.seq.MaxIndex()

	var matches []*models.LiveMatch

	for _, page := range p.seq.Pages() {
		matches = append(matches, page.matches...)
	}

	p.log.WithFields(logrus.Fields{
		"index": index,
		"count": len(matches),
	}).Debug("flushing matches")

	p.publishBusMatches(matches)

	if err := p.publishRedisIndex(index); err != nil {
		p.log.
			WithError(err).
			WithField("index", index).
			Error("error publishing live matches update")

		return
	}

	if err := p.clearRedisIndices(index); err != nil {
		p.log.
			WithError(err).
			WithField("index", index).
			Error("error caching live matches index")

		return
	}
}

func (p *Watcher) publishBusMatches(matches []*models.LiveMatch) {
	busmsg := &nsbus.LiveMatchesMessage{Matches: matches}
	p.bus.Pub(busmsg, nsbus.TopicLiveMatches)
}

func (p *Watcher) publishRedisIndex(index uint32) error {
	if err := p.redis.Set(nsrds.KeyLiveMatchesIndex, index, 0).Err(); err != nil {
		return err
	}

	rKey := nsrds.KeyLiveMatches(int(index))

	if err := p.redis.Publish(nsrds.TopicLiveMatchesUpdate, rKey).Err(); err != nil {
		return err
	}

	return nil
}

func (p *Watcher) clearRedisIndices(currentIndex uint32) error {
	skipKey := nsrds.KeyLiveMatches(int(currentIndex))
	iter := p.redis.Scan(0, nsrds.KeysLiveMatchesPattern, 0).Iterator()

	var keys []string

	for iter.Next() {
		key := iter.Val()

		if key == skipKey {
			continue
		}

		keys = append(keys, key)
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	return p.redis.Del(keys...).Err()
}
