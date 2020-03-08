package players

import (
	"github.com/faceit/go-steam/steamid"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdutil "github.com/13k/night-stalker/cmd/ns/internal/util"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsjson "github.com/13k/night-stalker/internal/json"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
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
	AccountID    nspb.AccountID    `json:"account_id,omitempty"`
	SteamID      nsjson.StringUint `json:"steamid,omitempty"`
	TeamID       models.TeamID     `json:"team_id,omitempty"`
	Name         string            `json:"name,omitempty"`
	PersonaName  string            `json:"personaname,omitempty"`
	Avatar       string            `json:"avatar,omitempty"`
	AvatarMedium string            `json:"avatarmedium,omitempty"`
	AvatarFull   string            `json:"avatarfull,omitempty"`
	ProfileURL   string            `json:"profileurl,omitempty"`
	CountryCode  string            `json:"loccountrycode,omitempty"`
	FantasyRole  nspb.FantasyRole  `json:"fantasy_role,omitempty"`
	IsLocked     bool              `json:"is_locked,omitempty"`
	LockedUntil  *nsjson.UnixTime  `json:"locked_until,omitempty"`
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
		log.WithError(err).Fatal("error")
	}

	if !resp.IsSuccess() {
		log.WithField("status", resp.Status()).Fatal("HTTP error")
	}

	log.WithField("count", len(result)).Info("importing pro players ...")

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
			l.WithError(err).Fatal("error")
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
			l.WithError(err).Fatal("error")
		}

		followed := &models.FollowedPlayer{
			AccountID: entry.AccountID,
			Label:     entry.Name,
		}

		followed, err = nscmdutil.FollowPlayer(db, followed, false)

		if err != nil {
			if err != nscmdutil.ErrFollowedPlayerAlreadyExists {
				tx.Rollback()
				l.WithError(err).Fatal("error")
			}

			l.Warn(err.Error())
			continue
		}

		l.WithField("label", followed.Label).Info("imported")
	}

	if err = tx.Commit().Error; err != nil {
		log.WithError(err).Fatal("error")
	}

	log.Info("done")
}
