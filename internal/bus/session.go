package bus

import (
	nsdota2 "github.com/13k/night-stalker/internal/dota2"
	nssteam "github.com/13k/night-stalker/internal/steam"
)

type SteamSessionChangeMessage struct {
	*nssteam.StateTransition

	IsReady bool
	Err     error
}

type DotaSessionChangeMessage struct {
	*nsdota2.StateTransition

	IsReady bool
	Err     error
}
