package cmddebug

import (
	"os"

	"github.com/docker/go-units"
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	nscmdlog "github.com/13k/night-stalker/cmd/ns/internal/logger"
)

var CmdPkger = &cobra.Command{
	Use:   "pkger",
	Short: "List all pkger embedded files",
	RunE:  debugPkger,
}

func init() {
	Cmd.AddCommand(CmdPkger)
}

func debugPkger(cmd *cobra.Command, args []string) error {
	log := nscmdlog.Instance()

	defer log.Close()

	err := pkger.Walk("/", func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		log.WithOFields(
			"is_dir", i.IsDir(),
			"size", i.Size(),
			"size_h", units.BytesSize(float64(i.Size())),
		).Info(p)

		return nil
	})

	if err != nil {
		return xerrors.Errorf("error walking pkger tree: %w", err)
	}

	return nil
}
