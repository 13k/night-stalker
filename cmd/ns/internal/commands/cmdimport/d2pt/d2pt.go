package d2pt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

const (
	baseURL       = "http://www.dota2protracker.com"
	dataPath      = "/static/search_items.json"
	playerPathFmt = "/player/%s"
)

var (
	accountIDRegex = regexp.MustCompile(`dotabuff\.com/players/(\d+)`)
)

var Cmd = &cobra.Command{
	Use:   "d2pt",
	Short: "Import players from dota2protracker",
	RunE:  run,
}

var (
	inputFile string
)

func init() {
	Cmd.Flags().StringVarP(&inputFile, "input", "i", "", "use given JSON file as input")
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
	accountID nspb.AccountID
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	client := resty.New().SetHostURL(baseURL)
	result, err := loadData(client, inputFile)

	if err != nil {
		return xerrors.Errorf("error loading d2pt data: %w", err)
	}

	workerPool, err := ants.NewPool(10)

	if err != nil {
		return xerrors.Errorf("error starting worker pool: %w", err)
	}

	defer workerPool.Release()

	dbs := nsdbda.NewSaver(db)

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

			followed := &nsm.FollowedPlayer{
				AccountID: p.accountID,
				Label:     p.label,
			}

			created, err := dbs.FollowPlayer(cmd.Context(), followed, &nsdbda.FollowPlayerOptions{
				Update: false,
			})

			if err != nil {
				if xerrors.Is(err, nsdbda.ErrPlayerAlreadyFollowed) {
					l.Warn(err.Error())
					return
				}

				l.WithError(err).Error("failed to follow player")
				return
			}

			msg := "updated followed player"

			if created {
				msg = "following player"
			}

			l.WithField("account_id", followed.AccountID).Info(msg)
		})

		if err != nil {
			return xerrors.Errorf("error enqueueing task for label %q: %w", label, err)
		}
	}

	wg.Wait()
	log.Info("done")

	return nil
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
		accountID: nspb.AccountID(accountID64),
	}

	return p, nil
}

func loadData(client *resty.Client, inputFile string) (*dataResponse, error) {
	if inputFile != "" {
		return loadLocalData(inputFile)
	}

	return loadRemoteData(client)
}

func loadLocalData(inputFile string) (*dataResponse, error) {
	result := &dataResponse{}
	data, err := ioutil.ReadFile(inputFile)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}

func loadRemoteData(client *resty.Client) (*dataResponse, error) {
	result := &dataResponse{}
	res, err := client.R().SetResult(result).Get(dataPath)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.Status())
	}

	return result, nil
}
