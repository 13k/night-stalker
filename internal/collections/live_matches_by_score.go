package collections

import (
	"sort"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchesByScore struct {
	LiveMatches

	index map[nspb.MatchID]*nsm.LiveMatch
}

func NewLiveMatchesByScore(matches ...*nsm.LiveMatch) *LiveMatchesByScore {
	s := &LiveMatchesByScore{
		LiveMatches: matches,
		index:       make(map[nspb.MatchID]*nsm.LiveMatch),
	}

	sort.Sort(s)

	for _, m := range s.LiveMatches {
		s.index[nspb.MatchID(m.MatchID)] = m
	}

	return s
}

func (s *LiveMatchesByScore) Less(i, j int) bool {
	return s.At(i).SortScore > s.At(j).SortScore
}

func (s *LiveMatchesByScore) At(i int) *nsm.LiveMatch {
	return s.LiveMatches[i]
}

func (s *LiveMatchesByScore) All() LiveMatches {
	return s.LiveMatches
}

// SearchIndex [O(log n)] performs a binary search for an index where the given match is or would be
// in sorted order by SortScore.
//
// It returns Len() if the match was not found.
func (s *LiveMatchesByScore) SearchIndex(match *nsm.LiveMatch) int {
	return sort.Search(s.Len(), func(i int) bool {
		return s.At(i).SortScore <= match.SortScore
	})
}

// FindIndex [O(n)] finds the index of a LiveMatch with the given matchID.
//
// It returns -1 if the matchID was not found.
func (s *LiveMatchesByScore) FindIndex(matchID nspb.MatchID) int {
	for i, match := range s.All() {
		if matchID == nspb.MatchID(match.MatchID) {
			return i
		}
	}

	return -1
}

// Add inserts a match in sorted order if it isn't present (matched by MatchID). If it's present, it
// updates the match (including SortScore, which will reposition the match) if the match changed.
//
// It returns the match index if the match was added or updated, otherwise returns -1.
func (s *LiveMatchesByScore) Add(match *nsm.LiveMatch) int {
	matchID := nspb.MatchID(match.MatchID)

	if m, ok := s.index[matchID]; ok {
		if m.Equal(match) {
			return -1
		}

		s.Remove(matchID)
	}

	i := s.SearchIndex(match)

	s.Insert(i, match)
	s.index[matchID] = match

	return i
}

func (s *LiveMatchesByScore) Remove(matchID nspb.MatchID) nspb.MatchID {
	match := s.index[matchID]

	if match == nil {
		return 0
	}

	i := s.FindIndex(nspb.MatchID(match.MatchID))

	if i < 0 {
		return 0
	}

	removed := s.LiveMatches.Remove(i)

	if removed == nil {
		return 0
	}

	removedMatchID := nspb.MatchID(removed.MatchID)

	delete(s.index, removedMatchID)

	return removedMatchID
}
