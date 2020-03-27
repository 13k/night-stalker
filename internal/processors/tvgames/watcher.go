package tvgames

import (
	"context"
	"time"

	"cirello.io/oversight"
	"github.com/jinzhu/gorm"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrt "github.com/13k/night-stalker/internal/runtime"
	"github.com/13k/night-stalker/models"
)

const (
	processorName = "tv_games"

	msgTypeFindTopSourceTVGames = d2pb.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames
)

type WatcherOptions struct {
	Log             *nslog.Logger
	Bus             *nsbus.Bus
	Interval        time.Duration
	ShutdownTimeout time.Duration
}

var _ nsproc.Processor = (*Watcher)(nil)

type Watcher struct {
	options           WatcherOptions
	discoveryPage     *discoveryPage
	ctx               context.Context
	log               *nslog.Logger
	db                *gorm.DB
	dota              *nsdota2.Client
	bus               *nsbus.Bus
	busSubTVGamesResp *nsbus.Subscription
}

func NewWatcher(options WatcherOptions) *Watcher {
	return &Watcher{
		options:       options,
		log:           options.Log.WithPackage(processorName),
		bus:           options.Bus,
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
		Start:    p.Start,
		Restart:  oversight.Transient(),
		Shutdown: shutdown,
	}
}

func (p *Watcher) Start(ctx context.Context) (err error) {
	defer nsrt.RecoverError(p.log, &err)

	err = p.start(ctx)

	if err != nil {
		p.handleError(err)
	}

	return err
}

func (p *Watcher) start(ctx context.Context) error {
	if err := p.setupContext(ctx); err != nil {
		return xerrors.Errorf("error setting up context: %w", err)
	}

	p.busSubscribe()

	return p.loop()
}

func (p *Watcher) stop(t *time.Ticker) {
	t.Stop()
	p.busUnsubscribe()
	p.ctx = nil
	p.log.Warn("stop")
}

func (p *Watcher) busSubscribe() {
	if p.busSubTVGamesResp == nil {
		p.busSubTVGamesResp = p.bus.Sub(nsbus.TopicGCDispatcherReceivedFindTopSourceTVGamesResponse)
	}
}

func (p *Watcher) busUnsubscribe() {
	if p.busSubTVGamesResp != nil {
		p.bus.Unsub(p.busSubTVGamesResp)
		p.busSubTVGamesResp = nil
	}
}

func (p *Watcher) setupContext(ctx context.Context) error {
	if p.db = nsctx.GetDB(ctx); p.db == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDatabase)
	}

	if p.dota = nsctx.GetDota(ctx); p.dota == nil {
		return xerrors.Errorf("processor context error: %w", nsproc.ErrProcessorContextDotaClient)
	}

	p.ctx = ctx

	return nil
}

func (p *Watcher) loop() error {
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
		case busmsg, ok := <-p.busSubTVGamesResp.C:
			if !ok {
				return xerrors.Errorf("bus error: %w", &nsbus.ErrSubscriptionExpired{
					Subscription: p.busSubTVGamesResp,
				})
			}

			if dspmsg, ok := busmsg.Payload.(*nsbus.GCDispatcherReceivedMessage); ok {
				if res, ok := dspmsg.Message.(*d2pb.CMsgGCToClientFindTopSourceTVGamesResponse); ok {
					p.handleFindTopSourceTVGamesResponse(res)
				}
			}
		}
	}
}

func (p *Watcher) tick() {
	if !p.dota.Session.IsReady() {
		return
	}

	if err := p.query(); err != nil {
		p.handleError(xerrors.Errorf("error querying tv games: %w", err))
	}
}

func (p *Watcher) query() error {
	if !p.dota.Session.IsReady() {
		return nil
	}

	if p.ctx.Err() != nil {
		return xerrors.Errorf("error querying: %w", p.ctx.Err())
	}

	page := &queryPage{
		start: p.discoveryPage.LastPageStart(),
	}

	if err := p.queryPage(page); err != nil {
		return xerrors.Errorf("error querying page: %w", err)
	}

	return nil
}

func (p *Watcher) queryPage(page *queryPage) error {
	if !p.dota.Session.IsReady() {
		return nil
	}

	if p.ctx.Err() != nil {
		return xerrors.Errorf("error querying page: %w", &errQueryPageFailure{
			Page: page,
			Err:  p.ctx.Err(),
		})
	}

	p.log.WithOFields(
		"index", page.index,
		"start", page.start,
	).Debug("requesting tv games")

	if err := p.busPubRequestTVGames(page); err != nil {
		return xerrors.Errorf("error querying page: %w", &errQueryPageFailure{
			Page: page,
			Err:  err,
		})
	}

	return nil
}

func (p *Watcher) handleFindTopSourceTVGamesResponse(msg *d2pb.CMsgGCToClientFindTopSourceTVGamesResponse) {
	go func() {
		if err := p.handleResponse(newQueryPage(msg)); err != nil {
			p.handleError(xerrors.Errorf("error handling response: %w", err))
		}
	}()
}

func (p *Watcher) handleResponse(page *queryPage) error {
	if page.psize == 0 || page.total == 0 {
		return xerrors.Errorf("empty response: %w", &errEmptyResponse{
			Page: page,
		})
	}

	l := p.log.WithOFields(
		"index", page.index,
		"start", page.start,
		"psize", page.psize,
		"total", page.total,
	)

	if p.discoveryPage.Empty() {
		l.Debug("received discovery response")

		p.discoveryPage.SetPage(page)

		if err := p.query(); err != nil {
			return xerrors.Errorf("error querying tv games: %w", err)
		}

		return nil
	}

	games := nscol.TVGames(page.res.GetGameList())
	cleaned := games.Clean()

	inProgress, err := p.filterFinished(cleaned)

	if err != nil {
		return xerrors.Errorf("error filtering finished tv games: %w", &errHandleResponseFailure{
			Page: page,
			Err:  err,
		})
	}

	liveMatches, err := p.saveGames(inProgress)

	if err != nil {
		return xerrors.Errorf("error saving tv games: %w", &errHandleResponseFailure{
			Page: page,
			Err:  err,
		})
	}

	deactivated := liveMatches.RemoveDeactivated()

	if len(liveMatches) > 0 {
		if err := p.busPubLiveMatchesAdd(liveMatches); err != nil {
			return xerrors.Errorf("error publishing live matches change: %w", &errHandleResponseFailure{
				Page: page,
				Err:  err,
			})
		}
	}

	if len(deactivated) > 0 {
		if err := p.busPubLiveMatchesRemove(deactivated); err != nil {
			return xerrors.Errorf("error publishing live matches change: %w", &errHandleResponseFailure{
				Page: page,
				Err:  err,
			})
		}
	}

	l.WithOFields(
		"games", len(games),
		"cleaned", len(games)-len(cleaned),
		"finished", len(cleaned)-len(inProgress),
		"in_progress", len(inProgress),
		"deactivated", len(deactivated),
		"live", len(liveMatches),
	).Debug("handled tv games")

	return nil
}

func (p *Watcher) saveGames(games nscol.TVGames) (nscol.LiveMatches, error) {
	liveMatches := make(nscol.LiveMatches, 0, len(games))

	for _, game := range games {
		liveMatch := models.LiveMatchDotaProto(game)
		errSave := &errSaveGameFailure{
			MatchID:  liveMatch.MatchID,
			ServerID: liveMatch.ServerID,
			LobbyID:  liveMatch.LobbyID,
		}

		if p.ctx.Err() != nil {
			errSave.Err = p.ctx.Err()
			return nil, xerrors.Errorf("error saving game: %w", errSave)
		}

		tx := p.db.Begin()

		dbres := tx.
			Where(models.LiveMatch{MatchID: liveMatch.MatchID}).
			Assign(liveMatch).
			FirstOrCreate(liveMatch)

		if dbres.Error != nil {
			tx.Rollback()

			errSave.Err = dbres.Error
			return nil, xerrors.Errorf("error saving game: %w", errSave)
		}

		for _, gamePlayer := range game.GetPlayers() {
			if p.ctx.Err() != nil {
				tx.Rollback()

				errSave.Err = p.ctx.Err()
				return nil, xerrors.Errorf("error saving game: %w", errSave)
			}

			livePlayer := models.NewLiveMatchPlayer(liveMatch, gamePlayer)

			criteria := &models.LiveMatchPlayer{
				LiveMatchID: livePlayer.LiveMatchID,
				AccountID:   livePlayer.AccountID,
			}

			dbres = tx.
				Where(criteria).
				Assign(livePlayer).
				FirstOrCreate(livePlayer)

			if dbres.Error != nil {
				tx.Rollback()

				errSave.Err = dbres.Error
				return nil, xerrors.Errorf("error saving game: %w", errSave)
			}
		}

		if err := tx.Commit().Error; err != nil {
			errSave.Err = err
			return nil, xerrors.Errorf("error saving game: %w", errSave)
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
		return nil, xerrors.Errorf("error filtering finished games: %w", err)
	}

	for _, matchID := range finishedMatchIDs {
		tvGames, _ = tvGames.RemoveByMatchID(matchID)
	}

	return tvGames, nil
}

func (p *Watcher) busPubRequestTVGames(page *queryPage) error {
	req := &d2pb.CMsgClientToGCFindTopSourceTVGames{
		GameListIndex: proto.Uint32(page.index),
		StartGame:     proto.Uint32(page.start),
	}

	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicGCDispatcherSend,
		Payload: &nsbus.GCDispatcherSendMessage{
			MsgType: msgTypeFindTopSourceTVGames,
			Message: req,
		},
	})
}

func (p *Watcher) busPubLiveMatchesAdd(liveMatches nscol.LiveMatches) error {
	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesAdd,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_COLLECTION_OP_ADD,
			Matches: liveMatches,
		},
	})
}

func (p *Watcher) busPubLiveMatchesRemove(liveMatches nscol.LiveMatches) error {
	return p.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicLiveMatchesRemove,
		Payload: &nsbus.LiveMatchesChangeMessage{
			Op:      nspb.CollectionOp_COLLECTION_OP_REMOVE,
			Matches: liveMatches,
		},
	})
}

func (p *Watcher) handleError(err error) {
	if e := (&errEmptyResponse{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"index", e.Page.index,
			"start", e.Page.start,
			"psize", e.Page.psize,
			"total", e.Page.total,
		).Warn("ignoring empty page")

		return
	}

	if e := (&errQueryPageFailure{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"index", e.Page.index,
			"start", e.Page.start,
			"psize", e.Page.psize,
			"total", e.Page.total,
		).WithError(e.Err).Error("error querying page")
	} else if e := (&errHandleResponseFailure{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"index", e.Page.index,
			"start", e.Page.start,
			"psize", e.Page.psize,
			"total", e.Page.total,
		).WithError(e.Err).Error("error handling response")
	} else if e := (&errSaveGameFailure{}); xerrors.As(err, &e) {
		p.log.WithOFields(
			"match_id", e.MatchID,
			"server_id", e.ServerID,
			"lobby_id", e.LobbyID,
		).WithError(e.Err).Error("error saving tv game")
	}

	p.log.Errorx(err)
}
