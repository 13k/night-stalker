package redis

import (
	"context"

	"github.com/go-redis/redis/v7"
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

type PubSub struct {
	*redis.PubSub
}

func NewPubSub(pubsub *redis.PubSub) *PubSub {
	return &PubSub{PubSub: pubsub}
}

func (ps *PubSub) Watch(ctx context.Context, handler func(*redis.Message)) {
	defer ps.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-ps.Channel():
			if !ok {
				return
			}

			go handler(msg)
		}
	}
}
