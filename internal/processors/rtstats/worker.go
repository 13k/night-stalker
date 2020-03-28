package rtstats

import (
	"context"
	"net/url"
	"strconv"
	"sync"

	"github.com/13k/geyser"
	gsdota2 "github.com/13k/geyser/dota2"
	"github.com/jinzhu/gorm"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsd2 "github.com/13k/night-stalker/internal/dota2"
	nserr "github.com/13k/night-stalker/internal/errors"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type worker struct {
	ctx          context.Context
	db           *gorm.DB
	api          *gsdota2.DOTA2MatchStats
	activeReqs   *sync.Map
	liveMatch    *models.LiveMatch
	results      *flusher
	errorHandler func(error)
}

func (w *worker) Run() {
	var err error
	var stats *models.LiveMatchStats

	defer func() {
		if v := recover(); v != nil {
			err = xerrors.Errorf("worker panic: %w", &errWorkerPanic{
				LiveMatch: w.liveMatch,
				Value:     v,
			})
		}

		if err != nil {
			w.errorHandler(xerrors.Errorf("worker error: %w", err))
		}
	}()

	stats, err = w.run()

	if err == nil && stats != nil {
		w.results.Add(stats)
	}
}

func (w *worker) run() (*models.LiveMatchStats, error) {
	if w.ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", w.ctx.Err())
	}

	_, skip := w.activeReqs.LoadOrStore(w.liveMatch.MatchID, true)

	if skip {
		return nil, xerrors.Errorf("request in progress: %w", &errRequestInProgress{
			LiveMatch: w.liveMatch,
		})
	}

	defer w.activeReqs.Delete(w.liveMatch.MatchID)

	pbmsg, err := w.requestMatchStats()

	if err != nil {
		return nil, xerrors.Errorf("error requesting API: %w", err)
	}

	if nspb.MatchID(pbmsg.GetMatch().GetMatchid()) != w.liveMatch.MatchID {
		return nil, xerrors.Errorf("invalid response: %w", &errInvalidResponse{
			LiveMatch: w.liveMatch,
			Result:    pbmsg,
		})
	}

	stats, err := w.createLiveMatchStats(w.liveMatch, pbmsg)

	if err != nil {
		return nil, xerrors.Errorf("error saving stats to database: %w", err)
	}

	return stats, nil
}

func (w *worker) requestMatchStats() (*d2pb.CMsgDOTARealtimeGameStatsTerse, error) {
	if w.ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", w.ctx.Err())
	}

	req, err := w.api.GetRealtimeStats()

	if err != nil {
		return nil, &errRequestFailure{
			LiveMatch: w.liveMatch,
			Err:       nserr.Wrap("error creating API request", err),
		}
	}

	headers := map[string]string{
		"Connection": "keep-alive",
		"User-Agent": nsd2.UserAgent,
	}

	params := url.Values{}
	params.Set("server_steam_id", strconv.FormatUint(w.liveMatch.ServerID.ToUint64(), 10))

	reqOptions := geyser.RequestOptions{
		Context: w.ctx,
		Params:  params,
		Headers: headers,
	}

	req.SetOptions(reqOptions)

	res, err := req.Execute()

	if err != nil {
		return nil, &errRequestFailure{
			LiveMatch: w.liveMatch,
			Request:   req,
			Response:  res,
			Err:       nserr.Wrap("error performing request", err),
		}
	}

	r := &response{Response: res}

	if !r.IsSuccess() {
		return nil, &errInvalidResponseStatus{
			LiveMatch: w.liveMatch,
			Request:   req,
			Response:  res,
		}
	}

	pbmsg, err := r.Parse()

	if err != nil {
		return nil, &errRequestFailure{
			LiveMatch: w.liveMatch,
			Request:   req,
			Response:  res,
			Err:       nserr.Wrap("error parsing response", err),
		}
	}

	return pbmsg, nil
}

func (w *worker) createLiveMatchStats(
	liveMatch *models.LiveMatch,
	result *d2pb.CMsgDOTARealtimeGameStatsTerse,
) (*models.LiveMatchStats, error) {
	if w.ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", w.ctx.Err())
	}

	stats := models.NewLiveMatchStats(liveMatch, result)

	for _, team := range result.GetTeams() {
		stats.Teams = append(stats.Teams, models.LiveMatchStatsTeamDotaProto(team))

		for _, player := range team.GetPlayers() {
			stats.Players = append(stats.Players, models.NewLiveMatchStatsPlayer(stats, player))
		}
	}

	for _, pickban := range result.GetMatch().GetPicks() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(false, pickban))
	}

	for _, pickban := range result.GetMatch().GetBans() {
		stats.Draft = append(stats.Draft, models.LiveMatchStatsPickBanDotaProto(true, pickban))
	}

	for _, building := range result.GetBuildings() {
		stats.Buildings = append(stats.Buildings, models.LiveMatchStatsBuildingDotaProto(building))
	}

	if err := w.db.Save(stats).Error; err != nil {
		return nil, &errStatsSaveFailure{
			LiveMatch: liveMatch,
			Err:       nserr.Wrap("error creating live match stats", err),
		}
	}

	return stats, nil
}
