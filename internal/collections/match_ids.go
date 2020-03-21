package collections

import (
	"strconv"
	"strings"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type MatchIDs []nspb.MatchID

func NewMatchIDs(s ...uint64) MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, id := range s {
		matchIDs[i] = nspb.MatchID(id)
	}

	return matchIDs
}

func NewMatchIDsFromString(s, sep string) (MatchIDs, error) {
	if len(s) == 0 {
		return nil, nil
	}

	ss := strings.Split(s, sep)
	matchIDs := make(MatchIDs, len(ss))

	for i, idStr := range ss {
		matchID, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			return nil, err
		}

		matchIDs[i] = nspb.MatchID(matchID)
	}

	return matchIDs, nil
}

func (s MatchIDs) AddUnique(ids ...nspb.MatchID) MatchIDs {
	if len(ids) == 0 {
		return s
	}

	unique := make(map[nspb.MatchID]bool)

	for _, sid := range s {
		unique[sid] = true
	}

	for _, id := range ids {
		if !unique[id] {
			s = append(s, id)
			unique[id] = true
		}
	}

	return s
}

func (s MatchIDs) Join(sep string) string {
	if len(s) == 0 {
		return ""
	}

	var b strings.Builder

	for i, id := range s {
		if i > 0 {
			b.WriteString(sep)
		}

		b.WriteString(strconv.FormatUint(uint64(id), 10))
	}

	return b.String()
}

func (s MatchIDs) ToUint64s() []uint64 {
	if s == nil {
		return nil
	}

	result := make([]uint64, len(s))

	for i, matchID := range s {
		result[i] = uint64(matchID)
	}

	return result
}

func (s MatchIDs) ToUint64Interfaces() []interface{} {
	if s == nil {
		return nil
	}

	result := make([]interface{}, len(s))

	for i, matchID := range s {
		result[i] = uint64(matchID)
	}

	return result
}
