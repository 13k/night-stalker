package web

import (
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
)

func (app *App) rdsLiveMatchIDs() (nscol.MatchIDs, error) {
	result := app.rds.ZRevRange(nsrds.KeyLiveMatches, 0, -1)

	if err := result.Err(); err != nil {
		err = xerrors.Errorf("error loading redis key %s: %w", nsrds.KeyLiveMatches, err)
		return nil, err
	}

	matchIDs := make(nscol.MatchIDs, len(result.Val()))

	if err := result.ScanSlice(&matchIDs); err != nil {
		err = xerrors.Errorf("error parsing live matches IDs: %w", err)
		return nil, err
	}

	return matchIDs, nil
}

func (app *App) seedLiveMatches() error {
	if app.matches != nil {
		return nil
	}

	matchIDs, err := app.rdsLiveMatchIDs()

	if err != nil {
		err = xerrors.Errorf("error loading live matches IDs: %w", err)
		return err
	}

	if len(matchIDs) == 0 {
		app.matches = nscol.NewLiveMatchesContainer()
		return nil
	}

	liveMatches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return err
	}

	app.matches = nscol.NewLiveMatchesContainer(liveMatches...)

	app.log.WithField("count", len(liveMatches)).Debug("seeded matches")

	return nil
}

func (app *App) subscribeLiveMatches() error {
	app.rdsSubLiveMatchesAll = app.rds.PSubscribe(nsrds.TopicPatternLiveMatchesAll)

	msg, err := app.rdsSubLiveMatchesAll.Receive()

	if err != nil {
		err = xerrors.Errorf("error subscribing to topic %s: %w", nsrds.TopicPatternLiveMatchesAll, err)
		return err
	}

	switch m := msg.(type) {
	case *redis.Subscription:
	case *redis.Pong:
	case *redis.Message:
		go app.handleLiveMatchesChange(m)
	default:
		return xerrors.Errorf(
			"received invalid message %T when subscribing to topic %s",
			m,
			nsrds.TopicPatternLiveMatchesAll,
		)
	}

	go app.watchLiveMatches()

	return nil
}

func (app *App) watchLiveMatches() {
	defer func() {
		app.rdsSubLiveMatchesAll.Close()
		app.log.Warn("stopped watching live matches")
	}()

	for {
		select {
		case <-app.ctx.Done():
			return
		case msg, ok := <-app.rdsSubLiveMatchesAll.Channel():
			if !ok {
				return
			}

			go app.handleLiveMatchesChange(msg)
		}
	}
}

func (app *App) handleLiveMatchesChange(rmsg *redis.Message) {
	l := app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	})

	l.Debug("received live matches change")

	matchIDs, err := nscol.NewMatchIDsFromString(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchesAdd:
		app.handleLiveMatchesAdd(matchIDs)
	case nsrds.TopicLiveMatchesRemove:
		app.handleLiveMatchesRemove(matchIDs)
	default:
		l.Warn("ignoring live matches change")
		return
	}
}

func (app *App) handleLiveMatchesAdd(matchIDs nscol.MatchIDs) {
	view, err := app.addLiveMatches(matchIDs)

	if err != nil || view == nil {
		return
	}

	app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesAdd,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_ADD,
			Change: view,
		},
	})
}

func (app *App) handleLiveMatchesRemove(matchIDs nscol.MatchIDs) {
	view, err := app.removeLiveMatches(matchIDs)

	if err != nil || view == nil {
		return
	}

	app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesRemove,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_REMOVE,
			Change: view,
		},
	})
}

func (app *App) addLiveMatches(matchIDs nscol.MatchIDs) (*nspb.LiveMatches, error) {
	l := app.log.WithField("match_ids", matchIDs)

	liveMatches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return nil, err
	}

	if len(liveMatches) == 0 {
		l.Debug("ignoring empty live matches")
		return nil, nil
	}

	view, err := app.createLiveMatchesView(liveMatches...)

	if err != nil {
		err = xerrors.Errorf("error creating LiveMatches view: %w", err)
		return nil, err
	}

	app.matches.Add(liveMatches...)

	return view, nil
}

func (app *App) removeLiveMatches(matchIDs nscol.MatchIDs) (*nspb.LiveMatches, error) {
	l := app.log.WithField("match_ids", matchIDs)

	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading LiveMatches view: %w", err)
		return nil, err
	}

	if len(view.Matches) == 0 {
		l.Debug("ignoring empty live matches")
		return nil, nil
	}

	app.matches.Remove(matchIDs...)

	return view, nil
}

func (app *App) subscribeLiveMatchStats() error {
	app.rdsSubLiveMatchStatsAll = app.rds.PSubscribe(nsrds.TopicPatternLiveMatchStatsAll)

	msg, err := app.rdsSubLiveMatchStatsAll.Receive()

	if err != nil {
		err = xerrors.Errorf("error subscribing to topic %s: %w", nsrds.TopicPatternLiveMatchStatsAll, err)
		return err
	}

	switch m := msg.(type) {
	case *redis.Subscription:
	case *redis.Pong:
	case *redis.Message:
		go app.handleLiveMatchStatsChange(m)
	default:
		return xerrors.Errorf(
			"received invalid message %T when subscribing to topic %s",
			m,
			nsrds.TopicPatternLiveMatchStatsAll,
		)
	}

	go app.watchLiveMatchStats()

	return nil
}

func (app *App) watchLiveMatchStats() {
	defer func() {
		app.rdsSubLiveMatchStatsAll.Close()
		app.log.Warn("stopped watching match stats updates")
	}()

	for {
		select {
		case <-app.ctx.Done():
			return
		case msg, ok := <-app.rdsSubLiveMatchStatsAll.Channel():
			if !ok {
				return
			}

			go app.handleLiveMatchStatsChange(msg)
		}
	}
}

func (app *App) handleLiveMatchStatsChange(rmsg *redis.Message) {
	l := app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	})

	l.Debug("received live match stats change")

	matchIDs, err := nscol.NewMatchIDsFromString(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchStatsAdd:
		app.handleLiveMatchStatsAdd(matchIDs)
	default:
		l.Warn("ignoring live match stats change")
		return
	}
}

func (app *App) handleLiveMatchStatsAdd(matchIDs nscol.MatchIDs) {
	l := app.log.WithField("match_ids", matchIDs)

	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		l.WithError(err).Error("error loading live matches view")
		return
	}

	if len(view.Matches) == 0 {
		l.Debug("ignoring empty live matches view")
		return
	}

	app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesUpdate,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_UPDATE,
			Change: view,
		},
	})
}
