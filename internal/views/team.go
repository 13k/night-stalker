package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func NewTeam(team *models.Team) *nspb.Team {
	return &nspb.Team{
		Id:      uint64(team.ID),
		Name:    team.Name,
		Tag:     team.Tag,
		LogoUrl: team.LogoURL,
	}
}
