package cmdfieldtypes

import (
	"fmt"

	"github.com/spf13/cobra"

	mg "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

var CmdFieldTypes = &cobra.Command{
	Use:   "fieldtypes <package>",
	Short: "List types for all struct fields",
	RunE:  cmdRun,
}

func cmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Help()
	}

	pkgName := args[0]
	pkgs, err := mg.Parse(pkgName)

	if err != nil {
		return fmt.Errorf("error parsing package %s: %w", pkgName, err)
	}

	if n := mg.PrintErrors(pkgs); n > 0 {
		return fmt.Errorf("found %d errors while parsing package %s", n, pkgName)
	}

	for _, p := range pkgs {
		for _, s := range p.Structs {
			for _, f := range s.Fields {
				cmd.Printf("%20[1]T   %[1]s\n", f.Type())
			}
		}
	}

	return nil
}
