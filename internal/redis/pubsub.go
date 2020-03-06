package redis

import (
	"context"

	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"
)

const (
	TopicLiveMatchesReplace       = "live_matches.replace"
	TopicLiveMatchesAdd           = "live_matches.add"
	TopicLiveMatchesUpdate        = "live_matches.update"
	TopicLiveMatchesRemove        = "live_matches.remove"
	TopicPatternLiveMatchesAll    = "live_matches.*"
	TopicLiveMatchStatsAdd        = "live_match_stats.add"
	TopicPatternLiveMatchStatsAll = "live_match_stats.*"
)

func PSubscribe(rds *redis.Client, topic string) (*redis.PubSub, error) {
	pubsub := rds.PSubscribe(topic)
	msg, err := pubsub.Receive()

	if err != nil {
		err = xerrors.Errorf("error subscribing to topic %s: %w", topic, err)
		return nil, err
	}

	switch m := msg.(type) {
	case *redis.Subscription:
		return pubsub, nil
	case error:
		err = xerrors.Errorf("error subscribing to topic %s: %w", topic, m)
		return nil, err
	default:
		err = xerrors.Errorf("received invalid message %T when subscribing to topic %s", m, topic)
		return nil, err
	}
}

func WatchPubsub(ctx context.Context, pubsub *redis.PubSub, handler func(*redis.Message)) {
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-pubsub.Channel():
			if !ok {
				return
			}

			go handler(msg)
		}
	}
}
