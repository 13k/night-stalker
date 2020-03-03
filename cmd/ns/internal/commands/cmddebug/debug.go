package cmddebug

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
	// nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	// "github.com/13k/night-stalker/models"
)

var Cmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug something",
	Run:   debug,
}

var CmdPkger = &cobra.Command{
	Use:   "pkger",
	Short: "List all pkger embedded files",
	Run:   debugPkger,
}

func init() {
	Cmd.AddCommand(CmdPkger)
}

func dumpf(format string, values ...interface{}) { //nolint: unused
	dumps := make([]interface{}, len(values))

	for i, v := range values {
		dumps[i] = spew.Sdump(v)
	}

	fmt.Printf(format, dumps...)
}

func debug(cmd *cobra.Command, args []string) {
	/*
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
	*/
}

func debugPkger(cmd *cobra.Command, args []string) {
	log, err := nscmdlog.New()

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
