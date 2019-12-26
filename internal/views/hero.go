package views

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

func HeroProto(m *models.Hero) *nspb.Hero {
	return &nspb.Hero{
		Id:               uint32(m.ID),
		Name:             m.Name,
		LocalizedName:    m.LocalizedName,
		ImageFullUrl:     m.ImageFullURL,
		ImageLargeUrl:    m.ImageLargeURL,
		ImageSmallUrl:    m.ImageSmallURL,
		ImagePortraitUrl: m.ImagePortraitURL,
	}
}
