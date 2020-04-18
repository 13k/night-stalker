package redis

import (
	"context"

	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nscol "github.com/13k/night-stalker/internal/collections"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nslog "github.com/13k/night-stalker/internal/logger"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	nsvs "github.com/13k/night-stalker/internal/views"
)

type PubSubOptions struct {
	Log        *nslog.Logger
	Bus        *nsbus.Bus
	Redis      *nsrds.Redis
	DataLoader *nsdbda.Loader
}

type PubSub struct {
	options *PubSubOptions
	log     *nslog.Logger
	bus     *nsbus.Bus
	rds     *nsrds.Redis
	ctx     context.Context
	dbl     *nsdbda.Loader
}

func NewPubSub(options *PubSubOptions) *PubSub {
	return &PubSub{
		options: options,
		log:     options.Log.WithPackage("redis"),
		bus:     options.Bus,
		rds:     options.Redis,
		dbl:     options.DataLoader,
	}
}

func (s *PubSub) Start(ctx context.Context) error {
	s.ctx = ctx

	if err := s.watchLiveMatches(); err != nil {
		return err
	}

	if err := s.watchLiveMatchStats(); err != nil {
		return err
	}

	return nil
}

func (s *PubSub) watchLiveMatches() error {
	pubsub, err := s.rds.PSubscribe(nsrds.TopicPatternLiveMatchesAll)

	if err != nil {
		return xerrors.Errorf("error subscribing to live matches: %w", err)
	}

	go pubsub.Watch(s.ctx, s.handleMessage)

	return nil
}

func (s *PubSub) watchLiveMatchStats() error {
	pubsub, err := s.rds.PSubscribe(nsrds.TopicPatternLiveMatchStatsAll)

	if err != nil {
		return xerrors.Errorf("error subscribing to live matches: %w", err)
	}

	go pubsub.Watch(s.ctx, s.handleMessage)

	return nil
}

func (s *PubSub) handleMessage(rmsg *redis.Message) {
	l := s.log.WithOFields(
		"channel", rmsg.Channel,
		"pattern", rmsg.Pattern,
	)

	var op nspb.CollectionOp

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchesAdd:
		op = nspb.CollectionOp_COLLECTION_OP_ADD
	case nsrds.TopicLiveMatchesUpdate:
		op = nspb.CollectionOp_COLLECTION_OP_UPDATE
	case nsrds.TopicLiveMatchesRemove:
		op = nspb.CollectionOp_COLLECTION_OP_REMOVE
	case nsrds.TopicLiveMatchStatsAdd:
		op = nspb.CollectionOp_COLLECTION_OP_UPDATE
	default:
		l.Trace("ignored message from unknown channel")
		return
	}

	matchIDs, err := nscol.NewMatchIDsStrings(rmsg.Payload, ",")

	if err != nil {
		l.WithField("payload", rmsg.Payload).WithError(err).Error("error parsing payload")
		return
	}

	matchIDs = matchIDs.Unique()
	lenBefore := len(matchIDs)

	if op != nspb.CollectionOp_COLLECTION_OP_REMOVE {
		// messages may arrive out of order (for example, a delayed stats add after a match finished)
		matchIDs, err = s.filterFinished(matchIDs)

		if err != nil {
			l.WithError(err).Error("error filtering finished match IDs")
			return
		}
	}

	lenAfter := len(matchIDs)

	if lenAfter != lenBefore {
		l.WithOFields(
			"before", lenBefore,
			"after", lenAfter,
		).Trace("filtered match IDs")
	}

	if len(matchIDs) == 0 {
		l.Trace("ignored empty live matches change")
		return
	}

	l.WithField("count", lenAfter).Trace("received live matches change")

	var liveMatchesView *nspb.LiveMatches

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchesAdd:
		liveMatchesView, err = s.loadLiveMatchesView(matchIDs)
	case nsrds.TopicLiveMatchesUpdate:
		liveMatchesView, err = s.loadLiveMatchesView(matchIDs)
	case nsrds.TopicLiveMatchesRemove:
		liveMatchesView = nsvs.NewShallowLiveMatches(matchIDs)
	case nsrds.TopicLiveMatchStatsAdd:
		liveMatchesView, err = s.loadLiveMatchesView(matchIDs)
	}

	if err != nil {
		l.WithError(err).Error("error handling live matches change")
		return
	}

	if liveMatchesView == nil {
		l.Trace("ignored empty live matches view")
		return
	}

	view := &nspb.LiveMatchesChange{Op: op, Change: liveMatchesView}

	if err := s.busPubWebLiveMatchesChange(view); err != nil {
		l.WithError(err).Error("error publishing live matches change")
		return
	}

	l.WithOFields(
		"op", view.Op.String(),
		"match_ids", matchIDs.Join(","),
	).Trace("bus published live matches change")
}

func (s *PubSub) filterFinished(matchIDs nscol.MatchIDs) (nscol.MatchIDs, error) {
	if len(matchIDs) == 0 {
		return nil, nil
	}

	finishedMatchIDs, err := s.dbl.FindMatchIDs(s.ctx, nsdbda.MatchFilters{
		MatchIDs: matchIDs,
	})

	if err != nil {
		return nil, xerrors.Errorf("error finding match IDs: %w", err)
	}

	if len(finishedMatchIDs) == 0 {
		return matchIDs, nil
	}

	return matchIDs.Sub(finishedMatchIDs), nil
}

func (s *PubSub) loadLiveMatchesView(matchIDs nscol.MatchIDs) (*nspb.LiveMatches, error) {
	data, err := s.dbl.LiveMatchesData(s.ctx, &nsdbda.LiveMatchesParams{
		MatchIDs:         matchIDs,
		WithFollowedOnly: true,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading live matches data: %w", err)
	}

	view, err := nsvs.NewLiveMatches(data)

	if err != nil {
		return nil, xerrors.Errorf("error creating live matches view: %w", err)
	}

	return view, nil
}

func (s *PubSub) busPubWebLiveMatchesChange(view *nspb.LiveMatchesChange) error {
	var topic string

	switch view.Op {
	case nspb.CollectionOp_COLLECTION_OP_ADD:
		topic = nsbus.TopicWebLiveMatchesAdd
	case nspb.CollectionOp_COLLECTION_OP_UPDATE:
		topic = nsbus.TopicWebLiveMatchesUpdate
	case nspb.CollectionOp_COLLECTION_OP_REMOVE:
		topic = nsbus.TopicWebLiveMatchesRemove
	}

	err := s.bus.Pub(nsbus.Message{
		Topic:   topic,
		Payload: view,
	})

	if err != nil {
		return xerrors.Errorf("error publishing live matches change: %w", err)
	}

	return nil
}
