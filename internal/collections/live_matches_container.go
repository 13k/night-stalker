package collections

import (
	"sync"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

type LiveMatchesContainer struct {
	mtx     sync.RWMutex
	matches *LiveMatchesByScore
}

func NewLiveMatchesContainer(matches ...*nsm.LiveMatch) *LiveMatchesContainer {
	return &LiveMatchesContainer{matches: NewLiveMatchesByScore(matches...)}
}

func (m *LiveMatchesContainer) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	return m.matches.Len()
}

func (m *LiveMatchesContainer) All() LiveMatches {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if m.matches == nil {
		return nil
	}

	matches := make(LiveMatches, len(m.matches.All()))
	copy(matches, m.matches.All())

	return matches
}

func (m *LiveMatchesContainer) Add(matches ...*nsm.LiveMatch) LiveMatches {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.matches == nil {
		m.matches = &LiveMatchesByScore{}
	}

	change := LiveMatches{}

	for _, match := range matches {
		if addedIdx := m.matches.Add(match); addedIdx >= 0 {
			change.Push(m.matches.At(addedIdx))
		}
	}

	return change
}

func (m *LiveMatchesContainer) Remove(matchIDs ...nspb.MatchID) MatchIDs {
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
