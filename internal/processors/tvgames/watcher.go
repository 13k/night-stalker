package tvgames

import (
	"context"
	"fmt"
	"time"

	"cirello.io/oversight"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/proto"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsctx "github.com/13k/night-stalker/internal/context"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nserr "github.com/13k/night-stalker/internal/errors"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsproc "github.com/13k/night-stalker/internal/processors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrt "github.com/13k/night-stalker/internal/runtime"
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
	db                *nsdb.DB
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
		return &errQueryPageFailure{
			Page: page,
			Err:  nserr.Wrap("error querying page", p.ctx.Err()),
		}
	}

	p.log.WithOFields(
		"index", page.index,
		"start", page.start,
	).Debug("requesting tv games")

	if err := p.busPubRequestTVGames(page); err != nil {
		return &errQueryPageFailure{
			Page: page,
			Err:  nserr.Wrap("error querying page", err),
		}
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
		return &errHandleResponseFailure{
			Page: page,
			Err:  nserr.Wrap("error filtering finished tv games", err),
		}
	}

	liveMatches, err := p.saveGames(inProgress)

	if err != nil {
		return &errHandleResponseFailure{
			Page: page,
			Err:  nserr.Wrap("error saving tv games", err),
		}
	}

	deactivated := liveMatches.RemoveDeactivated()

	if len(liveMatches) > 0 {
		if err := p.busPubLiveMatchesAdd(liveMatches); err != nil {
			return &errHandleResponseFailure{
				Page: page,
				Err:  nserr.Wrap("error publishing live matches change", err),
			}
		}
	}

	if len(deactivated) > 0 {
		if err := p.busPubLiveMatchesRemove(deactivated); err != nil {
			return &errHandleResponseFailure{
				Page: page,
				Err:  nserr.Wrap("error publishing live matches change", err),
			}
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
	dbs := nsdbda.NewSaver(p.db)
	liveMatches := make(nscol.LiveMatches, len(games))

	for i, game := range games {
		liveMatch, err := dbs.UpsertLiveMatchProto(p.ctx, game)

		if err != nil {
			return nil, &errSaveGameFailure{
				MatchID:  nspb.MatchID(game.GetMatchId()),
				ServerID: nspb.SteamID(game.GetServerSteamId()),
				LobbyID:  nspb.LobbyID(game.GetLobbyId()),
				Err:      nserr.Wrap("error saving live match", err),
			}
		}

		liveMatches[i] = liveMatch
	}

	return liveMatches, nil
}

func (p *Watcher) filterFinished(tvGames nscol.TVGames) (nscol.TVGames, error) {
	dbl := nsdbda.NewLoader(p.db)

	finishedMatchIDs, err := dbl.FindMatchIDs(p.ctx, nsdbda.MatchFilters{
		MatchIDs: tvGames.MatchIDs(),
	})

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

	msg := fmt.Sprintf("%s error", processorName)
	l := p.log

	if e := (&errQueryPageFailure{}); xerrors.As(err, &e) {
		msg = "error querying page"
		l = l.WithOFields(
			"index", e.Page.index,
			"start", e.Page.start,
			"psize", e.Page.psize,
			"total", e.Page.total,
		)
	} else if e := (&errHandleResponseFailure{}); xerrors.As(err, &e) {
		msg = "error handling response"
		l = l.WithOFields(
			"index", e.Page.index,
			"start", e.Page.start,
			"psize", e.Page.psize,
			"total", e.Page.total,
		)
	} else if e := (&errSaveGameFailure{}); xerrors.As(err, &e) {
		msg = "error saving tv game"
		l = l.WithOFields(
			"match_id", e.MatchID,
			"server_id", e.ServerID,
			"lobby_id", e.LobbyID,
		)
	}

	l.WithError(err).Error(msg)
}
