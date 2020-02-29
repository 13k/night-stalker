package heroes

import (
	"fmt"
	"strings"

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

const (
	fmtHeroImageURL = "http://cdn.dota2.com/apps/dota2/images/heroes/%s_%s.%s"
)

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

var imageURLFormats = map[string]struct {
	Suffix string
	Ext    string
}{
	"full": {
		Suffix: "full",
		Ext:    "png",
	},
	"large": {
		Suffix: "lg",
		Ext:    "png",
	},
	"small": {
		Suffix: "sb",
		Ext:    "png",
	},
	"portrait": {
		Suffix: "vert",
		Ext:    "jpg",
	},
}

func createImageURLs(name string) map[string]string {
	shortName := strings.TrimPrefix(name, "npc_dota_hero_")
	urls := make(map[string]string)

	for fname, fcfg := range imageURLFormats {
		urls[fname] = fmt.Sprintf(fmtHeroImageURL, shortName, fcfg.Suffix, fcfg.Ext)
	}

	return urls
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
		imageURLs := createImageURLs(entry.Name)
		hero := &models.Hero{
			ID:               entry.ID,
			Name:             entry.Name,
			LocalizedName:    entry.LocalizedName,
			ImageFullURL:     imageURLs["full"],
			ImageLargeURL:    imageURLs["large"],
			ImageSmallURL:    imageURLs["small"],
			ImagePortraitURL: imageURLs["portrait"],
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
