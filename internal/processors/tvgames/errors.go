package tvgames

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type errQueryPageFailure struct {
	Page *queryPage
	Err  error
}

func (*errQueryPageFailure) Error() string {
	return "query error"
}

func (err *errQueryPageFailure) Unwrap() error {
	return err.Err
}

type errEmptyResponse struct {
	Page *queryPage
}

func (*errEmptyResponse) Error() string {
	return "empty query response"
}

type errHandleResponseFailure struct {
	Page *queryPage
	Err  error
}

func (*errHandleResponseFailure) Error() string {
	return "error handling query response"
}

func (err *errHandleResponseFailure) Unwrap() error {
	return err.Err
}

type errSaveGameFailure struct {
	MatchID  nspb.MatchID
	ServerID nspb.SteamID
	LobbyID  nspb.LobbyID
	Err      error
}

func (*errSaveGameFailure) Error() string {
	return "error saving tv game"
}

func (err *errSaveGameFailure) Unwrap() error {
	return err.Err
}
