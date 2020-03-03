package heroes

import (
	"github.com/13k/geyser"
	"github.com/paralin/go-dota2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	"github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "heroes",
	Short: "Import heroes from Steam API",
	Run:   run,
}

type response struct {
	Result *result `json:"result"`
}

type result struct {
	Heroes []*resultHero `json:"heroes,omitempty"`
	Status int           `json:"status,omitempty"`
	Count  int           `json:"count,omitempty"`
}

type resultHero struct {
	ID            models.HeroID `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	LocalizedName string        `json:"localized_name,omitempty"`
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

	apiKey := viper.GetString("steam.api_key")

	if apiKey == "" {
		log.Fatal("steam.api_key is required")
	}

	log.Info("fetching heroes ...")

	client, err := geyser.New(
		geyser.WithKey(apiKey),
		geyser.WithLanguage("en_US"),
	)

	if err != nil {
		log.WithError(err).Fatal()
	}

	econdota2, err := client.EconDOTA2(dota2.AppID)

	if err != nil {
		log.WithError(err).Fatal()
	}

	req, err := econdota2.GetHeroes()

	result := &response{}
	req.SetResult(result)

	if err != nil {
		log.WithError(err).Fatal()
	}

	resp, err := req.Execute()

	if err != nil {
		log.WithError(err).Fatal()
	}

	if !resp.IsSuccess() {
		log.WithFields(logrus.Fields{
			"code":   resp.StatusCode(),
			"status": resp.StatusCode(),
		}).Fatal("HTTP error")
	}

	log.Info("importing heroes ...")

	tx := db.Begin()

	for _, entry := range result.Result.Heroes {
		l := log.WithField("name", entry.Name)
		hero := &models.Hero{
			ID:            entry.ID,
			Name:          entry.Name,
			LocalizedName: entry.LocalizedName,
		}

		result := tx.Assign(hero).FirstOrCreate(hero)

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
