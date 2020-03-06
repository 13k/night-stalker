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

func (app *App) watchLiveMatches() error {
	pubsub, err := nsrds.PSubscribe(app.rds, nsrds.TopicPatternLiveMatchesAll)

	if err != nil {
		err = xerrors.Errorf("error subscribing to live matches: %w", err)
		return err
	}

	go nsrds.WatchPubsub(app.ctx, pubsub, app.handleLiveMatchesChange)

	return nil
}

func (app *App) handleLiveMatchesChange(rmsg *redis.Message) {
	l := app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	})

	matchIDs, err := nscol.NewMatchIDsFromString(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	if len(matchIDs) == 0 {
		return
	}

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchesAdd:
		err = app.handleLiveMatchesAdd(matchIDs)
	case nsrds.TopicLiveMatchesRemove:
		err = app.handleLiveMatchesRemove(matchIDs)
	default:
		return
	}

	if err != nil {
		l.WithError(err).Error("error handling live matches change")
		return
	}
}

func (app *App) handleLiveMatchesAdd(matchIDs nscol.MatchIDs) error {
	view, err := app.addLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error adding live matches: %w", err)
		return err
	}

	if view == nil {
		return nil
	}

	if err := app.busPubWebLiveMatchesAdd(view); err != nil {
		err = xerrors.Errorf("error publishing live matches add: %w", err)
		return err
	}

	return nil
}

func (app *App) handleLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	view, err := app.removeLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error removing live matches: %w", err)
		return err
	}

	if view == nil {
		return nil
	}

	if err := app.busPubWebLiveMatchesRemove(view); err != nil {
		err = xerrors.Errorf("error publishing live matches remove: %w", err)
		return err
	}

	return nil
}

func (app *App) addLiveMatches(matchIDs nscol.MatchIDs) (*nspb.LiveMatches, error) {
	liveMatches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return nil, err
	}

	if len(liveMatches) == 0 {
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
	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading LiveMatches view: %w", err)
		return nil, err
	}

	if len(view.Matches) == 0 {
		return nil, nil
	}

	app.matches.Remove(matchIDs...)

	return view, nil
}

func (app *App) watchLiveMatchStats() error {
	pubsub, err := nsrds.PSubscribe(app.rds, nsrds.TopicPatternLiveMatchStatsAll)

	if err != nil {
		err = xerrors.Errorf("error subscribing to live matches: %w", err)
		return err
	}

	go nsrds.WatchPubsub(app.ctx, pubsub, app.handleLiveMatchStatsChange)

	return nil
}

func (app *App) handleLiveMatchStatsChange(rmsg *redis.Message) {
	l := app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	})

	matchIDs, err := nscol.NewMatchIDsFromString(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	if len(matchIDs) == 0 {
		return
	}

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchStatsAdd:
		err = app.handleLiveMatchStatsAdd(matchIDs)
	default:
		return
	}

	if err != nil {
		l.WithError(err).Error("error handling live match stats change")
		return
	}
}

func (app *App) handleLiveMatchStatsAdd(matchIDs nscol.MatchIDs) error {
	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading live matches view: %w", err)
		return err
	}

	if len(view.Matches) == 0 {
		return nil
	}

	if err := app.busPubWebLiveMatchesUpdate(view); err != nil {
		err = xerrors.Errorf("error publishing live matches update: %w", err)
		return err
	}

	return nil
}

func (app *App) busPubWebLiveMatchesAdd(view *nspb.LiveMatches) error {
	return app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesAdd,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_ADD,
			Change: view,
		},
	})
}

func (app *App) busPubWebLiveMatchesRemove(view *nspb.LiveMatches) error {
	return app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesRemove,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_REMOVE,
			Change: view,
		},
	})
}

func (app *App) busPubWebLiveMatchesUpdate(view *nspb.LiveMatches) error {
	return app.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesUpdate,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_UPDATE,
			Change: view,
		},
	})
}
