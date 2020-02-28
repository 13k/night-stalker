package db

import (
	"github.com/jinzhu/gorm"

	nscol "github.com/13k/night-stalker/internal/collections"
	"github.com/13k/night-stalker/models"
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

func (f *PlayerFilters) Filter(scope *gorm.DB, model models.Model) *gorm.DB {
	if len(f.MatchIDs) > 0 {
		scope = In(scope, model, "match_id", f.MatchIDs)
	}

	if len(f.AccountIDs) > 0 {
		scope = In(scope, model, "account_id", f.AccountIDs)
	}

	if len(f.HeroIDs) > 0 {
		scope = In(scope, model, "hero_id", f.HeroIDs)
	}

	return scope
}
