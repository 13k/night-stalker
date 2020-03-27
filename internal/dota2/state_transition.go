package dota2

import (
	"github.com/paralin/go-dota2/state"
)

type StateTransition struct {
	Previous       state.Dota2State
	Next           state.Dota2State
	ReadyToUnready bool
	UnreadyToReady bool
}

func NewStateTransition(prev, next state.Dota2State) *StateTransition {
	return &StateTransition{
		Previous:       prev,
		Next:           next,
		ReadyToUnready: prev.IsReady() && !next.IsReady(),
		UnreadyToReady: !prev.IsReady() && next.IsReady(),
	}
}
