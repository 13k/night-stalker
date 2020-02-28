package views

import (
	"errors"
	"fmt"
)

var (
	ErrMissingMatchID         = errors.New("missing match ID")
	ErrInconsistentMatchIDs   = errors.New("inconsistent match IDs")
	ErrMissingAccountID       = errors.New("missing account ID")
	ErrInconsistentAccountIDs = errors.New("inconsistent account IDs")
	ErrMissingHeroID          = errors.New("missing hero ID")
)

type ErrAssociationNotEagerLoaded struct {
	Association string
}

func NewErrAssociationNotEagerLoaded(association string) *ErrAssociationNotEagerLoaded {
	return &ErrAssociationNotEagerLoaded{Association: association}
}

func (err *ErrAssociationNotEagerLoaded) Error() string {
	return fmt.Sprintf("association not eager-loaded: %s", err.Association)
}
