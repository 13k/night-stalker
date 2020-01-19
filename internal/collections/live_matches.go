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

func NewLiveMatches() *LiveMatches {
	return &LiveMatches{matches: NewLiveMatchesByScore()}
}

func (m *LiveMatches) Reset() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.matches = nil
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

func (m *LiveMatches) Add(matches ...*models.LiveMatch) int {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.matches == nil {
		m.matches = &LiveMatchesByScore{}
	}

	var change int

	for _, match := range matches {
		if m.matches.Add(match) {
			change++
		}
	}

	return change
}

func (m *LiveMatches) Remove(matchIDs ...nspb.MatchID) int {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.matches == nil {
		return 0
	}

	var change int

	for _, matchID := range matchIDs {
		if m.matches.Remove(matchID) {
			change++
		}
	}

	return change
}
