package cmdapi

import (
	"github.com/spf13/cobra"
)

const (
	defaultFilenameSuffix = "_model"
	defaultRecvName       = "m"
)

var (
	flagListOnly        bool
	flagOutput          string
	flagFilenameSuffix  string
	flagDefaultRecvName string
)

var CmdAPI = &cobra.Command{
	Use:           "api <package>",
	Short:         "Generate methods to satisfy the Model API",
	RunE:          cmdRun,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	CmdAPI.Flags().BoolVarP(
		&flagListOnly,
		"list",
		"l",
		false,
		`only list files that would be generated`,
	)

	CmdAPI.Flags().StringVarP(
		&flagOutput,
		"output",
		"o",
		"",
		`output directory (empty means same directory of source file, "-" prints to stderr instead)`,
	)

	CmdAPI.Flags().StringVarP(
		&flagFilenameSuffix,
		"suffix",
		"s",
		defaultFilenameSuffix,
		`generated filename suffix`,
	)

	CmdAPI.Flags().StringVarP(
		&flagDefaultRecvName,
		"recv",
		"r",
		defaultRecvName,
		`method receiver name if it cannot be inferred from existing struct methods`,
	)
}

func cmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Usage()
	}

	return Run(&Options{
		Cmd:             cmd,
		Package:         args[0],
		Output:          flagOutput,
		ListOnly:        flagListOnly,
		FilenameSuffix:  flagFilenameSuffix,
		DefaultRecvName: flagDefaultRecvName,
	})
}
