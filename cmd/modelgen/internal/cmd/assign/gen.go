package cmdassign

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/13k/night-stalker/cmd/modelgen/internal/common"
	g "github.com/13k/night-stalker/cmd/modelgen/internal/parser"
)

type genOpts struct {
	*Options

	p     *g.Package
	s     *g.Struct
	fname string
}

type genState struct {
	imports map[string]string
}

func gen(opts *genOpts) error {
	s, p, fname := opts.s, opts.p, opts.fname

	state := &genState{
		imports: make(map[string]string),
	}

	fields, err := parseFields(s)

	if err != nil {
		return fmt.Errorf("error parsing struct %s fields: %w", s.Name, err)
	}

	if len(fields) == 0 {
		return nil
	}

	for _, f := range fields {
		for pkg, pkgAlias := range f.op.pkgs {
			if sAlias := state.imports[pkg]; sAlias != "" && sAlias != pkgAlias {
				return fmt.Errorf(
					"conflicting package aliases %q and %q for package %q",
					sAlias,
					pkgAlias,
					pkg,
				)
			}

			state.imports[pkg] = pkgAlias
		}
	}

	recv := s.GetFirstRecvName(opts.DefaultRecvName)
	b := &bytes.Buffer{}

	if err = execTplAssign(b, p, state, s, recv, fields); err != nil {
		return err
	}

	formatted, err := common.FormatCode(b.Bytes())

	if err != nil {
		return err
	}

	if opts.Output == "-" {
		opts.Cmd.PrintErrln(strings.Repeat("-", 80))
	}

	opts.Cmd.PrintErrln(fname)

	if opts.Output == "-" {
		opts.Cmd.PrintErrln(strings.Repeat("-", 80))
		opts.Cmd.Println(string(formatted))
		return nil
	}

	return common.CreateFile(fname, formatted)
}
