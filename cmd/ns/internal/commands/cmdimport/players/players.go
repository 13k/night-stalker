package players

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

const (
	apiURL  = "https://api.opendota.com"
	apiPath = "/api/proPlayers"
)

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Import players from OpenDota API",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	apiKey := v.GetString(v.KeyOpendotaAPIKey)

	log.Info("fetching players ...")

	client := resty.New().SetHostURL(apiURL)

	if apiKey != "" {
		client.SetQueryParam("api_key", apiKey)
	}

	result := apiResult{}

	res, err := client.R().
		SetResult(&result).
		Get(apiPath)

	if err != nil {
		return xerrors.Errorf("error fetching players: %w", err)
	}

	if !res.IsSuccess() {
		return xerrors.Errorf("HTTP error: %s", res.Status())
	}

	dbs := nsdbda.NewSaver(db)

	log.
		WithField("count", len(result)).
		Info("importing players ...")

	for _, entry := range result {
		r, err := dbs.ImportPlayer(cmd.Context(), &nsdbda.ImportPlayerData{
			AccountID:       entry.AccountID,
			SteamID:         nspb.SteamID(entry.SteamID.Uint64()),
			Label:           entry.Name,
			Name:            entry.Name,
			PersonaName:     entry.PersonaName,
			AvatarURL:       entry.Avatar,
			AvatarMediumURL: entry.AvatarMedium,
			AvatarFullURL:   entry.AvatarFull,
			ProfileURL:      entry.ProfileURL,
			CountryCode:     entry.CountryCode,
			TeamID:          entry.TeamID,
			IsLocked:        entry.IsLocked,
			LockedUntil:     nssql.NullTimeFromUnixJSON(entry.LockedUntil),
			FantasyRole:     entry.FantasyRole,
		})

		if err != nil {
			return xerrors.Errorf("error importing player %d: %w", entry.AccountID, err)
		}

		msg := "updated"

		if r.Created {
			msg = "imported"
		}

		log.WithOFields(
			"account_id", r.FollowedPlayer.AccountID,
			"label", r.FollowedPlayer.Label,
		).Info(msg)
	}

	log.Info("done")

	return nil
}
