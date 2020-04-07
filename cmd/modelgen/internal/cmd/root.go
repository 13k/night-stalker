package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cmdapi "github.com/13k/night-stalker/cmd/modelgen/internal/cmd/api"
	cmdassign "github.com/13k/night-stalker/cmd/modelgen/internal/cmd/assign"
	cmdfieldtypes "github.com/13k/night-stalker/cmd/modelgen/internal/cmd/fieldtypes"
	cmdinspect "github.com/13k/night-stalker/cmd/modelgen/internal/cmd/inspect"
)

var CmdRoot = &cobra.Command{
	Use:   "modelgen <command>",
	Short: "Model utilities",
	RunE:  cmdRun,
}

func init() {
	CmdRoot.AddCommand(cmdapi.CmdAPI)
	CmdRoot.AddCommand(cmdassign.CmdAssign)
	CmdRoot.AddCommand(cmdfieldtypes.CmdFieldTypes)
	CmdRoot.AddCommand(cmdinspect.CmdInspect)
}

func cmdRun(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func Execute() {
	if err := CmdRoot.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}
}
