package views

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

func NewHero(h *models.Hero) *nspb.Hero {
	return &nspb.Hero{
		Id:                 uint64(h.ID),
		Name:               h.Name,
		Slug:               h.Slug,
		LocalizedName:      h.LocalizedName,
		Aliases:            h.Aliases,
		Roles:              h.Roles,
		RoleLevels:         h.RoleLevels,
		Complexity:         int64(h.Complexity),
		Legs:               int64(h.Legs),
		AttributePrimary:   h.AttributePrimary,
		AttackCapabilities: h.AttackCapabilities,
	}
}
