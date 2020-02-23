package collections

import (
	"sync"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

type LiveMatches struct {
	mtx     sync.RWMutex
	matches *LiveMatchesByScore
}

func NewLiveMatches(matches ...*models.LiveMatch) *LiveMatches {
	return &LiveMatches{matches: NewLiveMatchesByScore(matches...)}
}

func (m *LiveMatches) Reset() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.matches = nil
}

func (m *LiveMatches) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	return m.matches.Len()
}

func (m *LiveMatches) All() LiveMatchesSlice {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if m.matches == nil {
		return nil
	}

	matches := make(LiveMatchesSlice, len(m.matches.All()))
	copy(matches, m.matches.All())

	return matches
}

func (m *LiveMatches) Add(matches ...*models.LiveMatch) LiveMatchesSlice {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.matches == nil {
		m.matches = &LiveMatchesByScore{}
	}

	change := LiveMatchesSlice{}

	for _, match := range matches {
		if addedIdx := m.matches.Add(match); addedIdx >= 0 {
			change.Push(m.matches.At(addedIdx))
		}
	}

	return change
}

func (m *LiveMatches) Remove(matchIDs ...nspb.MatchID) MatchIDs {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.matches == nil {
		return nil
	}

	var change MatchIDs

	for _, matchID := range matchIDs {
		if removedMatchID := m.matches.Remove(matchID); removedMatchID != 0 {
			change = append(change, removedMatchID)
		}
	}

	return change
}
