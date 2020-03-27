package leagues

import (
	"github.com/paralin/go-dota2/protocol"
	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdgeyser "github.com/13k/night-stalker/cmd/ns/internal/geyser"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "leagues",
	Short: "Import leagues from Dota2 WebAPI",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	api, err := nscmdgeyser.NewDota2()

	if err != nil {
		log.WithError(err).Fatal("error creating API client")
	}

	apiLeague, err := api.DOTA2League()

	if err != nil {
		log.WithError(err).Fatal("error creating API client")
	}

	req, err := apiLeague.GetLeagueInfoList()

	if err != nil {
		log.WithError(err).Fatal("error creating API request")
	}

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	log.Info("fetching leagues ...")

	result := &protocol.CMsgDOTALeagueInfoList{}
	req.SetResult(result)

	res, err := req.Execute()

	if err != nil {
		log.WithError(err).Fatal("error requesting leagues")
	}

	if !res.IsSuccess() {
		log.WithField("status", res.Status()).Fatal("error requesting leagues")
	}

	log.
		WithField("count", len(result.GetInfos())).
		Info("importing leagues ...")

	for _, info := range result.GetInfos() {
		l := log.WithField("league_id", info.GetLeagueId())

		league := &models.League{
			ID:             nspb.LeagueID(info.GetLeagueId()),
			Name:           info.GetName(),
			Tier:           nspb.LeagueTier(info.GetTier()),
			Region:         nspb.LeagueRegion(info.GetRegion()),
			Status:         nspb.LeagueStatus(info.GetStatus()),
			TotalPrizePool: info.GetTotalPrizePool(),
			LastActivityAt: nssql.NullTimeUnix(int64(info.GetMostRecentActivity())),
			StartAt:        nssql.NullTimeUnix(int64(info.GetStartTimestamp())),
			FinishAt:       nssql.NullTimeUnix(int64(info.GetEndTimestamp())),
		}

		result := db.
			Where(&models.League{ID: league.ID}).
			Assign(league).
			FirstOrCreate(league)

		if err = result.Error; err != nil {
			l.WithError(err).Fatal("error")
		}

		l.WithField("name", league.Name).Info("imported")
	}

	log.Info("done")
}
