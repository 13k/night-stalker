package leagues

import (
	d2pb "github.com/paralin/go-dota2/protocol"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdgeyser "github.com/13k/night-stalker/cmd/ns/internal/geyser"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
)

var Cmd = &cobra.Command{
	Use:   "leagues",
	Short: "Import leagues from Dota2 WebAPI",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	api, err := nscmdgeyser.NewDota2()

	if err != nil {
		return xerrors.Errorf("error creating API client: %w", err)
	}

	apiLeague, err := api.DOTA2League()

	if err != nil {
		return xerrors.Errorf("error creating API client: %w", err)
	}

	req, err := apiLeague.GetLeagueInfoList()

	if err != nil {
		return xerrors.Errorf("error creating API request: %w", err)
	}

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	log.Info("fetching leagues ...")

	result := &d2pb.CMsgDOTALeagueInfoList{}
	req.SetResult(result)

	res, err := req.Execute()

	if err != nil {
		return xerrors.Errorf("error requesting leagues: %w", err)
	}

	if !res.IsSuccess() {
		return xerrors.Errorf("HTTP error: %s", res.Status())
	}

	dbs := nsdbda.NewSaver(db)

	log.
		WithField("count", len(result.GetInfos())).
		Info("importing leagues ...")

	for _, pb := range result.GetInfos() {
		league, created, err := dbs.UpsertLeagueProto(cmd.Context(), pb)

		if err != nil {
			return xerrors.Errorf("error saving league %d: %w", league.ID, err)
		}

		msg := "updated"

		if created {
			msg = "imported"
		}

		log.WithOFields(
			"id", league.ID,
			"name", league.Name,
		).Info(msg)
	}

	log.Info("done")

	return nil
}
