package cmddebug

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmddb "github.com/13k/night-stalker/cmd/ns/internal/db"
	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
)

var Cmd = &cobra.Command{
	Use:   "debug [subcommand]",
	Short: "Debug",
	RunE:  debug,
}

func dumpf(format string, values ...interface{}) { //nolint: unused
	dumps := make([]interface{}, len(values))

	for i, v := range values {
		dumps[i] = spew.Sdump(v)
	}

	fmt.Printf(format, dumps...)
}

func debug(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	db, err := nscmddb.Connect(log)

	if err != nil {
		return xerrors.Errorf("error connecting to database: %w", err)
	}

	defer db.Close()

	log.Info("This is a template command")

	return nil
}
