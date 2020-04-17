package players

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	v "github.com/13k/night-stalker/cmd/ns/internal/viper"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
	nsm "github.com/13k/night-stalker/models"
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

	log.
		WithField("count", len(result)).
		Info("importing players ...")

	for _, entry := range result {
		player := &nsm.Player{
			AccountID:       entry.AccountID,
			SteamID:         nspb.SteamID(entry.SteamID.Uint64()),
			TeamID:          nsm.ID(entry.TeamID),
			Name:            entry.Name,
			PersonaName:     entry.PersonaName,
			AvatarURL:       entry.Avatar,
			AvatarMediumURL: entry.AvatarMedium,
			AvatarFullURL:   entry.AvatarFull,
			ProfileURL:      entry.ProfileURL,
			CountryCode:     entry.CountryCode,
			IsLocked:        entry.IsLocked,
			LockedUntil:     nssql.NullTimeFromUnixJSON(entry.LockedUntil),
			FantasyRole:     entry.FantasyRole,
		}

		q := db.
			Q().
			Select().
			Eq(nsm.PlayerTable.Col("account_id"), player.AccountID)

		created, err := db.M().Upsert(cmd.Context(), player, q)

		if err != nil {
			return xerrors.Errorf("error importing player %d: %w", entry.AccountID, err)
		}

		msg := "updated"

		if created {
			msg = "imported"
		}

		log.WithOFields(
			"account_id", player.AccountID,
			"name", player.Name,
		).Info(msg)
	}

	log.Info("done")

	return nil
}
