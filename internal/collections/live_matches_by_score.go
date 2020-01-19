package collections

import (
	"sort"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatchesByScore struct {
	LiveMatchesSlice

	index map[nspb.MatchID]*models.LiveMatch
}

func NewLiveMatchesByScore(matches ...*models.LiveMatch) *LiveMatchesByScore {
	s := &LiveMatchesByScore{
		LiveMatchesSlice: matches,
		index:            make(map[nspb.MatchID]*models.LiveMatch),
	}

	sort.Sort(s)

	for _, m := range s.LiveMatchesSlice {
		s.index[m.MatchID] = m
	}

	return s
}

func (s *LiveMatchesByScore) Less(i, j int) bool {
	return s.LiveMatchesSlice[i].SortScore > s.LiveMatchesSlice[j].SortScore
}

func (s *LiveMatchesByScore) At(i int) *models.LiveMatch {
	return s.LiveMatchesSlice[i]
}

func (s *LiveMatchesByScore) Get(matchID nspb.MatchID) *models.LiveMatch {
	if m, ok := s.index[matchID]; ok {
		return m
	}

	return nil
}

func (s *LiveMatchesByScore) All() LiveMatchesSlice {
	return s.LiveMatchesSlice
}

func (s *LiveMatchesByScore) SearchIndex(match *models.LiveMatch) int {
	return sort.Search(len(s.LiveMatchesSlice), func(i int) bool {
		return s.LiveMatchesSlice[i].SortScore <= match.SortScore
	})
}

func (s *LiveMatchesByScore) safeIsAt(i int, matchID nspb.MatchID) bool {
	return s.LiveMatchesSlice[i].MatchID == matchID
}

func (s *LiveMatchesByScore) IsAt(i int, matchID nspb.MatchID) bool {
	if i < 0 || i >= len(s.LiveMatchesSlice) {
		return false
	}

	return s.safeIsAt(i, matchID)
}

// Add inserts a match in sorted order if it isn't present (matched by MatchID). If it's present but
// out of order (SortScore possibly changed), it removes the existing match and inserts the given
// match into the correct position.
//
// It returns true if the match was added or its position changed.
func (s *LiveMatchesByScore) Add(match *models.LiveMatch) bool {
	if m, ok := s.index[match.MatchID]; ok {
		if m.SortScore == match.SortScore {
			return false
		}

		s.Remove(m.MatchID)
	}

	i := s.SearchIndex(match)

	s.LiveMatchesSlice.Insert(i, match)
	s.index[match.MatchID] = match

	return true
}

func (s *LiveMatchesByScore) Remove(matchID nspb.MatchID) bool {
	if m, ok := s.index[matchID]; ok {
		i := s.SearchIndex(m)
		_ = s.LiveMatchesSlice.Remove(i)
		delete(s.index, matchID)
		return true
	}

	return false
}
