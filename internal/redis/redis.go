package redis

import (
	"github.com/go-redis/redis/v7"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nserr "github.com/13k/night-stalker/internal/errors"
)

type Redis struct {
	*redis.Client
}

func NewRedis(c *redis.Client) *Redis {
	return &Redis{Client: c}
}

func (r *Redis) PSubscribe(pattern string) (*PubSub, error) {
	pubsub := r.Client.PSubscribe(pattern)
	msg, err := pubsub.Receive()

	if err != nil {
		err = xerrors.Errorf("error subscribing to topic %s: %w", pattern, err)
		return nil, err
	}

	switch m := msg.(type) {
	case *redis.Subscription:
		break
	case error:
		err = xerrors.Errorf("error subscribing to topic %s: %w", pattern, m)
		return nil, err
	default:
		err = xerrors.Errorf("received invalid message %T when subscribing to topic %s", m, pattern)
		return nil, err
	}

	return NewPubSub(pubsub), nil
}

func (r *Redis) LiveMatchIDs() (nscol.MatchIDs, error) {
	rcmd := r.ZRevRange(KeyLiveMatches, 0, -1)

	if err := rcmd.Err(); err != nil {
		return nil, &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatches,
			Err: nserr.Wrap("error fetching cached live matches IDs", err),
		}
	}

	matchIDs := make([]uint64, len(rcmd.Val()))

	if err := rcmd.ScanSlice(&matchIDs); err != nil {
		return nil, &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatches,
			Err: nserr.Wrap("error parsing live match IDs", err),
		}
	}

	return nscol.NewMatchIDs(matchIDs...), nil
}

func (r *Redis) AddLiveMatches(liveMatches nscol.LiveMatches) error {
	zValues := LiveMatchesToZValues(liveMatches)
	rcmd := r.ZAdd(KeyLiveMatches, zValues...)

	if err := rcmd.Err(); err != nil {
		return &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatches,
			Err: nserr.Wrap("error adding live matches IDs", err),
		}
	}

	zValues = LiveMatchesToZValuesByTime(liveMatches)
	rcmd = r.ZAdd(KeyLiveMatchesByTime, zValues...)

	if err := rcmd.Err(); err != nil {
		return &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatchesByTime,
			Err: nserr.Wrap("error adding live matches IDs", err),
		}
	}

	return nil
}

func (r *Redis) RemoveLiveMatches(matchIDs nscol.MatchIDs) error {
	ifaceMatchIDs := matchIDs.ToUint64Interfaces()

	rcmd := r.ZRem(KeyLiveMatches, ifaceMatchIDs...)

	if err := rcmd.Err(); err != nil {
		return &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatches,
			Err: nserr.Wrap("error removing live matches IDs", err),
		}
	}

	rcmd = r.ZRem(KeyLiveMatchesByTime, ifaceMatchIDs...)

	if err := rcmd.Err(); err != nil {
		return &ErrCommandFailure{
			Cmd: rcmd,
			Key: KeyLiveMatchesByTime,
			Err: nserr.Wrap("error removing live matches IDs", err),
		}
	}

	return nil
}

func (r *Redis) PubLiveMatchesAdd(liveMatches nscol.LiveMatches) error {
	if len(liveMatches) == 0 {
		return nil
	}

	rcmd := r.Publish(TopicLiveMatchesAdd, liveMatches.MatchIDs().Join(","))

	if err := rcmd.Err(); err != nil {
		return &ErrPubsubFailure{
			Cmd:   rcmd,
			Topic: TopicLiveMatchesAdd,
			Err:   nserr.Wrap("error publishing live matches change", err),
		}
	}

	return nil
}

func (r *Redis) PubLiveMatchesRemove(matchIDs nscol.MatchIDs) error {
	if len(matchIDs) == 0 {
		return nil
	}

	rcmd := r.Publish(TopicLiveMatchesRemove, matchIDs.Join(","))

	if err := rcmd.Err(); err != nil {
		return &ErrPubsubFailure{
			Cmd:   rcmd,
			Topic: TopicLiveMatchesRemove,
			Err:   nserr.Wrap("error publishing live matches removal", err),
		}
	}

	return nil
}

func (r *Redis) PubLiveMatchStatsAdd(stats nscol.LiveMatchStats) error {
	if len(stats) == 0 {
		return nil
	}

	rcmd := r.Publish(TopicLiveMatchStatsAdd, stats.MatchIDs().Join(","))

	if err := rcmd.Err(); err != nil {
		return &ErrPubsubFailure{
			Cmd:   rcmd,
			Topic: TopicLiveMatchStatsAdd,
			Err:   nserr.Wrap("error publishing live match stats adding", err),
		}
	}

	return nil
}
