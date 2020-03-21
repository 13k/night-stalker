package players

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdutil "github.com/13k/night-stalker/cmd/ns/internal/util"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
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

	log.WithField("count", len(result)).Info("importing pro players ...")

	for _, entry := range result {
		l := log.WithField("account_id", entry.AccountID)
		tx := db.Begin()

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

			tx.Rollback()
			l.Warn(err.Error())

			continue
		}

		player := &models.Player{
			AccountID:       entry.AccountID,
			SteamID:         nspb.SteamID(entry.SteamID.Uint64()),
			Name:            entry.Name,
			PersonaName:     entry.PersonaName,
			AvatarURL:       entry.Avatar,
			AvatarMediumURL: entry.AvatarMedium,
			AvatarFullURL:   entry.AvatarFull,
			ProfileURL:      entry.ProfileURL,
			CountryCode:     entry.CountryCode,
		}

		dbres := tx.
			Where(models.Player{AccountID: entry.AccountID}).
			Assign(player).
			FirstOrCreate(player)

		if err = dbres.Error; err != nil {
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

		dbres = tx.
			Where(models.ProPlayer{AccountID: entry.AccountID}).
			Assign(pro).
			FirstOrCreate(pro)

		if err = dbres.Error; err != nil {
			tx.Rollback()
			l.WithError(err).Fatal("error")
		}

		if err = tx.Commit().Error; err != nil {
			log.WithError(err).Fatal("error")
		}

		l.WithField("label", followed.Label).Info("imported")
	}

	log.Info("done")
}
