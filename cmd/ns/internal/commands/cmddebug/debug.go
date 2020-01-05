package cmddebug

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/ns/internal/db"
	"github.com/13k/night-stalker/cmd/ns/internal/logger"
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug something",
	Run:   debugDB,
}

var CmdPkger = &cobra.Command{
	Use:   "pkger",
	Short: "List all pkger embedded files",
	Run:   debugPkger,
}

func init() {
	Cmd.AddCommand(CmdPkger)
}

func debugPkger(cmd *cobra.Command, args []string) {
	log, err := logger.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	dir := "/"

	err = pkger.Walk(dir, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		log.WithFields(logrus.Fields{
			"dir":  i.IsDir(),
			"size": i.Size(),
		}).Info(p)

		return nil
	})

	if err != nil {
		log.WithError(err).Fatal()
	}
}

func debugDB(cmd *cobra.Command, args []string) {
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

	type livePlayerResult struct {
		*models.LiveMatchPlayer
		MatchID nspb.MatchID
	}

	var livePlayers []*models.LiveMatchPlayer

	err = db.
		Debug().
		Where(&models.LiveMatchPlayer{AccountID: 140288368}).
		Preload("LiveMatch").
		Find(&livePlayers).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.WithError(err).Fatal("database live players")
	}

	// livePlayers := make(map[nspb.MatchID]*models.LiveMatchPlayer)

	// for _, result := range livePlayerResults {
	// 	livePlayers[result.MatchID] = result.LiveMatchPlayer
	// }

	log.Println(spew.Sdump(livePlayers))
}
