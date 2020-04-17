package views

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

func NewTeam(team *nsm.Team) *nspb.Team {
	return &nspb.Team{
		Id:      uint64(team.ID),
		Name:    team.Name,
		Tag:     team.Tag,
		LogoUrl: team.LogoURL,
	}
}
