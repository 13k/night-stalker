package tvgames

import (
	nserr "github.com/13k/night-stalker/internal/errors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type errQueryPageFailure struct {
	*nserr.Err
	Page *queryPage
}

type errEmptyResponse struct {
	Page *queryPage
}

func (*errEmptyResponse) Error() string {
	return "empty query response"
}

type errHandleResponseFailure struct {
	*nserr.Err
	Page *queryPage
}

type errSaveGameFailure struct {
	*nserr.Err
	MatchID  nspb.MatchID
	ServerID nspb.SteamID
	LobbyID  nspb.LobbyID
}
