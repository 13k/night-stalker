package collections

import (
	"github.com/13k/night-stalker/models"
)

// LiveMatchesSlice is a slice of LiveMatch.
//
// Methods based on https://github.com/golang/go/wiki/SliceTricks
type LiveMatchesSlice []*models.LiveMatch

func (s LiveMatchesSlice) Len() int      { return len(s) }
func (s LiveMatchesSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s *LiveMatchesSlice) Unshift(match *models.LiveMatch) {
	*s = append(LiveMatchesSlice{match}, *s...)
}

func (s *LiveMatchesSlice) Push(match *models.LiveMatch) {
	*s = append(*s, match)
}

func (s *LiveMatchesSlice) Insert(i int, match *models.LiveMatch) {
	if i < 0 || i > len(*s) {
		return
	}

	if i == 0 {
		s.Unshift(match)
		return
	}

	if i == len(*s) {
		s.Push(match)
		return
	}

	*s = append(*s, nil)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = match
}

func (s *LiveMatchesSlice) Shift() (match *models.LiveMatch) {
	if len(*s) < 1 {
		return
	}

	match, *s = (*s)[0], (*s)[1:]

	return
}

func (s *LiveMatchesSlice) Pop() (match *models.LiveMatch) {
	if len(*s) < 1 {
		return
	}

	match, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return
}

func (s *LiveMatchesSlice) Remove(i int) *models.LiveMatch {
	if i < 0 || i >= len(*s) {
		return nil
	}

	if i == 0 {
		return s.Shift()
	}

	if i == len(*s)-1 {
		return s.Pop()
	}

	match := (*s)[i]

	copy((*s)[i:], (*s)[i+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]

	return match
}

func (s LiveMatchesSlice) Batches(batchSize int) []LiveMatchesSlice {
	if s == nil {
		return nil
	}

	if len(s) == 0 {
		return []LiveMatchesSlice{}
	}

	matches := s
	batches := make([]LiveMatchesSlice, 0, (len(matches)+batchSize-1)/batchSize)

	for batchSize < len(matches) {
		matches, batches = matches[batchSize:], append(batches, matches[0:batchSize:batchSize])
	}

	batches = append(batches, matches)

	return batches
}
