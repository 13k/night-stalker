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

	go pubsub.Watch(s.ctx, s.handleLiveMatchesChange)

	return nil
}

func (s *PubSub) handleLiveMatchesChange(rmsg *redis.Message) {
	l := s.log.WithOFields(
		"channel", rmsg.Channel,
		"pattern", rmsg.Pattern,
		"payload", rmsg.Payload,
	)

	matchIDs, err := nscol.NewMatchIDsStrings(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	if len(matchIDs) == 0 {
		l.Trace("ignored empty live matches change")
		return
	}

	l.Trace("received live matches change")

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchesAdd:
		err = s.handleLiveMatchesAdd(matchIDs)
	case nsrds.TopicLiveMatchesRemove:
		err = s.handleLiveMatchesRemove(matchIDs)
	default:
		return
	}

	if err != nil {
		l.WithError(err).Error("error handling live matches change")
		return
	}
}

func (s *PubSub) handleLiveMatchesAdd(matchIDs nscol.MatchIDs) error {
	view, err := s.loadLiveMatchesView(matchIDs)

	if err != nil {
		return xerrors.Errorf("error adding live matches: %w", err)
	}

	if view == nil {
		return nil
	}

	if err := s.busPubWebLiveMatchesAdd(view); err != nil {
		return xerrors.Errorf("error publishing live matches add: %w", err)
	}

	return nil
}

func (s *PubSub) handleLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	view := nsvs.NewShallowLiveMatches(matchIDs)

	if view == nil {
		return nil
	}

	if err := s.busPubWebLiveMatchesRemove(view); err != nil {
		return xerrors.Errorf("error publishing live matches remove: %w", err)
	}

	return nil
}

func (s *PubSub) watchLiveMatchStats() error {
	pubsub, err := s.rds.PSubscribe(nsrds.TopicPatternLiveMatchStatsAll)

	if err != nil {
		return xerrors.Errorf("error subscribing to live matches: %w", err)
	}

	go pubsub.Watch(s.ctx, s.handleLiveMatchStatsChange)

	return nil
}

func (s *PubSub) handleLiveMatchStatsChange(rmsg *redis.Message) {
	l := s.log.WithOFields(
		"channel", rmsg.Channel,
		"pattern", rmsg.Pattern,
		"payload", rmsg.Payload,
	)

	matchIDs, err := nscol.NewMatchIDsStrings(rmsg.Payload, ",")

	if err != nil {
		l.WithError(err).Error("error parsing payload")
		return
	}

	if len(matchIDs) == 0 {
		l.Trace("ignored empty live match stats change")
		return
	}

	l.Trace("received live match stats change")

	switch rmsg.Channel {
	case nsrds.TopicLiveMatchStatsAdd:
		err = s.handleLiveMatchStatsAdd(matchIDs)
	default:
		return
	}

	if err != nil {
		l.WithError(err).Error("error handling live match stats change")
		return
	}
}

func (s *PubSub) handleLiveMatchStatsAdd(matchIDs nscol.MatchIDs) error {
	view, err := s.loadLiveMatchesView(matchIDs)

	if err != nil {
		return xerrors.Errorf("error loading live matches view: %w", err)
	}

	if view == nil {
		return nil
	}

	if err := s.busPubWebLiveMatchesUpdate(view); err != nil {
		return xerrors.Errorf("error publishing live matches update: %w", err)
	}

	return nil
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

func (s *PubSub) busPubWebLiveMatchesAdd(view *nspb.LiveMatches) error {
	err := s.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesAdd,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_ADD,
			Change: view,
		},
	})

	if err != nil {
		return xerrors.Errorf("error publishing live matches add: %w", err)
	}

	s.log.WithOFields().Trace("published live matches add")

	return nil
}

func (s *PubSub) busPubWebLiveMatchesRemove(view *nspb.LiveMatches) error {
	err := s.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesRemove,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_REMOVE,
			Change: view,
		},
	})

	if err != nil {
		return xerrors.Errorf("error publishing live matches remove: %w", err)
	}

	s.log.WithOFields().Trace("published live matches remove")

	return nil
}

func (s *PubSub) busPubWebLiveMatchesUpdate(view *nspb.LiveMatches) error {
	err := s.bus.Pub(nsbus.Message{
		Topic: nsbus.TopicWebLiveMatchesUpdate,
		Payload: &nspb.LiveMatchesChange{
			Op:     nspb.CollectionOp_COLLECTION_OP_UPDATE,
			Change: view,
		},
	})

	if err != nil {
		return xerrors.Errorf("error publishing live matches update: %w", err)
	}

	s.log.WithOFields().Trace("published live matches update")

	return nil
}
