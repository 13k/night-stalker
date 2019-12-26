package cmddebug

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/ns/internal/db"
	"github.com/13k/night-stalker/cmd/ns/internal/logger"
	nslog "github.com/13k/night-stalker/internal/logger"
	"github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug something",
	Run:   run,
}

var CmdPkger = &cobra.Command{
	Use:   "pkger",
	Short: "List all pkger embedded files",
	Run:   debugPkger,
}

func init() {
	Cmd.AddCommand(CmdPkger)
}

func run(cmd *cobra.Command, args []string) {
	log, err := logger.New()

	if err != nil {
		panic(err)
	}

	defer log.Close()

	debugDB(log)
}

func debugDB(log *nslog.Logger) {
	db, err := db.Connect()

	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	defer db.Close()

	matchID := uint64(123)

	liveMatch := &models.LiveMatch{}

	result := db.Debug().
		Where("match_id = ?", matchID).
		Find(liveMatch)

	if err := result.Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		log.WithError(err).Fatal()
	}

	log.Printf("%#v", liveMatch)
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
