package cmdfollow

import (
	"strconv"

	"github.com/spf13/cobra"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	nscmdutil "github.com/13k/night-stalker/cmd/ns/internal/util"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	"github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "follow [flags] <account_id> <label>",
	Short: "Add a player to be stalked",
	Run:   run,
	Args:  cobra.ExactArgs(2),
}

var (
	flagUpdate bool
)

func init() {
	Cmd.Flags().BoolVarP(&flagUpdate, "update", "u", false, "update label of existing player")
}

func run(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	label := args[1]
	accountID64, err := strconv.ParseUint(args[0], 10, 32)

	if err != nil {
		log.WithError(err).WithField("account_id", args[0]).Fatal("invalid account_id value")
	}

	accountID := nspb.AccountID(accountID64)

	db, err := nscmddb.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	followed := &models.FollowedPlayer{
		AccountID: accountID,
		Label:     label,
	}

	followed, err = nscmdutil.FollowPlayer(db, followed, flagUpdate)

	if err != nil {
		log.WithError(err).Fatal("error")
	}

	log.WithOFields(
		"account_id", followed.AccountID,
		"label", followed.Label,
	).Info("following player")
}
