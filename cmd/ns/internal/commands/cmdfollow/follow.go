package cmdfollow

import (
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsm "github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "follow [flags] <account_id> <label>",
	Short: "Add a player to be stalked",
	RunE:  run,
	Args:  cobra.ExactArgs(2),
}

var (
	flagUpdate bool
)

func init() {
	Cmd.Flags().BoolVarP(&flagUpdate, "update", "u", false, "update label of existing player")
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return cmd.Usage()
	}

	log := nscmdlog.Instance()

	defer log.Close()

	label := args[1]
	accountID64, err := strconv.ParseUint(args[0], 10, 32)

	if err != nil {
		return xerrors.Errorf("invalid accountID value %q: %w", args[0], err)
	}

	accountID := nspb.AccountID(accountID64)

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	dbs := nsdbda.NewSaver(db)

	followed := &nsm.FollowedPlayer{
		AccountID: accountID,
		Label:     label,
	}

	created, err := dbs.FollowPlayer(cmd.Context(), followed, &nsdbda.FollowPlayerOptions{
		Update: flagUpdate,
	})

	if err != nil {
		return xerrors.Errorf("error following player: %w", err)
	}

	msg := "updated followed player"

	if created {
		msg = "following player"
	}

	log.WithOFields(
		"account_id", followed.AccountID,
		"label", followed.Label,
	).Info(msg)

	return nil
}
