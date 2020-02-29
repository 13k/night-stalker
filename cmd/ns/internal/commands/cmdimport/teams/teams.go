package teams

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsjson "github.com/13k/night-stalker/internal/json"
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

type response []*responseEntry

type responseEntry struct {
	TeamID        models.TeamID    `json:"team_id,omitempty"`
	Name          string           `json:"name,omitempty"`
	Tag           string           `json:"tag,omitempty"`
	Rating        float32          `json:"rating,omitempty"`
	Wins          uint32           `json:"wins,omitempty"`
	Losses        uint32           `json:"losses,omitempty"`
	LogoURL       string           `json:"logo_url,omitempty"`
	LastMatchTime *nsjson.UnixTime `json:"last_match_time,omitempty"`
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

	apiKey := viper.GetString("opendota.api_key")

	log.Info("fetching teams ...")

	client := resty.New().SetHostURL(apiURL)

	if apiKey != "" {
		client.SetQueryParam("api_key", apiKey)
	}

	result := response{}

	resp, err := client.R().
		SetResult(&result).
		Get(apiPath)

	if err != nil {
		log.WithError(err).Fatal()
	}

	if !resp.IsSuccess() {
		log.WithFields(logrus.Fields{
			"code":   resp.StatusCode(),
			"status": resp.StatusCode(),
		}).Fatal("HTTP error")
	}

	log.WithField("count", len(result)).Info("importing teams ...")

	tx := db.Begin()

	for _, entry := range result {
		l := log.WithField("team_id", entry.TeamID)
		team := &models.Team{
			ID:      entry.TeamID,
			Name:    entry.Name,
			Tag:     entry.Tag,
			Rating:  entry.Rating,
			Wins:    entry.Wins,
			Losses:  entry.Losses,
			LogoURL: entry.LogoURL,
		}

		if entry.LastMatchTime != nil {
			team.LastMatchTime = entry.LastMatchTime.Time
		}

		result := tx.Assign(team).FirstOrCreate(team)

		if err = result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Fatal()
		}

		l.Info("imported")
	}

	if err = tx.Commit().Error; err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("done")
}
