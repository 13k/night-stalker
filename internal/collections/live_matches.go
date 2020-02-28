package collections

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatches []*models.LiveMatch

func (s LiveMatches) Len() int      { return len(s) }
func (s LiveMatches) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s LiveMatches) MatchIDs() MatchIDs {
	if s == nil {
		return nil
	}

	matchIDs := make(MatchIDs, len(s))

	for i, liveMatch := range s {
		matchIDs[i] = liveMatch.MatchID
	}

	return matchIDs
}

func (s LiveMatches) KeyByMatchID() map[nspb.MatchID]*models.LiveMatch {
	if s == nil {
		return nil
	}

	m := make(map[nspb.MatchID]*models.LiveMatch)

	for _, liveMatch := range s {
		m[liveMatch.MatchID] = liveMatch
	}

	return m
}

func (s *LiveMatches) Unshift(liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	*s = append(LiveMatches{liveMatch}, *s...)
}

func (s *LiveMatches) Push(liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	*s = append(*s, liveMatch)
}

func (s *LiveMatches) Insert(i int, liveMatch *models.LiveMatch) {
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

func (s *LiveMatches) Shift() (liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	if len(*s) < 1 {
		return
	}

	liveMatch, *s = (*s)[0], (*s)[1:]

	return
}

func (s *LiveMatches) Pop() (liveMatch *models.LiveMatch) {
	if s == nil {
		return
	}

	if len(*s) < 1 {
		return
	}

	liveMatch, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return
}

func (s *LiveMatches) Remove(i int) *models.LiveMatch {
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

func (s *LiveMatches) RemoveDeactivated() LiveMatches {
	if s == nil {
		return nil
	}

	if *s == nil {
		return nil
	}

	removed := LiveMatches{}
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

func (s LiveMatches) Batches(batchSize int) []LiveMatches {
	if s == nil {
		return nil
	}

	if len(s) == 0 {
		return []LiveMatches{}
	}

	liveMatches := s
	batches := make([]LiveMatches, 0, (len(liveMatches)+batchSize-1)/batchSize)

	for batchSize < len(liveMatches) {
		liveMatches, batches = liveMatches[batchSize:], append(batches, liveMatches[0:batchSize:batchSize])
	}

	batches = append(batches, liveMatches)

	return batches
}
