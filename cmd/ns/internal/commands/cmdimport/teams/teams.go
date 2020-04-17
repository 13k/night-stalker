package teams

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsm "github.com/13k/night-stalker/models"
)

const (
	apiURL  = "https://api.opendota.com"
	apiPath = "/api/teams"
)

var Cmd = &cobra.Command{
	Use:   "teams",
	Short: "Import teams from OpenDota API",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	apiKey := v.GetString(v.KeyOpendotaAPIKey)

	log.Info("fetching teams ...")

	client := resty.New().SetHostURL(apiURL)

	if apiKey != "" {
		client.SetQueryParam("api_key", apiKey)
	}

	result := make(apiResult, 0)

	res, err := client.R().
		SetResult(&result).
		Get(apiPath)

	if err != nil {
		return xerrors.Errorf("error fetching teams: %w", err)
	}

	if !res.IsSuccess() {
		return xerrors.Errorf("HTTP error: %s", res.Status())
	}

	log.
		WithField("count", len(result)).
		Info("importing teams ...")

	for _, entry := range result {
		team := &nsm.Team{
			ID:            nsm.ID(entry.TeamID),
			Name:          entry.Name,
			Tag:           entry.Tag,
			Rating:        entry.Rating,
			Wins:          entry.Wins,
			Losses:        entry.Losses,
			LogoURL:       entry.LogoURL,
			LastMatchTime: nssql.NullTimeFromUnixJSON(entry.LastMatchTime),
		}

		q := db.
			Q().
			Select().
			Eq(nsm.TeamTable.PK(), entry.TeamID)

		created, err := db.M().Upsert(cmd.Context(), team, q)

		if err != nil {
			return xerrors.Errorf("error saving team %d: %w", entry.TeamID, err)
		}

		msg := "updated"

		if created {
			msg = "imported"
		}

		log.WithOFields(
			"id", entry.TeamID,
			"tag", entry.Tag,
		).Info(msg)
	}

	log.Info("done")

	return nil
}
