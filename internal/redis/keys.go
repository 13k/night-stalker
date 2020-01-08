package redis

import (
	"fmt"
	"strconv"
)

const (
	KeyLiveMatchesIndex    = "live_matches_index"
	KeysLiveMatchesPattern = "live_matches:*"
	fmtKeyLiveMatches      = "live_matches:%d"
)

func KeyLiveMatches(index int) string {
	return fmt.Sprintf(fmtKeyLiveMatches, index)
}

func KeyLiveMatchesString(index string) (string, error) {
	i, err := ParseKeyLiveMatchesIndex(index)

	if err != nil {
		return "", err
	}

	return KeyLiveMatches(i), nil
}

func ParseKeyLiveMatchesIndex(index string) (int, error) {
	i, err := strconv.ParseInt(index, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(i), nil
}
