package cmdinspect

import (
	"github.com/spf13/cobra"
)

var CmdInspect = &cobra.Command{
	Use:   "inspect <package>",
	Short: "Inspect package code",
	RunE:  cmdRun,
}

func cmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Usage()
	}

	return Run(&Options{
		Package: args[0],
	})
}
