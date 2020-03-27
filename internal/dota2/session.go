package dota2

import (
	"sync"

	"github.com/paralin/go-dota2/state"
)

type Session struct {
	State          state.Dota2State
	LastTransition *StateTransition

	mtx sync.RWMutex
}

func (s *Session) IsReady() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.State.IsReady()
}

func (s *Session) StateChange(t *StateTransition) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.LastTransition = t
	s.State = t.Next
}
