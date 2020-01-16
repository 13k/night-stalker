package bus

import (
	nspb "github.com/13k/night-stalker/internal/protocol"

	"github.com/13k/night-stalker/models"
)

type LiveMatchesMessage struct {
	Matches []*models.LiveMatch
}

type LiveMatchesChangeMessage struct {
	Change *nspb.LiveMatchesChange
}
