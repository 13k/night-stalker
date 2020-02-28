package redis

import (
	"github.com/go-redis/redis/v7"

	nscol "github.com/13k/night-stalker/internal/collections"
)

func LiveMatchesToZValues(liveMatches nscol.LiveMatches) []*redis.Z {
	zValues := make([]*redis.Z, len(liveMatches))

	for i, liveMatch := range liveMatches {
		zValues[i] = &redis.Z{
			Score:  liveMatch.SortScore,
			Member: liveMatch.MatchID,
		}
	}

	return zValues
}

func LiveMatchesToZValuesByTime(liveMatches nscol.LiveMatches) []*redis.Z {
	zValues := make([]*redis.Z, len(liveMatches))

	for i, liveMatch := range liveMatches {
		var activateTimeUnix int64

		if liveMatch.ActivateTime != nil {
			activateTimeUnix = liveMatch.ActivateTime.UTC().Unix()
		}

		zValues[i] = &redis.Z{
			Member: liveMatch.MatchID,
			Score:  float64(activateTimeUnix),
		}
	}

	return zValues
}
