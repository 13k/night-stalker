package steam

import (
	"sync"

	"github.com/faceit/go-steam"
)

type Session struct {
	State          State
	LastTransition *StateTransition

	c   *steam.Client
	mtx sync.RWMutex
}

func (s *Session) IsReady() bool {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.State.Ready
}

func (s *Session) StateChange(t *StateTransition) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.LastTransition = t
	s.State = t.Next

	go s.c.Emit(&ClientStateChanged{
		Transition: t,
	})
}
