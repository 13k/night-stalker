package rtstats

import (
	"github.com/13k/geyser"
	"github.com/go-resty/resty/v2"
	d2pb "github.com/paralin/go-dota2/protocol"

	"github.com/13k/night-stalker/models"
)

type errWorkerSubmitFailure struct {
	LiveMatch *models.LiveMatch
	Err       error
}

func (*errWorkerSubmitFailure) Error() string {
	return "worker submit error"
}

func (err *errWorkerSubmitFailure) Unwrap() error {
	return err.Err
}

type errWorkerPanic struct {
	LiveMatch *models.LiveMatch
	Value     interface{}
}

func (*errWorkerPanic) Error() string {
	return "recovered worker panic"
}

type errRequestInProgress struct {
	LiveMatch *models.LiveMatch
}

func (*errRequestInProgress) Error() string {
	return "request in progress"
}

type errInvalidResponse struct {
	LiveMatch *models.LiveMatch
	Result    *d2pb.CMsgDOTARealtimeGameStatsTerse
}

func (*errInvalidResponse) Error() string {
	return "invalid response"
}

type errRequestFailure struct {
	LiveMatch *models.LiveMatch
	Request   *geyser.Request
	Response  *resty.Response
	Err       error
}

func (*errRequestFailure) Error() string {
	return "request error"
}

func (err *errRequestFailure) Unwrap() error {
	return err.Err
}

type errStatsSaveFailure struct {
	LiveMatch *models.LiveMatch
	Err       error
}

func (*errStatsSaveFailure) Error() string {
	return "database error"
}

func (err *errStatsSaveFailure) Unwrap() error {
	return err.Err
}
