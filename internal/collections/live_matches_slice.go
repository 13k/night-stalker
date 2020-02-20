package collections

import (
	"github.com/go-redis/redis/v7"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

// LiveMatchesSlice is a slice of LiveMatch.
//
// Methods based on https://github.com/golang/go/wiki/SliceTricks
type LiveMatchesSlice []*models.LiveMatch

func (s LiveMatchesSlice) Len() int      { return len(s) }
func (s LiveMatchesSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s *LiveMatchesSlice) Unshift(liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	*s = append(LiveMatchesSlice{liveMatch}, *s...)
}

func (s *LiveMatchesSlice) Push(liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	*s = append(*s, liveMatch)
}

func (s *LiveMatchesSlice) Insert(i int, liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	if i < 0 || i > len(*s) {
		return
	}

	if i == 0 {
		s.Unshift(liveMatch)
		return
	}

	if i == len(*s) {
		s.Push(liveMatch)
		return
	}

	*s = append(*s, nil)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = liveMatch
}

func (s *LiveMatchesSlice) Shift() (liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	if len(*s) < 1 {
		return
	}

	liveMatch, *s = (*s)[0], (*s)[1:]

	return
}

func (s *LiveMatchesSlice) Pop() (liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	if len(*s) < 1 {
		return
	}

	liveMatch, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return
}

func (s *LiveMatchesSlice) Remove(i int) *models.LiveMatch {
	if s == nil {
		return nil
	}

	if i < 0 || i >= len(*s) {
		return nil
	}

	if i == 0 {
		return s.Shift()
	}

	if i == len(*s)-1 {
		return s.Pop()
	}

	liveMatch := (*s)[i]

	copy((*s)[i:], (*s)[i+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]

	return liveMatch
}

func (s *LiveMatchesSlice) RemoveByMatchID(matchID nspb.MatchID) *models.LiveMatch {
	if s == nil {
		return nil
	}

	return s.Remove(s.FindIndexByMatchID(matchID))
}

func (s *LiveMatchesSlice) RemoveDeactivated() LiveMatchesSlice {
	if s == nil {
		return nil
	}

	if *s == nil {
		return nil
	}

	removed := LiveMatchesSlice{}
	var indices []int

	for i, liveMatch := range *s {
		if liveMatch.DeactivateTime != nil {
			// items will be sequentially removed below, so account for previously removed indices
			indices = append(indices, i-len(indices))
		}
	}

	for _, i := range indices {
		removed = append(removed, s.Remove(i))
	}

	return removed
}

func (s LiveMatchesSlice) FindIndexByMatchID(matchID nspb.MatchID) int {
	for i, liveMatch := range s {
		if liveMatch.MatchID == matchID {
			return i
		}
	}

	return -1
}

func (s LiveMatchesSlice) Batches(batchSize int) []LiveMatchesSlice {
	if s == nil {
		return nil
	}

	if len(s) == 0 {
		return []LiveMatchesSlice{}
	}

	liveMatches := s
	batches := make([]LiveMatchesSlice, 0, (len(liveMatches)+batchSize-1)/batchSize)

	for batchSize < len(liveMatches) {
		liveMatches, batches = liveMatches[batchSize:], append(batches, liveMatches[0:batchSize:batchSize])
	}

	batches = append(batches, liveMatches)

	return batches
}

func (s LiveMatchesSlice) MatchIDs() MatchIDs {
	matchIDs := make(MatchIDs, len(s))

	for i, liveMatch := range s {
		matchIDs[i] = liveMatch.MatchID
	}

	return matchIDs
}

func (s LiveMatchesSlice) ToRedisZValues() []*redis.Z {
	result := make([]*redis.Z, len(s))

	for i, liveMatch := range s {
		result[i] = &redis.Z{
			Score:  liveMatch.SortScore,
			Member: liveMatch.MatchID,
		}
	}

	return result
}
