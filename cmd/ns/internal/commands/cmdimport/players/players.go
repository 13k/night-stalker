package players

import (
	"github.com/13k/night-stalker/cmd/ns/internal/db"
	"github.com/13k/night-stalker/cmd/ns/internal/logger"
	"github.com/13k/night-stalker/cmd/ns/internal/util"
	nsjson "github.com/13k/night-stalker/internal/json"
	nsproto "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
	"github.com/faceit/go-steam/steamid"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	apiURL  = "https://api.opendota.com"
	apiPath = "/api/proPlayers"
)

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Import players from OpenDota API",
	Run:   run,
}

type response []*responseEntry

type responseEntry struct {
	AccountID    nsproto.AccountID   `json:"account_id,omitempty"`
	SteamID      nsjson.StringUint   `json:"steamid,omitempty"`
	TeamID       models.TeamID       `json:"team_id,omitempty"`
	Name         string              `json:"name,omitempty"`
	PersonaName  string              `json:"personaname,omitempty"`
	Avatar       string              `json:"avatar,omitempty"`
	AvatarMedium string              `json:"avatarmedium,omitempty"`
	AvatarFull   string              `json:"avatarfull,omitempty"`
	ProfileURL   string              `json:"profileurl,omitempty"`
	CountryCode  string              `json:"loccountrycode,omitempty"`
	FantasyRole  nsproto.FantasyRole `json:"fantasy_role,omitempty"`
	IsLocked     bool                `json:"is_locked,omitempty"`
	LockedUntil  *nsjson.UnixTime    `json:"locked_until,omitempty"`
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

	apiKey := viper.GetString("opendota_api_key")

	log.Info("fetching pro players ...")

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

	log.
		WithField("count", len(result)).
		Info("importing pro players ...")

	tx := db.Begin()

	for _, entry := range result {
		l := log.WithField("account_id", entry.AccountID)

		player := &models.Player{
			AccountID:       entry.AccountID,
			SteamID:         steamid.SteamId(entry.SteamID.Uint64()),
			Name:            entry.Name,
			PersonaName:     entry.PersonaName,
			AvatarURL:       entry.Avatar,
			AvatarMediumURL: entry.AvatarMedium,
			AvatarFullURL:   entry.AvatarFull,
			ProfileURL:      entry.ProfileURL,
			CountryCode:     entry.CountryCode,
		}

		result := tx.
			Where(models.Player{AccountID: entry.AccountID}).
			Attrs(player).
			FirstOrCreate(player)

		if err = result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Fatal()
		}

		pro := &models.ProPlayer{
			AccountID:   entry.AccountID,
			TeamID:      entry.TeamID,
			IsLocked:    entry.IsLocked,
			FantasyRole: entry.FantasyRole,
		}

		if entry.LockedUntil != nil {
			pro.LockedUntil = entry.LockedUntil.Time
		}

		result = tx.
			Where(models.ProPlayer{AccountID: entry.AccountID}).
			Attrs(pro).
			FirstOrCreate(pro)

		if err = result.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Fatal()
		}

		var followed *models.FollowedPlayer

		followed, err = util.FollowPlayer(db, entry.AccountID, entry.Name, false)

		if err != nil {
			if err != util.ErrFollowedPlayerAlreadyExists {
				tx.Rollback()
				l.WithError(err).Fatal()
			}

			l.Warn(err.Error())
			continue
		}

		l.WithField("label", followed.Label).Info("imported")
	}

	if err = tx.Commit().Error; err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("done")
}
