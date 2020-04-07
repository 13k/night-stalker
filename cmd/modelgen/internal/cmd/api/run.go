package cmdapi

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/13k/night-stalker/cmd/modelgen/internal/common"
	mg "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type Options struct {
	Cmd             *cobra.Command
	Package         string
	Output          string
	ListOnly        bool
	FilenameSuffix  string
	DefaultRecvName string
}

func Run(opts *Options) error {
	pkgs, err := mg.Parse(opts.Package)

	if err != nil {
		return fmt.Errorf("error parsing package %s: %w", opts.Package, err)
	}

	for _, p := range pkgs {
		for _, s := range p.Structs {
			if !s.IsModelStruct() {
				continue
			}

			fname := common.GenFilename(s.Position().Filename, opts.FilenameSuffix)

			if opts.Output != "" && opts.Output != "-" {
				fname = filepath.Join(opts.Output, filepath.Base(fname))
			}

			if opts.ListOnly {
				opts.Cmd.PrintErrln(fname)
				continue
			}

			genOpts := &genOpts{
				Options: opts,
				p:       p,
				s:       s,
				fname:   fname,
			}

			if err := gen(genOpts); err != nil {
				return err
			}
		}
	}

	return nil
}
