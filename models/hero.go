package models

import (
	"strings"

	"github.com/lib/pq"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsstr "github.com/13k/night-stalker/internal/strings"
)

var HeroModel Model = (*Hero)(nil)

type HeroID uint64

type Hero struct {
	ID                 HeroID             `gorm:"column:id;primary_key"`
	Name               string             `gorm:"column:name;size:255;unique_index;not null"`
	Slug               string             `gorm:"column:slug;size:255;unique_index;not null"`
	LocalizedName      string             `gorm:"column:localized_name;size:255;not null"`
	Aliases            pq.StringArray     `gorm:"column:aliases"`
	Roles              nssql.HeroRoles    `gorm:"column:roles"`
	RoleLevels         pq.Int64Array      `gorm:"column:role_levels"`
	Complexity         int                `gorm:"column:complexity"`
	Legs               int                `gorm:"column:legs"`
	AttributePrimary   nspb.DotaAttribute `gorm:"column:attribute_primary"`
	AttackCapabilities nspb.DotaUnitCap   `gorm:"column:attack_capabilities"`
	Timestamps
}

func (*Hero) TableName() string {
	return "heroes"
}

func (h *Hero) BeforeCreate() error {
	if h.Slug == "" {
		h.Slug = nsstr.Slugify(strings.Replace(h.Name, "npc_dota_hero_", "", 1))
	}

	return nil
}
