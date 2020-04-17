package dataaccess

import (
	"errors"
)

var (
	ErrMissingMatchID         = errors.New("missing match ID")
	ErrInconsistentMatchIDs   = errors.New("inconsistent match IDs")
	ErrMissingAccountID       = errors.New("missing account ID")
	ErrInconsistentAccountIDs = errors.New("inconsistent account IDs")
	ErrMissingHeroID          = errors.New("missing hero ID")
	ErrEmptyMatchIDs          = errors.New("empty match IDs")
	ErrEmptyAccountIDs        = errors.New("empty account IDs")
	ErrEmptyFilters           = errors.New("empty filters")
)
