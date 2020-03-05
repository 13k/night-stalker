package tvgames

import (
	"context"
	"runtime/debug"
	"time"

	"cirello.io/oversight"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
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
	processorName = "tv_games"

	msgTypeFindTopSourceTVGames = protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames
)

type WatcherOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Watcher)(nil)

type Watcher struct {
	options        WatcherOptions
	discoveryPage  *discoveryPage
	ctx            context.Context
	log            *nslog.Logger
	db             *gorm.DB
	bus            *nsbus.Bus
	busTVGamesResp <-chan nsbus.Message
}

func NewWatcher(options WatcherOptions) *Watcher {
	p := &Watcher{
		options:       options,
		log:           options.Log.WithPackage(processorName),
		bus:           options.Bus,
		discoveryPage: &discoveryPage{},
	}

	p.busSubscribe()

	return p
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
	if err := p.setupContext(ctx); err != nil {
		return err
	}

	p.busSubscribe()

	return p.loop()
}

func (p *Watcher) busSubscribe() {
	if p.busTVGamesResp == nil {
		p.busTVGamesResp = p.bus.Sub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse)
	}
}

func (p *Watcher) busUnsubscribe() {
	if p.busTVGamesResp != nil {
		p.bus.Unsub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse, p.busTVGamesResp)
		p.busTVGamesResp = nil
	}
}

func (p *Watcher) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return nsproc.ErrProcessorContextDatabase
	}

	p.ctx = ctx

	return nil
}

func (p *Watcher) loop() error {
	defer func() {
		if err := recover(); err != nil {
			p.log.WithField("error", err).Error("recovered panic")
			p.log.Error(string(debug.Stack()))
		}
	}()

	t := time.NewTicker(p.options.Interval)

	defer p.stop(t)

	p.log.Info("start")
	p.tick()

	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-t.C:
			p.tick()
		case busmsg, ok := <-p.busTVGamesResp:
			if !ok {
				return nil
			}

			if dspmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dspmsg.Message.(*protocol.CMsgGCToClientFindTopSourceTVGamesResponse); ok {
					p.handleFindTopSourceTVGamesResponse(res)
				}
			}
		}
	}
}

func (p *Watcher) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.log.Warn("stop")
}

func (p *Watcher) tick() {
	if err := p.query(); err != nil {
		p.log.WithError(err).Error("error querying tv games")
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

	p.log.WithFields(logrus.Fields{
		"index": page.index,
		"start": page.start,
	}).Debug("requesting tv games")

	p.busPubRequestTVGames(page)

	return nil
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

	if page.psize == 0 || page.total == 0 {
		l.Warn("ignoring empty page")
		return
	}

	if p.discoveryPage.Empty() {
		l.Debug("received discovery response")

		p.discoveryPage.SetPage(page)

		if err := p.query(); err != nil {
			l.WithError(err).Error("error requesting tv games")
			return
		}

		return
	}

	games := nscol.TVGames(page.res.GetGameList())
	originalLen := len(games)
	games = games.Clean()
	cleanLen := len(games)
	cleanCount := originalLen - cleanLen
	games, err := p.filterFinished(games)

	if err != nil {
		l.WithError(err).Error("error filtering finished")
		return
	}

	finalLen := len(games)
	finishedCount := cleanLen - finalLen
	l = l.WithField("count", finalLen)

	if originalLen != finalLen {
		l.WithFields(logrus.Fields{
			"cleaned":  cleanCount,
			"finished": finishedCount,
		}).Debug("filtered tv games")
	}

	l.Debug("received tv games")

	liveMatches, err := p.saveGames(games)

	if err != nil {
		l.WithError(err).Error("error saving response")
		return
	}

	deactivated := liveMatches.RemoveDeactivated()

	if len(liveMatches) > 0 {
		p.busPubLiveMatchesAdd(liveMatches)
	}

	if len(deactivated) > 0 {
		p.busPubLiveMatchesRemove(deactivated)
	}
}

func (p *Watcher) saveGames(games nscol.TVGames) (nscol.LiveMatches, error) {
	liveMatches := make(nscol.LiveMatches, 0, len(games))

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

			livePlayer := models.NewLiveMatchPlayer(liveMatch, gamePlayer)

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

func (p *Watcher) filterFinished(tvGames nscol.TVGames) (nscol.TVGames, error) {
	matchIDs := tvGames.MatchIDs()
	var finishedMatchIDs nscol.MatchIDs

	err := p.db.
		Model(models.MatchModel).
		Where("id IN (?)", matchIDs).
		Pluck("id", &finishedMatchIDs).
		Error

	if err != nil {
		return nil, err
	}

	for _, matchID := range finishedMatchIDs {
		tvGames, _ = tvGames.RemoveByMatchID(matchID)
	}

	return tvGames, nil
}

func (p *Watcher) busPubRequestTVGames(page *queryPage) {
	req := &protocol.CMsgClientToGCFindTopSourceTVGames{
		GameListIndex: proto.Uint32(page.index),
		StartGame:     proto.Uint32(page.start),
	}

	p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicGCDispatcherSend,
		Payload: &nsbus.GCDispatcherSendMessage{
			MsgType: msgTypeFindTopSourceTVGames,
			Message: req,
		},
	})

}

func (p *Watcher) busPubLiveMatchesAdd(liveMatches nscol.LiveMatches) {
	p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesAdd,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_COLLECTION_OP_ADD,
			Matches: liveMatches,
		},
	})
}

func (p *Watcher) busPubLiveMatchesRemove(liveMatches nscol.LiveMatches) {
	p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesRemove,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_COLLECTION_OP_REMOVE,
			Matches: liveMatches,
		},
	})
}
