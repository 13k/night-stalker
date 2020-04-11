package models

import (
	"strings"

	"github.com/lib/pq"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsstr "github.com/13k/night-stalker/internal/strings"
)

var HeroTable = NewTable("heroes")

type Hero struct {
	ID `db:"id"`

	Name               string             `db:"name"`
	Slug               string             `db:"slug"`
	LocalizedName      string             `db:"localized_name"`
	Aliases            pq.StringArray     `db:"aliases"`
	Roles              nssql.HeroRoles    `db:"roles"`
	RoleLevels         pq.Int64Array      `db:"role_levels"`
	Complexity         int                `db:"complexity"`
	Legs               int                `db:"legs"`
	AttributePrimary   nspb.DotaAttribute `db:"attribute_primary"`
	AttackCapabilities nspb.DotaUnitCap   `db:"attack_capabilities"`

	Timestamps
	SoftDelete
}

func (m *Hero) BeforeCreate() error {
	if m.Slug == "" {
		m.Slug = nsstr.Slugify(strings.Replace(m.Name, "npc_dota_hero_", "", 1))
	}

	return nil
}
