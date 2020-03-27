package steam

type StateTransition struct {
	Previous       State
	Next           State
	ReadyToUnready bool
	UnreadyToReady bool
}

func NewStateTransition(prev, next State) *StateTransition {
	return &StateTransition{
		Previous:       prev,
		Next:           next,
		ReadyToUnready: prev.Ready && !next.Ready,
		UnreadyToReady: !prev.Ready && next.Ready,
	}
}
