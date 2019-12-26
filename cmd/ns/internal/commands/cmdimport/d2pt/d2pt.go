package d2pt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/13k/night-stalker/cmd/ns/internal/db"
	"github.com/13k/night-stalker/cmd/ns/internal/logger"
	"github.com/13k/night-stalker/cmd/ns/internal/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	dataURL      = "http://www.dota2protracker.com/static/search.json"
	playerURLFmt = "http://www.dota2protracker.com/player/%s"
)

var (
	accountIDRegex = regexp.MustCompile(`dotabuff\.com/players/(\d+)`)
)

var Cmd = &cobra.Command{
	Use:   "d2pt",
	Short: "Import players from dota2protracker",
	Run:   run,
}

type response struct {
	Heroes  map[string]responseEntry `json:"heroes"`
	Players map[string]responseEntry `json:"players"`
}

type responseEntry struct {
	Aliases []string `json:"valid"`
}

func fetchAccountID(playerLabel string) (uint32, error) {
	url := fmt.Sprintf(playerURLFmt, playerLabel)

	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	matches := accountIDRegex.FindSubmatch(body)

	if matches == nil {
		return 0, fmt.Errorf("could not find account_id in %q", url)
	}

	accountIDStr := string(matches[1])
	accountID64, err := strconv.ParseUint(accountIDStr, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("invalid account_id value %q found in %q", accountIDStr, url)
	}

	return uint32(accountID64), nil
}

func run(cmd *cobra.Command, args []string) {
	log, err := logger.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	db, err := db.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	resp, err := http.Get(dataURL)

	if err != nil {
		log.WithError(err).Fatal("error requesting d2pt data")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.WithError(err).Fatal("error requesting d2pt data")
	}

	data := &response{}

	if err := json.Unmarshal(body, data); err != nil {
		log.WithError(err).Fatal("error parsing d2pt data")
	}

	for label := range data.Players {
		l := log.WithField("label", label)

		accountID, err := fetchAccountID(label)

		if err != nil {
			l.WithError(err).Error("failed to fetch account_id")
			continue
		}

		followed, err := util.FollowPlayer(db, accountID, label, false)

		if err != nil {
			l.WithError(err).Error("failed to follow player")
			continue
		}

		log.WithFields(logrus.Fields{
			"account_id": followed.AccountID,
			"label":      followed.Label,
		}).Info("following player")
	}
}
