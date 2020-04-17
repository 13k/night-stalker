package rtstats

import (
	"fmt"

	"github.com/13k/geyser"
	"github.com/go-resty/resty/v2"
	d2pb "github.com/paralin/go-dota2/protocol"

	nserr "github.com/13k/night-stalker/internal/errors"
	nsm "github.com/13k/night-stalker/models"
)

type errWorkerSubmitFailure struct {
	*nserr.Err
	LiveMatch *nsm.LiveMatch
}

type errWorkerPanic struct {
	LiveMatch *nsm.LiveMatch
	Value     interface{}
}

func (*errWorkerPanic) Error() string {
	return "recovered worker panic"
}

type errRequestInProgress struct {
	LiveMatch *nsm.LiveMatch
}

func (*errRequestInProgress) Error() string {
	return "request in progress"
}

type errRequestFailure struct {
	*nserr.Err
	LiveMatch *nsm.LiveMatch
	Request   *geyser.Request
	Response  *resty.Response
}

type errInvalidResponseStatus struct {
	LiveMatch *nsm.LiveMatch
	Request   *geyser.Request
	Response  *resty.Response
}

func (err *errInvalidResponseStatus) Error() string {
	return fmt.Sprintf("invalid response status %d", err.Response.StatusCode())
}

type errInvalidResponse struct {
	LiveMatch *nsm.LiveMatch
	Result    *d2pb.CMsgDOTARealtimeGameStatsTerse
}

func (*errInvalidResponse) Error() string {
	return "invalid response"
}

type errStatsSaveFailure struct {
	*nserr.Err
	LiveMatch *nsm.LiveMatch
}
