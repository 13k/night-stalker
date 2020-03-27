package teams

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nssql "github.com/13k/night-stalker/internal/sql"
	"github.com/13k/night-stalker/models"
)

const (
	apiURL  = "https://api.opendota.com"
	apiPath = "/api/teams"
)

var Cmd = &cobra.Command{
	Use:   "teams",
	Short: "Import teams from OpenDota API",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	apiKey := v.GetString(v.KeyOpendotaAPIKey)

	log.Info("fetching teams ...")

	client := resty.New().SetHostURL(apiURL)

	if apiKey != "" {
		client.SetQueryParam("api_key", apiKey)
	}

	result := apiResult{}

	res, err := client.R().
		SetResult(&result).
		Get(apiPath)

	if err != nil {
		log.WithError(err).Fatal("error")
	}

	if !res.IsSuccess() {
		log.WithField("status", res.Status()).Fatal("HTTP error")
	}

	log.WithField("count", len(result)).Info("importing teams ...")

	for _, entry := range result {
		l := log.WithField("team_id", entry.TeamID)

		team := &models.Team{
			ID:            entry.TeamID,
			Name:          entry.Name,
			Tag:           entry.Tag,
			Rating:        entry.Rating,
			Wins:          entry.Wins,
			Losses:        entry.Losses,
			LogoURL:       entry.LogoURL,
			LastMatchTime: nssql.NullTimeFromUnixJSON(entry.LastMatchTime),
		}

		dbres := db.
			Where(&models.Team{ID: entry.TeamID}).
			Assign(team).
			FirstOrCreate(team)

		if err = dbres.Error; err != nil {
			l.WithError(err).Fatal("error")
		}

		l.Info("imported")
	}

	log.Info("done")
}
