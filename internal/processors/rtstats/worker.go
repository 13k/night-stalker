package rtstats

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/13k/geyser"
	gsdota2 "github.com/13k/geyser/dota2"
	"github.com/jinzhu/gorm"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"

	nsd2 "github.com/13k/night-stalker/internal/dota2"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

type worker struct {
	db         *gorm.DB
	api        *gsdota2.DOTA2MatchStats
	activeReqs *sync.Map
	liveMatch  *models.LiveMatch
}

func (w *worker) Run(ctx context.Context) (*models.LiveMatchStats, error) {
	if ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", ctx.Err())
	}

	_, skip := w.activeReqs.LoadOrStore(w.liveMatch.MatchID, true)

	if skip {
		return nil, xerrors.Errorf("request in progress: %w", &errRequestInProgress{
			LiveMatch: w.liveMatch,
		})
	}

	defer w.activeReqs.Delete(w.liveMatch.MatchID)

	result, err := w.requestMatchStats(ctx)

	if err != nil {
		return nil, xerrors.Errorf("error requesting API: %w", err)
	}

	if nspb.MatchID(result.GetMatch().GetMatchid()) != w.liveMatch.MatchID {
		return nil, xerrors.Errorf("invalid response: %w", &errInvalidResponse{
			LiveMatch: w.liveMatch,
			Result:    result,
		})
	}

	stats, err := w.createLiveMatchStats(ctx, w.liveMatch, result)

	if err != nil {
		return nil, xerrors.Errorf("error saving stats to database: %w", err)
	}

	return stats, nil
}

func (w *worker) requestMatchStats(ctx context.Context) (*d2pb.CMsgDOTARealtimeGameStatsTerse, error) {
	if ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", ctx.Err())
	}

	req, err := w.api.GetRealtimeStats()

	if err != nil {
		return nil, xerrors.Errorf("error creating API request: %w", &errRequestFailure{
			LiveMatch: w.liveMatch,
			Err:       err,
		})
	}

	headers := map[string]string{
		"Connection": "keep-alive",
		"User-Agent": nsd2.UserAgent,
	}

	params := url.Values{}
	params.Set("server_steam_id", strconv.FormatUint(w.liveMatch.ServerID.ToUint64(), 10))

	reqOptions := geyser.RequestOptions{
		Context: ctx,
		Params:  params,
		Headers: headers,
	}

	result := &apiResult{}
	req.SetOptions(reqOptions).SetResult(result)

	res, err := req.Execute()

	if err != nil {
		return nil, xerrors.Errorf("error performing request: %w", &errRequestFailure{
			LiveMatch: w.liveMatch,
			Request:   req,
			Response:  res,
			Err:       err,
		})
	}

	if !res.IsSuccess() {
		return nil, xerrors.Errorf("error performing request: %w", &errRequestFailure{
			LiveMatch: w.liveMatch,
			Request:   req,
			Response:  res,
			Err:       fmt.Errorf("invalid response status: %s", res.Status()),
		})
	}

	return result.ToProto(), nil
}

func (w *worker) createLiveMatchStats(
	ctx context.Context,
	liveMatch *models.LiveMatch,
	result *d2pb.CMsgDOTARealtimeGameStatsTerse,
) (*models.LiveMatchStats, error) {
	if ctx.Err() != nil {
		return nil, xerrors.Errorf("worker error: %w", ctx.Err())
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
		return nil, xerrors.Errorf("error creating live match stats: %w", &errStatsSaveFailure{
			LiveMatch: liveMatch,
			Err:       err,
		})
	}

	return stats, nil
}
