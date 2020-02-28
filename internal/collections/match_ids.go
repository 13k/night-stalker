package collections

import (
	"strconv"
	"strings"

	nspb "github.com/13k/night-stalker/internal/protocol"
)

type MatchIDs []nspb.MatchID

func NewMatchIDsFromString(s, sep string) (MatchIDs, error) {
	if len(s) == 0 {
		return nil, nil
	}

	var err error

	ss := strings.Split(s, sep)
	matchIDs := make(MatchIDs, len(ss))

	for i, idStr := range ss {
		matchIDs[i], err = strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			return nil, err
		}
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

		b.WriteString(strconv.FormatUint(id, 10))
	}

	return b.String()
}

func (s MatchIDs) ToInterfaces() []interface{} {
	if len(s) == 0 {
		return nil
	}

	result := make([]interface{}, len(s))

	for i, matchID := range s {
		result[i] = matchID
	}

	return result
}
