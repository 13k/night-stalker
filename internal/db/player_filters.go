package db

import (
	"errors"

	"github.com/jinzhu/gorm"

	nscol "github.com/13k/night-stalker/internal/collections"
	"github.com/13k/night-stalker/models"
)

var (
	ErrEmptyPlayerFilters = errors.New("empty filters")
)

type PlayerFilters struct {
	MatchIDs   nscol.MatchIDs
	AccountIDs nscol.AccountIDs
	HeroIDs    nscol.HeroIDs
}

func (f *PlayerFilters) Empty() bool {
	if f == nil {
		return true
	}

	return len(f.MatchIDs) == 0 && len(f.AccountIDs) == 0 && len(f.HeroIDs) == 0
}

func (f *PlayerFilters) Validate() error {
	if f.Empty() {
		return ErrEmptyPlayerFilters
	}

	return nil
}

func (f *PlayerFilters) Filter(scope *gorm.DB, model models.Model) *gorm.DB {
	switch l := len(f.MatchIDs); {
	case l > 1:
		scope = In(scope, model, "match_id", f.MatchIDs)
	case l == 1:
		scope = Eq(scope, model, "match_id", f.MatchIDs[0])
	}

	switch l := len(f.AccountIDs); {
	case l > 1:
		scope = In(scope, model, "account_id", f.AccountIDs)
	case l == 1:
		scope = Eq(scope, model, "account_id", f.AccountIDs[0])
	}

	switch l := len(f.HeroIDs); {
	case l > 1:
		scope = In(scope, model, "hero_id", f.HeroIDs)
	case l == 1:
		scope = Eq(scope, model, "hero_id", f.HeroIDs[0])
	}

	return scope
}
