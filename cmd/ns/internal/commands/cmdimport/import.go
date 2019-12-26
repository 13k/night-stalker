package cmdimport

import (
	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport/d2pt"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport/heroes"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport/players"
	"github.com/13k/night-stalker/cmd/ns/internal/commands/cmdimport/teams"
)

var Cmd = &cobra.Command{
	Use:   "import <source>",
	Short: "Import data",
	Run:   run,
}

func init() {
	Cmd.AddCommand(d2pt.Cmd)
	Cmd.AddCommand(heroes.Cmd)
	Cmd.AddCommand(players.Cmd)
	Cmd.AddCommand(teams.Cmd)
}

func run(cmd *cobra.Command, args []string) {
	if err := cmd.Usage(); err != nil {
		panic(err)
	}
}
