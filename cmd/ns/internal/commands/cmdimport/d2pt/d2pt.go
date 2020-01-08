package d2pt

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/ns/internal/db"
	"github.com/13k/night-stalker/cmd/ns/internal/logger"
	"github.com/13k/night-stalker/cmd/ns/internal/util"
)

const (
	baseURL       = "http://www.dota2protracker.com"
	dataPath      = "/static/search.json"
	playerPathFmt = "/player/%s"
)

var (
	accountIDRegex = regexp.MustCompile(`dotabuff\.com/players/(\d+)`)
)

var Cmd = &cobra.Command{
	Use:   "d2pt",
	Short: "Import players from dota2protracker",
	Run:   run,
}

type dataResponse struct {
	Heroes  map[string]dataResponseEntry `json:"heroes"`
	Players map[string]dataResponseEntry `json:"players"`
}

type dataResponseEntry struct {
	Aliases []string `json:"valid"`
}

type fetchError struct {
	label   string
	err     error
	message string
}

func (err *fetchError) Error() string {
	return err.message
}

type player struct {
	label     string
	accountID uint32
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

	client := resty.New().
		SetHostURL(baseURL)

	result := &dataResponse{}

	_, err = client.R().
		SetResult(result).
		Get(dataPath)

	if err != nil {
		log.WithError(err).Fatal("error requesting d2pt data")
	}

	workerPool, err := ants.NewPool(10)

	if err != nil {
		log.WithError(err).Fatal("error starting worker pool")
	}

	defer workerPool.Release()

	var wg sync.WaitGroup

	for label := range result.Players {
		err := workerPool.Submit(func() {
			defer wg.Done()
			wg.Add(1)

			l := log.WithField("label", label)
			p, ferr := fetchAccountID(client, label)

			if ferr != nil {
				l.WithError(ferr.err).Error(ferr.message)
				return
			}

			followed, err := util.FollowPlayer(db, p.accountID, p.label, false)

			if err != nil {
				if err == util.ErrFollowedPlayerAlreadyExists {
					l.Warn(err.Error())
					return
				}

				l.WithError(err).Error("failed to follow player")
				return
			}

			l.WithField("account_id", followed.AccountID).Info("following player")
		})

		if err != nil {
			log.
				WithError(err).
				WithField("label", label).
				Error("error queueing task")
		}
	}

	wg.Wait()
	log.Info("done")
}

func fetchAccountID(client *resty.Client, label string) (*player, *fetchError) {
	path := fmt.Sprintf(playerPathFmt, label)

	resp, err := client.R().
		SetDoNotParseResponse(true).
		Get(path)

	if err != nil {
		ferr := &fetchError{
			label:   label,
			err:     err,
			message: fmt.Sprintf("error requesting %s: %s", path, err.Error()),
		}

		return nil, ferr
	}

	defer resp.RawBody().Close()

	body, err := ioutil.ReadAll(resp.RawBody())

	if err != nil {
		ferr := &fetchError{
			label:   label,
			err:     err,
			message: fmt.Sprintf("error reading response body: %s", err.Error()),
		}

		return nil, ferr
	}

	matches := accountIDRegex.FindSubmatch(body)

	if matches == nil {
		ferr := &fetchError{
			label:   label,
			err:     err,
			message: fmt.Sprintf("could not find account_id in %s", path),
		}

		return nil, ferr
	}

	accountIDStr := string(matches[1])
	accountID64, err := strconv.ParseUint(accountIDStr, 10, 32)

	if err != nil {
		ferr := &fetchError{
			label:   label,
			err:     err,
			message: fmt.Sprintf("invalid account_id value %q found in %s", accountIDStr, path),
		}

		return nil, ferr
	}

	p := &player{
		label:     label,
		accountID: uint32(accountID64),
	}

	return p, nil
}
